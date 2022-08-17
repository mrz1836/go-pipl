package pipl

import "context"

// SearchService is the search services
type SearchService interface {
	Search(ctx context.Context, searchPerson *Person) (*Response, error)
	SearchAllPossiblePeople(ctx context.Context, searchPerson *Person) (*Response, error)
	SearchByPointer(ctx context.Context, searchPointer string) (*Response, error)
}

// ClientInterface is the client interface
type ClientInterface interface {
	SearchService
	HTTPClient() HTTPInterface
	UserAgent() string
}
