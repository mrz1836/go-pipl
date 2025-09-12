package pipl

import (
	"crypto/md5" //nolint:gosec // MD5 is used by PIPL API for email hashing
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// loadResponseData loads a good forged response JSON (6/8/2019)
func loadResponseData(filename string) (response *Response, err error) {
	// Open our jsonFile
	var jsonFile *os.File
	if jsonFile, err = os.Open("responses/" + filename); err != nil { //nolint:gosec // Safe test file inclusion
		return
	}

	// Defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		_ = jsonFile.Close()
	}()

	// Read our opened xmlFile as a byte array.
	var byteValue []byte
	if byteValue, err = io.ReadAll(jsonFile); err != nil {
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
	require.NoError(t, err)

	// Set the defaults (these are in the JSON file)
	var personID GUID = "f4a7d898-6fc1-4a24-b043-43eb292a6fd5"

	// ==================================================================================================================

	require.Equal(t, "0", response.SearchID)
	require.Equal(t, 1, response.PersonsCount)
	require.Equal(t, 3, response.VisibleSources)
	require.Equal(t, 4, response.AvailableSources)
	require.Equal(t, http.StatusOK, response.HTTPStatusCode)

	// ==================================================================================================================

	// Test query parameters (hash of email)
	require.Equal(t, testEmail, response.Query.Emails[0].Address)
	require.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(testEmail))), response.Query.Emails[0].AddressMD5) //nolint:gosec // Testing PIPL API MD5 behavior

	// ==================================================================================================================

	// Test available data
	require.Equal(t, 1, response.AvailableData.Premium.DOBs)
	require.Equal(t, 1, response.AvailableData.Premium.Genders)
	require.Equal(t, 1, response.AvailableData.Premium.LandlinePhones)
	require.Equal(t, 1, response.AvailableData.Premium.Languages)
	require.Equal(t, 1, response.AvailableData.Premium.OriginCountries)
	require.Equal(t, 1, response.AvailableData.Premium.Phones)
	require.Equal(t, 1, response.AvailableData.Premium.UserIDs)
	require.Equal(t, 2, response.AvailableData.Premium.Addresses)
	require.Equal(t, 2, response.AvailableData.Premium.Educations)
	require.Equal(t, 2, response.AvailableData.Premium.Images)
	require.Equal(t, 2, response.AvailableData.Premium.Usernames)
	require.Equal(t, 3, response.AvailableData.Premium.Ethnicities)
	require.Equal(t, 3, response.AvailableData.Premium.Jobs)
	require.Equal(t, 3, response.AvailableData.Premium.Names)
	require.Equal(t, 3, response.AvailableData.Premium.SocialProfiles)
	require.Equal(t, 4, response.AvailableData.Premium.Emails)
	require.Equal(t, 6, response.AvailableData.Premium.Relationships)

	// ==================================================================================================================

	// Test person struct and data (base)
	require.Equal(t, personID, response.Person.ID)
	require.InEpsilon(t, float32(response.PersonsCount), response.Person.Match, 0.001)
	require.Equal(t, "1906090343183157724859920008073008866", response.Person.SearchPointer)

	// ==================================================================================================================

	// Test person struct and data (names)
	require.Equal(t, len(response.Person.Names), response.AvailableData.Premium.Names)

	// Name number 1
	require.Equal(t, "Kal", response.Person.Names[0].First)
	require.Equal(t, "El", response.Person.Names[0].Last)
	require.Equal(t, response.Person.Names[0].First+" "+response.Person.Names[0].Last, response.Person.Names[0].Display)

	// Name number 2
	require.Equal(t, "Clark", response.Person.Names[1].First)
	require.Equal(t, "Joseph", response.Person.Names[1].Middle)
	require.Equal(t, "Kent", response.Person.Names[1].Last)
	require.Equal(t, response.Person.Names[1].First+" "+response.Person.Names[1].Middle+" "+response.Person.Names[1].Last, response.Person.Names[1].Display)

	// Name number 3
	require.Equal(t, "The red blue blur", response.Person.Names[2].Display)

	// ==================================================================================================================

	// Test person struct and data (emails)
	require.Equal(t, len(response.Person.Emails), response.AvailableData.Premium.Emails)

	// Test email 1
	require.Equal(t, "work", response.Person.Emails[0].Type)
	require.False(t, response.Person.Emails[0].EmailProvider)
	require.Equal(t, "clark.kent@thedailyplanet.com", response.Person.Emails[0].Address)
	require.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[0].Address))), response.Person.Emails[0].AddressMD5) //nolint:gosec // Testing PIPL API MD5 behavior

	// Test email 2
	require.True(t, response.Person.Emails[1].Disposable)
	require.False(t, response.Person.Emails[1].EmailProvider)
	require.Equal(t, "ck242@guerrillamail.com", response.Person.Emails[1].Address)
	require.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[1].Address))), response.Person.Emails[1].AddressMD5) //nolint:gosec // Testing PIPL API MD5 behavior

	// Test email 3
	require.Equal(t, "personal", response.Person.Emails[2].Type)
	require.True(t, response.Person.Emails[2].EmailProvider)
	require.Equal(t, "clark@gmail.com", response.Person.Emails[2].Address)
	require.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[2].Address))), response.Person.Emails[2].AddressMD5) //nolint:gosec // Testing PIPL API MD5 behavior

	// Test email 4
	require.True(t, response.Person.Emails[3].Disposable)
	require.False(t, response.Person.Emails[3].EmailProvider)
	require.Equal(t, "clark.kent@example.com", response.Person.Emails[3].Address)
	require.Equal(t, fmt.Sprintf("%x", md5.Sum([]byte(response.Person.Emails[3].Address))), response.Person.Emails[3].AddressMD5) //nolint:gosec // Testing PIPL API MD5 behavior

	// ==================================================================================================================

	// Test person struct and data (usernames)
	require.Equal(t, len(response.Person.Usernames), response.AvailableData.Premium.Usernames)

	// Test username 1
	require.Equal(t, "superman@facebook", response.Person.Usernames[0].Content)

	// Test username 2
	require.Equal(t, "@ckent", response.Person.Usernames[1].Content)

	// ==================================================================================================================

	// Test person struct and data (phones)
	require.Equal(t, len(response.Person.Phones), response.AvailableData.Premium.Phones)

	// Test phone 1
	require.Equal(t, "home_phone", response.Person.Phones[0].Type)
	require.Equal(t, 1, response.Person.Phones[0].CountryCode)
	require.Equal(t, int64(9785550145), response.Person.Phones[0].Number)
	require.Equal(t, "978-555-0145", response.Person.Phones[0].Display)
	require.Equal(t, "+1 978-555-0145", response.Person.Phones[0].DisplayInternational)

	// ==================================================================================================================

	// Test person struct and data (gender)
	require.Equal(t, genderMale, response.Person.Gender.Content)

	// ==================================================================================================================

	// Test person struct and data (dob)
	require.Equal(t, "32 years old", response.Person.DateOfBirth.Display)
	require.Equal(t, "1986-01-01", response.Person.DateOfBirth.DateRange.Start)
	require.Equal(t, "1987-05-13", response.Person.DateOfBirth.DateRange.End)

	// ==================================================================================================================

	// Test person struct and data (languages)
	require.Equal(t, len(response.Person.Languages), response.AvailableData.Premium.Languages)

	// Test language 1
	require.Equal(t, DefaultCountry, response.Person.Languages[0].Region)
	require.Equal(t, DefaultLanguage, response.Person.Languages[0].Language)
	require.Equal(t, DefaultDisplayLanguage, response.Person.Languages[0].Display)

	// ==================================================================================================================

	// Test person struct and data (ethnicities)
	require.Equal(t, len(response.Person.Ethnicities), response.AvailableData.Premium.Ethnicities)

	// Test ethnicities
	require.Equal(t, "other", response.Person.Ethnicities[0].Content)
	require.Equal(t, "american_indian", response.Person.Ethnicities[1].Content)
	require.Equal(t, "white", response.Person.Ethnicities[2].Content)

	// ==================================================================================================================

	// Test person struct and data (origin_countries)
	require.Equal(t, len(response.Person.OriginCountries), response.AvailableData.Premium.OriginCountries)

	// Test countries
	require.Equal(t, DefaultCountry, response.Person.OriginCountries[0].Country)

	// ==================================================================================================================

	// Test person struct and data (addresses)
	require.Equal(t, len(response.Person.Addresses), response.AvailableData.Premium.Addresses)

	// Test address 1
	require.Equal(t, "2005-02-12", response.Person.Addresses[0].ValidSince)
	require.Equal(t, "work", response.Person.Addresses[0].Type)
	require.Equal(t, DefaultCountry, response.Person.Addresses[0].Country)
	require.Equal(t, testState, response.Person.Addresses[0].State)
	require.Equal(t, "Metropolis", response.Person.Addresses[0].City)
	require.Equal(t, "Broadway", response.Person.Addresses[0].Street)
	require.Equal(t, "1000", response.Person.Addresses[0].House)
	require.Equal(t, "355", response.Person.Addresses[0].Apartment)
	require.Equal(t, "1000-355 Broadway, Metropolis, Kansas", response.Person.Addresses[0].Display)

	// Test address 2
	require.Equal(t, "1999-02-01", response.Person.Addresses[1].ValidSince)
	require.Equal(t, "home", response.Person.Addresses[1].Type)
	require.Equal(t, DefaultCountry, response.Person.Addresses[1].Country)
	require.Equal(t, testState, response.Person.Addresses[1].State)
	require.Equal(t, testCity, response.Person.Addresses[1].City)
	require.Equal(t, testStreet, response.Person.Addresses[1].Street)
	require.Equal(t, testHouseNumber, response.Person.Addresses[1].House)
	require.Equal(t, testApartment, response.Person.Addresses[1].Apartment)
	require.Equal(t, "66605", response.Person.Addresses[1].ZipCode)
	require.Equal(t, "10-1 Hickory Lane, Smallville, Kansas", response.Person.Addresses[1].Display)

	// ==================================================================================================================

	// Test person struct and data (jobs)
	require.Equal(t, len(response.Person.Jobs), response.AvailableData.Premium.Jobs)

	// Test job 1
	require.Equal(t, "Field Reporter", response.Person.Jobs[0].Title)
	require.Equal(t, "The Daily Planet", response.Person.Jobs[0].Organization)
	require.Equal(t, "Journalism", response.Person.Jobs[0].Industry)
	require.Equal(t, "Field Reporter at The Daily Planet (2000-2012)", response.Person.Jobs[0].Display)
	require.Equal(t, "2000-12-08", response.Person.Jobs[0].DateRange.Start)
	require.Equal(t, "2012-10-09", response.Person.Jobs[0].DateRange.End)

	// Test job 2
	require.Equal(t, "Junior Reporter", response.Person.Jobs[1].Title)
	require.Equal(t, "The Daily Planet", response.Person.Jobs[1].Organization)
	require.Equal(t, "Journalism", response.Person.Jobs[1].Industry)
	require.Equal(t, "1999-10-10", response.Person.Jobs[1].DateRange.Start)
	require.Equal(t, "2000-10-10", response.Person.Jobs[1].DateRange.End)
	require.Equal(t, "Junior Reporter at The Daily Planet (1999-2000)", response.Person.Jobs[1].Display)

	// Test job 3
	require.Equal(t, "Top Reporter", response.Person.Jobs[2].Title)
	require.Equal(t, "The Daily Planet", response.Person.Jobs[2].Organization)
	require.Equal(t, "Reporting", response.Person.Jobs[2].Industry)
	require.Equal(t, "Top Reporter at The Daily Planet", response.Person.Jobs[2].Display)

	// ==================================================================================================================

	// Test person struct and data (educations)
	require.Equal(t, len(response.Person.Educations), response.AvailableData.Premium.Educations)

	// Test education #1
	require.Equal(t, "B.Sc Advanced Science", response.Person.Educations[0].Degree)
	require.Equal(t, "Metropolis University", response.Person.Educations[0].School)
	require.Equal(t, "B.Sc Advanced Science from Metropolis University (2005-2008)", response.Person.Educations[0].Display)
	require.Equal(t, "2005-09-01", response.Person.Educations[0].DateRange.Start)
	require.Equal(t, "2008-05-14", response.Person.Educations[0].DateRange.End)

	// Test education #2
	require.Equal(t, "Smallville High", response.Person.Educations[1].School)
	require.Equal(t, "2001-09-01", response.Person.Educations[1].DateRange.Start)
	require.Equal(t, "2005-06-01", response.Person.Educations[1].DateRange.End)
	require.Equal(t, "Smallville High (2001-2005)", response.Person.Educations[1].Display)

	// ==================================================================================================================

	// Test person struct and data (relationships)
	require.Equal(t, len(response.Person.Relationships), response.AvailableData.Premium.Relationships)

	// ==================================================================================================================

	// Test person struct and data (user_ids)
	require.Equal(t, len(response.Person.UserIDs), response.AvailableData.Premium.UserIDs)

	// Test user id #1
	require.Equal(t, "11231@facebook", response.Person.UserIDs[0].Content)

	// ==================================================================================================================

	// Test person struct and data (images)
	require.Equal(t, len(response.Person.Images), response.AvailableData.Premium.Images)

	// Test images
	require.Equal(t, "https://vignette1.wikia.nocookie.net/smallville/images/e/ea/Buddies_forever.jpg", response.Person.Images[0].URL)
	require.Equal(t, "https://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg", response.Person.Images[1].URL)

	// ==================================================================================================================

	// Test person struct and data (urls) or (social profiles)
	require.Equal(t, len(response.Person.URLs), response.AvailableData.Premium.SocialProfiles)

	// Test url #1
	require.Equal(t, "edc6aa8fa3f211cfad7c12a0ba5b32f4", response.Person.URLs[0].SourceID)
	require.Equal(t, "linkedin.com", response.Person.URLs[0].Domain)
	require.Equal(t, "LinkedIn", response.Person.URLs[0].Name)
	require.Equal(t, "professional_and_business", response.Person.URLs[0].Category)
	require.Equal(t, "https://linkedin.com/clark.kent", response.Person.URLs[0].URL)

	// Test url #2
	require.Equal(t, "5d836a4acc55922e49fc709c7a39e233", response.Person.URLs[1].SourceID)
	require.Equal(t, "facebook.com", response.Person.URLs[1].Domain)
	require.Equal(t, "Facebook", response.Person.URLs[1].Name)
	require.Equal(t, "personal_profiles", response.Person.URLs[1].Category)
	require.Equal(t, "https://facebook.com/superman", response.Person.URLs[1].URL)

	// Test url #3
	require.Equal(t, "linkedin.com", response.Person.URLs[2].Domain)
	require.Equal(t, "professional_and_business", response.Person.URLs[2].Category)
	require.Equal(t, "https://www.linkedin.com/pub/superman/20/7a/365", response.Person.URLs[2].URL)
}

// Test_PersonNotFoundResponse test a person not found response JSON
func Test_PersonNotFoundResponse(t *testing.T) {
	// Load the response data
	response, err := loadResponseData("response_not_found.json")
	require.NoError(t, err)

	require.Equal(t, 0, response.PersonsCount)
	require.Equal(t, 0, response.AvailableSources)
	require.Equal(t, http.StatusOK, response.HTTPStatusCode)
	require.Empty(t, response.Error)
}

// Test_BadKeyResponse test a bad api key response JSON
func Test_BadKeyResponse(t *testing.T) {
	// Load the response data
	response, err := loadResponseData("response_bad_key.json")
	require.NoError(t, err)

	require.Equal(t, 0, response.PersonsCount)
	require.Equal(t, 0, response.AvailableSources)
	require.Equal(t, http.StatusForbidden, response.HTTPStatusCode)
	require.Equal(t, "Unrecognized API key", response.Error)
}

// Test_PackageErrorResponse test a bad package response JSON
func Test_PackageErrorResponse(t *testing.T) {
	// Load the response data
	response, err := loadResponseData("response_package_error.json")
	require.NoError(t, err)

	require.Equal(t, 0, response.PersonsCount)
	require.Equal(t, 0, response.AvailableSources)
	require.Equal(t, http.StatusBadRequest, response.HTTPStatusCode)
	require.Equal(t, "Your data package does not contain email", response.Error)
}
