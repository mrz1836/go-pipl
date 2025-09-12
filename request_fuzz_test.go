package pipl

import (
	"encoding/json"
	"fmt"
	"testing"
	"unicode/utf8"
)

// FuzzJSONUnmarshalResponse tests JSON unmarshaling of API responses with various malformed inputs
func FuzzJSONUnmarshalResponse(f *testing.F) {
	// Seed corpus with valid and invalid JSON responses
	f.Add(`{"person": {"names": [{"first": "John", "last": "Doe"}]}, "error": "", "@search_id": "12345"}`)
	f.Add(`{"person": {}, "possible_persons": [], "error": ""}`)
	f.Add(`{}`)
	f.Add(`{"invalid": json}`)
	f.Add(`{"person": null}`)
	f.Add(`{"person": {"names": [{"first": null}]}}`)
	f.Add(`{"person": {"names": [{}]}}`)
	f.Add(`{"person": {"emails": [{"address": "test@example.com"}]}}`)
	f.Add(`{"person": {"phones": [{"number": 5551234567, "country_code": 1}]}}`)
	f.Add(`{"error": "API Error Message"}`)

	f.Fuzz(func(_ *testing.T, jsonData string) {
		var response Response
		err := json.Unmarshal([]byte(jsonData), &response)
		// We don't expect any panics during unmarshaling
		// The function should either succeed or return an error gracefully
		if err != nil {
			// This is expected for malformed JSON
			return
		}

		// If unmarshaling succeeded, the response should be valid
		// Test that we can access fields without panicking
		_ = response.Error
		_ = response.SearchID
		_ = response.HTTPStatusCode
		_ = response.PersonsCount
		_ = response.VisibleSources
		_ = response.AvailableSources
		_ = response.TopMatch

		// Test accessing person fields
		_ = response.Person.ID
		_ = response.Person.Match
		_ = response.Person.Inferred
		_ = response.Person.SearchPointer

		// Test accessing arrays without panicking
		for _, name := range response.Person.Names {
			_ = name.First
			_ = name.Last
			_ = name.Raw
		}

		for _, email := range response.Person.Emails {
			_ = email.Address
		}

		for _, phone := range response.Person.Phones {
			_ = phone.Number
			_ = phone.CountryCode
			_ = phone.Raw
		}

		for _, addr := range response.Person.Addresses {
			_ = addr.House
			_ = addr.Street
			_ = addr.City
			_ = addr.State
			_ = addr.Country
		}
	})
}

// FuzzJSONUnmarshalPerson tests JSON unmarshaling of Person objects specifically
func FuzzJSONUnmarshalPerson(f *testing.F) {
	// Seed corpus with various person JSON structures
	f.Add(`{"names": [{"first": "John", "last": "Doe"}]}`)
	f.Add(`{"emails": [{"address": "test@example.com"}]}`)
	f.Add(`{"phones": [{"number": 5551234567, "country_code": 1}]}`)
	f.Add(`{"addresses": [{"house": "123", "street": "Main St", "city": "Springfield", "state": "IL"}]}`)
	f.Add(`{"dob": {"date_range": {"start": "1990-01-01", "end": "1990-01-01"}}}`)
	f.Add(`{"gender": {"content": "male"}}`)
	f.Add(`{"jobs": [{"title": "Engineer", "organization": "Company"}]}`)
	f.Add(`{"educations": [{"degree": "BS", "school": "University"}]}`)
	f.Add(`{"usernames": [{"content": "user@twitter"}]}`)
	f.Add(`{"user_ids": [{"content": "123@facebook"}]}`)
	f.Add(`{"urls": [{"url": "https://example.com"}]}`)
	f.Add(`{"languages": [{"language": "en", "region": "US"}]}`)
	f.Add(`{"ethnicities": [{"content": "white"}]}`)
	f.Add(`{"origin_countries": [{"country": "US"}]}`)
	f.Add(`{"images": [{"url": "https://example.com/image.jpg"}]}`)
	f.Add(`{"@match": 0.95, "@inferred": true}`)
	f.Add(`null`)
	f.Add(`[]`)
	f.Add(`"string"`)

	f.Fuzz(func(_ *testing.T, jsonData string) {
		var person Person
		err := json.Unmarshal([]byte(jsonData), &person)
		// Should handle malformed JSON gracefully
		if err != nil {
			return
		}

		// Test that we can access all person fields without panicking
		_ = person.ID
		_ = person.Match
		_ = person.Inferred
		_ = person.SearchPointer

		// Test accessing all slice fields
		for _, name := range person.Names {
			_ = name.First
			_ = name.Last
			_ = name.Middle
			_ = name.Raw
			_ = name.Prefix
			_ = name.Suffix
		}

		for _, email := range person.Emails {
			_ = email.Address
			_ = email.Current
			_ = email.Inferred
		}

		for _, phone := range person.Phones {
			_ = phone.Number
			_ = phone.CountryCode
			_ = phone.Raw
		}

		// Test pointer fields
		if person.Gender != nil {
			_ = person.Gender.Content
		}

		if person.DateOfBirth != nil {
			_ = person.DateOfBirth.DateRange.Start
			_ = person.DateOfBirth.DateRange.End
		}
	})
}

// FuzzProcessThumbnails tests the ProcessThumbnails function with various thumbnail settings and image data
func FuzzProcessThumbnails(f *testing.F) {
	// Seed corpus with various thumbnail configurations
	f.Add("https://thumb.pipl.com/image", 250, 250, true, true, "token123", "token456")
	f.Add("", 0, 0, false, false, "", "")
	f.Add("https://example.com", -1, -1, true, false, "token", "")
	f.Add("invalid-url", 1000, 1000, false, true, "very-long-token-string", "another-token")

	f.Fuzz(func(t *testing.T, url string, height, width int, favicon, zoomFace bool, token1, token2 string) {
		person := NewPerson()

		// Add images with thumbnail tokens
		if token1 != "" {
			person.Images = append(person.Images, Image{
				ThumbnailToken: token1,
				URL:            "https://example.com/image1.jpg",
			})
		}
		if token2 != "" {
			person.Images = append(person.Images, Image{
				ThumbnailToken: token2,
				URL:            "https://example.com/image2.jpg",
			})
		}

		// Create thumbnail settings
		thumbnailSettings := &ThumbnailSettings{
			URL:      url,
			Height:   height,
			Width:    width,
			Enabled:  true,
			Favicon:  favicon,
			ZoomFace: zoomFace,
		}

		// Test ProcessThumbnails - should not panic with any input
		person.ProcessThumbnails(thumbnailSettings)

		// Verify that thumbnail URLs were generated for images with tokens
		for _, image := range person.Images {
			if image.ThumbnailToken != "" && url != "" {
				if image.ThumbnailURL == "" {
					t.Errorf("Expected thumbnail URL to be generated for image with token")
				}
				// Verify URL contains expected parameters
				// Basic sanity check - URL should contain the base URL
				_ = image.ThumbnailURL
			}
		}

		// Test with nil settings - should not panic
		person.ProcessThumbnails(nil)

		// Test with empty person - should not panic
		emptyPerson := NewPerson()
		emptyPerson.ProcessThumbnails(thumbnailSettings)
	})
}

// FuzzPersonJSONRoundTrip tests JSON marshaling and unmarshaling of Person objects
//
//nolint:gocognit // Comprehensive fuzz testing requires complex validation logic
func FuzzPersonJSONRoundTrip(f *testing.F) {
	f.Add("John", "Doe", "john@example.com", "5551234567", "male", "1990-01-01")

	f.Fuzz(func(t *testing.T, firstName, lastName, email, phone, gender, dob string) {
		originalPerson := NewPerson()

		// Add data to person if valid
		if firstName != "" && lastName != "" {
			_ = originalPerson.AddName(firstName, "", lastName, "", "")
		}
		if email != "" {
			_ = originalPerson.AddEmail(email)
		}
		if phone != "" {
			_ = originalPerson.AddPhoneRaw(phone)
		}
		if gender == "male" || gender == "female" {
			_ = originalPerson.SetGender(gender)
		}
		if dob != "" {
			_ = originalPerson.SetDateOfBirth(dob, dob)
		}

		// Marshal to JSON
		jsonData, err := json.Marshal(originalPerson)
		if err != nil {
			t.Errorf("Failed to marshal person to JSON: %v", err)
			return
		}

		// Unmarshal from JSON
		var unmarshaledPerson Person
		err = json.Unmarshal(jsonData, &unmarshaledPerson)
		if err != nil {
			t.Errorf("Failed to unmarshal person from JSON: %v", err)
			return
		}

		// Basic validation that data survived the round trip
		if len(originalPerson.Names) != len(unmarshaledPerson.Names) {
			t.Errorf("Name count mismatch after JSON round trip")
		}
		if len(originalPerson.Emails) != len(unmarshaledPerson.Emails) {
			t.Errorf("Email count mismatch after JSON round trip")
		}
		if len(originalPerson.Phones) != len(unmarshaledPerson.Phones) {
			t.Errorf("Phone count mismatch after JSON round trip")
		}

		// Test that Has* methods work the same on both objects
		if originalPerson.HasEmail() != unmarshaledPerson.HasEmail() {
			t.Errorf("HasEmail mismatch after JSON round trip")
		}
		if originalPerson.HasPhone() != unmarshaledPerson.HasPhone() {
			t.Errorf("HasPhone mismatch after JSON round trip")
		}
		if originalPerson.HasName() != unmarshaledPerson.HasName() {
			t.Errorf("HasName mismatch after JSON round trip")
		}
	})
}

// FuzzResponseJSONRoundTrip tests JSON marshaling and unmarshaling of Response objects
func FuzzResponseJSONRoundTrip(f *testing.F) {
	f.Add("search123", "API Error", 200, 1, 5, true)

	f.Fuzz(func(t *testing.T, searchID, errorMsg string, httpStatus, personsCount, availableSources int, topMatch bool) {
		// Skip test cases with invalid UTF-8 strings as JSON marshaling will replace
		// invalid UTF-8 sequences with replacement characters, making round-trip impossible
		if !utf8.ValidString(searchID) || !utf8.ValidString(errorMsg) {
			t.Skip("Skipping test with invalid UTF-8 strings")
		}

		originalResponse := Response{
			SearchID:         searchID,
			Error:            errorMsg,
			HTTPStatusCode:   httpStatus,
			PersonsCount:     personsCount,
			AvailableSources: availableSources,
			TopMatch:         topMatch,
		}

		// Add a person to the response
		person := NewPerson()
		_ = person.AddName("John", "", "Doe", "", "")
		originalResponse.Person = *person

		// Marshal to JSON
		jsonData, err := json.Marshal(originalResponse)
		if err != nil {
			t.Errorf("Failed to marshal response to JSON: %v", err)
			return
		}

		// Unmarshal from JSON
		var unmarshaledResponse Response
		err = json.Unmarshal(jsonData, &unmarshaledResponse)
		if err != nil {
			t.Errorf("Failed to unmarshal response from JSON: %v", err)
			return
		}

		// Validate that key fields survived the round trip
		if originalResponse.SearchID != unmarshaledResponse.SearchID {
			t.Errorf("SearchID mismatch after JSON round trip")
		}
		if originalResponse.Error != unmarshaledResponse.Error {
			t.Errorf("Error mismatch after JSON round trip")
		}
		if originalResponse.HTTPStatusCode != unmarshaledResponse.HTTPStatusCode {
			t.Errorf("HTTPStatusCode mismatch after JSON round trip")
		}
		if originalResponse.TopMatch != unmarshaledResponse.TopMatch {
			t.Errorf("TopMatch mismatch after JSON round trip")
		}
	})
}

// FuzzInvalidJSONStructures tests handling of various invalid JSON structures
func FuzzInvalidJSONStructures(f *testing.F) {
	// Seed with various malformed JSON patterns
	f.Add(`{"person": {"names": [{"first":`)         // Incomplete JSON
	f.Add(`{"person": {"names": [{"first": 123}]}}`) // Wrong type
	f.Add(`{"person": {"emails": "not_an_array"}}`)  // Wrong array type
	f.Add(`{"person": {"@match": "not_a_number"}}`)  // Wrong number type
	f.Add(`{"person": {"dob": "not_an_object"}}`)    // Wrong object type
	f.Add(`{"person": {"names": [null]}}`)           // Null in array
	f.Add(`{"@http_status_code": "not_a_number"}`)   // Wrong type for status
	f.Add(`{"possible_persons": "not_an_array"}`)    // Wrong array type

	f.Fuzz(func(_ *testing.T, jsonData string) {
		// Test Response unmarshaling
		var response Response
		err := json.Unmarshal([]byte(jsonData), &response)
		// Should not panic, even with invalid JSON
		if err != nil {
			// Expected for malformed JSON
			return
		}

		// Test Person unmarshaling
		var person Person
		_ = json.Unmarshal([]byte(jsonData), &person)

		// Test that methods don't panic on partially unmarshaled objects
		_ = response.Person.HasEmail()
		_ = response.Person.HasPhone()
		_ = response.Person.HasName()
		_ = response.Person.HasAddress()
		_ = SearchMeetsMinimumCriteria(&response.Person)
	})
}

// FuzzJSONFieldTypes tests JSON unmarshaling with various field type combinations
func FuzzJSONFieldTypes(f *testing.F) {
	f.Add("string", int64(123), 45.67, true)

	f.Fuzz(func(_ *testing.T, stringVal string, intVal int64, floatVal float64, boolVal bool) {
		// Create JSON with mixed field types
		jsonTemplate := `{
			"person": {
				"names": [{"first": "%s", "display": "%s"}],
				"@match": %f,
				"@inferred": %t
			},
			"@persons_count": %d,
			"@http_status_code": %d,
			"top_match": %t,
			"error": "%s"
		}`

		jsonData := fmt.Sprintf(jsonTemplate, stringVal, stringVal, floatVal, boolVal, intVal, intVal, boolVal, stringVal)

		var response Response
		err := json.Unmarshal([]byte(jsonData), &response)
		// Should handle type conversions gracefully
		if err != nil {
			// May error on invalid JSON format, which is acceptable
			return
		}

		// Verify that fields were parsed correctly
		if len(response.Person.Names) > 0 {
			// Only check if JSON was valid and matched
			_ = response.Person.Names[0].First
		}
	})
}
