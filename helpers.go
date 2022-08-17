package pipl

import (
	"fmt"
	"strings"
)

const genderMale = "male"
const genderFemale = "female"

// SearchMeetsMinimumCriteria is used internally by Search to do some very
// basic verification that the verify that search object has enough terms to
// meet the requirements for a search.
// From Pipl documentation:
//
//	"The minimal requirement to run a search is to have at least one full
//	name, email, phone, username, user_id, URL or a single valid US address
//	(down to a house number). We can’t search for a job title or location
//	alone. We’re not a directory and can't provide bulk lists of people,
//	rather we specialize in identity resolution of single individuals."
func SearchMeetsMinimumCriteria(searchPerson *Person) bool {

	// If an email is found, that meets minimum criteria
	if searchPerson.HasEmail() {
		return true
	}

	// If a phone is found, that meets minimum criteria
	if searchPerson.HasPhone() {
		return true
	}

	// If a userID is found, that meets minimum criteria
	if searchPerson.HasUserID() {
		return true
	}

	// If a username is found, that meets minimum criteria
	if searchPerson.HasUsername() {
		return true
	}

	// If a URL is found, that meets minimum criteria
	if searchPerson.HasURL() {
		return true
	}

	// If a full name is found, that meets minimum criteria
	if searchPerson.HasName() {
		return true
	}

	// If an address is found, that meets minimum criteria
	if searchPerson.HasAddress() {
		return true
	}

	// Did not meet criteria, fail
	return false
}

// NewPerson makes a new blank person object to be filled with terms
func NewPerson() *Person {
	return new(Person)
}

// AddName adds a name to the search object. For well-defined names.
//
// Source: https://docs.pipl.com/reference#name
//
// Note: Values are assumed to be sanitized already
//
// Plan: All Plans
func (p *Person) AddName(firstName, middleName, lastName, prefix, suffix string) error {

	// We need a first or last name
	if len(firstName) == 0 && len(lastName) == 0 {
		return ErrMissingFirstLastName
	}

	// Start a new name
	newName := new(Name)
	newName.First = firstName
	newName.Middle = middleName
	newName.Last = lastName
	newName.Prefix = prefix
	newName.Suffix = suffix
	p.Names = append(p.Names, *newName)
	return nil
}

// AddNameRaw can be used when you're unsure how to handle breaking down the name in
// question into its constituent parts. Basically, let Pipl handle parsing it.
//
// Source: https://docs.pipl.com/reference#name
//
// Note: Values are assumed to be sanitized already
//
// Plan: All Plans
func (p *Person) AddNameRaw(fullName string) error {

	// Do we have a valid name?
	if len(fullName) <= 5 {
		return ErrNameTooShort
	}

	// Start the name
	newName := new(Name)
	newName.Raw = fullName
	p.Names = append(p.Names, *newName)
	return nil
}

// AddEmail appends an email address to the specified search object
//
// Source: https://docs.pipl.com/reference#email
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Business Plan
func (p *Person) AddEmail(emailAddress string) (err error) {

	// No Email or missing @ sign
	if len(emailAddress) == 0 || !strings.Contains(emailAddress, "@") {
		return ErrInvalidEmailAddress
	}

	// Add the email
	newEmail := new(Email)
	newEmail.Address = emailAddress
	p.Emails = append(p.Emails, *newEmail)
	return nil
}

// AddUsername appends a username to the specified search object
//
// Source: https://docs.pipl.com/reference#username
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Social Plan
func (p *Person) AddUsername(username string, serviceProvider string) error {

	// Must be greater than 3 characters
	if len(username) <= 3 {
		return ErrUserNameTooShort
	}

	// Must be greater than 2 characters
	if len(serviceProvider) <= 2 {
		return ErrServiceProviderTooShort
	}

	// Accepted provider?
	if !isAcceptedValue(strings.ToLower(serviceProvider), &AllowedServiceProviders) {
		return ErrServiceProviderNotAccepted
	}

	// Add username
	newUsername := new(Username)
	newUsername.Content = username + "@" + serviceProvider
	p.Usernames = append(p.Usernames, *newUsername)
	return nil
}

// AddUserID appends a user ID to the specified search object
//
// Source: https://docs.pipl.com/reference#user-id
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Social Plan
func (p *Person) AddUserID(userID, serviceProvider string) error {

	// Must be greater than 2 characters
	if len(userID) <= 2 {
		return ErrUserIDTooShort
	}

	// Must be greater than 2 characters
	if len(serviceProvider) <= 2 {
		return ErrServiceProviderTooShort
	}

	// Accepted provider?
	if !isAcceptedValue(strings.ToLower(serviceProvider), &AllowedServiceProviders) {
		return ErrServiceProviderNotAccepted
	}

	// Set the user id
	newUserID := new(UserID)
	newUserID.Content = userID + "@" + serviceProvider
	p.UserIDs = append(p.UserIDs, *newUserID)
	return nil
}

// AddURL appends a URL to the specified search object (website, social, etc)
//
// Source: https://docs.pipl.com/reference#url
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Social Plan
func (p *Person) AddURL(url string) error {
	// Must be greater than 5 characters
	if len(url) <= 5 {
		return ErrURLTooShort
	}

	// Set the url
	newURL := new(URL)
	newURL.URL = url
	p.URLs = append(p.URLs, *newURL)
	return nil
}

// AddPhone appends a phone to the specified search object
//
// Source: https://docs.pipl.com/reference#phone
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Landline phones are available in all plans, mobile phones
//
//	in the BUSINESS plan only.
func (p *Person) AddPhone(phoneNumber int64, countryCode int) error {

	// Phone is required (min length unknown)
	if phoneNumber == 0 {
		return ErrInvalidPhoneNumber
	}

	// Country code is required
	// International call country code. See ITU-T Recommendation E.164
	if countryCode == 0 {
		return ErrMissingCountryCode
	}

	// Set phone
	newPhone := new(Phone)
	newPhone.Number = phoneNumber
	newPhone.CountryCode = countryCode
	p.Phones = append(p.Phones, *newPhone)
	return nil
}

// AddPhoneRaw appends a phone to the specified search object
//
// Source:https://docs.pipl.com/reference#phone
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Landline phones are available in all plans, mobile phones
//
//	in the BUSINESS plan only.
func (p *Person) AddPhoneRaw(phoneNumber string) error {

	// Phone is required (min length unknown)
	if len(phoneNumber) < 4 {
		return ErrInvalidPhoneNumber
	}

	// Set the phone
	newPhone := new(Phone)
	newPhone.Raw = phoneNumber
	p.Phones = append(p.Phones, *newPhone)
	return nil
}

// AddAddress appends an address to the specified search object
//
// Source: https://docs.pipl.com/reference#address
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) AddAddress(house, street, apartment, city, state, country, poBox string) error {

	// Must have a number and street
	if len(house) == 0 || len(street) == 0 {
		return ErrMissingNumberOrStreet
	}

	// Must have a city or state
	if len(city) == 0 && len(state) == 0 {
		return ErrMissingCityOrState
	}

	// Force the country to Default (for searching only USA)
	// Alpha-2 ISO 3166 country code
	// https://docs.pipl.com/reference#address
	if country != DefaultCountry {
		country = DefaultCountry
	}

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
	return nil
}

// AddAddressRaw can be used when many of the address parts are missing, or
// you're unsure how to split it up. Let Pipl handle parsing.
//
// Source: https://docs.pipl.com/reference#address
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) AddAddressRaw(fullAddress string) error {

	// Must have a minimum length
	if len(fullAddress) < 5 {
		return ErrAddressTooShort
	}

	// Set the address
	newAddress := new(Address)
	newAddress.Raw = fullAddress
	p.Addresses = append(p.Addresses, *newAddress)
	return nil
}

// AddJob appends a job entry to the specified search object
//
// Source: https://docs.pipl.com/reference#job
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Business Plan
func (p *Person) AddJob(title, organization, industry, dateRangeStart, dateRangeEnd string) error {

	// Title required
	if len(title) == 0 || len(organization) == 0 {
		return ErrMissingTitleOrOrganization
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
	return nil
}

// AddEducation appends an education entry to the specified search object
//
// Source: https://docs.pipl.com/reference#education
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Business Plan
func (p *Person) AddEducation(degree, school, dateRangeStart, dateRangeEnd string) error {

	// School is required
	if len(school) == 0 {
		return ErrMissingSchool
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
	return nil
}

// AddLanguage appends a language to the specified search object.
// Language is a 2 character language code (e.g. "en")
// Region is a country code (e.g "US") Alpha-2 ISO 3166 country code
// Language objects are not used by the API for search.
//
// Source: https://docs.pipl.com/reference#language
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) AddLanguage(languageCode, regionCode string) error {

	// Invalid language code (2 char only)
	if len(languageCode) >= 3 {
		return ErrInvalidLanguageCode
	}

	// Invalid region code (4 is a guess)
	if len(regionCode) >= 4 {
		return ErrInvalidRegionCode
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
	return nil
}

// AddEthnicity appends an ethnicity to the specified search object
// Ethnicity data is not yet available. (6/8/19)
//
// Source: https://docs.pipl.com/reference#ethinicity
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) AddEthnicity(ethnicity string) error {

	// No ethnicity
	if len(ethnicity) == 0 {
		return ErrInvalidEthnicity
	}

	// Accepted provider?
	if !isAcceptedValue(strings.ToLower(ethnicity), &AllowedEthnicities) {
		return ErrEthnicityNotAccepted
	}

	// Add ethnicity
	newEthnicity := new(Ethnicity)
	newEthnicity.Content = ethnicity
	p.Ethnicities = append(p.Ethnicities, *newEthnicity)
	return nil
}

// AddOriginCountry appends an origin country to the specified search object
// Alpha-2 ISO 3166 country code
//
// Source: https://docs.pipl.com/reference#origin-country
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) AddOriginCountry(countryCode string) error {

	// No code
	if len(countryCode) == 0 {
		return ErrInvalidCountryCode
	}

	// todo: accepted list of countries?

	// Set country
	newCountry := new(OriginCountry)
	newCountry.Country = countryCode
	p.OriginCountries = append(p.OriginCountries, *newCountry)
	return nil
}

// AddRelationship appends a relationship entry to the specified search object
//
// Source: https://docs.pipl.com/reference#relationship
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: Social Plan
func (p *Person) AddRelationship(relationship Relationship) (err error) {
	// todo: missing validations

	// Set relationship
	p.Relationships = append(p.Relationships, relationship)
	return
}

// SetGender sets the gender of the specified search object
//
// Source: https://docs.pipl.com/reference#gender
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) SetGender(gender string) error {

	// Invalid gender
	if gender != genderMale && gender != genderFemale {
		return ErrInvalidGender
	}

	// Set gender
	newGender := new(Gender)
	newGender.Content = gender
	p.Gender = newGender
	return nil
}

// SetDateOfBirth sets the DOB of the specified search object
// DOB string format: "YYYY-MM-DD"
// DOB string format: "YYYY-MM-DD"
// Set both Start and End to the same date if known
//
// Source: https://docs.pipl.com/reference#date-of-birth
//
// Note: Values are assumed to be sanitized/validated already
//
// Plan: All Plans
func (p *Person) SetDateOfBirth(startDate, endDate string) error {

	// No start date
	if len(startDate) == 0 || len(endDate) == 0 {
		return ErrMissingBirthDate
	}

	// Rough validation check
	if !strings.Contains(startDate, "-") || len(startDate) != 10 {
		return ErrInvalidStartOfBirthDate
	}

	// Rough validation check
	if !strings.Contains(endDate, "-") || len(endDate) != 10 {
		return ErrInvalidEndOfBirthDate
	}

	// todo: test date compared to today (cannot be future date)
	// todo: regex on the YYYY-MM-DD

	// Add the date range for dob
	newDOB := new(DateOfBirth)
	newDOB.DateRange.Start = startDate
	newDOB.DateRange.End = endDate
	p.DateOfBirth = newDOB
	return nil
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

// HasURL returns true if the person has an url
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
func (p *Person) ProcessThumbnails(thumbnailSettings *ThumbnailSettings) {

	// No images to process or no client
	if len(p.Images) == 0 || thumbnailSettings == nil {
		return
	}

	// Loop all images
	for index, image := range p.Images {
		if image.ThumbnailToken != "" {
			p.Images[index].ThumbnailURL = fmt.Sprintf("%s?height=%d&width=%d&favicon=%t&zoom_face=%t&tokens=%s",
				thumbnailSettings.URL,
				thumbnailSettings.Height,
				thumbnailSettings.Width,
				thumbnailSettings.Favicon,
				thumbnailSettings.ZoomFace,
				image.ThumbnailToken,
			)
		}
	}
}

// isAcceptedValue tests a value against allowed values
func isAcceptedValue(testValue string, allowedValues *[]string) (success bool) {
	// Check that the value is an allowed value (case-sensitive)
	for _, value := range *allowedValues {
		if testValue == value {
			success = true
			return
		}
	}
	return
}
