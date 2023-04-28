package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/Siddheshk02/securely/controllers"
	"github.com/gorilla/mux"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func Auth0signup() error {

	// Create a new Gorilla router
	r := mux.NewRouter()

	// Handle the login page
	r.HandleFunc("/login", controllers.HandleLogin).Methods("GET").Host("dev-securely.us.auth0.com")

	// Handle the OAuth callback from the Google provider
	r.HandleFunc("/login/callback", controllers.HandleCallback).Methods("GET").Host("dev-securely.us.auth0.com")

	open.Start("https://dev-securely.us.auth0.com/login")
	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":0", r))

	return nil

}

func Signup() error {
	r := mux.NewRouter()
	r.HandleFunc("/google/login", controllers.GoogleLogin).Methods("GET")
	r.HandleFunc("/google/callback", controllers.GoogleCallback).Methods("GET")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	//ch = 0

	//browser.OpenURL("http://localhost:8080/google/login")

	open.Start("http://localhost:8080/google/login")
	//http.ListenAndServe(":8080", r)
	http.Serve(l, r)
	return err
}

func Check() bool {
	var token *oauth2.Token

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	// Construct the path to the token file in the user's home directory
	tokenPath := currentUser.HomeDir + "/securely/token.json"

	if !controllers.FileExists(tokenPath) {
		file, err := ioutil.ReadFile(tokenPath)
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
			os.Remove(tokenPath)
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
	// .Host("securelee.tech")
	r.HandleFunc("/google/login", controllers.UserGoogleLogin).Methods("GET")
	r.HandleFunc("/google/callback", controllers.UserGoogleCallback).Methods("GET")
	//l, err := net.Listen("tcp", ":9091")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	//ch = 0
	// server := &http.Server{
	// 	Addr:    ":443",
	// 	Handler: r,
	// }

	// Start the server
	open.Start("http://localhost:8080/google/login")

	// log.Printf("Starting server on port 443")
	// err := server.ListenAndServeTLS("certificate.crt", "private.key")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//browser.OpenURL("http://localhost:8080/google/login")

	// open.Start("https://securelee.tech/google/login")
	//http.ListenAndServe(":8080", r)
	// err := http.ListenAndServeTLS(":443", "certificate.crt", "private.key", r)
	// if err != nil {
	// 	log.Fatal("ListenAndServeTLS: ", err)
	// }

	http.Serve(l, r)
	return nil
}

func UserCheck() bool {
	var token *oauth2.Token

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	// Construct the path to the token file in the user's home directory
	tokenPath := currentUser.HomeDir + "/securely/usertoken.json"

	if !controllers.UserFileExists(tokenPath) {
		file, err := ioutil.ReadFile(tokenPath)
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
			os.Remove(tokenPath)
			return true
		} else {
			return false
		}

	} else {
		return true
	}

}
