// Package main demonstrates how to use custom search options with the PIPL API.
package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-pipl"
)

func main() {
	customSearchOptions := &pipl.SearchOptions{
		Search: &pipl.SearchParameters{
			MinimumProbability: 0.5,
			MinimumMatch:       0.1,
			InferPersons:       true,
			HideSponsored:      true,
			LiveFeeds:          true,
			TopMatch:           true,
			Pretty:             true,
		},
		Thumbnail: &pipl.ThumbnailSettings{
			Height:   75,
			Width:    75,
			Enabled:  true,
			Favicon:  true,
			ZoomFace: true,
		},
	}

	c := pipl.NewClient(
		pipl.WithAPIKey(os.Getenv("PIPL_API_KEY")),
		pipl.WithSearchOptions(customSearchOptions),
	)

	log.Println("client loaded:", c.UserAgent())
}
