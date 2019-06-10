package pipl

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

//======================================================================================================================
// Testing variables
var testImage = "https://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg"
var testThumbnailToken = "AE2861B242686E7BD0CB4D9049298EB7D18FEF66D950E8AB78BCD3F484345CE74536C19A85D0BA3D32DC9E7D1878CD4D341254E7AD129255C6983E6E154C4530A0DAAF665EA325FC0206F8B1D7E0B6B7AD9EBF71FCF610D57D"

//======================================================================================================================
// Response Parsing and Expected Values

// loadGoodResponseData loads a good forged response JSON (6/8/2019)
func loadGoodResponseData() (response *Response, rawJSON string, err error) {

	// Open our jsonFile
	var jsonFile *os.File
	jsonFile, err = os.Open("response_success.json")
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

	// Test email 1
	if response.Person.Emails[0].Type != "work" {
		t.Fatalf("expected: %s, got: %s", "work", response.Person.Emails[0].Type)
	}
	if response.Person.Emails[0].EmailProvider {
		t.Fatalf("provider should be false")
	}
	if response.Person.Emails[0].Address != "clark.kent@thedailyplanet.com" {
		t.Fatalf("expected: %s, got: %s", "clark.kent@thedailyplanet.com", response.Person.Emails[0].Address)
	}
	// Test the md5 digest
	emailDigest = fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[0].Address)))
	if response.Person.Emails[0].AddressMD5 != emailDigest {
		t.Fatalf("expected: %s, got: %s", emailDigest, response.Person.Emails[0].AddressMD5)
	}

	// Test email 2
	if !response.Person.Emails[1].Disposable {
		t.Fatalf("disposable should be true")
	}
	if response.Person.Emails[1].EmailProvider {
		t.Fatalf("provider should be false")
	}
	if response.Person.Emails[1].Address != "ck242@guerrillamail.com" {
		t.Fatalf("expected: %s, got: %s", "ck242@guerrillamail.com", response.Person.Emails[1].Address)
	}
	// Test the md5 digest
	emailDigest = fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[1].Address)))
	if response.Person.Emails[1].AddressMD5 != emailDigest {
		t.Fatalf("expected: %s, got: %s", emailDigest, response.Person.Emails[1].AddressMD5)
	}

	// Test email 3
	if response.Person.Emails[2].Type != "personal" {
		t.Fatalf("expected: %s, got: %s", "personal", response.Person.Emails[2].Type)
	}
	if !response.Person.Emails[2].EmailProvider {
		t.Fatalf("provider should be true")
	}
	if response.Person.Emails[2].Address != "clark@gmail.com" {
		t.Fatalf("expected: %s, got: %s", "clark@gmail.com", response.Person.Emails[2].Address)
	}
	// Test the md5 digest
	emailDigest = fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[2].Address)))
	if response.Person.Emails[2].AddressMD5 != emailDigest {
		t.Fatalf("expected: %s, got: %s", emailDigest, response.Person.Emails[2].AddressMD5)
	}

	// Test email 4
	if !response.Person.Emails[3].Disposable {
		t.Fatalf("disposable should be true")
	}
	if response.Person.Emails[3].EmailProvider {
		t.Fatalf("provider should be false")
	}
	if response.Person.Emails[3].Address != "clark.kent@example.com" {
		t.Fatalf("expected: %s, got: %s", "clark.kent@example.com", response.Person.Emails[3].Address)
	}
	// Test the md5 digest
	emailDigest = fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[3].Address)))
	if response.Person.Emails[3].AddressMD5 != emailDigest {
		t.Fatalf("expected: %s, got: %s", emailDigest, response.Person.Emails[3].AddressMD5)
	}

	//==================================================================================================================

	// Test person struct and data (usernames)
	if len(response.Person.Usernames) != response.AvailableData.Premium.Usernames {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Usernames, len(response.Person.Usernames))
	}

	// Test username 1
	if response.Person.Usernames[0].Content != "superman@facebook" {
		t.Fatalf("expected: %s, got: %s", "superman@facebook", response.Person.Usernames[0].Content)
	}

	// Test username 2
	if response.Person.Usernames[1].Content != "@ckent" {
		t.Fatalf("expected: %s, got: %s", "@ckent", response.Person.Usernames[1].Content)
	}

	//==================================================================================================================

	// Test person struct and data (phones)
	if len(response.Person.Phones) != response.AvailableData.Premium.Phones {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Phones, len(response.Person.Phones))
	}

	// Test phone 1
	if response.Person.Phones[0].Type != "home_phone" {
		t.Fatalf("expected: %s, got: %s", "home_phone", response.Person.Phones[0].Type)
	}
	if response.Person.Phones[0].CountryCode != 1 {
		t.Fatalf("expected: %d, got: %d", 1, response.Person.Phones[0].CountryCode)
	}
	if response.Person.Phones[0].Number != 9785550145 {
		t.Fatalf("expected: %d, got: %d", 9785550145, response.Person.Phones[0].Number)
	}
	if response.Person.Phones[0].Display != "978-555-0145" {
		t.Fatalf("expected: %s, got: %s", "978-555-0145", response.Person.Phones[0].Display)
	}
	if response.Person.Phones[0].DisplayInternational != "+1 978-555-0145" {
		t.Fatalf("expected: %s, got: %s", "+1 978-555-0145", response.Person.Phones[0].DisplayInternational)
	}

	//==================================================================================================================

	// Test person struct and data (gender)
	// Only ONE!

	if response.Person.Gender.Content != "male" {
		t.Fatalf("expected superman to be male")
	}

	//==================================================================================================================

	// Test person struct and data (dob)
	// Only ONE!

	if response.Person.DateOfBirth.Display != "32 years old" {
		t.Fatalf("expected: %s, got: %s", "32 years old", response.Person.DateOfBirth.Display)
	}
	if response.Person.DateOfBirth.DateRange.Start != "1986-01-01" {
		t.Fatalf("expected: %s, got: %s", "1986-01-01", response.Person.DateOfBirth.DateRange.Start)
	}
	if response.Person.DateOfBirth.DateRange.End != "1987-05-13" {
		t.Fatalf("expected: %s, got: %s", "1987-05-13", response.Person.DateOfBirth.DateRange.End)
	}

	//==================================================================================================================

	// Test person struct and data (languages)
	if len(response.Person.Languages) != response.AvailableData.Premium.Languages {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Languages, len(response.Person.Languages))
	}

	// Test language 1
	if response.Person.Languages[0].Region != "US" {
		t.Fatalf("expected: %s, got: %s", "US", response.Person.Languages[0].Region)
	}
	if response.Person.Languages[0].Language != "en" {
		t.Fatalf("expected: %s, got: %s", "en", response.Person.Languages[0].Language)
	}
	if response.Person.Languages[0].Display != "en_US" {
		t.Fatalf("expected: %s, got: %s", "en_US", response.Person.Languages[0].Display)
	}

	//==================================================================================================================

	// Test person struct and data (ethnicities)
	if len(response.Person.Ethnicities) != response.AvailableData.Premium.Ethnicities {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Ethnicities, len(response.Person.Ethnicities))
	}

	// Test ethnicities
	if response.Person.Ethnicities[0].Content != "other" {
		t.Fatalf("expected: %s, got: %s", "other", response.Person.Ethnicities[0].Content)
	}
	if response.Person.Ethnicities[1].Content != "american_indian" {
		t.Fatalf("expected: %s, got: %s", "american_indian", response.Person.Ethnicities[1].Content)
	}
	if response.Person.Ethnicities[2].Content != "white" {
		t.Fatalf("expected: %s, got: %s", "white", response.Person.Ethnicities[2].Content)
	}

	//==================================================================================================================

	// Test person struct and data (origin_countries)
	if len(response.Person.OriginCountries) != response.AvailableData.Premium.OriginCountries {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.OriginCountries, len(response.Person.OriginCountries))
	}

	// Test countries
	if response.Person.OriginCountries[0].Country != "US" {
		t.Fatalf("expected: %s, got: %s", "US", response.Person.OriginCountries[0].Country)
	}

	//==================================================================================================================

	// Test person struct and data (addresses)
	if len(response.Person.Addresses) != response.AvailableData.Premium.Addresses {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Addresses, len(response.Person.Addresses))
	}

	// Test address 1
	if response.Person.Addresses[0].ValidSince != "2005-02-12" {
		t.Fatalf("expected: %s, got: %s", "2005-02-12", response.Person.Addresses[0].ValidSince)
	}
	if response.Person.Addresses[0].Type != "work" {
		t.Fatalf("expected: %s, got: %s", "work", response.Person.Addresses[0].Type)
	}
	if response.Person.Addresses[0].Country != "US" {
		t.Fatalf("expected: %s, got: %s", "US", response.Person.Addresses[0].Country)
	}
	if response.Person.Addresses[0].State != "KS" {
		t.Fatalf("expected: %s, got: %s", "KS", response.Person.Addresses[0].State)
	}
	if response.Person.Addresses[0].City != "Metropolis" {
		t.Fatalf("expected: %s, got: %s", "Metropolis", response.Person.Addresses[0].City)
	}
	if response.Person.Addresses[0].Street != "Broadway" {
		t.Fatalf("expected: %s, got: %s", "Broadway", response.Person.Addresses[0].Street)
	}
	if response.Person.Addresses[0].House != "1000" {
		t.Fatalf("expected: %s, got: %s", "1000", response.Person.Addresses[0].House)
	}
	if response.Person.Addresses[0].Apartment != "355" {
		t.Fatalf("expected: %s, got: %s", "355", response.Person.Addresses[0].Apartment)
	}
	if response.Person.Addresses[0].Display != "1000-355 Broadway, Metropolis, Kansas" {
		t.Fatalf("expected: %s, got: %s", "1000-355 Broadway, Metropolis, Kansas", response.Person.Addresses[0].Display)
	}

	// todo: add address #2

	//==================================================================================================================

	// Test person struct and data (jobs)
	if len(response.Person.Jobs) != response.AvailableData.Premium.Jobs {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Jobs, len(response.Person.Jobs))
	}

	// Test job 1
	if response.Person.Jobs[0].Title != "Field Reporter" {
		t.Fatalf("expected: %s, got: %s", "Field Reporter", response.Person.Jobs[0].Title)
	}
	if response.Person.Jobs[0].Organization != "The Daily Planet" {
		t.Fatalf("expected: %s, got: %s", "The Daily Planet", response.Person.Jobs[0].Organization)
	}
	if response.Person.Jobs[0].Industry != "Journalism" {
		t.Fatalf("expected: %s, got: %s", "Journalism", response.Person.Jobs[0].Industry)
	}
	if response.Person.Jobs[0].Display != "Field Reporter at The Daily Planet (2000-2012)" {
		t.Fatalf("expected: %s, got: %s", "Field Reporter at The Daily Planet (2000-2012)", response.Person.Jobs[0].Display)
	}
	if response.Person.Jobs[0].DateRange.Start != "2000-12-08" {
		t.Fatalf("expected: %s, got: %s", "2000-12-08", response.Person.Jobs[0].DateRange.Start)
	}
	if response.Person.Jobs[0].DateRange.End != "2012-10-09" {
		t.Fatalf("expected: %s, got: %s", "2012-10-09", response.Person.Jobs[0].DateRange.End)
	}

	// todo: add job #2 and #3

	//==================================================================================================================

	// Test person struct and data (educations)
	if len(response.Person.Educations) != response.AvailableData.Premium.Educations {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Educations, len(response.Person.Educations))
	}

	// Test education #1
	if response.Person.Educations[0].Degree != "B.Sc Advanced Science" {
		t.Fatalf("expected: %s, got: %s", "B.Sc Advanced Science", response.Person.Educations[0].Degree)
	}
	if response.Person.Educations[0].School != "Metropolis University" {
		t.Fatalf("expected: %s, got: %s", "Metropolis University", response.Person.Educations[0].School)
	}
	if response.Person.Educations[0].Display != "B.Sc Advanced Science from Metropolis University (2005-2008)" {
		t.Fatalf("expected: %s, got: %s", "B.Sc Advanced Science from Metropolis University (2005-2008)", response.Person.Educations[0].Display)
	}
	if response.Person.Educations[0].DateRange.Start != "2005-09-01" {
		t.Fatalf("expected: %s, got: %s", "2005-09-01", response.Person.Educations[0].DateRange.Start)
	}
	if response.Person.Educations[0].DateRange.End != "2008-05-14" {
		t.Fatalf("expected: %s, got: %s", "2008-05-14", response.Person.Educations[0].DateRange.End)
	}

	// todo: add education #2

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

	// Test user id #1
	if response.Person.UserIDs[0].Content != "11231@facebook" {
		t.Fatalf("expected: %s, got: %s", "11231@facebook", response.Person.UserIDs[0].Content)
	}

	//==================================================================================================================

	// Test person struct and data (images)
	if len(response.Person.Images) != response.AvailableData.Premium.Images {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Images, len(response.Person.Images))
	}

	// Test image #1
	if response.Person.Images[0].URL != "http://vignette1.wikia.nocookie.net/smallville/images/e/ea/Buddies_forever.jpg" {
		t.Fatalf("expected: %s, got: %s", "http://vignette1.wikia.nocookie.net/smallville/images/e/ea/Buddies_forever.jpg", response.Person.Images[0].URL)
	}

	// Test image #2
	if response.Person.Images[1].URL != "http://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg" {
		t.Fatalf("expected: %s, got: %s", "http://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg", response.Person.Images[1].URL)
	}

	//==================================================================================================================

	// Test person struct and data (urls) or (social profiles)
	if len(response.Person.URLs) != response.AvailableData.Premium.SocialProfiles {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.SocialProfiles, len(response.Person.URLs))
	}

	// Test url #1
	if response.Person.URLs[0].SourceID != "edc6aa8fa3f211cfad7c12a0ba5b32f4" {
		t.Fatalf("expected: %s, got: %s", "edc6aa8fa3f211cfad7c12a0ba5b32f4", response.Person.URLs[0].SourceID)
	}
	if response.Person.URLs[0].Domain != "linkedin.com" {
		t.Fatalf("expected: %s, got: %s", "linkedin.com", response.Person.URLs[0].Domain)
	}
	if response.Person.URLs[0].Name != "LinkedIn" {
		t.Fatalf("expected: %s, got: %s", "LinkedIn", response.Person.URLs[0].Name)
	}
	if response.Person.URLs[0].Category != "professional_and_business" {
		t.Fatalf("expected: %s, got: %s", "professional_and_business", response.Person.URLs[0].Category)
	}
	if response.Person.URLs[0].URL != "http://linkedin.com/clark.kent" {
		t.Fatalf("expected: %s, got: %s", "http://linkedin.com/clark.kent", response.Person.URLs[0].URL)
	}

	// todo: add url #2 and #3
}

//======================================================================================================================
// Helper Methods

// TestNewPerson testing new person function
func TestNewPerson(t *testing.T) {
	person := NewPerson()
	if reflect.TypeOf(person).String() != "*pipl.Person" {
		t.Fatal("expected type to be *pipl.Person")
	}
}

// BenchmarkNewPerson benchmarks the NewPerson method
func BenchmarkNewPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewPerson()
	}
}

// TestAddName test adding a name to a person object
func TestAddName(t *testing.T) {
	person := NewPerson()
	person.AddName("clark", "ryan", "kent", "mr", "jr")
	if len(person.Names) == 0 {
		t.Fatal("expected a name in this person object")
	}
	if person.Names[0].First != "clark" {
		t.Fatalf("expected value to be clark, got %s", person.Names[0].First)
	}
	if person.Names[0].Middle != "ryan" {
		t.Fatalf("expected value to be ryan, got %s", person.Names[0].Middle)
	}
	if person.Names[0].Last != "kent" {
		t.Fatalf("expected value to be kent, got %s", person.Names[0].Last)
	}
	if person.Names[0].Prefix != "mr" {
		t.Fatalf("expected value to be mr, got %s", person.Names[0].Prefix)
	}
	if person.Names[0].Suffix != "jr" {
		t.Fatalf("expected value to be jr, got %s", person.Names[0].Suffix)
	}
}

//ExamplePerson_AddName example using AddName()
func ExamplePerson_AddName() {
	person := NewPerson()
	person.AddName("clark", "ryan", "kent", "mr", "jr")
	fmt.Println(person.Names[0].First + " " + person.Names[0].Last)
	// Output: clark kent
}

// BenchmarkAddName benchmarks the AddName method
func BenchmarkAddName(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddName("clark", "ryan", "kent", "mr", "jr")
	}
}

// TestAddNameRaw test adding a raw name to a person object
func TestAddNameRaw(t *testing.T) {
	person := NewPerson()
	person.AddNameRaw("clark ryan kent")
	if len(person.Names) == 0 {
		t.Fatal("expected a name in this person object")
	}
	if person.Names[0].Raw != "clark ryan kent" {
		t.Fatalf("expected value to be clark ryan kent, got %s", person.Names[0].Raw)
	}
}

//ExamplePerson_AddNameRaw example using AddNameRaw()
func ExamplePerson_AddNameRaw() {
	person := NewPerson()
	person.AddNameRaw("clark kent")
	fmt.Println(person.Names[0].Raw)
	// Output: clark kent
}

// BenchmarkAddNameRaw benchmarks the AddNameRaw method
func BenchmarkAddNameRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddNameRaw("clark ryan kent")
	}
}

// TestAddEmail test adding an email to a person object
func TestAddEmail(t *testing.T) {
	person := NewPerson()
	person.AddEmail("clarkkent@gmail.com")
	if len(person.Emails) == 0 {
		t.Fatal("expected an email in this person object")
	}
	if person.Emails[0].Address != "clarkkent@gmail.com" {
		t.Fatalf("expected value to be clarkkent@gmail.com, got %s", person.Emails[0].Address)
	}
}

//ExamplePerson_AddEmail example using AddEmail()
func ExamplePerson_AddEmail() {
	person := NewPerson()
	person.AddEmail("clarkkent@gmail.com")
	fmt.Println(person.Emails[0].Address)
	// Output:clarkkent@gmail.com
}

// BenchmarkAddEmail benchmarks the AddEmail method
func BenchmarkAddEmail(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddEmail("clarkkent@gmail.com")
	}
}

// TestAddUsername test adding an username to a person object
func TestAddUsername(t *testing.T) {
	person := NewPerson()
	person.AddUsername("clarkkent")
	if len(person.Usernames) == 0 {
		t.Fatal("expected a username in this person object")
	}
	if person.Usernames[0].Content != "clarkkent" {
		t.Fatalf("expected value to be clarkkent, got %s", person.Usernames[0].Content)
	}
}

//ExamplePerson_AddUsername example using AddUsername()
func ExamplePerson_AddUsername() {
	person := NewPerson()
	person.AddUsername("clarkkent")
	fmt.Println(person.Usernames[0].Content)
	// Output:clarkkent
}

// BenchmarkAddUsername benchmarks the AddUsername method
func BenchmarkAddUsername(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddUsername("clarkkent")
	}
}

// TestAddPhone test adding a phone to a person object
func TestAddPhone(t *testing.T) {
	person := NewPerson()
	person.AddPhone(9785550145, 1)
	if len(person.Phones) == 0 {
		t.Fatal("expected a phone in this person object")
	}
	if person.Phones[0].Number != 9785550145 {
		t.Fatalf("expected value to be 9785550145, got %d", person.Phones[0].Number)
	}
	if person.Phones[0].CountryCode != 1 {
		t.Fatalf("expected value to be 1, got %d", person.Phones[0].CountryCode)
	}
}

//ExamplePerson_AddPhone example using AddPhone()
func ExamplePerson_AddPhone() {
	person := NewPerson()
	person.AddPhone(9785550145, 1)
	fmt.Println(person.Phones[0].Number)
	// Output:9785550145
}

// BenchmarkAddPhone benchmarks the AddPhone method
func BenchmarkAddPhone(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddPhone(9785550145, 1)
	}
}

// TestAddPhoneRaw test adding a phone to a person object
func TestAddPhoneRaw(t *testing.T) {
	person := NewPerson()
	person.AddPhoneRaw("9785550145")
	if len(person.Phones) == 0 {
		t.Fatal("expected a phone in this person object")
	}
	if person.Phones[0].Raw != "9785550145" {
		t.Fatalf("expected value to be 9785550145, got %s", person.Phones[0].Raw)
	}
}

//ExamplePerson_AddPhoneRaw example using AddPhoneRaw()
func ExamplePerson_AddPhoneRaw() {
	person := NewPerson()
	person.AddPhoneRaw("9785550145")
	fmt.Println(person.Phones[0].Raw)
	// Output:9785550145
}

// BenchmarkAddPhoneRaw benchmarks the AddPhoneRaw method
func BenchmarkAddPhoneRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddPhoneRaw("9785550145")
	}
}

// TestSetGender test setting a gender on a person object
func TestSetGender(t *testing.T) {
	person := NewPerson()
	person.SetGender("male")
	if person.Gender.Content != "male" {
		t.Fatalf("expected value to be male, got %s", person.Gender.Content)
	}
	person.SetGender("female")
	if person.Gender.Content != "female" {
		t.Fatalf("expected value to be female, got %s", person.Gender.Content)
	}

	// This is a case where gender is NOT male/female
	person.SetGender("other")
	if person.Gender.Content != "male" {
		t.Fatalf("expected value to be male, got %s", person.Gender.Content)
	}
}

//ExamplePerson_SetGender example using SetGender()
func ExamplePerson_SetGender() {
	person := NewPerson()
	person.SetGender("male")
	fmt.Println(person.Gender.Content)
	// Output:male
}

// BenchmarkSetGender benchmarks the SetGender method
func BenchmarkSetGender(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.SetGender("male")
	}
}

// TestSetDateOfBirth test setting a DOB on a person object
func TestSetDateOfBirth(t *testing.T) {
	person := NewPerson()
	person.SetDateOfBirth("1987-01-01")
	if person.DateOfBirth.DateRange.Start != "1987-01-01" {
		t.Fatalf("expected value to be 1987-01-01, got %s", person.DateOfBirth.DateRange.Start)
	}
	if person.DateOfBirth.DateRange.End != "1987-01-01" {
		t.Fatalf("expected value to be 1987-01-01, got %s", person.DateOfBirth.DateRange.End)
	}
}

//ExamplePerson_SetDateOfBirth example using SetDateOfBirth()
func ExamplePerson_SetDateOfBirth() {
	person := NewPerson()
	person.SetDateOfBirth("1987-01-01")
	fmt.Println(person.DateOfBirth.DateRange.Start)
	// Output:1987-01-01
}

// BenchmarkSetDateOfBirth benchmarks the SetDateOfBirth method
func BenchmarkSetDateOfBirth(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.SetDateOfBirth("1987-01-01")
	}
}

// TestAddLanguage test adding a language to a person object
func TestAddLanguage(t *testing.T) {
	person := NewPerson()
	person.AddLanguage("en", "US")
	if len(person.Languages) == 0 {
		t.Fatal("expected a language in this person object")
	}
	if person.Languages[0].Language != "en" {
		t.Fatalf("expected value to be en, got %s", person.Languages[0].Language)
	}
	if person.Languages[0].Region != "US" {
		t.Fatalf("expected value to be US, got %s", person.Languages[0].Region)
	}
}

//ExamplePerson_AddLanguage example using AddLanguage()
func ExamplePerson_AddLanguage() {
	person := NewPerson()
	person.AddLanguage("en", "US")
	fmt.Println(person.Languages[0].Language)
	// Output:en
}

// BenchmarkAddLanguage benchmarks the AddLanguage method
func BenchmarkAddLanguage(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddLanguage("en", "US")
	}
}

// TestAddEthnicity test adding a ethnicity to a person object
func TestAddEthnicity(t *testing.T) {
	person := NewPerson()
	person.AddEthnicity("white")
	if len(person.Ethnicities) == 0 {
		t.Fatal("expected an ethnicity in this person object")
	}
	if person.Ethnicities[0].Content != "white" {
		t.Fatalf("expected value to be white, got %s", person.Ethnicities[0].Content)
	}
}

//ExamplePerson_AddEthnicity example using AddEthnicity()
func ExamplePerson_AddEthnicity() {
	person := NewPerson()
	person.AddEthnicity("white")
	fmt.Println(person.Ethnicities[0].Content)
	// Output:white
}

// BenchmarkAddEthnicity benchmarks the AddEthnicity method
func BenchmarkAddEthnicity(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddEthnicity("white")
	}
}

// TestAddOriginCountry test adding a an origin country to a person object
func TestAddOriginCountry(t *testing.T) {
	person := NewPerson()
	person.AddOriginCountry("US")
	if len(person.OriginCountries) == 0 {
		t.Fatal("expected an origin country in this person object")
	}
	if person.OriginCountries[0].Country != "US" {
		t.Fatalf("expected value to be US, got %s", person.OriginCountries[0].Country)
	}
}

//ExamplePerson_AddOriginCountry example using AddOriginCountry()
func ExamplePerson_AddOriginCountry() {
	person := NewPerson()
	person.AddOriginCountry("US")
	fmt.Println(person.OriginCountries[0].Country)
	// Output:US
}

// BenchmarkAddOriginCountry benchmarks the AddOriginCountry method
func BenchmarkAddOriginCountry(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddOriginCountry("US")
	}
}

// TestAddAddress test adding an address to a person object
func TestAddAddress(t *testing.T) {
	person := NewPerson()
	person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", "US", "123")
	if len(person.Addresses) == 0 {
		t.Fatal("expected an address in this person object")
	}
	if person.Addresses[0].House != "10" {
		t.Fatalf("expected value to be 10, got %s", person.Addresses[0].House)
	}
	if person.Addresses[0].Street != "Hickory Lane" {
		t.Fatalf("expected value to be Hickory Lane, got %s", person.Addresses[0].Street)
	}
	if person.Addresses[0].Apartment != "1" {
		t.Fatalf("expected value to be 1, got %s", person.Addresses[0].Apartment)
	}
	if person.Addresses[0].City != "Smallville" {
		t.Fatalf("expected value to be Smallville, got %s", person.Addresses[0].City)
	}
	if person.Addresses[0].State != "KS" {
		t.Fatalf("expected value to be KS, got %s", person.Addresses[0].State)
	}
	if person.Addresses[0].Country != "US" {
		t.Fatalf("expected value to be US, got %s", person.Addresses[0].Country)
	}
	if person.Addresses[0].POBox != "123" {
		t.Fatalf("expected value to be 123, got %s", person.Addresses[0].POBox)
	}
}

//ExamplePerson_AddAddress example using AddAddress()
func ExamplePerson_AddAddress() {
	person := NewPerson()
	person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", "US", "123")
	fmt.Println(person.Addresses[0].House + " " + person.Addresses[0].Street + ", " + person.Addresses[0].City + " " + person.Addresses[0].State)
	// Output:10 Hickory Lane, Smallville KS
}

// BenchmarkAddAddress benchmarks the AddAddress method
func BenchmarkAddAddress(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", "US", "123")
	}
}

// TestAddAddressRaw test adding a an address to a person object
func TestAddAddressRaw(t *testing.T) {
	person := NewPerson()
	person.AddAddressRaw("10 Hickory Lane, Kansas, USA")
	if len(person.Addresses) == 0 {
		t.Fatal("expected an address in this person object")
	}
	if person.Addresses[0].Raw != "10 Hickory Lane, Kansas, USA" {
		t.Fatalf("expected value to be 10 Hickory Lane, Kansas, USA, got %s", person.Addresses[0].Raw)
	}
}

//ExamplePerson_AddAddressRaw example using AddAddressRaw()
func ExamplePerson_AddAddressRaw() {
	person := NewPerson()
	person.AddAddressRaw("10 Hickory Lane, Kansas, USA")
	fmt.Println(person.Addresses[0].Raw)
	// Output:10 Hickory Lane, Kansas, USA
}

// BenchmarkAddAddressRaw benchmarks the AddAddressRaw method
func BenchmarkAddAddressRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddAddressRaw("10 Hickory Lane, Kansas, USA")
	}
}

// TestAddJob test adding a job to a person object
func TestAddJob(t *testing.T) {
	person := NewPerson()
	person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
	if len(person.Jobs) == 0 {
		t.Fatal("expected a job in this person object")
	}
	if person.Jobs[0].Title != "reporter" {
		t.Fatalf("expected value to be reporter, got %s", person.Jobs[0].Title)
	}
	if person.Jobs[0].Organization != "daily post" {
		t.Fatalf("expected value to be daily post, got %s", person.Jobs[0].Organization)
	}
	if person.Jobs[0].Industry != "news" {
		t.Fatalf("expected value to be news, got %s", person.Jobs[0].Industry)
	}
	if person.Jobs[0].DateRange.Start != "2010-01-01" {
		t.Fatalf("expected value to be 2010-01-01, got %s", person.Jobs[0].DateRange.Start)
	}
	if person.Jobs[0].DateRange.End != "2011-01-01" {
		t.Fatalf("expected value to be 2011-01-01, got %s", person.Jobs[0].DateRange.End)
	}
}

//ExamplePerson_AddJob example using AddJob()
func ExamplePerson_AddJob() {
	person := NewPerson()
	person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
	fmt.Println(person.Jobs[0].Title + " at " + person.Jobs[0].Organization + " in " + person.Jobs[0].Industry)
	// Output:reporter at daily post in news
}

// BenchmarkAddJob benchmarks the AddJob method
func BenchmarkAddJob(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
	}
}

// TestAddEducation test adding a an education to a person object
func TestAddEducation(t *testing.T) {
	person := NewPerson()
	person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
	if len(person.Educations) == 0 {
		t.Fatal("expected an education in this person object")
	}
	if person.Educations[0].Degree != "masters" {
		t.Fatalf("expected value to be masters, got %s", person.Educations[0].Degree)
	}
	if person.Educations[0].School != "fau" {
		t.Fatalf("expected value to be fau, got %s", person.Educations[0].School)
	}
	if person.Educations[0].DateRange.Start != "2010-01-01" {
		t.Fatalf("expected value to be 2010-01-01, got %s", person.Educations[0].DateRange.Start)
	}
	if person.Educations[0].DateRange.End != "2011-01-01" {
		t.Fatalf("expected value to be 2011-01-01, got %s", person.Educations[0].DateRange.End)
	}
}

//ExamplePerson_AddEducation example using AddEducation()
func ExamplePerson_AddEducation() {
	person := NewPerson()
	person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
	fmt.Println(person.Educations[0].Degree + " from " + person.Educations[0].School)
	// Output:masters from fau
}

// BenchmarkAddEducation benchmarks the AddEducation method
func BenchmarkAddEducation(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
	}
}

// TestAddUserID test adding a user id to a person object
func TestAddUserID(t *testing.T) {
	person := NewPerson()
	person.AddUserID("clarkkent")
	if len(person.UserIDs) == 0 {
		t.Fatal("expected a user id in this person object")
	}
	if person.UserIDs[0].Content != "clarkkent" {
		t.Fatalf("expected value to be clarkkent, got %s", person.UserIDs[0].Content)
	}
}

//ExamplePerson_AddUserID example using AddUserID()
func ExamplePerson_AddUserID() {
	person := NewPerson()
	person.AddUserID("clarkkent")
	fmt.Println(person.UserIDs[0].Content)
	// Output:clarkkent
}

// BenchmarkAddUserID benchmarks the AddUserID method
func BenchmarkAddUserID(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddUserID("clarkkent")
	}
}

// TestAddURL test adding a url to a person object
func TestAddURL(t *testing.T) {
	person := NewPerson()
	person.AddURL("https://twitter.com/clarkkent")
	if len(person.URLs) == 0 {
		t.Fatal("expected a url in this person object")
	}
	if person.URLs[0].URL != "https://twitter.com/clarkkent" {
		t.Fatalf("expected value to be https://twitter.com/clarkkent, got %s", person.URLs[0].URL)
	}
}

//ExamplePerson_AddURL example using AddURL()
func ExamplePerson_AddURL() {
	person := NewPerson()
	person.AddURL("https://twitter.com/clarkkent")
	fmt.Println(person.URLs[0].URL)
	// Output:https://twitter.com/clarkkent
}

// BenchmarkAddURL benchmarks the AddURL method
func BenchmarkAddURL(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		person.AddURL("https://twitter.com/clarkkent")
	}
}

// TestPerson_ProcessThumbnails test processing images for thumbnails
func TestPerson_ProcessThumbnails(t *testing.T) {

	// Create the client
	client, err := NewClient("1234567890")
	if err != nil {
		t.Fatal(err)
	}

	// Create person and image
	person := NewPerson()
	image := new(Image)
	image.URL = testImage
	image.ThumbnailToken = testThumbnailToken
	person.Images = append(person.Images, *image)

	// Process using defaults
	person.ProcessThumbnails(client)

	// Test for url
	if len(person.Images[0].ThumbnailURL) == 0 {
		t.Fatal("this url should not be empty")
	}

	// Does it have the right width
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("width=%d", ThumbnailWidth)) {
		t.Fatal("expected value is not the same", ThumbnailWidth)
	}

	// Does it have the right height
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("height=%d", ThumbnailHeight)) {
		t.Fatal("expected value is not the same", ThumbnailHeight)
	}

	// Does it have the right favicon
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("favicon=%t", client.ThumbnailSettings.Favicon)) {
		t.Fatal("expected value is not the same", client.ThumbnailSettings.Favicon)
	}

	// Does it have the right zoom face
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("zoom_face=%t", client.ThumbnailSettings.ZoomFace)) {
		t.Fatal("expected value is not the same", client.ThumbnailSettings.ZoomFace)
	}

	// Does it have the right token
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("tokens=%s", person.Images[0].ThumbnailToken)) {
		t.Fatal("expected value is not the same", person.Images[0].ThumbnailToken)
	}
}

//ExamplePerson_ProcessThumbnails example using ProcessThumbnails()
func ExamplePerson_ProcessThumbnails() {
	client, _ := NewClient("1234567890")
	person := NewPerson()

	image := new(Image)
	image.URL = testImage
	image.ThumbnailToken = testThumbnailToken
	person.Images = append(person.Images, *image)

	person.ProcessThumbnails(client)
	fmt.Println(person.Images[0].ThumbnailURL)
	// Output: https://thumb.pipl.com/image?height=250&width=250&favicon=false&zoom_face=false&tokens=AE2861B242686E7BD0CB4D9049298EB7D18FEF66D950E8AB78BCD3F484345CE74536C19A85D0BA3D32DC9E7D1878CD4D341254E7AD129255C6983E6E154C4530A0DAAF665EA325FC0206F8B1D7E0B6B7AD9EBF71FCF610D57D
}

// BenchmarkProcessThumbnails benchmarks the ProcessThumbnails method
func BenchmarkProcessThumbnails(b *testing.B) {
	client, _ := NewClient("1234567890")
	person := NewPerson()

	image := new(Image)
	image.URL = testImage
	image.ThumbnailToken = testThumbnailToken
	person.Images = append(person.Images, *image)

	for i := 0; i < b.N; i++ {
		person.ProcessThumbnails(client)
	}
}

// todo: test AddRelationship()

//======================================================================================================================
// Pipl Core Methods

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
	person.AddNameRaw("john smith")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true", person.Names[0].Raw)
	}

	// Reset
	person = new(Person)

	// Just first (missing last)
	person.AddName("john", "", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Reset
	person = new(Person)

	// Just last (missing first)
	person.AddName("", "", "smith", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Reset
	person = new(Person)

	// Test both first and last name
	person.AddName("john", "", "smith", "", "")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test email address
	person.AddEmail("clarkkent@gmail.com")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test phone with country code
	person.AddPhone(9785550145, 1)
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test phone without country code
	person.AddPhone(9785550145, 0)
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Reset
	person = new(Person)

	// Test phone RAW
	person.AddPhoneRaw("9785550145")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test user id
	person.AddUserID("clarkkent123")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test usernames
	person.AddUsername("clarkkent")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Test urls
	person.AddURL("https://twitter.com/clarkkent")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	// Reset
	person = new(Person)

	// Partial address
	person.AddAddress("10", "", "", "", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Partial address
	person.AddAddress("10", "Hickory Lane", "", "", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Partial address
	person.AddAddress("10", "Hickory Lane", "", "Smallville", "", "", "")
	if SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return false")
	}

	// Full address
	person.AddAddress("10", "Hickory Lane", "", "Smallville", "KS", "", "")
	if !SearchMeetsMinimumCriteria(person) {
		t.Fatal("method should return true")
	}

	//person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", "US", "123")
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
func TestSearchByPerson(t *testing.T) {
	// Skip tis test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Create a new client object to handle your queries (supply an API Key)
	client, _ := NewClient("your-api-key")

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
	results, err := client.Search(searchObject)
	if err != nil {
		t.Fatal(err)
	}

	// Do we match?
	if results.Person.Names[0].First != "Jeff" {
		t.Fatal("uh oh! Jeff wasn't found!")
	}
}
