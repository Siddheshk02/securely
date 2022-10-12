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

	"github.com/hashicorp/vault/shamir"
)

var shares int
var threshold int
var n [][]byte
var gcm cipher.AEAD

func Admin(filepath string) (err error) {
	pass := "Hello, World!!"
	fmt.Println("Enter the Number of Shares and the Threshold you want to create: ")
	fmt.Scanf("%d", &shares)
	fmt.Scanf("%d", &threshold)

	c := sha256.New()

	c.Write([]byte(pass))
	key := c.Sum(nil)

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

	err = ioutil.WriteFile("Admin Files/secret.bin", secret, 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("User Files/secret.bin", secret, 0777)
	if err != nil {
		log.Fatal(err)
	}

	n, err = shamir.Split(secret, shares, threshold)
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < shares; x++ {
		path := "Admin Files/share" + strconv.Itoa(x+1) + ".txt"
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

	return err
}
