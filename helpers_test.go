package pipl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewPerson testing new person function
func TestNewPerson(t *testing.T) {
	person := NewPerson()
	require.Equal(t, "*pipl.Person", reflect.TypeOf(person).String())
}

// BenchmarkNewPerson benchmarks the NewPerson method
func BenchmarkNewPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewPerson()
	}
}

// TestAddName test adding a name to a person object
func TestAddName(t *testing.T) {
	t.Parallel()

	person := NewPerson()

	// Test missing first and last
	err := person.AddName("", "", "", "mr", "jr")
	require.Error(t, err)

	// Reset
	person = NewPerson()
	err = person.AddName("clark", "ryan", "kent", "mr", "jr")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Names))
	require.Equal(t, "clark", person.Names[0].First)
	require.Equal(t, "ryan", person.Names[0].Middle)
	require.Equal(t, "kent", person.Names[0].Last)
	require.Equal(t, "mr", person.Names[0].Prefix)
	require.Equal(t, "jr", person.Names[0].Suffix)
}

// ExamplePerson_AddName example using AddName()
func ExamplePerson_AddName() {
	person := NewPerson()
	_ = person.AddName("clark", "ryan", "kent", "mr", "jr")
	fmt.Println(person.Names[0].First + " " + person.Names[0].Last)
	// Output: clark kent
}

// BenchmarkAddName benchmarks the AddName method
func BenchmarkAddName(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddName("clark", "ryan", "kent", "mr", "jr")
	}
}

// TestAddNameRaw test adding a raw name to a person object
func TestAddNameRaw(t *testing.T) {
	t.Parallel()

	// Test too short
	person := NewPerson()
	err := person.AddNameRaw("clark")
	require.Error(t, err)

	// Reset
	person = NewPerson()
	err = person.AddNameRaw("clark ryan kent")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Names))
	require.Equal(t, "clark ryan kent", person.Names[0].Raw)
}

// ExamplePerson_AddNameRaw example using AddNameRaw()
func ExamplePerson_AddNameRaw() {
	person := NewPerson()
	_ = person.AddNameRaw("clark kent")
	fmt.Println(person.Names[0].Raw)
	// Output: clark kent
}

// BenchmarkAddNameRaw benchmarks the AddNameRaw method
func BenchmarkAddNameRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddNameRaw("clark ryan kent")
	}
}

// TestAddEmail test adding an email to a person object
func TestAddEmail(t *testing.T) {
	t.Parallel()

	// Invalid email
	person := NewPerson()
	err := person.AddEmail("clarkkent")
	require.Error(t, err)

	// Empty email
	person = NewPerson()
	err = person.AddEmail("")
	require.Error(t, err)

	// Valid email
	person = NewPerson()
	err = person.AddEmail("clarkkent@gmail.com")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Emails))
	require.Equal(t, "clarkkent@gmail.com", person.Emails[0].Address)
}

// ExamplePerson_AddEmail example using AddEmail()
func ExamplePerson_AddEmail() {
	person := NewPerson()
	_ = person.AddEmail("clarkkent@gmail.com")
	fmt.Println(person.Emails[0].Address)
	// Output:clarkkent@gmail.com
}

// BenchmarkAddEmail benchmarks the AddEmail method
func BenchmarkAddEmail(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddEmail("clarkkent@gmail.com")
	}
}

// TestAddUsername test adding a username to a person object
func TestAddUsername(t *testing.T) {
	t.Parallel()

	// Bad user id and service provider
	person := NewPerson()
	err := person.AddUsername("c", "x")
	require.Error(t, err)

	// Unknown service provider
	person = NewPerson()
	err = person.AddUsername("clarkkent", "notFound")
	require.Error(t, err)

	// Reset
	person = NewPerson()
	err = person.AddUsername("clarkkent", "twitter")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Usernames))
	require.Equal(t, "clarkkent@twitter", person.Usernames[0].Content)
}

// ExamplePerson_AddUsername example using AddUsername()
func ExamplePerson_AddUsername() {
	person := NewPerson()
	_ = person.AddUsername("clarkkent", "twitter")
	fmt.Println(person.Usernames[0].Content)
	// Output:clarkkent@twitter
}

// BenchmarkAddUsername benchmarks the AddUsername method
func BenchmarkAddUsername(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddUsername("clarkkent", "twitter")
	}
}

// TestAddUserID test adding a user id to a person object
func TestAddUserID(t *testing.T) {
	t.Parallel()

	// Bad user id and service provider
	person := NewPerson()
	err := person.AddUserID("c", "x")
	require.Error(t, err)

	// Unknown service provider
	person = NewPerson()
	err = person.AddUserID("clarkkent", "notFound")
	require.Error(t, err)

	// Reset
	person = NewPerson()
	err = person.AddUserID("clarkkent", "twitter")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.UserIDs))
	require.Equal(t, "clarkkent@twitter", person.UserIDs[0].Content)
}

// ExamplePerson_AddUserID example using AddUserID()
func ExamplePerson_AddUserID() {
	person := NewPerson()
	_ = person.AddUserID("clarkkent", "twitter")
	fmt.Println(person.UserIDs[0].Content)
	// Output:clarkkent@twitter
}

// BenchmarkAddUserID benchmarks the AddUserID method
func BenchmarkAddUserID(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddUserID("clarkkent", "twitter")
	}
}

// TestAddPhone test adding a phone to a person object
func TestAddPhone(t *testing.T) {
	t.Parallel()

	// Missing both phone and country code
	person := NewPerson()
	err := person.AddPhone(0, 0)
	require.Error(t, err)

	// Missing country code
	person = NewPerson()
	err = person.AddPhone(9785550145, 0)
	require.Error(t, err)

	// Missing phone
	person = NewPerson()
	err = person.AddPhone(0, 1)
	require.Error(t, err)

	// Valid phone
	person = NewPerson()
	err = person.AddPhone(9785550145, 1)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Phones))
	require.Equal(t, int64(9785550145), person.Phones[0].Number)
	require.Equal(t, 1, person.Phones[0].CountryCode)
}

// ExamplePerson_AddPhone example using AddPhone()
func ExamplePerson_AddPhone() {
	person := NewPerson()
	_ = person.AddPhone(9785550145, 1)
	fmt.Println(person.Phones[0].Number)
	// Output:9785550145
}

// BenchmarkAddPhone benchmarks the AddPhone method
func BenchmarkAddPhone(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddPhone(9785550145, 1)
	}
}

// TestAddPhoneRaw test adding a phone to a person object
func TestAddPhoneRaw(t *testing.T) {
	t.Parallel()

	// Too short
	person := NewPerson()
	err := person.AddPhoneRaw("12")
	require.Error(t, err)

	// Reset / Valid phone
	person = NewPerson()
	err = person.AddPhoneRaw("19785550145")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Phones))
	require.Equal(t, "19785550145", person.Phones[0].Raw)
}

// ExamplePerson_AddPhoneRaw example using AddPhoneRaw()
func ExamplePerson_AddPhoneRaw() {
	person := NewPerson()
	_ = person.AddPhoneRaw("9785550145")
	fmt.Println(person.Phones[0].Raw)
	// Output:9785550145
}

// BenchmarkAddPhoneRaw benchmarks the AddPhoneRaw method
func BenchmarkAddPhoneRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddPhoneRaw("9785550145")
	}
}

// TestSetGender test setting a gender on a person object
func TestSetGender(t *testing.T) {
	t.Parallel()

	// Missing
	person := NewPerson()
	err := person.SetGender("")
	require.Error(t, err)

	// Invalid
	person = NewPerson()
	err = person.SetGender("binary")
	require.Error(t, err)

	// Valid values
	person = NewPerson()
	err = person.SetGender("male")
	require.NoError(t, err)
	require.Equal(t, "male", person.Gender.Content)

	err = person.SetGender("female")
	require.NoError(t, err)
	require.Equal(t, "female", person.Gender.Content)
}

// ExamplePerson_SetGender example using SetGender()
func ExamplePerson_SetGender() {
	person := NewPerson()
	_ = person.SetGender("male")
	fmt.Println(person.Gender.Content)
	// Output:male
}

// BenchmarkSetGender benchmarks the SetGender method
func BenchmarkSetGender(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.SetGender("male")
	}
}

// TestSetDateOfBirth test setting a DOB on a person object
func TestSetDateOfBirth(t *testing.T) {
	t.Parallel()

	// Missing dates
	person := NewPerson()
	err := person.SetDateOfBirth("", "")
	require.Error(t, err)

	// Missing dates
	person = NewPerson()
	err = person.SetDateOfBirth("1981-01-01", "")
	require.Error(t, err)

	// Invalid start dates
	person = NewPerson()
	err = person.SetDateOfBirth("19810101", "1987-01-31")
	require.Error(t, err)

	// Invalid start dates
	person = NewPerson()
	err = person.SetDateOfBirth("1987-01-01", "19870131")
	require.Error(t, err)

	// Valid dates
	person = NewPerson()
	err = person.SetDateOfBirth("1987-01-01", "1987-01-31")
	require.NoError(t, err)
	require.Equal(t, "1987-01-01", person.DateOfBirth.DateRange.Start)
	require.Equal(t, "1987-01-31", person.DateOfBirth.DateRange.End)
}

// ExamplePerson_SetDateOfBirth example using SetDateOfBirth()
func ExamplePerson_SetDateOfBirth() {
	person := NewPerson()
	_ = person.SetDateOfBirth("1987-01-01", "1987-01-01")
	fmt.Println(person.DateOfBirth.DateRange.Start)
	// Output:1987-01-01
}

// BenchmarkSetDateOfBirth benchmarks the SetDateOfBirth method
func BenchmarkSetDateOfBirth(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.SetDateOfBirth("1987-01-01", "1987-01-01")
	}
}

// TestAddLanguage test adding a language to a person object
func TestAddLanguage(t *testing.T) {
	t.Parallel()

	// Invalid language code
	person := NewPerson()
	err := person.AddLanguage("wrong", DefaultCountry)
	require.Error(t, err)

	// Invalid country
	person = NewPerson()
	err = person.AddLanguage(DefaultLanguage, "wrong")
	require.Error(t, err)

	person = NewPerson()
	err = person.AddLanguage(DefaultLanguage, DefaultCountry)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Languages))
	require.Equal(t, DefaultLanguage, person.Languages[0].Language)
	require.Equal(t, DefaultCountry, person.Languages[0].Region)
	require.Equal(t, DefaultLanguage+"_"+DefaultCountry, person.Languages[0].Display)
}

// ExamplePerson_AddLanguage example using AddLanguage()
func ExamplePerson_AddLanguage() {
	person := NewPerson()
	_ = person.AddLanguage(DefaultLanguage, DefaultCountry)
	fmt.Println(person.Languages[0].Display)
	// Output:en_US
}

// BenchmarkAddLanguage benchmarks the AddLanguage method
func BenchmarkAddLanguage(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddLanguage(DefaultLanguage, DefaultCountry)
	}
}

// TestAddEthnicity test adding an ethnicity to a person object
func TestAddEthnicity(t *testing.T) {
	t.Parallel()

	// Missing value
	person := NewPerson()
	err := person.AddEthnicity("")
	require.Error(t, err)

	// Invalid value
	person = NewPerson()
	err = person.AddEthnicity("unknown")
	require.Error(t, err)

	person = NewPerson()
	err = person.AddEthnicity("white")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Ethnicities))
	require.Equal(t, "white", person.Ethnicities[0].Content)
}

// ExamplePerson_AddEthnicity example using AddEthnicity()
func ExamplePerson_AddEthnicity() {
	person := NewPerson()
	_ = person.AddEthnicity("white")
	fmt.Println(person.Ethnicities[0].Content)
	// Output:white
}

// BenchmarkAddEthnicity benchmarks the AddEthnicity method
func BenchmarkAddEthnicity(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddEthnicity("white")
	}
}

// TestAddOriginCountry test adding an origin country to a person object
func TestAddOriginCountry(t *testing.T) {
	t.Parallel()

	person := NewPerson()
	err := person.AddOriginCountry("")
	require.Error(t, err)

	person = NewPerson()
	err = person.AddOriginCountry(DefaultCountry)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.OriginCountries))
	require.Equal(t, DefaultCountry, person.OriginCountries[0].Country)
}

// ExamplePerson_AddOriginCountry example using AddOriginCountry()
func ExamplePerson_AddOriginCountry() {
	person := NewPerson()
	_ = person.AddOriginCountry(DefaultCountry)
	fmt.Println(person.OriginCountries[0].Country)
	// Output:US
}

// BenchmarkAddOriginCountry benchmarks the AddOriginCountry method
func BenchmarkAddOriginCountry(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddOriginCountry(DefaultCountry)
	}
}

// TestAddAddress test adding an address to a person object
func TestAddAddress(t *testing.T) {
	t.Parallel()

	// Missing number and street
	person := NewPerson()
	err := person.AddAddress("", "", "1", "Smallville", "KS", DefaultCountry, "123")
	require.Error(t, err)

	// Missing city and state
	person = NewPerson()
	err = person.AddAddress("10", "Hickory Lane", "1", "", "", DefaultCountry, "123")
	require.Error(t, err)

	// Valid address
	person = NewPerson()
	err = person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", DefaultCountry, "123")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Addresses))
	require.Equal(t, "10", person.Addresses[0].House)
	require.Equal(t, "Hickory Lane", person.Addresses[0].Street)
	require.Equal(t, "1", person.Addresses[0].Apartment)
	require.Equal(t, "Smallville", person.Addresses[0].City)
	require.Equal(t, "KS", person.Addresses[0].State)
	require.Equal(t, DefaultCountry, person.Addresses[0].Country)
	require.Equal(t, "123", person.Addresses[0].POBox)
}

// ExamplePerson_AddAddress example using AddAddress()
func ExamplePerson_AddAddress() {
	person := NewPerson()
	_ = person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", DefaultCountry, "123")
	fmt.Println(person.Addresses[0].House + " " + person.Addresses[0].Street + ", " + person.Addresses[0].City + " " + person.Addresses[0].State)
	// Output:10 Hickory Lane, Smallville KS
}

// BenchmarkAddAddress benchmarks the AddAddress method
func BenchmarkAddAddress(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", DefaultCountry, "123")
	}
}

// TestAddAddressRaw test adding an address to a person object
func TestAddAddressRaw(t *testing.T) {
	t.Parallel()

	// Too short
	person := NewPerson()
	err := person.AddAddressRaw("10")
	require.Error(t, err)

	// Valid address
	person = NewPerson()
	err = person.AddAddressRaw("10 Hickory Lane, Kansas, " + DefaultCountry)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Addresses))
	require.Equal(t, "10 Hickory Lane, Kansas, "+DefaultCountry, person.Addresses[0].Raw)
}

// ExamplePerson_AddAddressRaw example using AddAddressRaw()
func ExamplePerson_AddAddressRaw() {
	person := NewPerson()
	_ = person.AddAddressRaw("10 Hickory Lane, Kansas, " + DefaultCountry)
	fmt.Println(person.Addresses[0].Raw)
	// Output:10 Hickory Lane, Kansas, US
}

// BenchmarkAddAddressRaw benchmarks the AddAddressRaw method
func BenchmarkAddAddressRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddAddressRaw("10 Hickory Lane, Kansas, " + DefaultCountry)
	}
}

// TestAddJob test adding a job to a person object
func TestAddJob(t *testing.T) {
	t.Parallel()

	// Missing title
	person := NewPerson()
	err := person.AddJob("", "daily post", "news", "2010-01-01", "2011-01-01")
	require.Error(t, err)

	// Missing organization
	person = NewPerson()
	err = person.AddJob("reporter", "", "news", "2010-01-01", "2011-01-01")
	require.Error(t, err)

	// Valid job
	person = NewPerson()
	err = person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Jobs))
	require.Equal(t, "reporter", person.Jobs[0].Title)
	require.Equal(t, "daily post", person.Jobs[0].Organization)
	require.Equal(t, "news", person.Jobs[0].Industry)
	require.Equal(t, "2010-01-01", person.Jobs[0].DateRange.Start)
	require.Equal(t, "2011-01-01", person.Jobs[0].DateRange.End)
}

// ExamplePerson_AddJob example using AddJob()
func ExamplePerson_AddJob() {
	person := NewPerson()
	_ = person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
	fmt.Println(person.Jobs[0].Title + " at " + person.Jobs[0].Organization + " in " + person.Jobs[0].Industry)
	// Output:reporter at daily post in news
}

// BenchmarkAddJob benchmarks the AddJob method
func BenchmarkAddJob(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
	}
}

// TestAddEducation test adding an education to a person object
func TestAddEducation(t *testing.T) {
	t.Parallel()

	// Missing school
	person := NewPerson()
	err := person.AddEducation("masters", "", "2010-01-01", "2011-01-01")
	require.Error(t, err)

	// Valid education
	person = NewPerson()
	err = person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.Educations))
	require.Equal(t, "masters", person.Educations[0].Degree)
	require.Equal(t, "fau", person.Educations[0].School)
	require.Equal(t, "2010-01-01", person.Educations[0].DateRange.Start)
	require.Equal(t, "2011-01-01", person.Educations[0].DateRange.End)
}

// ExamplePerson_AddEducation example using AddEducation()
func ExamplePerson_AddEducation() {
	person := NewPerson()
	_ = person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
	fmt.Println(person.Educations[0].Degree + " from " + person.Educations[0].School)
	// Output:masters from fau
}

// BenchmarkAddEducation benchmarks the AddEducation method
func BenchmarkAddEducation(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
	}
}

// TestAddURL test adding an url to a person object
func TestAddURL(t *testing.T) {
	t.Parallel()

	person := NewPerson()
	err := person.AddURL("http")
	require.Error(t, err)

	// Reset
	person = NewPerson()
	err = person.AddURL("https://twitter.com/clarkkent")
	require.NoError(t, err)
	require.NotEqual(t, 0, len(person.URLs))
	require.Equal(t, "https://twitter.com/clarkkent", person.URLs[0].URL)
}

// ExamplePerson_AddURL example using AddURL()
func ExamplePerson_AddURL() {
	person := NewPerson()
	_ = person.AddURL("https://twitter.com/clarkkent")
	fmt.Println(person.URLs[0].URL)
	// Output:https://twitter.com/clarkkent
}

// BenchmarkAddURL benchmarks the AddURL method
func BenchmarkAddURL(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddURL("https://twitter.com/clarkkent")
	}
}

// TestPerson_ProcessThumbnails test processing images for thumbnails
func TestPerson_ProcessThumbnails(t *testing.T) {
	t.Parallel()

	// Create the client
	client, err := NewClient("1234567890", nil)
	require.NoError(t, err)

	// Create person and image
	person := NewPerson()
	image := new(Image)
	image.URL = testImage
	image.ThumbnailToken = testThumbnailToken
	person.Images = append(person.Images, *image)

	// Process using defaults
	person.ProcessThumbnails(client)

	// Test for url
	require.NotEqual(t, 0, person.Images[0].ThumbnailURL)
	require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("width=%d", ThumbnailWidth))
	require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("height=%d", ThumbnailHeight))
	require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("favicon=%t", client.Parameters.Thumbnail.Favicon))
	require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("zoom_face=%t", client.Parameters.Thumbnail.ZoomFace))
	require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("tokens=%s", person.Images[0].ThumbnailToken))
}

// ExamplePerson_ProcessThumbnails example using ProcessThumbnails()
func ExamplePerson_ProcessThumbnails() {
	client, _ := NewClient("1234567890", nil)
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
	client, _ := NewClient("1234567890", nil)
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
