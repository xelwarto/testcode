package main

import (
	"fmt"
	//"os"

	//"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

  //"google.golang.org/api/people/v1"
)

func main() {
  conf := &oauth2.Config {
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_PASSWORD",
		Scopes: []string{ "profile", "email" },
		Endpoint: google.Endpoint,
		RedirectURL: "REDIRECT_URI",
	}

	fmt.Printf("URL: %v\n", conf.AuthCodeURL("blah"))

	//ctx := context.Background()
	/*token, err := conf.Exchange(ctx, "CODE")
	if err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }

  token := &oauth2.Token {
		AccessToken: "TOKEN",
	}

	fmt.Printf("Access Token: %v\n", token.AccessToken)
	fmt.Printf("Refresh Token: %v\n", token.RefreshToken)

  */
}
