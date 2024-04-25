package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("./my-test-project-370312-903e2fe6d4b8.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.CredentialsFromJSONWithParams(ctx, b, google.CredentialsParams{
		Scopes:  []string{admin.AdminDirectoryUserReadonlyScope, admin.AdminDirectoryUserScope},
		Subject: "atsushi.miyazaki@the-udon-empire.com",
	})
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := admin.NewService(ctx, option.WithCredentials(config))
	if err != nil {
		log.Fatalf("Unable to retrieve directory Client %v", err)
	}

	// fmt.Println(srv.Customers)
	// r, err := srv.Users.Get("atsushi.miyazaki@the-udon-empire.com").Do()
	r, err := srv.Users.Get("taro.test@the-udon-empire.com").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve users in domain: %v", err)
	}

	v, _ := json.Marshal(r)
	fmt.Println(string(v))
}
