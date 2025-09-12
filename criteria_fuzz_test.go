package pipl

import (
	"testing"
)

// FuzzSearchMeetsMinimumCriteria tests the SearchMeetsMinimumCriteria function with various person configurations
//
//nolint:gocognit // Comprehensive fuzz testing requires complex validation logic
func FuzzSearchMeetsMinimumCriteria(f *testing.F) {
	// Seed corpus with various person configurations
	f.Add("john@example.com", "", "", "", "", "", "", "", "", "", "", "", "", "")
	f.Add("", "5551234567", "", "", "", "", "", "", "", "", "", "", "", "")
	f.Add("", "", "johndoe", "twitter", "", "", "", "", "", "", "", "", "", "")
	f.Add("", "", "", "", "user123", "facebook", "", "", "", "", "", "", "", "")
	f.Add("", "", "", "", "", "", "https://example.com", "", "", "", "", "", "", "")
	f.Add("", "", "", "", "", "", "", "John", "Doe", "", "", "", "", "")
	f.Add("", "", "", "", "", "", "", "", "", "123 Main St", "Springfield", "IL", "US", "")
	f.Add("", "", "", "", "", "", "", "", "", "", "", "", "", "John Doe Raw")

	f.Fuzz(func(t *testing.T, email, phone, username, usernameProvider, userID, userIDProvider, url, firstName, lastName, house, city, state, country, rawName string) {
		person := NewPerson()

		// Add email if provided
		if email != "" {
			_ = person.AddEmail(email)
		}

		// Add phone if provided
		if phone != "" {
			_ = person.AddPhoneRaw(phone)
		}

		// Add username if both username and provider provided
		if username != "" && usernameProvider != "" {
			_ = person.AddUsername(username, usernameProvider)
		}

		// Add userID if both userID and provider provided
		if userID != "" && userIDProvider != "" {
			_ = person.AddUserID(userID, userIDProvider)
		}

		// Add URL if provided
		if url != "" {
			_ = person.AddURL(url)
		}

		// Add structured name if both first and last name provided
		if firstName != "" && lastName != "" {
			_ = person.AddName(firstName, "", lastName, "", "")
		}

		// Add raw name if provided
		if rawName != "" {
			_ = person.AddNameRaw(rawName)
		}

		// Add address if house, city, and state provided
		if house != "" && city != "" && state != "" {
			_ = person.AddAddress(house, "Main St", "", city, state, country, "")
		}

		result := SearchMeetsMinimumCriteria(person)

		// Manually check if criteria should be met
		expectedResult := person.HasEmail() || person.HasPhone() || person.HasUserID() ||
			person.HasUsername() || person.HasURL() || person.HasName() ||
			person.HasAddress()

		if result != expectedResult {
			t.Errorf("SearchMeetsMinimumCriteria returned %v, expected %v for person with: "+
				"hasEmail=%v, hasPhone=%v, hasUserID=%v, hasUsername=%v, hasURL=%v, hasName=%v, hasAddress=%v",
				result, expectedResult,
				person.HasEmail(), person.HasPhone(), person.HasUserID(),
				person.HasUsername(), person.HasURL(), person.HasName(), person.HasAddress())
		}

		// Test edge case: empty person should not meet criteria
		if !person.HasEmail() && !person.HasPhone() && !person.HasUserID() &&
			!person.HasUsername() && !person.HasURL() && !person.HasName() && !person.HasAddress() {
			if result {
				t.Errorf("Empty person should not meet minimum criteria")
			}
		}
	})
}

// FuzzIsAcceptedValue tests the isAcceptedValue helper function with various inputs
//
//nolint:gocognit // Comprehensive fuzz testing requires complex validation logic
func FuzzIsAcceptedValue(f *testing.F) {
	// Seed corpus with known values and edge cases
	f.Add("twitter", true)  // Valid service provider
	f.Add("facebook", true) // Valid service provider
	f.Add("invalid", false) // Invalid service provider
	f.Add("", false)        // Empty string
	f.Add("TWITTER", false) // Case sensitive test
	f.Add("white", true)    // Valid ethnicity (testing with ethnicity list)
	f.Add("black", true)    // Valid ethnicity
	f.Add("purple", false)  // Invalid ethnicity

	f.Fuzz(func(t *testing.T, testValue string, useServiceProviders bool) {
		var allowedValues *[]string

		if useServiceProviders {
			allowedValues = &AllowedServiceProviders
		} else {
			allowedValues = &AllowedEthnicities
		}

		result := isAcceptedValue(testValue, allowedValues)

		// Manually check if value should be accepted
		expectedResult := false
		for _, value := range *allowedValues {
			if testValue == value {
				expectedResult = true
				break
			}
		}

		if result != expectedResult {
			t.Errorf("isAcceptedValue returned %v for value %q, expected %v", result, testValue, expectedResult)
		}

		// Test that empty string is never accepted (unless it's actually in the allowed list)
		if testValue == "" && result {
			found := false
			for _, value := range *allowedValues {
				if value == "" {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Empty string should not be accepted unless explicitly in allowed values")
			}
		}
	})
}

// FuzzPersonHasMethods tests all the Has* methods on Person struct
//
//nolint:gocognit // Comprehensive fuzz testing requires complex validation logic
func FuzzPersonHasMethods(f *testing.F) {
	f.Add("test@example.com", "5551234567", "123", "Main St", "Springfield", "IL", "John", "Doe", "https://example.com")

	f.Fuzz(func(t *testing.T, email, phone, house, street, city, state, firstName, lastName, url string) {
		person := NewPerson()

		// Test HasEmail
		var emailAdded bool
		if email != "" {
			err := person.AddEmail(email)
			emailAdded = err == nil
		}
		hasEmailResult := person.HasEmail()
		expectedHasEmail := emailAdded

		if hasEmailResult != expectedHasEmail {
			t.Errorf("HasEmail returned %v, expected %v", hasEmailResult, expectedHasEmail)
		}

		// Test HasPhone
		var phoneAdded bool
		if phone != "" {
			err := person.AddPhoneRaw(phone)
			phoneAdded = err == nil
		}
		hasPhoneResult := person.HasPhone()
		expectedHasPhone := phoneAdded

		if hasPhoneResult != expectedHasPhone {
			t.Errorf("HasPhone returned %v, expected %v", hasPhoneResult, expectedHasPhone)
		}

		// Test HasAddress
		var addressAdded bool
		if house != "" && street != "" && city != "" && state != "" {
			err := person.AddAddress(house, street, "", city, state, "US", "")
			addressAdded = err == nil
		}
		hasAddressResult := person.HasAddress()
		expectedHasAddress := addressAdded

		if hasAddressResult != expectedHasAddress {
			t.Errorf("HasAddress returned %v, expected %v", hasAddressResult, expectedHasAddress)
		}

		// Test HasName
		var nameAdded bool
		if firstName != "" && lastName != "" {
			err := person.AddName(firstName, "", lastName, "", "")
			nameAdded = err == nil
		}
		hasNameResult := person.HasName()
		expectedHasName := nameAdded

		if hasNameResult != expectedHasName {
			t.Errorf("HasName returned %v, expected %v", hasNameResult, expectedHasName)
		}

		// Test HasURL
		var urlAdded bool
		if url != "" {
			err := person.AddURL(url)
			urlAdded = err == nil
		}
		hasURLResult := person.HasURL()
		expectedHasURL := urlAdded

		if hasURLResult != expectedHasURL {
			t.Errorf("HasURL returned %v, expected %v", hasURLResult, expectedHasURL)
		}
	})
}

// FuzzPersonHasUsernameAndUserID tests HasUsername and HasUserID methods
//
//nolint:gocognit // Comprehensive fuzz testing requires complex validation logic
func FuzzPersonHasUsernameAndUserID(f *testing.F) {
	f.Add("johndoe", "twitter", "123456", "facebook")
	f.Add("", "twitter", "123456", "facebook")
	f.Add("johndoe", "", "123456", "facebook")
	f.Add("johndoe", "twitter", "", "facebook")
	f.Add("johndoe", "twitter", "123456", "")

	f.Fuzz(func(t *testing.T, username, usernameProvider, userID, userIDProvider string) {
		person := NewPerson()

		// Test HasUsername
		if username != "" && usernameProvider != "" {
			// Only add if both are provided and provider is valid
			for _, allowed := range AllowedServiceProviders {
				if usernameProvider == allowed {
					_ = person.AddUsername(username, usernameProvider)
					break
				}
			}
		}
		hasUsernameResult := person.HasUsername()

		// Test HasUserID
		if userID != "" && userIDProvider != "" {
			// Only add if both are provided and provider is valid
			for _, allowed := range AllowedServiceProviders {
				if userIDProvider == allowed {
					_ = person.AddUserID(userID, userIDProvider)
					break
				}
			}
		}
		hasUserIDResult := person.HasUserID()

		// Verify that Has* methods return true only if the data was successfully added
		if len(person.Usernames) > 0 && !hasUsernameResult {
			t.Errorf("HasUsername returned false but usernames exist")
		}
		if len(person.Usernames) == 0 && hasUsernameResult {
			t.Errorf("HasUsername returned true but no usernames exist")
		}

		if len(person.UserIDs) > 0 && !hasUserIDResult {
			t.Errorf("HasUserID returned false but user IDs exist")
		}
		if len(person.UserIDs) == 0 && hasUserIDResult {
			t.Errorf("HasUserID returned true but no user IDs exist")
		}
	})
}

// FuzzPersonWithMultipleEntries tests person objects with multiple entries of the same type
func FuzzPersonWithMultipleEntries(f *testing.F) {
	f.Add("john@example.com", "jane@example.org", "+1-555-123-4567", "555-987-6543")

	f.Fuzz(func(t *testing.T, email1, email2, phone1, phone2 string) {
		person := NewPerson()

		// Add multiple emails if they're different and valid
		if email1 != "" {
			_ = person.AddEmail(email1)
		}
		if email2 != "" && email2 != email1 {
			_ = person.AddEmail(email2)
		}

		// Add multiple phones if they're different and valid
		if phone1 != "" {
			_ = person.AddPhoneRaw(phone1)
		}
		if phone2 != "" && phone2 != phone1 {
			_ = person.AddPhoneRaw(phone2)
		}

		// Test that Has* methods work correctly with multiple entries
		hasEmailResult := person.HasEmail()
		expectedHasEmail := email1 != "" || email2 != ""

		if hasEmailResult != expectedHasEmail {
			t.Errorf("HasEmail with multiple entries returned %v, expected %v", hasEmailResult, expectedHasEmail)
		}

		hasPhoneResult := person.HasPhone()
		expectedHasPhone := phone1 != "" || phone2 != ""

		if hasPhoneResult != expectedHasPhone {
			t.Errorf("HasPhone with multiple entries returned %v, expected %v", hasPhoneResult, expectedHasPhone)
		}

		// Test SearchMeetsMinimumCriteria with multiple entries
		criteriaResult := SearchMeetsMinimumCriteria(person)
		expectedCriteria := hasEmailResult || hasPhoneResult

		if criteriaResult != expectedCriteria {
			t.Errorf("SearchMeetsMinimumCriteria with multiple entries returned %v, expected %v", criteriaResult, expectedCriteria)
		}
	})
}

// FuzzPersonWithEmptyFields tests person objects with empty or whitespace-only fields
func FuzzPersonWithEmptyFields(f *testing.F) {
	f.Add("   ", "\t", "\n", "  \t\n  ")

	f.Fuzz(func(t *testing.T, email, phone, name, address string) {
		person := NewPerson()

		// Attempt to add fields with only whitespace
		_ = person.AddEmail(email)
		_ = person.AddPhoneRaw(phone)
		if len(name) > 5 { // AddNameRaw has minimum length requirement
			_ = person.AddNameRaw(name)
		}
		if len(address) > 5 { // AddAddressRaw has minimum length requirement
			_ = person.AddAddressRaw(address)
		}

		// Test that empty/whitespace fields are handled appropriately
		hasEmailResult := person.HasEmail()
		hasPhoneResult := person.HasPhone()
		hasNameResult := person.HasName()
		hasAddressResult := person.HasAddress()

		// SearchMeetsMinimumCriteria should handle edge cases gracefully
		criteriaResult := SearchMeetsMinimumCriteria(person)
		expectedCriteria := hasEmailResult || hasPhoneResult || hasNameResult || hasAddressResult

		if criteriaResult != expectedCriteria {
			t.Errorf("SearchMeetsMinimumCriteria with whitespace fields returned %v, expected %v", criteriaResult, expectedCriteria)
		}
	})
}
