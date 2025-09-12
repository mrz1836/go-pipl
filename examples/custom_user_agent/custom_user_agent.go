// Package main demonstrates how to use a custom user agent with the PIPL API.
package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-pipl"
)

func main() {
	customUserAgent := "my-custom-app v1.2.3"

	c := pipl.NewClient(
		pipl.WithAPIKey(os.Getenv("PIPL_API_KEY")),
		pipl.WithUserAgent(customUserAgent),
	)

	log.Println("client loaded:", c.UserAgent())
}
