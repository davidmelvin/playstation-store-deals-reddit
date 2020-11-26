package main

import (
	"context"
	"fmt"

	"log"

	"github.com/jzelinskie/geddit"
	"github.com/kelseyhightower/envconfig"
)

type Credentials struct {
	ClientID     string `envconfig:"REDDIT_CLIENT_ID"`
	ClientSecret string `envconfig:"REDDIT_CLIENT_SECRET"`
	Username     string `envconfig:"REDDIT_USERNAME"`
	Password     string `envconfig:"REDDIT_PASSWORD"`
}

func getSubmissions(ctx context.Context) ([]*geddit.Submission, error) {
	// adapted from https://gist.github.com/sergeyklay/9f51bcd6905788382b5deb790c68eac1
	// config := clientcredentials.Config{
	// 	ClientID:     os.Getenv("REDDIT_CLIENT_ID"),
	// 	ClientSecret: os.Getenv("REDDIT_CLIENT_SECRET"),
	// 	TokenURL:     "https://www.reddit.com/api/v1/access_token",
	// }

	// client := config.Client(ctx)

	// resp, err := client.Get("https://oauth.reddit.com/api/v1/me")
	// if err != nil {
	// 	fmt.Printf("Error making submissions request: %s", err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Error reading response body bytes: %s", err)
	// }
	// defer resp.Body.Close()

	// fmt.Println("body", body)

	var creds Credentials
	err := envconfig.Process("REDDIT", &creds)
	if err != nil {
		log.Fatalf("cannot get env vars for credentials: %s", err)
	}
	fmt.Printf("creds: %#v", creds)

	o, err := geddit.NewOAuthSession(
		creds.ClientID,
		creds.ClientSecret,
		"User agent: golang v0.1 https://github.com/davidmelvin/playstation-store-deals-reddit (by /u/GamingDealsBot)",
		"http://localhost:8080",
	)
	if err != nil {
		log.Fatalf("Unable to create geddit oauth session: %s\n", err)
	}
	err = o.LoginAuth(creds.Username, creds.Password)
	if err != nil {
		log.Fatalf("Unable to get auth token: %s\n", err)
	}

	subOpts := geddit.ListingOptions{
		Limit: 10,
	}

	submissions, err := o.SubredditSubmissions("ps4Deals", geddit.NewSubmissions, subOpts)
	if err != nil {
		log.Printf("Unable to get subreddit posts: %s\n", err)
	}

	fmt.Println(submissions)

	return []*geddit.Submission{}, nil
}
