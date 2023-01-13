package controllers

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Siddheshk02/securely/config"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	temp := config.GoogleConfig()
	url := temp.AuthCodeURL("randomstate")

	//c.Status(fiber.StatusSeeOther)
	//open.Run(url)
	//http.Redirect(w, r, url, 302)
	/*if err := http.Redirect(w, r, url, http.StatusSeeOther); err != nil {
		panic(fmt.Errorf("failed to open browser for authentication %s", err.Error()))
	}*/
	http.Redirect(w, r, url, http.StatusSeeOther)
	//return open.Json(url)
	//return err
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
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

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Fatalf("User Data Fetch Failed %v", err)
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("JSON Parsing Failed %v", err)
	}

	//return r.sendstring(userData)
	//fmt.Println(string(userData))
	print(string(userData))
}
