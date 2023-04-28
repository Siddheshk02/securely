package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/Siddheshk02/securely/controllers"
	"github.com/Siddheshk02/securely/database"
	"github.com/Siddheshk02/securely/mail"
	"github.com/Siddheshk02/securely/storage"

	"github.com/hashicorp/vault/shamir"
)

func Admin(filep string) (err error) {
	var shares int
	var gcm cipher.AEAD

	pass := "Hello, World!!"
	fmt.Print("\nEnter the Number of Shares you want to create: ")
	fmt.Scan(&shares)
	threshold := int(math.Ceil(float64(shares * 2 / 3)))

	// fmt.Scan(&threshold)

	// dirname, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// st := []byte{byte(shares), byte(threshold)}
	// str := dirname + "/securely/info.txt"

	// fmt.Println("test110")
	// //te := ioutil.WriteFile(str, st, 0777)
	// fmt.Println("test111")

	// if te != nil {
	// 	log.Fatal(err)
	// }

	c := sha256.New()

	c.Write([]byte(pass))
	key := c.Sum(nil)

	loc := filepath.Base(filep)
	//str1 := dirname + "/securely/key.bin"

	data, err := controllers.Whoami()
	if err != nil {
		fmt.Println("Error while getting the User Data. Please Try Again.")
		return
	}
	//fmt.Println("test")

	temp := storage.Files(key, data, loc, 0, 0, 1)
	//temp := ioutil.WriteFile("Admin/key.bin", key, 0777)
	if temp != nil {
		log.Fatal(err)
	}

	//fmt.Println("test1")
	text, err := os.ReadFile(filep)
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

	//str2 := dirname + "/securely/" + filepath

	// err = ioutil.WriteFile(loc, secret, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = ioutil.WriteFile("User/secret.key", secret, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	n, err := shamir.Split(secret, shares, threshold)
	if err != nil {
		log.Fatal(err)
	}

	//str3 := dirname + "/securely/main.txt"

	file, err := os.Create("Admin/main.txt")
	if err != nil {
		log.Fatal("os.Create", err)
	}

	for l := 0; l < shares; l++ {
		fmt.Fprintln(file, n[l])
	}

	var sharelist []string = make([]string, shares)
	for x := 0; x < shares; x++ {
		//path := dirname + "/securely/share" + strconv.Itoa(x+1) + ".txt"
		//path := "Admin/share" + strconv.Itoa(x+1) + ".txt"
		sharelist[x] = base64.StdEncoding.EncodeToString(n[x])

		// file1 := ioutil.WriteFile(path, n[x], 0777)
		// if file1 != nil {
		// 	log.Fatal(err)
		// }

	}

	//storage.Upload(loc, data, shares, threshold)
	storage.SharesFile(n, data, loc, shares)
	storage.Files(secret, data, loc, shares, threshold, 3)

	fmt.Println("Select the Users who can access this file using Shares : ")
	fmt.Println(" ")
	database.AddUser(data, shares, sharelist, loc)

	return err

}

func User(filename string, admin string, user []byte) (err error) {

label:

	// fmt.Print("Enter the Number of Secret Shares you want to enter: ")
	// j := 0
	// fmt.Scan(&j)

	shares, threshold := database.Read(filename, admin)
	// fmt.Println(shares, threshold)

	//b, err := os.ReadFile("Admin/info.txt")

	// shares := int(b[0])
	// threshold := int(b[1])

	// if j < threshold {
	// 	fmt.Printf("Please enter the minimum number of Shares i.e. %d\n", threshold)
	// 	fmt.Println(" ")

	// 	goto label
	// } else if j > shares {
	// 	fmt.Println("Exceeded the Number of Shares!!")

	// 	goto label
	// }
	n, key, ciphertext, err := storage.ReadShares(filename, admin, shares)
	if err != nil {
		fmt.Println("Error while reading the shares")
		return
	}

	var parts [][]byte

	//parts = make([][]byte, threshold)
	// fmt.Println(len(parts))
	//var parts [][]byte
	a := 53
	// fmt.Println("test1")
	// fmt.Println(parts[0][0])

	for i := 0; i < threshold; i++ {
		// parts[i] = make([]byte, a)
		fmt.Print("Enter the Secret Share: ")
		// // var f, g byte
		// // fmt.Scan(f)
		// // parts[i][0] = f
		// // fmt.Scan(g)
		// // parts[i][0] = g
		short := make([]byte, 53)
		for x := 0; x < a; x++ {
			fmt.Scan(&short[x])

		}
		parts = append(parts, short)
		fmt.Println(" ")

	}
	// fmt.Println("test2")
	// //fmt.Println("Check 1")

	// //Checking for duplicate share entered by User
	// for i := 0; i < threshold; i++ {
	// 	fmt.Println(parts[i])
	// }
	// for i := 0; i < threshold; i++ {
	// 	parts[i] = make([]byte, a)
	// }

	// // Read the input array from the user
	// fmt.Println("Enter the byte array elements:")
	// for i := 0; i < threshold; i++ {
	// 	for j := 0; j < a; j++ {
	// 		fmt.Printf("parts[%d][%d]: ", i, j)
	// 		fmt.Scanf("%d", &parts[i][j])
	// 	}
	// }

	// Print the input array
	// fmt.Println("Input byte array:")
	// for i := 0; i < threshold; i++ {
	// 	fmt.Println(parts[i])
	// }

	var con string

	//fmt.Println("Check 2")
	for h := 0; h < threshold-1; h++ {
		//fmt.Println("Check10")

		if (parts[h][0] == parts[h+1][0]) && (parts[h][1] == parts[h+1][1]) {
			//fmt.Println("Check11")
			fmt.Println("Share ", h+1, " is repeated")
			fmt.Print("Do you Want to continue? Yes/No :")
			fmt.Scan(&con)
			fmt.Println(" ")
			if con == "Yes" {
				goto label
			} else {
				err := errors.New("Process Stopped!")
				return err
			}

		}
	}

	//fmt.Println("Check 3")
	//Checkinh if The Share entered is Valid or not - start
	boolean := true
	var loc int

	//var n [100][49]byte
	// for o := 0; o < shares; o++ {

	// 	filetemp := "Admin/share" + strconv.Itoa(o+1) + ".txt"

	// 	abc, err := os.ReadFile(filetemp)
	// 	xyz := string(abc)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for eg := 0; eg < 49; eg++ {
	// 		n[o][eg] = xyz[eg]
	// 	}

	// }

	for i := 0; i < threshold; i++ {
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
	//Checking if The Share entered is Valid or not - end

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
			err := errors.New("Process Stopped!")
			return err
		}
	}

	//Decrypting the file
	// ciphertext, err := os.ReadFile(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// str1 := "Admin/key.bin"
	// key, err := os.ReadFile(str1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

	fmt.Print("Enter file path to save the downloaded file: ")
	var savePath string
	fmt.Scan(&savePath)

	err = ioutil.WriteFile(savePath+"/"+filename, plaintext, 0777)
	if err != nil {
		log.Panic(err)
	}

	var result map[string]interface{}
	json.Unmarshal(user, &result)
	name := result["name"].(string)
	email := result["email"].(string)

	err = database.AccessLogs(admin, filename, name, email)

	ademail := database.AdminEmail(admin, filename)
	mail.MailAdmin(admin, ademail, filename, name)

	return nil
}
