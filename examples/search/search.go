// Package main demonstrates how to search for a person using the PIPL API.
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

	// Create a person object to search
	searchPerson := new(pipl.Person)
	err := searchPerson.AddEmail("clark.kent@example.com")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Submit the search request
	var response *pipl.Response
	response, err = c.Search(context.Background(), searchPerson)
	if err != nil {
		log.Fatalln(err.Error())
	} else if response == nil {
		log.Fatalln("person not found")
	}

	log.Println("found person:", response.Person.Names[0].Display)
}
