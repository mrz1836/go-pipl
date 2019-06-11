package pipl

import (
	"fmt"
	"strings"
)

// NewPerson makes a new blank person object to be filled with terms
func NewPerson() *Person {
	return new(Person)
}

// AddName adds a name to the search object. For well defined names.
//
// Source: https://docs.pipl.com/reference#name
//
// Note: Values are assumed to be sanitized already
//
// Plan: All Plans
func (p *Person) AddName(firstName, middleName, lastName, prefix, suffix string) (err error) {

	// We need a first or last name
	if len(firstName) == 0 && len(lastName) == 0 {
		err = fmt.Errorf("first name and last name are missing")
		return
	}

	// Start a new name
	newName := new(Name)
	newName.First = firstName
	newName.Middle = middleName
	newName.Last = lastName
	newName.Prefix = prefix
	newName.Suffix = suffix
	p.Names = append(p.Names, *newName)
	return
}

// AddNameRaw can be used when you're unsure how to handle breaking down the name in
// question into its constituent parts. Basically, let Pipl handle parsing it.
// https://docs.pipl.com/reference#name
// Note: Values are assumed to be sanitized already
// Plan: All Plans
func (p *Person) AddNameRaw(fullName string) (err error) {

	// Do we have a valid name?
	if len(fullName) <= 5 {
		err = fmt.Errorf("name is too short, minimum of 5 characters: %s", fullName)
		return
	}

	// Start the name
	newName := new(Name)
	newName.Raw = fullName
	p.Names = append(p.Names, *newName)
	return
}

// AddEmail appends an email address to the specified search object
// https://docs.pipl.com/reference#email
// Note: Values are assumed to be sanitized/validated already
// Plan: Business Plan
func (p *Person) AddEmail(emailAddress string) (err error) {

	// No Email or missing @ sign
	if len(emailAddress) == 0 || !strings.Contains(emailAddress, "@") {
		err = fmt.Errorf("email is invalid or empty: %s", emailAddress)
		return
	}

	// Add the email
	newEmail := new(Email)
	newEmail.Address = emailAddress
	p.Emails = append(p.Emails, *newEmail)
	return
}

// AddUsername appends a username to the specified search object
// https://docs.pipl.com/reference#username
// Note: Values are assumed to be sanitized/validated already
// Plan: Social Plan
func (p *Person) AddUsername(username string) (err error) {

	// Must be greater than 3 characters
	if len(username) <= 3 {
		err = fmt.Errorf("username is too short: %s", username)
		return
	}

	// Add username
	newUsername := new(Username)
	newUsername.Content = username
	p.Usernames = append(p.Usernames, *newUsername)
	return
}

// AddUserID appends a user ID to the specified search object
// https://docs.pipl.com/reference#user-id
// Note: Values are assumed to be sanitized/validated already
// Plan: Social Plan
// Format: user-id, service
func (p *Person) AddUserID(userID, serviceProvider string) (err error) {

	// Must be greater than 2 characters
	if len(userID) <= 2 {
		err = fmt.Errorf("user_id is too short: %s", userID)
		return
	}

	// Must be greater than 2 characters
	if len(serviceProvider) <= 2 {
		err = fmt.Errorf("service_provider is too short: %s", userID)
		return
	}

	// Accepted provider?
	if !isAcceptedValue(strings.ToLower(serviceProvider), &AllowedServiceProviders) {
		err = fmt.Errorf("service_provider is not accepted: %s", serviceProvider)
		return
	}

	// Set the user id
	newUserID := new(UserID)
	newUserID.Content = userID + "@" + serviceProvider
	p.UserIDs = append(p.UserIDs, *newUserID)
	return
}

// AddURL appends a URL to the specified search object (website, social, etc)
// https://docs.pipl.com/reference#url
// Note: Values are assumed to be sanitized/validated already
// Plan: Social Plan
func (p *Person) AddURL(url string) (err error) {
	// Must be greater than 5 characters
	if len(url) <= 5 {
		err = fmt.Errorf("url is too short: %s", url)
		return
	}

	// Set the url
	newURL := new(URL)
	newURL.URL = url
	p.URLs = append(p.URLs, *newURL)
	return
}

// AddPhone appends a phone to the specified search object
// https://docs.pipl.com/reference#phone
// Note: Values are assumed to be sanitized/validated already
// Plan: Landline phones are available in all plans, mobile phones
// 		 in the BUSINESS plan only.
func (p *Person) AddPhone(phoneNumber, countryCode int) (err error) {

	// Phone is required (min length unknown)
	if phoneNumber == 0 {
		err = fmt.Errorf("invalid phone number: %d", phoneNumber)
		return
	}

	// Country code is required
	//International call country code. See ITU-T Recommendation E.164
	if countryCode == 0 {
		err = fmt.Errorf("missing country code: %d", phoneNumber)
		return
	}

	// Set phone
	newPhone := new(Phone)
	newPhone.Number = phoneNumber
	newPhone.CountryCode = countryCode
	p.Phones = append(p.Phones, *newPhone)
	return
}

// AddPhoneRaw appends a phone to the specified search object
// https://docs.pipl.com/reference#phone
// Note: Values are assumed to be sanitized/validated already
// Plan: Landline phones are available in all plans, mobile phones
// 		 in the BUSINESS plan only.
func (p *Person) AddPhoneRaw(phoneNumber string) (err error) {

	// Phone is required (min length unknown)
	if len(phoneNumber) < 4 {
		err = fmt.Errorf("invalid phone number: %s", phoneNumber)
		return
	}

	// Set the phone
	newPhone := new(Phone)
	newPhone.Raw = phoneNumber
	p.Phones = append(p.Phones, *newPhone)
	return
}

// AddAddress appends an address to the specified search object
// https://docs.pipl.com/reference#address
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) AddAddress(house, street, apartment, city, state, country, poBox string) (err error) {

	// Must have a number and street
	if len(house) == 0 || len(street) == 0 {
		err = fmt.Errorf("missing number or street")
		return
	}

	// Must have a city or state
	if len(city) == 0 && len(state) == 0 {
		err = fmt.Errorf("missing have a city or state")
		return
	}

	// Force the country to Default (for searching only USA)
	// Alpha-2 ISO 3166 country code
	// https://docs.pipl.com/reference#address
	country = DefaultCountry

	// Set the address
	newAddress := new(Address)
	newAddress.Apartment = apartment
	newAddress.City = city
	newAddress.Country = country
	newAddress.House = house
	newAddress.POBox = poBox
	newAddress.State = state
	newAddress.Street = street
	p.Addresses = append(p.Addresses, *newAddress)
	return
}

// AddAddressRaw can be used when many of the address parts are missing, or
// you're unsure how to split it up. Let Pipl handle parsing.
// https://docs.pipl.com/reference#address
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) AddAddressRaw(fullAddress string) (err error) {

	// Must have a minimum length
	if len(fullAddress) < 5 {
		err = fmt.Errorf("address is too short")
		return
	}

	// Set the address
	newAddress := new(Address)
	newAddress.Raw = fullAddress
	p.Addresses = append(p.Addresses, *newAddress)
	return
}

// AddJob appends a job entry to the specified search object
// https://docs.pipl.com/reference#job
// Note: Values are assumed to be sanitized/validated already
// Plan: Business Plan
func (p *Person) AddJob(title, organization, industry, dateRangeStart, dateRangeEnd string) (err error) {

	// Title required
	if len(title) == 0 || len(organization) == 0 {
		err = fmt.Errorf("missing required title or organization")
		return
	}

	// Add job
	newJob := new(Job)
	newJob.Title = title
	newJob.Organization = organization
	newJob.Industry = industry

	// todo: same test for dates (like DOB)
	// Only add if found
	if len(dateRangeStart) > 0 {
		newJob.DateRange.Start = dateRangeStart
	}
	if len(dateRangeEnd) > 0 {
		newJob.DateRange.End = dateRangeEnd
	}

	p.Jobs = append(p.Jobs, *newJob)
	return
}

// AddEducation appends an education entry to the specified search object
// https://docs.pipl.com/reference#education
// Note: Values are assumed to be sanitized/validated already
// Plan: Business Plan
func (p *Person) AddEducation(degree, school, dateRangeStart, dateRangeEnd string) (err error) {

	// School is required
	if len(school) == 0 {
		err = fmt.Errorf("missing required parameter: school")
		return
	}

	// Set the education
	newEducation := new(Education)
	newEducation.Degree = degree
	newEducation.School = school

	// todo: same test for dates (like DOB)
	if len(dateRangeStart) > 0 {
		newEducation.DateRange.Start = dateRangeStart
	}
	if len(dateRangeEnd) > 0 {
		newEducation.DateRange.End = dateRangeEnd
	}
	p.Educations = append(p.Educations, *newEducation)
	return
}

// AddLanguage appends a language to the specified search object.
// Language is a 2 character language code (e.g. "en")
// Region is a country code (e.g "US") Alpha-2 ISO 3166 country code
// Language objects are not used by the API for search.
// https://docs.pipl.com/reference#language
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) AddLanguage(languageCode, regionCode string) (err error) {

	// Invalid language code (2 char only)
	if len(languageCode) >= 3 {
		err = fmt.Errorf("invalud lanuage code: %s", languageCode)
		return
	}

	// Invalid region code (4 is a guess)
	if len(regionCode) >= 4 {
		err = fmt.Errorf("invalud region code: %s", regionCode)
		return
	}

	// todo: validate accepted languages and regions

	// Set the language
	newLanguage := new(Language)
	newLanguage.Language = languageCode
	newLanguage.Region = regionCode

	// Add the display if both are set
	if len(languageCode) > 0 && len(regionCode) > 0 {
		newLanguage.Display = languageCode + "_" + regionCode
	}
	p.Languages = append(p.Languages, *newLanguage)
	return
}

// AddEthnicity appends an ethnicity to the specified search object
// Ethnicity data is not yet available. (6/8/19)
// https://docs.pipl.com/reference#ethinicity
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) AddEthnicity(ethnicity string) (err error) {

	// No ethnicity
	if len(ethnicity) == 0 {
		err = fmt.Errorf("invalid ethnicity")
		return
	}

	// Accepted provider?
	if !isAcceptedValue(strings.ToLower(ethnicity), &AllowedEthnicities) {
		err = fmt.Errorf("ethnicity is not accepted: %s", ethnicity)
		return
	}

	// Add ethnicity
	newEthnicity := new(Ethnicity)
	newEthnicity.Content = ethnicity
	p.Ethnicities = append(p.Ethnicities, *newEthnicity)
	return
}

// AddOriginCountry appends an origin country to the specified search object
// Alpha-2 ISO 3166 country code
// https://docs.pipl.com/reference#origin-country
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) AddOriginCountry(countryCode string) (err error) {

	// No code
	if len(countryCode) == 0 {
		err = fmt.Errorf("invalid country code")
		return
	}

	// todo: accepted list of countries?

	// Set country
	newCountry := new(OriginCountry)
	newCountry.Country = countryCode
	p.OriginCountries = append(p.OriginCountries, *newCountry)
	return
}

// AddRelationship appends a relationship entry to the specified search object
// https://docs.pipl.com/reference#relationship
// Note: Values are assumed to be sanitized/validated already
// Plan: Social Plan
func (p *Person) AddRelationship(relationship Relationship) (err error) {
	// todo: missing validations

	// Set relationship
	p.Relationships = append(p.Relationships, relationship)
	return
}

// SetGender sets the gender of the specified search object
// https://docs.pipl.com/reference#gender
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) SetGender(gender string) (err error) {

	// Invalid gender
	if gender != "male" && gender != "female" {
		err = fmt.Errorf("invalid gender: %s", gender)
	}

	// Set gender
	newGender := new(Gender)
	newGender.Content = gender
	p.Gender = newGender
	return
}

// SetDateOfBirth sets the DOB of the specified search object
// DOB string format: "YYYY-MM-DD"
// Set both Start and End to the same date if known
// https://docs.pipl.com/reference#date-of-birth
// Note: Values are assumed to be sanitized/validated already
// Plan: All Plans
func (p *Person) SetDateOfBirth(startDate, endDate string) (err error) {

	// No start date
	if len(startDate) == 0 || len(endDate) == 0 {
		err = fmt.Errorf("missing start or end date of birth")
		return
	}

	// Rough validation check
	if !strings.Contains(startDate, "-") || len(startDate) != 10 {
		err = fmt.Errorf("invalid start date of birth")
		return
	}

	// Rough validation check
	if !strings.Contains(endDate, "-") || len(endDate) != 10 {
		err = fmt.Errorf("invalid end date of birth")
		return
	}

	// todo: test date compared to today (cannot be future date)
	// todo: regex on the YYYY-MM-DD

	// Add the date range for dob
	newDOB := new(DateOfBirth)
	newDOB.DateRange.Start = startDate
	newDOB.DateRange.End = endDate
	p.DateOfBirth = newDOB
	return
}

// HasEmail returns true if the person has an email address
func (p *Person) HasEmail() bool {
	if len(p.Emails) > 0 {
		for _, email := range p.Emails {
			if email.Address != "" {
				return true
			}
		}
	}
	return false
}

// HasPhone returns true if the person has a phone number
func (p *Person) HasPhone() bool {
	if len(p.Phones) > 0 {
		for _, phone := range p.Phones {
			if (phone.CountryCode != 0 && phone.Number != 0) || (phone.Raw != "") {
				return true
			}
		}
	}
	return false
}

// HasUserID returns true if the person has a user id
func (p *Person) HasUserID() bool {
	if len(p.UserIDs) > 0 {
		for _, userID := range p.UserIDs {
			if userID.Content != "" {
				return true
			}
		}
	}
	return false
}

// HasUsername returns true if the person has a username
func (p *Person) HasUsername() bool {
	if len(p.Usernames) > 0 {
		for _, username := range p.Usernames {
			if username.Content != "" {
				return true
			}
		}
	}
	return false
}

// HasURL returns true if the person has a url
func (p *Person) HasURL() bool {
	if len(p.URLs) > 0 {
		for _, u := range p.URLs {
			if u.URL != "" {
				return true
			}
		}
	}
	return false
}

// HasName returns true if the person has a name (minimum required)
func (p *Person) HasName() bool {
	if len(p.Names) > 0 {
		for _, name := range p.Names {
			if (name.First != "" && name.Last != "") || (name.Raw != "") {
				return true
			}
		}
	}
	return false
}

// HasAddress returns true if the person has an address (minimum required)
func (p *Person) HasAddress() bool {
	if len(p.Addresses) > 0 {
		for _, address := range p.Addresses {
			if address.House != "" && address.Street != "" && address.City != "" && address.State != "" {
				return true
			}
		}
	}
	return false
}

// ProcessThumbnails checks for any images and adds thumbnail urls to the existing image
// Requires the client for now since the client has the current configuration
func (p *Person) ProcessThumbnails(c *Client) {

	// No images to process or no client
	if len(p.Images) == 0 || c == nil {
		return
	}

	// Loop all images
	for index, image := range p.Images {
		if image.ThumbnailToken != "" {
			p.Images[index].ThumbnailURL = fmt.Sprintf("%s?height=%d&width=%d&favicon=%t&zoom_face=%t&tokens=%s",
				c.ThumbnailSettings.URL,
				c.ThumbnailSettings.Height,
				c.ThumbnailSettings.Width,
				c.ThumbnailSettings.Favicon,
				c.ThumbnailSettings.ZoomFace,
				image.ThumbnailToken,
			)
		}
	}
}

// isAcceptedValue tests a value against allowed values
func isAcceptedValue(testValue string, allowedValues *[]string) (success bool) {
	//Check that the value is an allowed value (case sensitive)
	for _, value := range *allowedValues {
		if testValue == value {
			success = true
			return
		}
	}
	return
}
