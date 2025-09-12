package pipl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

	t.Run("missing first and last", func(t *testing.T) {
		err := person.AddName("", "", "", "mr", "jr")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingFirstLastName)
	})

	t.Run("valid name", func(t *testing.T) {
		person = NewPerson()
		err := person.AddName(testFirstName, testMiddleName, testLastName, "mr", "jr")
		require.NoError(t, err)
		require.NotEmpty(t, person.Names)
		require.Equal(t, testFirstName, person.Names[0].First)
		require.Equal(t, testMiddleName, person.Names[0].Middle)
		require.Equal(t, testLastName, person.Names[0].Last)
		require.Equal(t, "mr", person.Names[0].Prefix)
		require.Equal(t, "jr", person.Names[0].Suffix)
	})
}

// ExamplePerson_AddName example using AddName()
func ExamplePerson_AddName() {
	person := NewPerson()
	_ = person.AddName(testFirstName, testMiddleName, testLastName, "mr", "jr")
	fmt.Println(person.Names[0].First + " " + person.Names[0].Last)
	// Output: clark kent
}

// BenchmarkAddName benchmarks the AddName method
func BenchmarkAddName(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddName(testFirstName, testMiddleName, testLastName, "mr", "jr")
	}
}

// TestAddNameRaw test adding a raw name to a person object
func TestAddNameRaw(t *testing.T) {
	t.Parallel()

	t.Run("too short", func(t *testing.T) {
		person := NewPerson()
		err := person.AddNameRaw(testFirstName)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrNameTooShort)
	})

	t.Run("valid name", func(t *testing.T) {
		person := NewPerson()
		err := person.AddNameRaw("clark ryan kent")
		require.NoError(t, err)
		require.NotEmpty(t, person.Names)
		require.Equal(t, "clark ryan kent", person.Names[0].Raw)
	})
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

	t.Run("invalid email", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEmail("clarkkent")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidEmailAddress)
	})

	t.Run("empty email", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEmail("")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidEmailAddress)
	})

	t.Run("valid email", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEmail(testEmailSecondary)
		require.NoError(t, err)
		require.NotEmpty(t, person.Emails)
		require.Equal(t, testEmailSecondary, person.Emails[0].Address)
	})
}

// ExamplePerson_AddEmail example using AddEmail()
func ExamplePerson_AddEmail() {
	person := NewPerson()
	_ = person.AddEmail(testEmailSecondary)
	fmt.Println(person.Emails[0].Address)
	// Output:clarkkent@gmail.com
}

// BenchmarkAddEmail benchmarks the AddEmail method
func BenchmarkAddEmail(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddEmail(testEmailSecondary)
	}
}

// TestAddUsername test adding a username to a person object
func TestAddUsername(t *testing.T) {
	t.Parallel()

	t.Run("bad user id and service provider", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUsername("c", "x")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUserNameTooShort)
	})

	t.Run("empty service provider", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUsername(testUserName, "")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrServiceProviderTooShort)
	})

	t.Run("unknown service provider", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUsername(testUserName, "notFound")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrServiceProviderNotAccepted)
	})

	t.Run("valid username", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUsername(testUserName, testUserNameServiceProvider)
		require.NoError(t, err)
		require.NotEmpty(t, person.Usernames)
		require.Equal(t, testUserName+"@"+testUserNameServiceProvider, person.Usernames[0].Content)
	})
}

// ExamplePerson_AddUsername example using AddUsername()
func ExamplePerson_AddUsername() {
	person := NewPerson()
	_ = person.AddUsername(testUserName, testUserNameServiceProvider)
	fmt.Println(person.Usernames[0].Content)
	// Output:clarkkent@twitter
}

// BenchmarkAddUsername benchmarks the AddUsername method
func BenchmarkAddUsername(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddUsername(testUserName, testUserNameServiceProvider)
	}
}

// TestAddUserID test adding a user id to a person object
func TestAddUserID(t *testing.T) {
	t.Parallel()

	t.Run("bad user id and service provider", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUserID("c", "x")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrUserIDTooShort)
	})

	t.Run("no service provider", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUserID(testUserName, "")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrServiceProviderTooShort)
	})

	t.Run("unknown service provider", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUserID(testUserName, "notFound")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrServiceProviderNotAccepted)
	})

	t.Run("valid user id", func(t *testing.T) {
		person := NewPerson()
		err := person.AddUserID(testUserName, testUserNameServiceProvider)
		require.NoError(t, err)
		require.NotEmpty(t, person.UserIDs)
		require.Equal(t, testUserName+"@"+testUserNameServiceProvider, person.UserIDs[0].Content)
	})
}

// ExamplePerson_AddUserID example using AddUserID()
func ExamplePerson_AddUserID() {
	person := NewPerson()
	_ = person.AddUserID(testUserName, testUserNameServiceProvider)
	fmt.Println(person.UserIDs[0].Content)
	// Output:clarkkent@twitter
}

// BenchmarkAddUserID benchmarks the AddUserID method
func BenchmarkAddUserID(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddUserID(testUserName, testUserNameServiceProvider)
	}
}

// TestAddPhone test adding a phone to a person object
func TestAddPhone(t *testing.T) {
	t.Parallel()

	t.Run("missing both phone and country code", func(t *testing.T) {
		person := NewPerson()
		err := person.AddPhone(0, 0)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidPhoneNumber)
	})

	t.Run("missing country code", func(t *testing.T) {
		person := NewPerson()
		err := person.AddPhone(testPhone, 0)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingCountryCode)
	})

	t.Run("missing phone", func(t *testing.T) {
		person := NewPerson()
		err := person.AddPhone(0, testPhoneCountryCode)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidPhoneNumber)
	})

	t.Run("valid phone", func(t *testing.T) {
		person := NewPerson()
		err := person.AddPhone(testPhone, testPhoneCountryCode)
		require.NoError(t, err)
		require.NotEmpty(t, person.Phones)
		require.Equal(t, testPhone, person.Phones[0].Number)
		require.Equal(t, 1, person.Phones[0].CountryCode)
	})
}

// ExamplePerson_AddPhone example using AddPhone()
func ExamplePerson_AddPhone() {
	person := NewPerson()
	_ = person.AddPhone(testPhone, testPhoneCountryCode)
	fmt.Println(person.Phones[0].Number)
	// Output:9785550145
}

// BenchmarkAddPhone benchmarks the AddPhone method
func BenchmarkAddPhone(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddPhone(testPhone, testPhoneCountryCode)
	}
}

// TestAddPhoneRaw test adding a phone to a person object
func TestAddPhoneRaw(t *testing.T) {
	t.Parallel()

	t.Run("too short", func(t *testing.T) {
		person := NewPerson()
		err := person.AddPhoneRaw("12")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidPhoneNumber)
	})

	t.Run("valid name", func(t *testing.T) {
		person := NewPerson()
		err := person.AddPhoneRaw(testPhoneRaw)
		require.NoError(t, err)
		require.NotEmpty(t, person.Phones)
		require.Equal(t, testPhoneRaw, person.Phones[0].Raw)
	})
}

// ExamplePerson_AddPhoneRaw example using AddPhoneRaw()
func ExamplePerson_AddPhoneRaw() {
	person := NewPerson()
	_ = person.AddPhoneRaw(testPhoneRaw)
	fmt.Println(person.Phones[0].Raw)
	// Output:19785550145
}

// BenchmarkAddPhoneRaw benchmarks the AddPhoneRaw method
func BenchmarkAddPhoneRaw(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddPhoneRaw(testPhoneRaw)
	}
}

// TestSetGender test setting a gender on a person object
func TestSetGender(t *testing.T) {
	t.Parallel()

	t.Run("missing gender", func(t *testing.T) {
		person := NewPerson()
		err := person.SetGender("")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidGender)
	})

	t.Run("invalid gender", func(t *testing.T) {
		person := NewPerson()
		err := person.SetGender("binary")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidGender)
	})

	t.Run("valid - male", func(t *testing.T) {
		person := NewPerson()
		err := person.SetGender(genderMale)
		require.NoError(t, err)
		require.Equal(t, genderMale, person.Gender.Content)
	})

	t.Run("valid - female", func(t *testing.T) {
		person := NewPerson()
		err := person.SetGender(genderFemale)
		require.NoError(t, err)
		require.Equal(t, genderFemale, person.Gender.Content)
	})
}

// ExamplePerson_SetGender example using SetGender()
func ExamplePerson_SetGender() {
	person := NewPerson()
	_ = person.SetGender(genderMale)
	fmt.Println(person.Gender.Content)
	// Output:male
}

// BenchmarkSetGender benchmarks the SetGender method
func BenchmarkSetGender(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.SetGender(genderMale)
	}
}

// TestSetDateOfBirth test setting a DOB on a person object
func TestSetDateOfBirth(t *testing.T) {
	t.Parallel()

	t.Run("missing both dates", func(t *testing.T) {
		person := NewPerson()
		err := person.SetDateOfBirth("", "")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingBirthDate)
	})

	t.Run("missing end date", func(t *testing.T) {
		person := NewPerson()
		err := person.SetDateOfBirth("1981-01-01", "")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingBirthDate)
	})

	t.Run("missing start date", func(t *testing.T) {
		person := NewPerson()
		err := person.SetDateOfBirth("", "1981-01-01")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingBirthDate)
	})

	t.Run("invalid start date", func(t *testing.T) {
		person := NewPerson()
		err := person.SetDateOfBirth("19810101", "1987-01-31")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidStartOfBirthDate)
	})

	t.Run("invalid end date", func(t *testing.T) {
		person := NewPerson()
		err := person.SetDateOfBirth("1987-01-01", "19870131")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidEndOfBirthDate)
	})

	t.Run("valid dates", func(t *testing.T) {
		person := NewPerson()
		err := person.SetDateOfBirth("1987-01-01", "1987-01-31")
		require.NoError(t, err)
		require.Equal(t, "1987-01-01", person.DateOfBirth.DateRange.Start)
		require.Equal(t, "1987-01-31", person.DateOfBirth.DateRange.End)
	})
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

	t.Run("invalid language code", func(t *testing.T) {
		person := NewPerson()
		err := person.AddLanguage("wrong", DefaultCountry)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidLanguageCode)
	})

	t.Run("invalid region code", func(t *testing.T) {
		person := NewPerson()
		err := person.AddLanguage(DefaultLanguage, "wrong")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidRegionCode)
	})

	t.Run("valid language", func(t *testing.T) {
		person := NewPerson()
		err := person.AddLanguage(DefaultLanguage, DefaultCountry)
		require.NoError(t, err)
		require.NotEmpty(t, person.Languages)
		require.Equal(t, DefaultLanguage, person.Languages[0].Language)
		require.Equal(t, DefaultCountry, person.Languages[0].Region)
		require.Equal(t, DefaultLanguage+"_"+DefaultCountry, person.Languages[0].Display)
	})
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

	t.Run("missing ethnicity", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEthnicity("")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidEthnicity)
	})

	t.Run("invalid ethnicity", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEthnicity("unknown")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrEthnicityNotAccepted)
	})

	t.Run("valid ethnicity", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEthnicity(testEthnicity)
		require.NoError(t, err)
		require.NotEmpty(t, person.Ethnicities)
		require.Equal(t, testEthnicity, person.Ethnicities[0].Content)
	})
}

// ExamplePerson_AddEthnicity example using AddEthnicity()
func ExamplePerson_AddEthnicity() {
	person := NewPerson()
	_ = person.AddEthnicity(testEthnicity)
	fmt.Println(person.Ethnicities[0].Content)
	// Output:white
}

// BenchmarkAddEthnicity benchmarks the AddEthnicity method
func BenchmarkAddEthnicity(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddEthnicity(testEthnicity)
	}
}

// TestAddOriginCountry test adding an origin country to a person object
func TestAddOriginCountry(t *testing.T) {
	t.Parallel()

	t.Run("invalid country code", func(t *testing.T) {
		person := NewPerson()
		err := person.AddOriginCountry("")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidCountryCode)
	})

	t.Run("valid country", func(t *testing.T) {
		person := NewPerson()
		err := person.AddOriginCountry(DefaultCountry)
		require.NoError(t, err)
		require.NotEmpty(t, person.OriginCountries)
		require.Equal(t, DefaultCountry, person.OriginCountries[0].Country)
	})
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

	t.Run("missing number and street", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddress("", "", testApartment, testCity, testState, DefaultCountry, testPOBox)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingNumberOrStreet)
	})

	t.Run("missing number", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddress("", testStreet, testApartment, testCity, testState, DefaultCountry, testPOBox)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingNumberOrStreet)
	})

	t.Run("missing street", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddress(testHouseNumber, "", testApartment, testCity, testState, DefaultCountry, testPOBox)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingNumberOrStreet)
	})

	t.Run("missing city and state", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddress(testHouseNumber, testStreet, testApartment, "", "", DefaultCountry, testPOBox)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingCityOrState)
	})

	t.Run("valid address", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddress(testHouseNumber, testStreet, testApartment, testCity, testState, DefaultCountry, testPOBox)
		require.NoError(t, err)
		require.NotEmpty(t, person.Addresses)
		require.Equal(t, testHouseNumber, person.Addresses[0].House)
		require.Equal(t, testStreet, person.Addresses[0].Street)
		require.Equal(t, testApartment, person.Addresses[0].Apartment)
		require.Equal(t, testCity, person.Addresses[0].City)
		require.Equal(t, testState, person.Addresses[0].State)
		require.Equal(t, DefaultCountry, person.Addresses[0].Country)
		require.Equal(t, testPOBox, person.Addresses[0].POBox)
	})
}

// ExamplePerson_AddAddress example using AddAddress()
func ExamplePerson_AddAddress() {
	person := NewPerson()
	_ = person.AddAddress(testHouseNumber, testStreet, testApartment, testCity, testState, DefaultCountry, testPOBox)
	fmt.Println(person.Addresses[0].House + " " + person.Addresses[0].Street + ", " + person.Addresses[0].City + " " + person.Addresses[0].State)
	// Output:10 Hickory Lane, Smallville KS
}

// BenchmarkAddAddress benchmarks the AddAddress method
func BenchmarkAddAddress(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddAddress(testHouseNumber, testStreet, testApartment, testCity, testState, DefaultCountry, testPOBox)
	}
}

// TestAddAddressRaw test adding an address to a person object
func TestAddAddressRaw(t *testing.T) {
	t.Parallel()

	t.Run("too short", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddressRaw(testHouseNumber)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrAddressTooShort)
	})

	t.Run("valid address", func(t *testing.T) {
		person := NewPerson()
		err := person.AddAddressRaw(testHouseNumber + " " + testStreet + ", Kansas, " + DefaultCountry)
		require.NoError(t, err)
		require.NotEmpty(t, person.Addresses)
		require.Equal(t, "10 Hickory Lane, Kansas, "+DefaultCountry, person.Addresses[0].Raw)
	})
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

	t.Run("missing title", func(t *testing.T) {
		person := NewPerson()
		err := person.AddJob("", "daily post", "news", "2010-01-01", "2011-01-01")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingTitleOrOrganization)
	})

	t.Run("missing organization", func(t *testing.T) {
		person := NewPerson()
		err := person.AddJob("reporter", "", "news", "2010-01-01", "2011-01-01")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingTitleOrOrganization)
	})

	t.Run("valid job", func(t *testing.T) {
		person := NewPerson()
		err := person.AddJob("reporter", "daily post", "news", "2010-01-01", "2011-01-01")
		require.NoError(t, err)
		require.NotEmpty(t, person.Jobs)
		require.Equal(t, "reporter", person.Jobs[0].Title)
		require.Equal(t, "daily post", person.Jobs[0].Organization)
		require.Equal(t, "news", person.Jobs[0].Industry)
		require.Equal(t, "2010-01-01", person.Jobs[0].DateRange.Start)
		require.Equal(t, "2011-01-01", person.Jobs[0].DateRange.End)
	})
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

	t.Run("missing school", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEducation("masters", "", "2010-01-01", "2011-01-01")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrMissingSchool)
	})

	t.Run("valid education", func(t *testing.T) {
		person := NewPerson()
		err := person.AddEducation("masters", "fau", "2010-01-01", "2011-01-01")
		require.NoError(t, err)
		require.NotEmpty(t, person.Educations)
		require.Equal(t, "masters", person.Educations[0].Degree)
		require.Equal(t, "fau", person.Educations[0].School)
		require.Equal(t, "2010-01-01", person.Educations[0].DateRange.Start)
		require.Equal(t, "2011-01-01", person.Educations[0].DateRange.End)
	})
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

// TestAddRelationship test adding a relationship to a person object
func TestAddRelationship(t *testing.T) {
	t.Parallel()

	t.Run("valid relationship", func(t *testing.T) {
		person := NewPerson()
		err := person.AddRelationship(Relationship{
			DateOfBirth: DateOfBirth{},
			Gender:      Gender{},
			Type:        "friend",
			Current:     false,
			Inferred:    false,
		})
		require.NoError(t, err)
		assert.Equal(t, "friend", person.Relationships[0].Type)
	})
}

// ExamplePerson_AddRelationship example using AddRelationship()
func ExamplePerson_AddRelationship() {
	person := NewPerson()
	_ = person.AddRelationship(Relationship{Type: "friend"})
	fmt.Println(person.Relationships[0].Type)
	// Output:friend
}

// BenchmarkAddRelationship benchmarks the AddRelationship method
func BenchmarkAddRelationship(b *testing.B) {
	person := NewPerson()
	relationship := Relationship{Type: "friend"}
	for i := 0; i < b.N; i++ {
		_ = person.AddRelationship(relationship)
	}
}

// TestAddURL test adding an url to a person object
func TestAddURL(t *testing.T) {
	t.Parallel()

	t.Run("too short", func(t *testing.T) {
		person := NewPerson()
		err := person.AddURL("http")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrURLTooShort)
	})

	t.Run("valid url", func(t *testing.T) {
		person := NewPerson()
		err := person.AddURL(testURL)
		require.NoError(t, err)
		require.NotEmpty(t, person.URLs)
		require.Equal(t, testURL, person.URLs[0].URL)
	})
}

// ExamplePerson_AddURL example using AddURL()
func ExamplePerson_AddURL() {
	person := NewPerson()
	_ = person.AddURL(testURL)
	fmt.Println(person.URLs[0].URL)
	// Output:https://twitter.com/clarkkent
}

// BenchmarkAddURL benchmarks the AddURL method
func BenchmarkAddURL(b *testing.B) {
	person := NewPerson()
	for i := 0; i < b.N; i++ {
		_ = person.AddURL(testURL)
	}
}

// TestPerson_ProcessThumbnails test processing images for thumbnails
func TestPerson_ProcessThumbnails(t *testing.T) {
	t.Parallel()

	t.Run("nil settings", func(t *testing.T) {
		// Create person and image
		person := NewPerson()
		require.NotNil(t, person)

		image := new(Image)
		image.URL = testImage
		image.ThumbnailToken = testThumbnailToken
		person.Images = append(person.Images, *image)

		person.ProcessThumbnails(nil)
		require.Empty(t, person.Images[0].ThumbnailURL)
	})

	t.Run("valid settings", func(t *testing.T) {
		// Create person and image
		person := NewPerson()
		image := new(Image)
		image.URL = testImage
		image.ThumbnailToken = testThumbnailToken
		person.Images = append(person.Images, *image)

		// Create settings
		settings := &ThumbnailSettings{
			URL:      thumbnailEndpoint,
			Height:   ThumbnailHeight,
			Width:    ThumbnailWidth,
			Enabled:  true,
			Favicon:  true,
			ZoomFace: true,
		}

		// Process using defaults
		person.ProcessThumbnails(settings)

		// Test for url
		require.NotEqual(t, 0, person.Images[0].ThumbnailURL)
		require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("width=%d", settings.Width))
		require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("height=%d", settings.Height))
		require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("favicon=%t", settings.Favicon))
		require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("zoom_face=%t", settings.ZoomFace))
		require.Contains(t, person.Images[0].ThumbnailURL, fmt.Sprintf("tokens=%s", person.Images[0].ThumbnailToken))
	})
}

// ExamplePerson_ProcessThumbnails example using ProcessThumbnails()
func ExamplePerson_ProcessThumbnails() {
	person := NewPerson()
	image := new(Image)
	image.URL = testImage
	image.ThumbnailToken = testThumbnailToken
	person.Images = append(person.Images, *image)

	person.ProcessThumbnails(&ThumbnailSettings{
		URL:      thumbnailEndpoint,
		Height:   ThumbnailHeight,
		Width:    ThumbnailWidth,
		Enabled:  true,
		Favicon:  true,
		ZoomFace: true,
	})
	fmt.Println(person.Images[0].ThumbnailURL)
	// Output: https://thumb.pipl.com/image?height=250&width=250&favicon=true&zoom_face=true&tokens=AE2861B242686E7BD0CB4D9049298EB7D18FEF66D950E8AB78BCD3F484345CE74536C19A85D0BA3D32DC9E7D1878CD4D341254E7AD129255C6983E6E154C4530A0DAAF665EA325FC0206F8B1D7E0B6B7AD9EBF71FCF610D57D
}

// BenchmarkProcessThumbnails benchmarks the ProcessThumbnails method
func BenchmarkProcessThumbnails(b *testing.B) {
	person := NewPerson()

	image := new(Image)
	image.URL = testImage
	image.ThumbnailToken = testThumbnailToken
	person.Images = append(person.Images, *image)

	settings := &ThumbnailSettings{
		URL:      thumbnailEndpoint,
		Height:   ThumbnailHeight,
		Width:    ThumbnailWidth,
		Enabled:  true,
		Favicon:  true,
		ZoomFace: true,
	}

	for i := 0; i < b.N; i++ {
		person.ProcessThumbnails(settings)
	}
}

// TestSearchMeetsMinimumCriteria test the minimum criteria for a search
//
//	This also tests: HasEmail, HasPhone, HasUserID, HasUsername, HasURL
//	HasName, HasAddress
func TestSearchMeetsMinimumCriteria(t *testing.T) {
	t.Parallel()

	t.Run("missing data", func(t *testing.T) {
		person := new(Person)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("raw name", func(t *testing.T) {
		person := new(Person)
		err := person.AddNameRaw(testFirstName + " " + testLastName)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("missing last name", func(t *testing.T) {
		person := new(Person)
		err := person.AddName(testFirstName, "", "", "", "")
		require.NoError(t, err)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("missing first name", func(t *testing.T) {
		person := new(Person)
		err := person.AddName("", "", testLastName, "", "")
		require.NoError(t, err)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("first and last name", func(t *testing.T) {
		person := new(Person)
		err := person.AddName(testFirstName, "", testLastName, "", "")
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("email address", func(t *testing.T) {
		person := new(Person)
		err := person.AddEmail(testEmailSecondary)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid phone number", func(t *testing.T) {
		person := new(Person)
		err := person.AddPhone(testPhone, testPhoneCountryCode)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("missing phone code", func(t *testing.T) {
		person := new(Person)
		err := person.AddPhone(testPhone, 0)
		require.Error(t, err)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid raw phone number", func(t *testing.T) {
		person := new(Person)
		err := person.AddPhoneRaw(testPhoneRaw)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid user id", func(t *testing.T) {
		person := new(Person)
		err := person.AddUserID(testUserName+"123", testUserNameServiceProvider)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid username", func(t *testing.T) {
		person := new(Person)
		err := person.AddUsername(testUserName, testUserNameServiceProvider)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid url", func(t *testing.T) {
		person := new(Person)
		err := person.AddURL(testURL)
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("partial address number", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress(testHouseNumber, "", "", "", "", "", "")
		require.Error(t, err)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("partial address street", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress(testHouseNumber, testStreet, "", "", "", "", "")
		require.Error(t, err)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("partial address city", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress(testHouseNumber, testStreet, "", testCity, "", "", "")
		require.NoError(t, err)
		require.False(t, SearchMeetsMinimumCriteria(person))
	})

	t.Run("full address", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress(testHouseNumber, testStreet, "", testCity, testState, "", "")
		require.NoError(t, err)
		require.True(t, SearchMeetsMinimumCriteria(person))
	})
}

// ExampleSearchMeetsMinimumCriteria example using SearchMeetsMinimumCriteria()
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
