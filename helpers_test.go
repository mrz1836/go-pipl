package pipl

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

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

	// Test missing first and last
	err := person.AddName("", "", "", "mr", "jr")
	if err == nil {
		t.Fatal("missing error, first and last are missing")
	}

	// Reset
	person = NewPerson()
	_ = person.AddName("clark", "ryan", "kent", "mr", "jr")
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

	// Test too short
	person := NewPerson()
	err := person.AddNameRaw("clark")
	if err == nil {
		t.Fatal("missing error, should have error for too short")
	}

	// Reset
	person = NewPerson()
	_ = person.AddNameRaw("clark ryan kent")
	if len(person.Names) == 0 {
		t.Fatal("expected a name in this person object")
	}
	if person.Names[0].Raw != "clark ryan kent" {
		t.Fatalf("expected value to be clark ryan kent, got %s", person.Names[0].Raw)
	}
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

	// Invalid email
	person := NewPerson()
	err := person.AddEmail("clarkkent")
	if err == nil {
		t.Fatal("should have failed, invalid email")
	}

	// Empty email
	person = NewPerson()
	err = person.AddEmail("")
	if err == nil {
		t.Fatal("should have failed, invalid email")
	}

	// Valid email
	person = NewPerson()
	_ = person.AddEmail("clarkkent@gmail.com")
	if len(person.Emails) == 0 {
		t.Fatal("expected an email in this person object")
	}
	if person.Emails[0].Address != "clarkkent@gmail.com" {
		t.Fatalf("expected value to be clarkkent@gmail.com, got %s", person.Emails[0].Address)
	}
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

// TestAddUsername test adding an username to a person object
func TestAddUsername(t *testing.T) {
	// Bad user id and service provider
	person := NewPerson()
	err := person.AddUsername("c", "x")
	if err == nil {
		t.Fatal("should have failed, username too short")
	}

	// Unknown service provider
	person = NewPerson()
	err = person.AddUsername("clarkkent", "notFound")
	if err == nil {
		t.Fatal("should have failed, service provider not found")
	}

	// Reset
	person = NewPerson()
	_ = person.AddUsername("clarkkent", "twitter")
	if len(person.Usernames) == 0 {
		t.Fatal("expected a username in this person object")
	}
	if person.Usernames[0].Content != "clarkkent@twitter" {
		t.Fatalf("expected value to be clarkkent, got %s", person.Usernames[0].Content)
	}
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

	// Bad user id and service provider
	person := NewPerson()
	err := person.AddUserID("c", "x")
	if err == nil {
		t.Fatal("should have failed, user id too short")
	}

	// Unknown service provider
	person = NewPerson()
	err = person.AddUserID("clarkkent", "notFound")
	if err == nil {
		t.Fatal("should have failed, service provider not known")
	}

	// Reset
	person = NewPerson()
	_ = person.AddUserID("clarkkent", "twitter")
	if len(person.UserIDs) == 0 {
		t.Fatal("expected a user id in this person object")
	}
	if person.UserIDs[0].Content != "clarkkent@twitter" {
		t.Fatalf("expected value to be clarkkent, got %s", person.UserIDs[0].Content)
	}
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

	// Missing both phone and country code
	person := NewPerson()
	err := person.AddPhone(0, 0)
	if err == nil {
		t.Fatal("should have failed, missing phone")
	}

	// Missing country code
	person = NewPerson()
	err = person.AddPhone(9785550145, 0)
	if err == nil {
		t.Fatal("should have failed, missing country code")
	}

	// Missing phone
	person = NewPerson()
	err = person.AddPhone(0, 1)
	if err == nil {
		t.Fatal("should have failed, missing phone")
	}

	// Valid phone
	person = NewPerson()
	_ = person.AddPhone(9785550145, 1)
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

	// Too short
	person := NewPerson()
	err := person.AddPhoneRaw("12")
	if err == nil {
		t.Fatal("should have failed, phone too short")
	}

	// Reset / Valid phone
	person = NewPerson()
	_ = person.AddPhoneRaw("19785550145")
	if len(person.Phones) == 0 {
		t.Fatal("expected a phone in this person object")
	}
	if person.Phones[0].Raw != "19785550145" {
		t.Fatalf("expected value to be 19785550145, got %s", person.Phones[0].Raw)
	}
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

	// Missing
	person := NewPerson()
	err := person.SetGender("")
	if err == nil {
		t.Fatal("should have failed, missing gender")
	}

	// Invalid
	person = NewPerson()
	err = person.SetGender("binary")
	if err == nil {
		t.Fatal("should have failed, invalid gender")
	}

	// Valid values
	person = NewPerson()
	_ = person.SetGender("male")
	if person.Gender.Content != "male" {
		t.Fatalf("expected value to be male, got %s", person.Gender.Content)
	}
	_ = person.SetGender("female")
	if person.Gender.Content != "female" {
		t.Fatalf("expected value to be female, got %s", person.Gender.Content)
	}
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

	// Missing dates
	person := NewPerson()
	err := person.SetDateOfBirth("", "")
	if err == nil {
		t.Fatal("should have failed, missing dates")
	}

	// Missing dates
	person = NewPerson()
	err = person.SetDateOfBirth("1981-01-01", "")
	if err == nil {
		t.Fatal("should have failed, missing dates")
	}

	// Invalid start dates
	person = NewPerson()
	err = person.SetDateOfBirth("19810101", "1987-01-31")
	if err == nil {
		t.Fatal("should have failed, invalid dates")
	}

	// Invalid start dates
	person = NewPerson()
	err = person.SetDateOfBirth("1987-01-01", "19870131")
	if err == nil {
		t.Fatal("should have failed, invalid dates")
	}

	// Valid dates
	person = NewPerson()
	_ = person.SetDateOfBirth("1987-01-01", "1987-01-31")
	if person.DateOfBirth.DateRange.Start != "1987-01-01" {
		t.Fatalf("expected value to be 1987-01-01, got %s", person.DateOfBirth.DateRange.Start)
	}
	if person.DateOfBirth.DateRange.End != "1987-01-31" {
		t.Fatalf("expected value to be 1987-01-31, got %s", person.DateOfBirth.DateRange.End)
	}
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

	// Invalid language code
	person := NewPerson()
	err := person.AddLanguage("wrong", DefaultCountry)
	if err == nil {
		t.Fatal("should have failed, invalid language code")
	}

	// Invalid country
	person = NewPerson()
	err = person.AddLanguage(DefaultLanguage, "wrong")
	if err == nil {
		t.Fatal("should have failed, invalid language code")
	}

	person = NewPerson()
	_ = person.AddLanguage(DefaultLanguage, DefaultCountry)
	if len(person.Languages) == 0 {
		t.Fatal("expected a language in this person object")
	}
	if person.Languages[0].Language != DefaultLanguage {
		t.Fatalf("expected value to be %s, got %s", DefaultLanguage, person.Languages[0].Language)
	}
	if person.Languages[0].Region != DefaultCountry {
		t.Fatalf("expected value to be %s, got %s", DefaultCountry, person.Languages[0].Region)
	}
	if person.Languages[0].Display != DefaultLanguage+"_"+DefaultCountry {
		t.Fatalf("expected value to be %s, got %s", DefaultLanguage+"_"+DefaultCountry, person.Languages[0].Region)
	}
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

// TestAddEthnicity test adding a ethnicity to a person object
func TestAddEthnicity(t *testing.T) {

	// Missing value
	person := NewPerson()
	err := person.AddEthnicity("")
	if err == nil {
		t.Fatal("should have failed, missing value")
	}

	// Invalid value
	person = NewPerson()
	err = person.AddEthnicity("unknown")
	if err == nil {
		t.Fatal("should have failed, missing value")
	}

	person = NewPerson()
	_ = person.AddEthnicity("white")
	if len(person.Ethnicities) == 0 {
		t.Fatal("expected an ethnicity in this person object")
	}
	if person.Ethnicities[0].Content != "white" {
		t.Fatalf("expected value to be white, got %s", person.Ethnicities[0].Content)
	}
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

// TestAddOriginCountry test adding a an origin country to a person object
func TestAddOriginCountry(t *testing.T) {
	person := NewPerson()
	err := person.AddOriginCountry("")
	if err == nil {
		t.Fatal("should have failed, missing country code")
	}

	person = NewPerson()
	_ = person.AddOriginCountry(DefaultCountry)
	if len(person.OriginCountries) == 0 {
		t.Fatal("expected an origin country in this person object")
	}
	if person.OriginCountries[0].Country != DefaultCountry {
		t.Fatalf("expected value to be %s, got %s", DefaultCountry, person.OriginCountries[0].Country)
	}
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

	// Missing number and street
	person := NewPerson()
	err := person.AddAddress("", "", "1", "Smallville", "KS", DefaultCountry, "123")
	if err == nil {
		t.Fatal("should have failed, missing number/street")
	}

	// Missing city and state
	person = NewPerson()
	err = person.AddAddress("10", "Hickory Lane", "1", "", "", DefaultCountry, "123")
	if err == nil {
		t.Fatal("should have failed, missing number/street")
	}

	// Valid address
	person = NewPerson()
	_ = person.AddAddress("10", "Hickory Lane", "1", "Smallville", "KS", DefaultCountry, "123")
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
	if person.Addresses[0].Country != DefaultCountry {
		t.Fatalf("expected value to be %s, got %s", DefaultCountry, person.Addresses[0].Country)
	}
	if person.Addresses[0].POBox != "123" {
		t.Fatalf("expected value to be 123, got %s", person.Addresses[0].POBox)
	}
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

// TestAddAddressRaw test adding a an address to a person object
func TestAddAddressRaw(t *testing.T) {

	// Too short
	person := NewPerson()
	err := person.AddAddressRaw("10")
	if err == nil {
		t.Fatal("should have failed, address too short")
	}

	// Valid address
	person = NewPerson()
	_ = person.AddAddressRaw("10 Hickory Lane, Kansas, " + DefaultCountry)
	if len(person.Addresses) == 0 {
		t.Fatal("expected an address in this person object")
	}
	if person.Addresses[0].Raw != "10 Hickory Lane, Kansas, "+DefaultCountry {
		t.Fatalf("expected value to be 10 Hickory Lane, Kansas, %s, got %s", DefaultCountry, person.Addresses[0].Raw)
	}
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

	// Missing title
	person := NewPerson()
	err := person.AddJob("", "daily post", "news", "2010-01-01", "2011-01-01")
	if err == nil {
		t.Fatal("should have failed, missing title")
	}

	// Missing organization
	person = NewPerson()
	err = person.AddJob("reporter", "", "news", "2010-01-01", "2011-01-01")
	if err == nil {
		t.Fatal("should have failed, missing organization")
	}

	// Valid job
	person = NewPerson()
	_ = person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
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

// TestAddEducation test adding a an education to a person object
func TestAddEducation(t *testing.T) {

	// Missing school
	person := NewPerson()
	err := person.AddEducation("masters", "", "2010-01-01", "2011-01-01")
	if err == nil {
		t.Fatal("should have failed, missing school")
	}

	// Valid education
	person = NewPerson()
	_ = person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
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

// TestAddURL test adding a url to a person object
func TestAddURL(t *testing.T) {
	person := NewPerson()
	err := person.AddURL("http")
	if err == nil {
		t.Fatal("should have returned an error, invalid url")
	}

	// Reset
	person = NewPerson()
	_ = person.AddURL("https://twitter.com/clarkkent")
	if len(person.URLs) == 0 {
		t.Fatal("expected a url in this person object")
	}
	if person.URLs[0].URL != "https://twitter.com/clarkkent" {
		t.Fatalf("expected value to be https://twitter.com/clarkkent, got %s", person.URLs[0].URL)
	}
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

	// Create the client
	client, err := NewClient("1234567890", nil)
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
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("favicon=%t", client.Parameters.Thumbnail.Favicon)) {
		t.Fatal("expected value is not the same", client.Parameters.Thumbnail.Favicon)
	}

	// Does it have the right zoom face
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("zoom_face=%t", client.Parameters.Thumbnail.ZoomFace)) {
		t.Fatal("expected value is not the same", client.Parameters.Thumbnail.ZoomFace)
	}

	// Does it have the right token
	if !strings.Contains(person.Images[0].ThumbnailURL, fmt.Sprintf("tokens=%s", person.Images[0].ThumbnailToken)) {
		t.Fatal("expected value is not the same", person.Images[0].ThumbnailToken)
	}
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
