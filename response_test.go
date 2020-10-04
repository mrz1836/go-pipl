package pipl

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// loadResponseData loads a good forged response JSON (6/8/2019)
func loadResponseData(filename string) (response *Response, err error) {

	// Open our jsonFile
	var jsonFile *os.File
	jsonFile, err = os.Open("responses/" + filename)
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
	// rawJSON = string(byteValue)

	// Set the response
	err = json.Unmarshal(byteValue, &response)
	return
}

// Test_GoodResponse test a good response JSON (expected)
func Test_GoodResponse(t *testing.T) {

	// Load the response data
	response, err := loadResponseData("response_success.json")
	if err != nil {
		t.Fatal(err)
	}

	// Set the defaults (these are in the JSON file)
	var testEmail = "clark.kent@example.com"
	var personID GUID = "f4a7d898-6fc1-4a24-b043-43eb292a6fd5"

	// ==================================================================================================================

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

	// ==================================================================================================================

	// Test query parameters (hash of email)
	if response.Query.Emails[0].Address != testEmail {
		t.Fatalf("expected: %s, got: %s", testEmail, response.Query.Emails[0].Address)
	}

	// Test the md5 digest
	emailDigest := fmt.Sprintf("%x", md5.Sum([]byte(testEmail)))
	if response.Query.Emails[0].AddressMD5 != emailDigest {
		t.Fatalf("expected: %s, got: %s", emailDigest, response.Query.Emails[0].AddressMD5)
	}

	// ==================================================================================================================

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

	// ==================================================================================================================

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

	// ==================================================================================================================

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

	// ==================================================================================================================

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

	// ==================================================================================================================

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

	// ==================================================================================================================

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

	// ==================================================================================================================

	// Test person struct and data (gender)
	// Only ONE!

	if response.Person.Gender.Content != "male" {
		t.Fatalf("expected superman to be male")
	}

	// ==================================================================================================================

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

	// ==================================================================================================================

	// Test person struct and data (languages)
	if len(response.Person.Languages) != response.AvailableData.Premium.Languages {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Languages, len(response.Person.Languages))
	}

	// Test language 1
	if response.Person.Languages[0].Region != DefaultCountry {
		t.Fatalf("expected: %s, got: %s", DefaultCountry, response.Person.Languages[0].Region)
	}
	if response.Person.Languages[0].Language != DefaultLanguage {
		t.Fatalf("expected: %s, got: %s", DefaultLanguage, response.Person.Languages[0].Language)
	}
	if response.Person.Languages[0].Display != DefaultDisplayLanguage {
		t.Fatalf("expected: %s, got: %s", DefaultDisplayLanguage, response.Person.Languages[0].Display)
	}

	// ==================================================================================================================

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

	// ==================================================================================================================

	// Test person struct and data (origin_countries)
	if len(response.Person.OriginCountries) != response.AvailableData.Premium.OriginCountries {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.OriginCountries, len(response.Person.OriginCountries))
	}

	// Test countries
	if response.Person.OriginCountries[0].Country != DefaultCountry {
		t.Fatalf("expected: %s, got: %s", DefaultCountry, response.Person.OriginCountries[0].Country)
	}

	// ==================================================================================================================

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
	if response.Person.Addresses[0].Country != DefaultCountry {
		t.Fatalf("expected: %s, got: %s", DefaultCountry, response.Person.Addresses[0].Country)
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

	// ==================================================================================================================

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

	// ==================================================================================================================

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

	// ==================================================================================================================

	// Test person struct and data (relationships)
	if len(response.Person.Relationships) != response.AvailableData.Premium.Relationships {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.Relationships, len(response.Person.Relationships))
	}

	// ==================================================================================================================

	// Test person struct and data (user_ids)
	if len(response.Person.UserIDs) != response.AvailableData.Premium.UserIDs {
		t.Fatalf("expected: %d, got: %d", response.AvailableData.Premium.UserIDs, len(response.Person.UserIDs))
	}

	// Test user id #1
	if response.Person.UserIDs[0].Content != "11231@facebook" {
		t.Fatalf("expected: %s, got: %s", "11231@facebook", response.Person.UserIDs[0].Content)
	}

	// ==================================================================================================================

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

	// ==================================================================================================================

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

// Test_PersonNotFoundResponse test a person not found response JSON
func Test_PersonNotFoundResponse(t *testing.T) {
	// Load the response data
	response, err := loadResponseData("response_not_found.json")
	if err != nil {
		t.Fatal(err)
	}

	if response.PersonsCount > 0 {
		t.Fatal("expected to be 0")
	}

	if response.AvailableSources > 0 {
		t.Fatal("expected to be 0")
	}

	if response.HTTPStatusCode != 200 {
		t.Fatal("expected to be 200")
	}

	if response.Error != "" {
		t.Fatal("expected to be empty")
	}
}

// Test_BadKeyResponse test a bad api key response JSON
func Test_BadKeyResponse(t *testing.T) {
	// Load the response data
	response, err := loadResponseData("response_bad_key.json")
	if err != nil {
		t.Fatal(err)
	}

	if response.PersonsCount > 0 {
		t.Fatal("expected to be 0")
	}

	if response.AvailableSources > 0 {
		t.Fatal("expected to be 0")
	}

	if response.HTTPStatusCode != 403 {
		t.Fatal("expected to be 403")
	}

	if response.Error != "Unrecognized API key" {
		t.Fatal("expected to be: Unrecognized API key")
	}
}

// Test_PackageErrorResponse test a bad package response JSON
func Test_PackageErrorResponse(t *testing.T) {
	// Load the response data
	response, err := loadResponseData("response_package_error.json")
	if err != nil {
		t.Fatal(err)
	}

	if response.PersonsCount > 0 {
		t.Fatal("expected to be 0")
	}

	if response.AvailableSources > 0 {
		t.Fatal("expected to be 0")
	}

	if response.HTTPStatusCode != 400 {
		t.Fatal("expected to be 400")
	}

	if response.Error != "Your data package does not contain email" {
		t.Fatal("expected to be: Your data package does not contain email")
	}
}
