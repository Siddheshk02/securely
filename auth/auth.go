package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Siddheshk02/securely/controllers"
	"github.com/gorilla/mux"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func Signup() error {
	r := mux.NewRouter()
	r.HandleFunc("/google/login", controllers.GoogleLogin).Methods("GET")
	r.HandleFunc("/google/callback", controllers.GoogleCallback).Methods("GET")
	l, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Fatal(err)
	}
	//ch = 0

	//browser.OpenURL("http://localhost:8080/google/login")

	open.Start("http://localhost:8088/google/login")
	//http.ListenAndServe(":8080", r)
	http.Serve(l, r)
	return err
}

func Check() bool {
	var token *oauth2.Token

	if !controllers.FileExists("token.json") {
		file, err := ioutil.ReadFile("token.json")
		if err != nil {
			//fmt.Printf("Error occured!, try again.")
			return true
		}
		err = json.Unmarshal(file, &token)
		if err != nil {
			//fmt.Printf("Error unmarshalling JSON: %v\n", err)
			return true
		}

		// Check if the token is expired
		if time.Now().After(token.Expiry) {
			//fmt.Printf("Token has expired\n")
			os.Remove("token.json")
			return true
		} else {
			return false
		}

	} else {
		return true
	}

}

func UserSignup() error {
	r := mux.NewRouter()
	r.HandleFunc("/google/login", controllers.UserGoogleLogin).Methods("GET")
	r.HandleFunc("/google/callback", controllers.UserGoogleCallback).Methods("GET")
	l, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Fatal(err)
	}
	//ch = 0

	//browser.OpenURL("http://localhost:8080/google/login")

	open.Start("http://localhost:8088/google/login")
	//http.ListenAndServe(":8080", r)
	http.Serve(l, r)
	return err
}

func UserCheck() bool {
	var token *oauth2.Token

	if !controllers.UserFileExists("usertoken.json") {
		file, err := ioutil.ReadFile("usertoken.json")
		if err != nil {
			//fmt.Printf("Error occured!, try again.")
			return true
		}
		err = json.Unmarshal(file, &token)
		if err != nil {
			//fmt.Printf("Error unmarshalling JSON: %v\n", err)
			return true
		}

		// Check if the token is expired
		if time.Now().After(token.Expiry) {
			//fmt.Printf("Token has expired\n")
			os.Remove("usertoken.json")
			return true
		} else {
			return false
		}

	} else {
		return true
	}

}
