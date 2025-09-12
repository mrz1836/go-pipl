package pipl

import "errors"

// ErrDoesNotMeetMinimumCriteria is when the minimum search criteria is not met
var ErrDoesNotMeetMinimumCriteria = errors.New("the search request submitted does not contain enough sufficient terms. " +
	"You must have one of the following: full name, email, phone, username, userID, url, or full street address")

// ErrInvalidSearchPointer is when the SEARCH_POINTER is not valid (IE: too short)
var ErrInvalidSearchPointer = errors.New("invalid search pointer")

// ErrInvalidEndOfBirthDate is when the END birthdate is invalid
var ErrInvalidEndOfBirthDate = errors.New("invalid end date of birth")

// ErrInvalidStartOfBirthDate is when the START birthdate is invalid
var ErrInvalidStartOfBirthDate = errors.New("invalid start date of birth")

// ErrMissingBirthDate is when the START or END is missing
var ErrMissingBirthDate = errors.New("missing start or end date of birth")

// ErrMissingFirstLastName is when the FIRST and LAST name is missing
var ErrMissingFirstLastName = errors.New("first name and last name are missing")

// ErrNameTooShort is when the NAME is too short
var ErrNameTooShort = errors.New("name is too short, minimum of 5 characters")

// ErrUserNameTooShort is when the USERNAME is too short
var ErrUserNameTooShort = errors.New("username is too short")

// ErrUserIDTooShort is when the USER_ID is too short
var ErrUserIDTooShort = errors.New("user_id is too short")

// ErrURLTooShort is when the URL is too short
var ErrURLTooShort = errors.New("url is too short")

// ErrServiceProviderTooShort is when the SERVICE_PROVIDER is too short
var ErrServiceProviderTooShort = errors.New("service_provider is too short")

// ErrServiceProviderNotAccepted is when the SERVICE_PROVIDER is not known or accepted
var ErrServiceProviderNotAccepted = errors.New("service_provider is not accepted")

// ErrInvalidEmailAddress is when the EMAIL is invalid or empty
var ErrInvalidEmailAddress = errors.New("email is invalid or empty")

// ErrInvalidPhoneNumber is when the PHONE is invalid or empty
var ErrInvalidPhoneNumber = errors.New("invalid phone number")

// ErrMissingCountryCode is when the COUNTRY_CODE is missing from the phone number
var ErrMissingCountryCode = errors.New("missing country code")

// ErrMissingNumberOrStreet is when the NUMBER or STREET is missing from an address
var ErrMissingNumberOrStreet = errors.New("missing number or street")

// ErrMissingCityOrState is when the CITY or STATE is missing from an address
var ErrMissingCityOrState = errors.New("missing have a city or state")

// ErrAddressTooShort is when the ADDRESS is too short
var ErrAddressTooShort = errors.New("address is too short")

// ErrMissingTitleOrOrganization is when the TITLE or ORGANIZATION is missing
var ErrMissingTitleOrOrganization = errors.New("missing required title or organization")

// ErrMissingSchool is when the SCHOOL is missing
var ErrMissingSchool = errors.New("missing required school")

// ErrInvalidLanguageCode is when the LANGUAGE_CODE is invalid
var ErrInvalidLanguageCode = errors.New("invalid language code")

// ErrInvalidRegionCode is when the REGION_CODE is invalid
var ErrInvalidRegionCode = errors.New("invalid region code")

// ErrInvalidEthnicity is when the ETHNICITY is invalid
var ErrInvalidEthnicity = errors.New("invalid ethnicity")

// ErrEthnicityNotAccepted is when the ETHNICITY is not found or accepted
var ErrEthnicityNotAccepted = errors.New("ethnicity is not accepted")

// ErrInvalidCountryCode is when the COUNTRY_CODE is invalid
var ErrInvalidCountryCode = errors.New("invalid country code")

// ErrInvalidGender is when the GENDER is invalid
var ErrInvalidGender = errors.New("invalid gender")

// ErrMissingRequest is when the HTTP request is missing
var ErrMissingRequest = errors.New("missing request")

// ErrRequestNotFound is when the request is not found
var ErrRequestNotFound = errors.New("request not found")

// ErrBadRequest is when the request is bad
var ErrBadRequest = errors.New("bad request")

// ErrAPIResponse is when the API returns an error response
var ErrAPIResponse = errors.New("API response error")
