package auth

import (
	"encoding/json"
	"fmt"
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
	// r.HandleFunc("/google/login", func(w http.ResponseWriter, r *http.Request) {
	// 	temp := config.GoogleConfig()
	// 	url := temp.AuthCodeURL("randomstate")

	// 	//c.Status(fiber.StatusSeeOther)
	// 	//open.Run(url)
	// 	//http.Redirect(w, r, url, 302)
	// 	/*if err := http.Redirect(w, r, url, http.StatusSeeOther); err != nil {
	// 		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
	// 	}*/
	// 	http.Redirect(w, r, url, http.StatusSeeOther)
	// })
	r.HandleFunc("/google/callback", controllers.GoogleCallback).Methods("GET")
	// r.HandleFunc("/google/callback", func(w http.ResponseWriter, r *http.Request) {
	// 	query := r.URL.Query()
	// 	state := query.Get("state")
	// 	if state != "randomstate" {
	// 		log.Fatal("States don't Match!!")
	// 	}

	// 	code := query.Get("code")

	// 	googlecon := config.GoogleConfig()

	// 	token, err := googlecon.Exchange(context.Background(), code)
	// 	if err != nil {
	// 		log.Fatalf("Code-Token Exchange Failed %v", err)
	// 	}

	// 	currentUser, err := user.Current()
	// 	if err != nil {
	// 		log.Fatal("Error occured!, try again.")
	// 	}

	// 	path := currentUser.HomeDir + "/securely"
	// 	err = os.MkdirAll(path, os.ModePerm)
	// 	if err != nil {
	// 		log.Fatal("Error occured!, try again.")
	// 	}

	// 	// Construct the path to the token file in the user's home directory
	// 	tokenPath := currentUser.HomeDir + "/securely/token.json"

	// 	file, err := os.Create(tokenPath)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	defer file.Close()

	// 	// encode token into json format
	// 	err = json.NewEncoder(file).Encode(token)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// temp = string(token.AccessToken)
	// 	// File, err := os.Create("TokenFile.txt")
	// 	// if err != nil {
	// 	// 	log.Fatal("ERROR! ", err)
	// 	// }

	// 	// defer File.Close()

	// 	//File.Write(string(token))
	// 	// File.WriteString(temp)

	// 	fmt.Println("Authentication successfully done, you are now logged in")

	// 	//accesstoken = token.AccessToken

	// 	/*resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// 	if err != nil {
	// 		log.Fatalf("User Data Fetch Failed %v", err)
	// 	}

	// 	userData, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatalf("JSON Parsing Failed %v", err)
	// 	}*/

	// 	//return r.sendstring(userData)
	// 	fmt.Fprint(w, "Authentication Successfully done. You can now get back to the CLI Terminal.")
	// 	//fmt.Fprint(w, string(userData))
	// 	//os.Exit(0)
	// 	defer func() {
	// 		var comp string
	// 		fmt.Println("\nEnter The Company/Organization Name : ")
	// 		fmt.Scanf("%s", &comp)

	// 		data, err := controllers.Whoami()
	// 		if err != nil {
	// 			fmt.Println("Error while getting the User Data. Please Try Again.")
	// 			return
	// 		}
	// 		abc := "admin"
	// 		ad := ""
	// 		database.DBconn(data, comp, ad, abc)
	// 		var result map[string]interface{}
	// 		json.Unmarshal(data, &result)

	// 		name := result["name"].(string)
	// 		email := result["email"].(string)

	// 		//fmt.Println(name, email)

	// 		err = mail.SendMail(name, email)
	// 	}()

	// 	return

	// })

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	//ch = 0

	//browser.OpenURL("http://localhost:8080/google/login")
	fmt.Scan(" ")

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
	// r.HandleFunc("/google/login", func(w http.ResponseWriter, r *http.Request) {
	// 	temp := config.GoogleConfig()
	// 	url := temp.AuthCodeURL("randomstate")

	// 	//c.Status(fiber.StatusSeeOther)
	// 	//open.Run(url)
	// 	//http.Redirect(w, r, url, 302)
	// 	/*if err := http.Redirect(w, r, url, http.StatusSeeOther); err != nil {
	// 		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
	// 	}*/
	// 	http.Redirect(w, r, url, http.StatusSeeOther)
	// })
	r.HandleFunc("/google/callback", controllers.UserGoogleCallback).Methods("GET")
	// r.HandleFunc("/google/callback", func(w http.ResponseWriter, r *http.Request) {
	// 	query := r.URL.Query()
	// 	state := query.Get("state")
	// 	if state != "randomstate" {
	// 		log.Fatal("States don't Match!!")
	// 	}

	// 	code := query.Get("code")

	// 	googlecon := config.GoogleConfig()

	// 	token, err := googlecon.Exchange(context.Background(), code)
	// 	if err != nil {
	// 		log.Fatalf("Code-Token Exchange Failed %v", err)
	// 	}

	// 	currentUser, err := user.Current()
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}

	// 	path := currentUser.HomeDir + "/securely"
	// 	err = os.MkdirAll(path, os.ModePerm)
	// 	if err != nil {
	// 		log.Fatal("Error occured!, try again.")
	// 	}

	// 	// Construct the path to the token file in the user's home directory
	// 	tokenPath := currentUser.HomeDir + "/securely/usertoken.json"

	// 	file, err := os.Create(tokenPath)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	defer file.Close()

	// 	// encode token into json format
	// 	err = json.NewEncoder(file).Encode(token)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	fmt.Println("Authentication successfully done, you are now logged in")

	// 	// fmt.Fprint(w, "Authentication Successfully done. You can now get back to the CLI Terminal.")
	// 	//fmt.Fprint(w, string(userData))
	// 	//os.Exit(0)
	// 	defer func() {
	// 		var comp, ad, dmin string
	// 		fmt.Println("\nEnter The Company/Organization Name and the Admin name : ")
	// 		fmt.Scanf("%s %s %s", &comp, &ad, &dmin)

	// 		data, err := controllers.WhoamiUser()
	// 		if err != nil {
	// 			fmt.Println("Error while getting the User Data. Please Try Again.")
	// 			return
	// 		}
	// 		abc := "user"
	// 		admin1 := ad + " " + dmin
	// 		database.DBconn(data, comp, admin1, abc)
	// 		var result map[string]interface{}
	// 		json.Unmarshal(data, &result)

	// 		name := result["name"].(string)
	// 		email := result["email"].(string)

	// 		// fmt.Println(name, email)

	// 		err = mail.SendMail(name, email)
	// 		fmt.Fprint(w, "Authentication Successfully done. You can now get back to the CLI Terminal.")
	// 	}()

	// 	return
	// })

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
