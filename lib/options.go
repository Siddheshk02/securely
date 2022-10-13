package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/vault/shamir"
)

var n [][]byte

func Admin(filepath string) (err error) {
	var shares int
	var threshold int
	//var n [][]byte
	var gcm cipher.AEAD

	pass := "Hello, World!!"
	fmt.Println("Enter the Number of Shares and the Threshold you want to create: ")
	fmt.Scanf("%d", &shares)
	fmt.Scanf("%d", &threshold)
	tem, err := os.Create("Admin/info.bin")
	//tem := ioutil.WriteFile("Admin/info.bin", shares, 0777)
	//tem = ioutil.WriteFile("Admin/info.bin", threshold, 0777)
	tem.WriteString(string(shares))
	tem.WriteString(string(threshold))
	if err != nil {
		log.Fatal(err)
	}

	c := sha256.New()

	c.Write([]byte(pass))
	key := c.Sum(nil)

	temp := ioutil.WriteFile("Admin/key.bin", key, 0777)
	if temp != nil {
		log.Fatal(err)
	}

	text, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err = cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	var secret []byte = gcm.Seal(nonce, nonce, text, nil)

	err = ioutil.WriteFile("Admin/secret.bin", secret, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("User/secret.bin", secret, 0777)
	if err != nil {
		log.Fatal(err)
	}

	n, err = shamir.Split(secret, shares, threshold)
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < shares; x++ {
		path := "Admin/share" + strconv.Itoa(x+1) + ".txt"
		file1, err := os.Create(path)
		if err != nil {
			log.Fatal("os.Create", err)
		}
		fmt.Fprintln(file1, n[x])
	}
	//fmt.Println("Shares created Successfully")
	return err

}

func User(filepath string) (err error) {
	//if len(os.Args) != 2 {
	//	fmt.Println("Please provide a filepath")
	//	return
	//}
	//filepath := os.Args[1]

	var j int = 0
label:

	fmt.Println("Enter the Number of Secret Shares you want to enter: ")
	fmt.Scanf(" %d", &j)

	b, err := ioutil.ReadFile("Admin/info.bin")

	lines := strings.Split(string(b), "\n")
	s := lines[len(lines)-2]
	z := lines[len(lines)-1]
	var shares, threshold int
	fmt.Sscanf(s, "%d", &shares)
	fmt.Sscanf(z, "%d", &threshold)
	//shares := strconv.Atoi(s)
	//threshold := strconv.Atoi()

	if j < threshold {
		fmt.Println("Please enter the minimum number of Shares i.e. 2")
		fmt.Println(" ")

		goto label
	} else if j > shares {
		fmt.Println("Exceeded the Number of Shares!!")

		goto label
	}

	var parts [10][51]byte
	a := 51

	for i := 0; i < j; i++ {
		fmt.Print("Enter the Secret Share: ")
		for x := 0; x < a; x++ {
			fmt.Scan(&parts[i][x])

		}
		fmt.Println(" ")

	}

	for h := 0; h < j; h++ {
		var con string
		if (parts[h][0] == parts[h+1][0]) && (parts[h][1] == parts[h+1][1]) {
			fmt.Println("Share ", h+1, " is repeated")
			fmt.Print("Do you Want to continue? Yes/No :")
			fmt.Scan(&con)
			fmt.Println(" ")
			if con == "Yes" {
				goto label
			} else {
				return
			}

		}
	}

	boolean := true
	var loc int

	for i := 0; i < j; i++ {
		boolean = true
		for x := 0; x < shares; x++ {
			if (parts[i][0] == n[x][0]) && (parts[i][1] == n[x][1]) {
				//var eg [60]byte = parts[i]
				//bool := bytes.Equal(eg, n[x])

				for p := 0; p < len(n[0]); p++ {
					if parts[i][p] != n[x][p] {
						boolean = false
						loc = i + 1
						goto label1
					}
				}
			}
		}
	}

label1:

	if boolean == false {
		var con string
		fmt.Println("Invalid Share ", loc)
		fmt.Print("Do you Want to continue? Yes/No :")
		fmt.Scan(&con)
		fmt.Println(" ")
		if con == "Yes" {
			goto label
		} else {
			return
		}
	}

	//fileloc := "C:/Users/user/go/src/github.com/Siddheshk02/secret-sharing/Users/secret.bin"

	ciphertext, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	key, err := os.ReadFile("Admin/key.bin")
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce1 := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce1, ciphertext, nil)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile("User/encrypted.txt", plaintext, 0777)
	if err != nil {
		log.Panic(err)
	}
	//fmt.Println("File Decrypted Successfully!!")

	return err
}
