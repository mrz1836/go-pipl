package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mrz1836/go-pipl"
)

func main() {
	customHTTPClient := &http.Client{}

	customHTTPClientOptions := pipl.DefaultHTTPOptions()

	c := pipl.NewClient(
		pipl.WithAPIKey(os.Getenv("PIPL_API_KEY")),
		pipl.WithHTTPClient(customHTTPClient),
		pipl.WithHTTPOptions(customHTTPClientOptions),
	)

	log.Println("client loaded:", c.UserAgent())
}
