package pipl

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// loadGoodResponseData loads a good forged response JSON (6/8/2019)
func loadGoodResponseData() (response *Response, rawJSON string, err error) {

	// Open our jsonFile
	var jsonFile *os.File
	jsonFile, err = os.Open("test-good-response.json")
	if err != nil {
		return
	}

	// Defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		_ = jsonFile.Close()
	}()

	// Read our opened xmlFile as a byte array.
	var byteValue []byte
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	// Set the JSON (for debugging)
	rawJSON = string(byteValue)

	// Set the response
	err = json.Unmarshal(byteValue, &response)
	return
}

// Test_GoodResponse test a good response JSON (expected)
func Test_GoodResponse(t *testing.T) {

	// Load the response data
	response, _, err := loadGoodResponseData()
	if err != nil {
		t.Fatal(err)
	}

	// Set the defaults (these are in the JSON file)
	var testEmail = "clark.kent@example.com"
	var personID GUID = "f4a7d898-6fc1-4a24-b043-43eb292a6fd5"

	//==================================================================================================================

	// Test status code
	if response.HTTPStatusCode != 200 {
		t.Fatalf("expected: 200, got: %d", response.HTTPStatusCode)
	}

	// Test visible sources
	if response.VisibleSources != 3 {
		t.Fatalf("expected: 3, got: %d", response.VisibleSources)
	}

	// Test available sources
	if response.AvailableSources != 4 {
		t.Fatalf("expected: 4, got: %d", response.AvailableSources)
	}

	// Test person count
	if response.PersonsCount != 1 {
		t.Fatalf("expected: 1, got: %d", response.PersonsCount)
	}

	// Test search id
	if response.SearchID != "0" {
		t.Fatalf("expected: 0, got: %s", response.SearchID)
	}

	//==================================================================================================================

	// Test query parameters (hash of email)
	if response.Query.Emails[0].Address != testEmail {
		t.Fatalf("expected: %s, got: %s", testEmail, response.Query.Emails[0].Address)
	}

	// Test the md5 digest
	emailDigest := fmt.Sprintf("%x", md5.Sum([]byte(testEmail)))
	if response.Query.Emails[0].AddressMD5 != emailDigest {
		t.Fatalf("expected: %s, got: %s", emailDigest, response.Query.Emails[0].AddressMD5)
	}

	//==================================================================================================================

	// Test available data
	if response.AvailableData.Premium.Relationships != 6 {
		t.Fatalf("expected: 6, got: %d", response.AvailableData.Premium.Relationships)
	}
	if response.AvailableData.Premium.Usernames != 2 {
		t.Fatalf("expected: 2, got: %d", response.AvailableData.Premium.Usernames)
	}
	if response.AvailableData.Premium.Jobs != 3 {
		t.Fatalf("expected: 3, got: %d", response.AvailableData.Premium.Jobs)
	}
	if response.AvailableData.Premium.Addresses != 2 {
		t.Fatalf("expected: 2, got: %d", response.AvailableData.Premium.Addresses)
	}
	if response.AvailableData.Premium.Ethnicities != 3 {
		t.Fatalf("expected: 3, got: %d", response.AvailableData.Premium.Ethnicities)
	}
	if response.AvailableData.Premium.Phones != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.Phones)
	}
	if response.AvailableData.Premium.LandlinePhones != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.LandlinePhones)
	}
	if response.AvailableData.Premium.Educations != 2 {
		t.Fatalf("expected: 2, got: %d", response.AvailableData.Premium.Educations)
	}
	if response.AvailableData.Premium.Languages != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.Languages)
	}
	if response.AvailableData.Premium.UserIDs != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.UserIDs)
	}
	if response.AvailableData.Premium.SocialProfiles != 3 {
		t.Fatalf("expected: 3, got: %d", response.AvailableData.Premium.SocialProfiles)
	}
	if response.AvailableData.Premium.Names != 3 {
		t.Fatalf("expected: 3, got: %d", response.AvailableData.Premium.Names)
	}
	if response.AvailableData.Premium.DOBs != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.DOBs)
	}
	if response.AvailableData.Premium.Images != 2 {
		t.Fatalf("expected: 2, got: %d", response.AvailableData.Premium.Images)
	}
	if response.AvailableData.Premium.Genders != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.Genders)
	}
	if response.AvailableData.Premium.Emails != 4 {
		t.Fatalf("expected: 4, got: %d", response.AvailableData.Premium.Emails)
	}
	if response.AvailableData.Premium.OriginCountries != 1 {
		t.Fatalf("expected: 1, got: %d", response.AvailableData.Premium.OriginCountries)
	}

	//==================================================================================================================

	// Test person struct and data (base)
	if response.Person.ID != personID {
		t.Fatalf("expected: %s, got: %s", personID, response.Person.ID)
	}
	if response.Person.Match != float32(response.PersonsCount) {
		t.Fatalf("expected: %d, got: %f", response.PersonsCount, response.Person.Match)
	}
	if response.Person.SearchPointer != "1906090343183157724859920008073008866" {
		t.Fatalf("expected: %s, got: %s", "1906090343183157724859920008073008866", response.Person.SearchPointer)
	}

	//==================================================================================================================

	// Test person struct and data (names)
	if len(response.Person.Names) != response.AvailableData.Premium.Names {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Names, len(response.Person.Names))
	}

	// Name number 1
	if response.Person.Names[0].First != "Kal" {
		t.Fatalf("expected: %s, got: %s", "Kal", response.Person.Names[0].First)
	}
	if response.Person.Names[0].Last != "El" {
		t.Fatalf("expected: %s, got: %s", "El", response.Person.Names[0].Last)
	}
	if response.Person.Names[0].Display != response.Person.Names[0].First+" "+response.Person.Names[0].Last {
		t.Fatalf("expected: %s, got: %s", response.Person.Names[0].First+" "+response.Person.Names[0].Last, response.Person.Names[0].Display)
	}

	// Name number 2
	if response.Person.Names[1].First != "Clark" {
		t.Fatalf("expected: %s, got: %s", "Clark", response.Person.Names[1].First)
	}
	if response.Person.Names[1].Middle != "Joseph" {
		t.Fatalf("expected: %s, got: %s", "Joseph", response.Person.Names[1].Middle)
	}
	if response.Person.Names[1].Last != "Kent" {
		t.Fatalf("expected: %s, got: %s", "Kent", response.Person.Names[1].Last)
	}
	if response.Person.Names[1].Display != response.Person.Names[1].First+" "+response.Person.Names[1].Middle+" "+response.Person.Names[1].Last {
		t.Fatalf("expected: %s, got: %s", response.Person.Names[1].First+" "+response.Person.Names[1].Middle+" "+response.Person.Names[1].Last, response.Person.Names[1].Display)
	}

	// Name number 3
	if response.Person.Names[2].Display != "The red blue blur" {
		t.Fatalf("expected: %s, got: %s", "The red blue blur", response.Person.Names[2].Display)
	}

	//==================================================================================================================

	// Test person struct and data (emails)
	if len(response.Person.Emails) != response.AvailableData.Premium.Emails {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Emails, len(response.Person.Emails))
	}

	//==================================================================================================================

	// Test person struct and data (usernames)
	if len(response.Person.Usernames) != response.AvailableData.Premium.Usernames {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Usernames, len(response.Person.Usernames))
	}

	//==================================================================================================================

	// Test person struct and data (phones)
	if len(response.Person.Phones) != response.AvailableData.Premium.Phones {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Phones, len(response.Person.Phones))
	}

	//==================================================================================================================

	// Test person struct and data (gender)
	// Only ONE!

	//==================================================================================================================

	// Test person struct and data (dob)
	// Only ONE!

	//==================================================================================================================

	// Test person struct and data (languages)
	if len(response.Person.Languages) != response.AvailableData.Premium.Languages {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Languages, len(response.Person.Languages))
	}

	//==================================================================================================================

	// Test person struct and data (ethnicities)
	if len(response.Person.Ethnicities) != response.AvailableData.Premium.Ethnicities {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Ethnicities, len(response.Person.Ethnicities))
	}

	//==================================================================================================================

	// Test person struct and data (origin_countries)
	if len(response.Person.OriginCountries) != response.AvailableData.Premium.OriginCountries {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.OriginCountries, len(response.Person.OriginCountries))
	}

	//==================================================================================================================

	// Test person struct and data (addresses)
	if len(response.Person.Addresses) != response.AvailableData.Premium.Addresses {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Addresses, len(response.Person.Addresses))
	}

	//==================================================================================================================

	// Test person struct and data (jobs)
	if len(response.Person.Jobs) != response.AvailableData.Premium.Jobs {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Jobs, len(response.Person.Jobs))
	}

	//==================================================================================================================

	// Test person struct and data (educations)
	if len(response.Person.Educations) != response.AvailableData.Premium.Educations {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Educations, len(response.Person.Educations))
	}

	//==================================================================================================================

	// Test person struct and data (relationships)
	if len(response.Person.Relationships) != response.AvailableData.Premium.Relationships {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Relationships, len(response.Person.Relationships))
	}

	//==================================================================================================================

	// Test person struct and data (user_ids)
	if len(response.Person.UserIDs) != response.AvailableData.Premium.UserIDs {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.UserIDs, len(response.Person.UserIDs))
	}

	//==================================================================================================================

	// Test person struct and data (images)
	if len(response.Person.Images) != response.AvailableData.Premium.Images {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Images, len(response.Person.Images))
	}

	//==================================================================================================================

	// Test person struct and data (urls) or (social profiles)
	if len(response.Person.URLs) != response.AvailableData.Premium.SocialProfiles {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.SocialProfiles, len(response.Person.URLs))
	}
}

// TestLiveSearch tests a live search using a real API key
func TestLiveSearch(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client := NewClient("your-api-key")

	// Set your match requirements if you have any. You don't pay for results that
	// don't satisfy your match requirements (but your returned results will be empty)
	client.SearchParameters.MatchRequirements = "name and phone"

	// Create a blank person to fill out with search terms
	searchObject := NewPerson()

	// Let's find out who this random guy is. We'll search by a username.
	searchObject.AddUsername("@jeffbezos")

	// And we'll add a full name
	searchObject.AddName("jeff", "preston", "bezos", "", "")

	// Some field types have a "raw" option that you can use if you're unsure
	// how to break it down. Pipl will attempt to parse it for you.
	// Generally you should use one or the other (AddX() or AddXRaw())
	searchObject.AddNameRaw("jeff preston bezos")

	// Launch the search (if you don't meet the minimum search criteria, an error
	// should be returned to you here stating such).
	results, err := client.SearchByPerson(searchObject)
	// Handle errors better than I do pls
	if err != nil {
		t.Fatal(err)
	}

	// When multiple PossiblePersons are returned, we get a "preview" of each of
	// each of them (< 100% match confidence)
	if results.PersonsCount > 1 {
		for _, person := range results.PossiblePersons {
			// In order to get the full info on each, we need to a follow up query
			// to pull a full person profile by search pointer
			searchPtr := person.SearchPointer
			ptrResults, err := client.SearchByPointer(searchPtr)
			if err != nil {
				t.Fatal(err)
			}
			ptrSummary, err := ptrResults.Summarize()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(ptrSummary)
		}
	} else if results.PersonsCount == 1 {
		// When a single result is returned from our search, we get a full
		// profile by default (100% match confidence)
		personSummary, err := results.Person.Summarize()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(personSummary)
	} else {
		t.Log("no results found")
	}
	t.Log("test complete")
}
