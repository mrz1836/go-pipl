package main

import (
	"context"
	"log"
	"os"

	"github.com/mrz1836/go-pipl"
)

func main() {

	// Load the client with your API key
	c := pipl.NewClient(
		pipl.WithAPIKey(os.Getenv("PIPL_API_KEY")),
	)

	// Submit the search request
	response, err := c.SearchByPointer(context.Background(), os.Getenv("PIPL_SEARCH_POINTER"))
	if err != nil {
		log.Fatalln(err.Error())
	} else if response == nil {
		log.Fatalln("person not found")
	}

	log.Println("found person:", response.Person.Names[0].Display)
}
