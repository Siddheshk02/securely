package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/Siddheshk02/securely/config"

	firebase "firebase.google.com/go"
)

var (
	app    *firebase.App
	client *auth.Client
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	conf := config.Auth0Config()
	url := conf.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Exchange the authorization code for a Firebase user
	query := r.URL.Query()
	state := query.Get("state")
	if state != "randomstate" {
		log.Fatal("States don't Match!!")
	}

	conf := config.Auth0Config()
	ctx := context.Background()
	code := r.URL.Query().Get("code")
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(token)
	//fmt.Fprintf(w, "User ID: %v\n", user.UID)
	fmt.Fprintf(w, "done")
}
