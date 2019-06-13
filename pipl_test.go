package pipl

import (
	"fmt"
	"os"
	"testing"
)

// Testing variables
var testImage = "https://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg"
var testThumbnailToken = "AE2861B242686E7BD0CB4D9049298EB7D18FEF66D950E8AB78BCD3F484345CE74536C19A85D0BA3D32DC9E7D1878CD4D341254E7AD129255C6983E6E154C4530A0DAAF665EA325FC0206F8B1D7E0B6B7AD9EBF71FCF610D57D"

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client, err := NewClient("1234567890")
	if err != nil {
		t.Fatal(err)
	}
	if client.SearchParameters.APIKey != "1234567890" {
		t.Fatalf("expected value 1234567890, got %s", client.SearchParameters.APIKey)
	}
	if client.SearchParameters.MinimumMatch != MinimumMatch {
		t.Fatalf("expected value %f, got %f", MinimumMatch, client.SearchParameters.MinimumMatch)
	}
	if client.SearchParameters.MinimumProbability != MinimumProbability {
		t.Fatalf("expected value %f, got %f", MinimumProbability, client.SearchParameters.MinimumProbability)
	}

	// todo: test changing these values in the SearchParameters
}

//ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient("1234567890")
	fmt.Println(client.SearchParameters.APIKey)
	// Output:1234567890
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient("1234567890")
	}
}

// TestSearchMeetsMinimumCriteria test the minimum criteria for a search
// 	This also tests: HasEmail, HasPhone, HasUserID, HasUsername, HasURL
//	HasName, HasAddress
func TestSearchMeetsMinimumCriteria(t *testing.T) {
	person := new(Person)

	// Missing data, should fail
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Raw name (good)
	_ = person.AddNameRaw("john smith")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true", person.Names[0].Raw)
	}

	// Reset
	person = new(Person)

	// Just first (missing last)
	_ = person.AddName("john", "", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Reset
	person = new(Person)

	// Just last (missing first)
	_ = person.AddName("", "", "smith", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Reset
	person = new(Person)

	// Test both first and last name
	_ = person.AddName("john", "", "smith", "", "")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test email address
	_ = person.AddEmail("clarkkent@gmail.com")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test phone with country code
	_ = person.AddPhone(9785550145, 1)
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test phone without country code
	_ = person.AddPhone(9785550145, 0)
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Reset
	person = new(Person)

	// Test phone RAW
	_ = person.AddPhoneRaw("19785550145")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test user id
	_ = person.AddUserID("clarkkent123", "twitter")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test usernames
	_ = person.AddUsername("clarkkent", "twitter")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test urls
	_ = person.AddURL("https://twitter.com/clarkkent")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Partial address
	_ = person.AddAddress("10", "", "", "", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Partial address
	_ = person.AddAddress("10", "Hickory Lane", "", "", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Partial address
	_ = person.AddAddress("10", "Hickory Lane", "", "Smallville", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Full address
	_ = person.AddAddress("10", "Hickory Lane", "", "Smallville", "KS", "", "")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}
}

//ExampleSearchMeetsMinimumCriteria example using SearchMeetsMinimumCriteria()
func ExampleSearchMeetsMinimumCriteria() {
	person := new(Person)
	if SearchMeetsMinimumCriteria(person) {
		fmt.Println("search meets minimum criteria")
	} else {
		fmt.Println("search does not meet minimum criteria")
	}
	// Output:search does not meet minimum criteria
}

// BenchmarkSearchMeetsMinimumCriteria benchmarks the SearchMeetsMinimumCriteria method
func BenchmarkSearchMeetsMinimumCriteria(b *testing.B) {
	person := new(Person)
	for i := 0; i < b.N; i++ {
		_ = SearchMeetsMinimumCriteria(person)
	}
}

// todo: test Search()

// todo: test SearchByPointer()

// todo: test SearchAllPossiblePeople()

//======================================================================================================================
// Full Pipl Integration Tests (-test.short to skip)

// TestSearchByPerson tests a live search using a real API key (if set)
// Tests a real pipl request
// Tests the debugging data returned from a request
func TestSearchByPerson(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Set the testing API key from ENV
	APIKey := os.Getenv("PIPL_CONTACT_API_KEY")

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(APIKey)
	if err != nil {
		t.Fatal(err)
	}

	// Set your match requirements if you have any. You don't pay for results that
	// don't satisfy your match requirements (but your returned results will be empty)
	client.SearchParameters.MatchRequirements = "name and phone"

	// Create a blank person to fill out with search terms
	searchObject := NewPerson()

	// Let's find out who this random guy is. We'll search by a username.
	err = searchObject.AddUsername("jeffbezos", "twitter")
	if err != nil {
		t.Fatal(err)
	}

	// And we'll add a full name
	err = searchObject.AddName("jeff", "preston", "bezos", "", "")
	if err != nil {
		t.Fatal(err)
	}

	// Some field types have a "raw" option that you can use if you're unsure
	// how to break it down. Pipl will attempt to parse it for you.
	// Generally you should use one or the other (AddX() or AddXRaw())
	err = searchObject.AddNameRaw("jeff preston bezos")
	if err != nil {
		t.Fatal(err)
	}

	// Launch the search (if you don't meet the minimum search criteria, an error
	// should be returned to you here stating such).
	results, err := client.Search(searchObject)
	if err != nil {
		t.Fatal(err)
	}

	// Test the debugging data
	if client.LastRequest.URL != "https://api.pipl.com/search/" {
		t.Fatal("url was not the same as expected", client.LastRequest.URL)
	}
	if client.LastRequest.Method != "POST" {
		t.Fatal("method was not the same as expected", client.LastRequest.Method)
	}
	if client.LastRequest.PostData != "key="+APIKey+"&match_requirements=name+and+phone&person=%7B%22names%22%3A%5B%7B%22first%22%3A%22jeff%22%2C%22last%22%3A%22bezos%22%2C%22middle%22%3A%22preston%22%7D%2C%7B%22raw%22%3A%22jeff+preston+bezos%22%7D%5D%2C%22usernames%22%3A%5B%7B%22content%22%3A%22jeffbezos%40twitter%22%7D%5D%7D&pretty=false&show_sources=all" {
		t.Fatal("post data was not the same as expected", client.LastRequest.PostData)
	}

	// Do we match?
	if results == nil {
		t.Fatal("uh oh! no results returned!")
	}

	if len(results.Person.Names) == 0 || results.Person.Names[0].First != "Jeffrey" {
		t.Fatal("uh oh! Jeff wasn't found!")
	}
}
