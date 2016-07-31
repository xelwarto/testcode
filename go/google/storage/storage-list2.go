package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

var bucket = "bucketname"

func main() {
	conf := &oauth2.Config {
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_PASSWORD",
		Scopes: []string{ "https://www.googleapis.com/auth/devstorage.read_write" },
		Endpoint: google.Endpoint,
		RedirectURL: "REDIRECT_URI",
	}

	fmt.Printf("URL: %v\n", conf.AuthCodeURL("blah"))

	ctx := context.Background()
	/*token, err := conf.Exchange(ctx, "CODE")
	if err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }*/

	token := &oauth2.Token {
		AccessToken: "TOKEN",
	}

	fmt.Printf("Access Token: %v\n", token.AccessToken)
	fmt.Printf("Refresh Token: %v\n", token.RefreshToken)

  client, err := storage.NewClient(ctx, cloud.WithTokenSource(conf.TokenSource(ctx, token)))
  if err != nil {
      fmt.Printf("Error: %v\n\n", err)
  }

	for {
    var query *storage.Query
    objects, err := client.Bucket(bucket).List(ctx, query)
    if err != nil {
        fmt.Printf("HERE Error: %v\n\n", err)
        os.Exit(1)
    }

    for _, obj := range objects.Results {
        fmt.Printf("object name: %s, size: %v\n", obj.Name, obj.Size)
    }
    // If there are more results, objects.Next will be non-nil.
    if objects.Next == nil {
        break
    }
    query = objects.Next
  }

}
