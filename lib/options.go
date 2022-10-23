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

func Admin(filepath string) (err error) {
	var shares int
	var threshold int
	var gcm cipher.AEAD

	pass := "Hello, World!!"
	fmt.Println("Enter the Number of Shares and the Threshold you want to create: ")
	fmt.Scanf("%d", &shares)
	fmt.Scanf("%d", &threshold)
	st := []byte{byte(shares), byte(threshold)}
	pa := "Admin/info.txt"
	tem, err := os.Create(pa)
	fmt.Fprintln(tem, st[0])
	fmt.Fprintln(tem, st[1])
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

	n, err := shamir.Split(secret, shares, threshold)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("Admin/main.txt")
	if err != nil {
		log.Fatal("os.Create", err)
	}

	for l := 0; l < shares; l++ {
		fmt.Fprintln(file, n[l])
	}

	for x := 0; x < shares; x++ {
		path := "Admin/share" + strconv.Itoa(x+1) + ".txt"
		file1, err := os.Create(path)
		if err != nil {
			log.Fatal("os.Create", err)
		}
		fmt.Fprintln(file1, n[x])
	}
	return err

}

func User(filepath string) (err error) {

label:
	var j int = 0
	fmt.Println("Enter the Number of Secret Shares you want to enter: ")
	fmt.Scanf("%d", &j)
	pa := "Admin/info.txt"

	b, err := ioutil.ReadFile(pa)

	lines := strings.Split(string(b), "\n")
	s := lines[len(lines)-2]
	z := lines[len(lines)-1]
	var shares, threshold int
	fmt.Sscanf(s, "%d", &shares)
	fmt.Sscanf(z, "%d", &threshold)

	if j < threshold {
		fmt.Println("Please enter the minimum number of Shares i.e. 2")
		fmt.Println(" ")

		goto label
	} else if j > shares {
		fmt.Println("Exceeded the Number of Shares!!")

		goto label
	}

	var parts [5][51]byte
	a := 51

	for i := 0; i < j; i++ {
		fmt.Print("Enter the Secret Share: ")
		for x := 0; x < a; x++ {
			fmt.Scan(&parts[i][x])

		}
		fmt.Println(" ")

	}

	for h := 0; h < j-1; h++ {
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

	var n [5][51]byte

	for o := 0; o < j; o++ {
		filetemp := "Admin/share" + strconv.Itoa(o+1) + ".txt"

		abc, err := ioutil.ReadFile(filetemp)
		xyz := string(abc)

		if err != nil {
			log.Fatal(err)
		}

		for eg := 0; eg < 51; eg++ {
			n[o][eg] = xyz[eg]
		}

	}

	for i := 0; i < j; i++ {
		boolean = true
		for x := 0; x < shares; x++ {
			if (parts[i][0] == n[x][0]) && (parts[i][1] == n[x][1]) {

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

	return err
}
