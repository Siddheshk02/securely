package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/Siddheshk02/securely/config"
	"github.com/Siddheshk02/securely/database"
	"github.com/Siddheshk02/securely/mail"
	"golang.org/x/oauth2"
)

func UserGoogleLogin(w http.ResponseWriter, r *http.Request) {
	temp := config.GoogleConfig()
	// fmt.Println(temp)
	url := temp.AuthCodeURL("randomstate")

	// fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	//return open.Json(url)
	//return err
}

func UserGoogleCallback(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	state := query.Get("state")
	if state != "randomstate" {
		log.Fatal("States don't Match!!")
	}

	code := query.Get("code")

	googlecon := config.GoogleConfig()

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Code-Token Exchange Failed %v", err)
	}

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	path := currentUser.HomeDir + "/securely"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	// Construct the path to the token file in the user's home directory
	tokenPath := currentUser.HomeDir + "/securely/usertoken.json"

	file, err := os.Create(tokenPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// encode token into json format
	err = json.NewEncoder(file).Encode(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Authentication successfully done, you are now logged in")

	fmt.Fprint(w, "Authentication Successfully done. You can now get back to the CLI Terminal.")
	//fmt.Fprint(w, string(userData))
	//os.Exit(0)
	defer func() {
		var comp, ad, dmin string
		fmt.Println("\nEnter The Company/Organization Name and the Admin name : ")
		fmt.Scanf("%s %s %s", &comp, &ad, &dmin)

		data, err := WhoamiUser()
		if err != nil {
			fmt.Println("Error while getting the User Data. Please Try Again.")
			return
		}
		abc := "user"
		admin1 := ad + " " + dmin
		database.DBconn(data, comp, admin1, abc)
		var result map[string]interface{}
		json.Unmarshal(data, &result)

		name := result["name"].(string)
		email := result["email"].(string)

		// fmt.Println(name, email)

		err = mail.SendMail(name, email)
	}()

	return

}

func WhoamiUser() ([]byte, error) {
	var token *oauth2.Token
	//temp := accesstoken
	// fmt.Println("1")
	// fmt.Println(temp)
	// fmt.Println("2")
	//file, err := os.ReadFile("TokenFile.txt")
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	// Construct the path to the token file in the user's home directory
	tokenPath := currentUser.HomeDir + "/securely/usertoken.json"

	file, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}
	err = json.Unmarshal(file, &token)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	if err != nil {
		log.Fatal("The User is not logged in.")
	}
	//fmt.Println(file)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Fatalf("User Data Fetch Failed %v", err)
	}
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("JSON Parsing Failed %v", err)
	}
	/*token1, err := oauth2.Token(context.Background(), token.AccessToken)
	if err != nil {
		return false
	}*/
	//fmt.Println(string(userData))
	return userData, err

}

func UserLogout() error {

	err := DeleteTokenUser()

	if err != nil {
		return err
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	// Construct the path to the token file in the user's home directory
	tokenPath := currentUser.HomeDir + "/securely/usertoken.json"

	err = os.Remove(tokenPath)
	if err != nil {
		return err
	}

	fmt.Println("Logout Successfully!!")
	return nil
}

func DeleteTokenUser() error {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	// Construct the path to the token file in the user's home directory
	tokenPath := currentUser.HomeDir + "/securely/usertoken.json"

	file, err := os.Open(tokenPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// decode token from json format
	var tokenRead *oauth2.Token
	err = json.NewDecoder(file).Decode(&tokenRead)
	if err != nil {
		log.Fatal(err)
	}

	googlecon := config.GoogleConfig()

	client := googlecon.Client(context.Background(), tokenRead)

	resp, err := client.Get("https://accounts.google.com/o/oauth2/revoke?token=" + tokenRead.AccessToken)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	return nil
}

func UserFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}
