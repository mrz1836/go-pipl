package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-pipl"
)

func main() {
	c := pipl.NewClient(
		pipl.WithAPIKey(os.Getenv("PIPL_API_KEY")),
	)

	log.Println("client loaded:", c.UserAgent())
}
