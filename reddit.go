package main

import (
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
)

func getSubmissions() ([]*geddit.Submission, error) {
	fmt.Println("here are your posts, ma'am")
	fmt.Printf("secret: %s", os.Getenv("REDDIT_CLIENT_SECRET"))

	return []*geddit.Submission{}, nil
}
