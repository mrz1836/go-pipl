package pipl

import (
	"strings"
	"testing"
)

// FuzzAddName tests the AddName function with various input combinations
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddName(f *testing.F) {
	// Seed corpus with valid and edge case inputs
	f.Add("John", "William", "Doe", "Mr.", "Jr.")
	f.Add("", "", "Smith", "", "")
	f.Add("Jane", "", "", "", "")
	f.Add("", "", "", "", "")
	f.Add("A", "B", "C", "Dr.", "III")
	f.Add("ÀáÂâÃãÄä", "ÒóÔôÕõÖö", "ÙúÛûÜü", "Señor", "Ñ")
	f.Add("José-María", "de la", "Cruz-Santos", "Don", "y Más")

	f.Fuzz(func(t *testing.T, firstName, middleName, lastName, prefix, suffix string) {
		person := NewPerson()
		err := person.AddName(firstName, middleName, lastName, prefix, suffix)

		// Should only error if both first and last names are empty
		if len(firstName) == 0 && len(lastName) == 0 {
			if err == nil {
				t.Errorf("Expected error when both first and last names are empty")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error with valid name inputs: %v", err)
			} else {
				// Verify the name was added correctly
				if len(person.Names) != 1 {
					t.Errorf("Expected 1 name, got %d", len(person.Names))
				} else {
					name := person.Names[0]
					if name.First != firstName || name.Middle != middleName ||
						name.Last != lastName || name.Prefix != prefix || name.Suffix != suffix {
						t.Errorf("Name fields not set correctly")
					}
				}
			}
		}
	})
}

// FuzzAddNameRaw tests the AddNameRaw function with various raw name inputs
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddNameRaw(f *testing.F) {
	f.Add("John William Doe Jr.")
	f.Add("Jane")
	f.Add("Short")
	f.Add("")
	f.Add("María José de la Cruz-Santos")
	f.Add("Dr. Richard P. Feynman III")
	f.Add("王小明") //nolint:gosmopolitan // Testing Unicode names
	f.Add("محمد عبدالله")

	f.Fuzz(func(t *testing.T, fullName string) {
		person := NewPerson()
		err := person.AddNameRaw(fullName)

		if len(fullName) <= 5 {
			if err == nil {
				t.Errorf("Expected error for short name: %q", fullName)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for valid name: %q, error: %v", fullName, err)
			} else {
				if len(person.Names) != 1 {
					t.Errorf("Expected 1 name, got %d", len(person.Names))
				} else if person.Names[0].Raw != fullName {
					t.Errorf("Raw name not set correctly: expected %q, got %q", fullName, person.Names[0].Raw)
				}
			}
		}
	})
}

// FuzzAddEmail tests the AddEmail function with various email formats
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddEmail(f *testing.F) {
	f.Add("user@example.com")
	f.Add("test@domain.org")
	f.Add("")
	f.Add("invalid-email")
	f.Add("@domain.com")
	f.Add("user@")
	f.Add("user@@domain.com")
	f.Add("üser@dömain.com")
	f.Add("user+tag@domain.co.uk")
	f.Add("very.long.email.address@very.long.domain.name.example.com")

	f.Fuzz(func(t *testing.T, emailAddress string) {
		person := NewPerson()
		err := person.AddEmail(emailAddress)

		hasAt := false
		for _, r := range emailAddress {
			if r == '@' {
				hasAt = true
				break
			}
		}

		if len(emailAddress) == 0 || !hasAt {
			if err == nil {
				t.Errorf("Expected error for invalid email: %q", emailAddress)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for potentially valid email: %q, error: %v", emailAddress, err)
			} else {
				if len(person.Emails) != 1 {
					t.Errorf("Expected 1 email, got %d", len(person.Emails))
				} else if person.Emails[0].Address != emailAddress {
					t.Errorf("Email address not set correctly")
				}
			}
		}
	})
}

// FuzzAddUsername tests the AddUsername function with various username/provider combinations
//
//nolint:gocognit,nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddUsername(f *testing.F) {
	f.Add("johndoe", "twitter")
	f.Add("user123", "facebook")
	f.Add("", "instagram")
	f.Add("user", "")
	f.Add("ab", "twitter")
	f.Add("validuser", "invalidprovider")
	f.Add("ユーザー名", "twitter") //nolint:gosmopolitan // Testing Unicode usernames
	f.Add("user.name", "linkedin")

	f.Fuzz(func(t *testing.T, username, serviceProvider string) {
		person := NewPerson()
		err := person.AddUsername(username, serviceProvider)

		shouldError := len(username) <= 3 || len(serviceProvider) <= 2

		// Check if provider is in allowed list
		if !shouldError {
			providerAllowed := false
			for _, allowed := range AllowedServiceProviders {
				if allowed == strings.ToLower(serviceProvider) {
					providerAllowed = true
					break
				}
			}
			shouldError = !providerAllowed
		}

		if shouldError {
			if err == nil {
				t.Errorf("Expected error for username: %q, provider: %q", username, serviceProvider)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for username: %q, provider: %q, error: %v", username, serviceProvider, err)
			} else {
				if len(person.Usernames) != 1 {
					t.Errorf("Expected 1 username, got %d", len(person.Usernames))
				} else {
					expected := username + "@" + serviceProvider
					if person.Usernames[0].Content != expected {
						t.Errorf("Username content not set correctly: expected %q, got %q", expected, person.Usernames[0].Content)
					}
				}
			}
		}
	})
}

// FuzzAddUserID tests the AddUserID function with various user ID/provider combinations
//
//nolint:gocognit,nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddUserID(f *testing.F) {
	f.Add("123456789", "twitter")
	f.Add("user123", "facebook")
	f.Add("", "instagram")
	f.Add("abc", "")
	f.Add("ab", "twitter")
	f.Add("validid", "invalidprovider")

	f.Fuzz(func(t *testing.T, userID, serviceProvider string) {
		person := NewPerson()
		err := person.AddUserID(userID, serviceProvider)

		shouldError := len(userID) <= 2 || len(serviceProvider) <= 2

		// Check if provider is in allowed list
		if !shouldError {
			providerAllowed := false
			for _, allowed := range AllowedServiceProviders {
				if allowed == strings.ToLower(serviceProvider) {
					providerAllowed = true
					break
				}
			}
			shouldError = !providerAllowed
		}

		if shouldError {
			if err == nil {
				t.Errorf("Expected error for userID: %q, provider: %q", userID, serviceProvider)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for userID: %q, provider: %q, error: %v", userID, serviceProvider, err)
			} else {
				if len(person.UserIDs) != 1 {
					t.Errorf("Expected 1 userID, got %d", len(person.UserIDs))
				} else {
					expected := userID + "@" + serviceProvider
					if person.UserIDs[0].Content != expected {
						t.Errorf("UserID content not set correctly: expected %q, got %q", expected, person.UserIDs[0].Content)
					}
				}
			}
		}
	})
}

// FuzzAddURL tests the AddURL function with various URL formats
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddURL(f *testing.F) {
	f.Add("https://example.com")
	f.Add("http://test.org")
	f.Add("")
	f.Add("short")
	f.Add("ftp://files.example.com/path")
	f.Add("mailto:user@example.com")
	f.Add("https://subdomain.example.com/path/to/resource?param=value#anchor")

	f.Fuzz(func(t *testing.T, url string) {
		person := NewPerson()
		err := person.AddURL(url)

		if len(url) <= 5 {
			if err == nil {
				t.Errorf("Expected error for short URL: %q", url)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for URL: %q, error: %v", url, err)
			} else {
				if len(person.URLs) != 1 {
					t.Errorf("Expected 1 URL, got %d", len(person.URLs))
				} else if person.URLs[0].URL != url {
					t.Errorf("URL not set correctly")
				}
			}
		}
	})
}

// FuzzAddPhone tests the AddPhone function with various phone number combinations
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddPhone(f *testing.F) {
	f.Add(int64(5551234567), 1)
	f.Add(int64(0), 1)
	f.Add(int64(123456789), 0)
	f.Add(int64(-1234567890), 1)
	f.Add(int64(9223372036854775807), 999)

	f.Fuzz(func(t *testing.T, phoneNumber int64, countryCode int) {
		person := NewPerson()
		err := person.AddPhone(phoneNumber, countryCode)

		if phoneNumber == 0 || countryCode == 0 {
			if err == nil {
				t.Errorf("Expected error for phone: %d, country: %d", phoneNumber, countryCode)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for phone: %d, country: %d, error: %v", phoneNumber, countryCode, err)
			} else {
				if len(person.Phones) != 1 {
					t.Errorf("Expected 1 phone, got %d", len(person.Phones))
				} else {
					phone := person.Phones[0]
					if phone.Number != phoneNumber || phone.CountryCode != countryCode {
						t.Errorf("Phone not set correctly")
					}
				}
			}
		}
	})
}

// FuzzAddPhoneRaw tests the AddPhoneRaw function with various phone string formats
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddPhoneRaw(f *testing.F) {
	f.Add("+1-555-123-4567")
	f.Add("555.123.4567")
	f.Add("123")
	f.Add("")
	f.Add("+44 20 7123 4567")
	f.Add("(555) 123-4567 ext. 1234")

	f.Fuzz(func(t *testing.T, phoneNumber string) {
		person := NewPerson()
		err := person.AddPhoneRaw(phoneNumber)

		if len(phoneNumber) < 4 {
			if err == nil {
				t.Errorf("Expected error for short phone: %q", phoneNumber)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for phone: %q, error: %v", phoneNumber, err)
			} else {
				if len(person.Phones) != 1 {
					t.Errorf("Expected 1 phone, got %d", len(person.Phones))
				} else if person.Phones[0].Raw != phoneNumber {
					t.Errorf("Phone raw not set correctly")
				}
			}
		}
	})
}

// FuzzAddAddress tests the AddAddress function with various address components
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddAddress(f *testing.F) {
	f.Add("123", "Main St", "Apt 1", "Springfield", "IL", "US", "")
	f.Add("", "Main St", "", "Springfield", "IL", "US", "")
	f.Add("123", "", "", "Springfield", "IL", "US", "")
	f.Add("123", "Main St", "", "", "", "US", "")
	f.Add("123", "Main St", "", "Springfield", "", "CA", "")

	f.Fuzz(func(t *testing.T, house, street, apartment, city, state, country, poBox string) {
		person := NewPerson()
		err := person.AddAddress(house, street, apartment, city, state, country, poBox)

		shouldError := (len(house) == 0 || len(street) == 0) || (len(city) == 0 && len(state) == 0)

		if shouldError {
			if err == nil {
				t.Errorf("Expected error for incomplete address")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for valid address, error: %v", err)
			} else {
				if len(person.Addresses) != 1 {
					t.Errorf("Expected 1 address, got %d", len(person.Addresses))
				} else {
					addr := person.Addresses[0]
					// Country should be forced to DefaultCountry
					if addr.Country != DefaultCountry {
						t.Errorf("Country not set to default: expected %q, got %q", DefaultCountry, addr.Country)
					}
				}
			}
		}
	})
}

// FuzzAddAddressRaw tests the AddAddressRaw function with various address strings
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddAddressRaw(f *testing.F) {
	f.Add("123 Main St, Springfield, IL 62701")
	f.Add("Short")
	f.Add("")
	f.Add("1600 Pennsylvania Avenue NW, Washington, DC 20500")

	f.Fuzz(func(t *testing.T, fullAddress string) {
		person := NewPerson()
		err := person.AddAddressRaw(fullAddress)

		if len(fullAddress) < 5 {
			if err == nil {
				t.Errorf("Expected error for short address: %q", fullAddress)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for address: %q, error: %v", fullAddress, err)
			} else {
				if len(person.Addresses) != 1 {
					t.Errorf("Expected 1 address, got %d", len(person.Addresses))
				} else if person.Addresses[0].Raw != fullAddress {
					t.Errorf("Address raw not set correctly")
				}
			}
		}
	})
}

// FuzzAddJob tests the AddJob function with various job information
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddJob(f *testing.F) {
	f.Add("Software Engineer", "Google", "Technology", "2020-01-01", "2023-12-31")
	f.Add("", "Company", "Industry", "", "")
	f.Add("Title", "", "Industry", "", "")
	f.Add("Manager", "Company", "", "2020-01-01", "")

	f.Fuzz(func(t *testing.T, title, organization, industry, dateStart, dateEnd string) {
		person := NewPerson()
		err := person.AddJob(title, organization, industry, dateStart, dateEnd)

		if len(title) == 0 || len(organization) == 0 {
			if err == nil {
				t.Errorf("Expected error for missing title or organization")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for valid job, error: %v", err)
			} else {
				if len(person.Jobs) != 1 {
					t.Errorf("Expected 1 job, got %d", len(person.Jobs))
				}
			}
		}
	})
}

// FuzzAddEducation tests the AddEducation function with various education information
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddEducation(f *testing.F) {
	f.Add("Bachelor of Science", "MIT", "2016-09-01", "2020-05-31")
	f.Add("", "Harvard", "2018-09-01", "2022-05-31")
	f.Add("PhD", "", "2020-09-01", "")

	f.Fuzz(func(t *testing.T, degree, school, dateStart, dateEnd string) {
		person := NewPerson()
		err := person.AddEducation(degree, school, dateStart, dateEnd)

		if len(school) == 0 {
			if err == nil {
				t.Errorf("Expected error for missing school")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for valid education, error: %v", err)
			} else {
				if len(person.Educations) != 1 {
					t.Errorf("Expected 1 education, got %d", len(person.Educations))
				}
			}
		}
	})
}

// FuzzAddLanguage tests the AddLanguage function with various language/region combinations
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddLanguage(f *testing.F) {
	f.Add("en", "US")
	f.Add("", "US")
	f.Add("en", "")
	f.Add("eng", "USA")
	f.Add("es", "MX")

	f.Fuzz(func(t *testing.T, languageCode, regionCode string) {
		person := NewPerson()
		err := person.AddLanguage(languageCode, regionCode)

		if len(languageCode) >= 3 || len(regionCode) >= 4 {
			if err == nil {
				t.Errorf("Expected error for invalid codes: lang=%q, region=%q", languageCode, regionCode)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for codes: lang=%q, region=%q, error: %v", languageCode, regionCode, err)
			} else {
				if len(person.Languages) != 1 {
					t.Errorf("Expected 1 language, got %d", len(person.Languages))
				}
			}
		}
	})
}

// FuzzAddEthnicity tests the AddEthnicity function with various ethnicity values
//
//nolint:gocognit,nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddEthnicity(f *testing.F) {
	f.Add("white")
	f.Add("black")
	f.Add("")
	f.Add("invalid_ethnicity")

	f.Fuzz(func(t *testing.T, ethnicity string) {
		person := NewPerson()
		err := person.AddEthnicity(ethnicity)

		if len(ethnicity) == 0 {
			if err == nil {
				t.Errorf("Expected error for empty ethnicity")
			}
		} else {
			// Check if ethnicity is in allowed list
			allowed := false
			for _, allowedEth := range AllowedEthnicities {
				if ethnicity == allowedEth {
					allowed = true
					break
				}
			}

			if !allowed {
				if err == nil {
					t.Errorf("Expected error for invalid ethnicity: %q", ethnicity)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for valid ethnicity: %q, error: %v", ethnicity, err)
				} else {
					if len(person.Ethnicities) != 1 {
						t.Errorf("Expected 1 ethnicity, got %d", len(person.Ethnicities))
					}
				}
			}
		}
	})
}

// FuzzSetGender tests the SetGender function with various gender values
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzSetGender(f *testing.F) {
	f.Add("male")
	f.Add("female")
	f.Add("")
	f.Add("other")
	f.Add("Male")
	f.Add("FEMALE")

	f.Fuzz(func(t *testing.T, gender string) {
		person := NewPerson()
		err := person.SetGender(gender)

		if gender != "male" && gender != "female" {
			if err == nil {
				t.Errorf("Expected error for invalid gender: %q", gender)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for valid gender: %q, error: %v", gender, err)
			} else {
				if person.Gender == nil {
					t.Errorf("Gender not set")
				} else if person.Gender.Content != gender {
					t.Errorf("Gender content not set correctly")
				}
			}
		}
	})
}

// FuzzSetDateOfBirth tests the SetDateOfBirth function with various date formats
//
//nolint:gocognit,nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzSetDateOfBirth(f *testing.F) {
	f.Add("1990-01-01", "1990-01-01")
	f.Add("", "1990-01-01")
	f.Add("1990-01-01", "")
	f.Add("invalid", "1990-01-01")
	f.Add("1990-01-01", "invalid")
	f.Add("90-1-1", "90-1-1")
	f.Add("1990-13-01", "1990-13-01")

	f.Fuzz(func(t *testing.T, startDate, endDate string) {
		person := NewPerson()
		err := person.SetDateOfBirth(startDate, endDate)

		shouldError := len(startDate) == 0 || len(endDate) == 0
		if !shouldError {
			// Check basic format validation
			if len(startDate) != 10 || len(endDate) != 10 {
				shouldError = true
			} else {
				hasStartDash := false
				hasEndDash := false
				for _, r := range startDate {
					if r == '-' {
						hasStartDash = true
						break
					}
				}
				for _, r := range endDate {
					if r == '-' {
						hasEndDash = true
						break
					}
				}
				shouldError = !hasStartDash || !hasEndDash
			}
		}

		if shouldError {
			if err == nil {
				t.Errorf("Expected error for invalid dates: start=%q, end=%q", startDate, endDate)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for dates: start=%q, end=%q, error: %v", startDate, endDate, err)
			} else {
				if person.DateOfBirth == nil {
					t.Errorf("Date of birth not set")
				}
			}
		}
	})
}

// FuzzAddOriginCountry tests the AddOriginCountry function with various country codes
//
//nolint:nestif // Comprehensive fuzz testing requires complex validation logic
func FuzzAddOriginCountry(f *testing.F) {
	f.Add("US")
	f.Add("GB")
	f.Add("")
	f.Add("USA")
	f.Add("XX")

	f.Fuzz(func(t *testing.T, countryCode string) {
		person := NewPerson()
		err := person.AddOriginCountry(countryCode)

		if len(countryCode) == 0 {
			if err == nil {
				t.Errorf("Expected error for empty country code")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for country code: %q, error: %v", countryCode, err)
			} else {
				if len(person.OriginCountries) != 1 {
					t.Errorf("Expected 1 origin country, got %d", len(person.OriginCountries))
				} else if person.OriginCountries[0].Country != countryCode {
					t.Errorf("Origin country not set correctly")
				}
			}
		}
	})
}
