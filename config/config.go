package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleLoginConfig oauth2.Config

func GoogleConfig() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	GoogleLoginConfig = oauth2.Config{
		RedirectURL: "http://localhost:8080/google/callback",
		ClientID:    os.Getenv("GOOGLE_CLIENT_ID"),
		//ClientID: id,
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		//ClientSecret: secret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return GoogleLoginConfig
}

var (
	auth0Domain       = os.Getenv("YOUR_AUTH0_DOMAIN")
	auth0ClientID     = os.Getenv("YOUR_AUTH0_CLIENT_ID")
	auth0ClientSecret = os.Getenv("YOUR_AUTH0_CLIENT_SECRET")
	auth0CallbackURL  = os.Getenv("YOUR_CALLBACK_URL")
	auth0Scopes       = []string{"openid", "profile", "email", "name"}
)

func Auth0Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     auth0ClientID,
		ClientSecret: auth0ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("https://%s/authorize", auth0Domain),
			TokenURL: fmt.Sprintf("https://%s/oauth/token", auth0Domain),
		},
		RedirectURL: auth0CallbackURL,
		Scopes:      auth0Scopes,
	}
}
