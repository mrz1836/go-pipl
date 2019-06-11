package pipl

// GUID is a unique format (but is just a string internally, since there's currently
// nothing all that fancy done with GUIDs). Additional guid-handling code may be
// added at a later date if needed.
type GUID string

// Name fields collectively define a possible name for a given person.
// If a search did not return information for a given field, it will be empty.
//
// Source: https://docs.pipl.com/reference#name
type Name struct {
	Current    bool   `json:"@current,omitempty"`
	Display    string `json:"display,omitempty"`
	First      string `json:"first,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Last       string `json:"last,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Middle     string `json:"middle,omitempty"`
	Prefix     string `json:"prefix,omitempty"`
	Raw        string `json:"raw,omitempty"`
	Suffix     string `json:"suffix,omitempty"`
	Type       string `json:"@type,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// Address fields collectively define a possible address for a given person
// If a search did not return information for a given field, it will be empty.
//
// Source: https://docs.pipl.com/reference#address
type Address struct {
	Apartment  string `json:"apartment,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Display    string `json:"display,omitempty"`
	House      string `json:"house,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	POBox      string `json:"po_box,omitempty"`
	Raw        string `json:"raw,omitempty"`
	State      string `json:"state,omitempty"`
	Street     string `json:"street,omitempty"`
	Type       string `json:"@type,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
	ZipCode    string `json:"zip_code,omitempty"`
}

// Phone fields collectively define a possible phone number for a given person
// If a search did not return information for a given field, it will be empty.
//
// Source: https://docs.pipl.com/reference#phone
type Phone struct {
	CountryCode          int    `json:"country_code,omitempty"`
	Current              bool   `json:"@current,omitempty"`
	Display              string `json:"display,omitempty"`
	DisplayInternational string `json:"display_international,omitempty"`
	Extension            int    `json:"extension,omitempty"`
	Inferred             bool   `json:"@inferred,omitempty"`
	LastSeen             string `json:"@last_seen,omitempty"`
	Number               int    `json:"number,omitempty"`
	Raw                  string `json:"raw,omitempty"`
	Type                 string `json:"@type,omitempty"`
	ValidSince           string `json:"@valid_since,omitempty"`
}

// Email fields collectively define a possible email address for a given person
// If a search did not return information for a given field, it will be empty.
//
// Source: https://docs.pipl.com/reference#email
type Email struct {
	Address       string `json:"address,omitempty"`
	AddressMD5    string `json:"address_md5,omitempty"`
	Current       bool   `json:"@current,omitempty"`
	Disposable    bool   `json:"@disposable,omitempty"`
	EmailProvider bool   `json:"@email_provider,omitempty"`
	Inferred      bool   `json:"@inferred,omitempty"`
	LastSeen      string `json:"@last_seen,omitempty"`
	Type          string `json:"@type,omitempty"`
	ValidSince    string `json:"@valid_since,omitempty"`
}

// Username fields collectively define a possible username used by a given person.
// If a search did not return information for a given field, it will be empty.
//
// Source: https://docs.pipl.com/reference#username
type Username struct {
	Content    string `json:"content,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// UserID fields collectively define a possible UserID used by a given person.
// If a search did not return information for a given field, it will be empty.
//
// Source: https://docs.pipl.com/reference#user-id
type UserID struct {
	Content    string `json:"content,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// AllowedServiceProviders is all the providers
//
// Source: https://docs.pipl.com/reference#section-list-of-known-services
//
// Note: all lowercase, case sensitive
var AllowedServiceProviders = []string{
	"aboutme",
	"badoo",
	"bebo",
	"classmates",
	"cpf",
	"cyworld",
	"delicious",
	"deviantart",
	"digg",
	"douban",
	"ebay",
	"facebook",
	"flavorsme",
	"flickr",
	"flixster",
	"foursquare",
	"freelancer",
	"friendster",
	"gaia",
	"github",
	"goodreads",
	"google",
	"gravatar",
	"habbo",
	"hi5",
	"hyves",
	"imgur",
	"instagram",
	"lastfm",
	"linkedin",
	"livejournal",
	"meetup",
	"myheritage",
	"mylife",
	"myspace",
	"myyearbook",
	"netlog",
	"ning",
	"odnoklassniki",
	"orkut",
	"pinterest",
	"quora",
	"qzone",
	"renren",
	"sonico",
	"soundcloud",
	"stumbleupon",
	"tagged",
	"tumblr",
	"twitter",
	"viadeo",
	"vkontakte",
	"weeworld",
	"xanga",
	"xing",
	"yelp",
	"youtube",
}

// DateRange specifies a range of time by a start and end date
//
// Source: https://docs.pipl.com/reference#date-range
type DateRange struct {
	Current    bool   `json:"@current,omitempty"`
	End        string `json:"end,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Start      string `json:"start,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// DateOfBirth specifies a possible DOB for a person.
//
// Source: https://docs.pipl.com/reference#date-of-birth
type DateOfBirth struct {
	Current    bool      `json:"@current,omitempty"`
	DateRange  DateRange `json:"date_range,omitempty"`
	Display    string    `json:"display,omitempty"`
	Inferred   bool      `json:"@inferred,omitempty"`
	LastSeen   string    `json:"@last_seen,omitempty"`
	ValidSince string    `json:"@valid_since,omitempty"`
}

// Image specifies a link to an image closely associated with the given person.
//
// Source: https://docs.pipl.com/reference#image
type Image struct {
	Current        bool   `json:"@current,omitempty"`
	Inferred       bool   `json:"@inferred,omitempty"`
	LastSeen       string `json:"@last_seen,omitempty"`
	ThumbnailToken string `json:"thumbnail_token,omitempty"`
	ThumbnailURL   string `json:"thumbnail_url,omitempty"`
	URL            string `json:"url,omitempty"`
	ValidSince     string `json:"@valid_since,omitempty"`
}

// Job specifies information about a possible occupation held by the given person.
//
// Source: https://docs.pipl.com/reference#job
type Job struct {
	Current      bool      `json:"@current,omitempty"`
	DateRange    DateRange `json:"date_range,omitempty"`
	Display      string    `json:"display,omitempty"`
	Industry     string    `json:"industry,omitempty"`
	Inferred     bool      `json:"@inferred,omitempty"`
	LastSeen     string    `json:"@last_seen,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Title        string    `json:"title,omitempty"`
	ValidSince   string    `json:"@valid_since,omitempty"`
}

// Education specifies a possible
//
// Source: https://docs.pipl.com/reference#education
type Education struct {
	Current    bool      `json:"@current,omitempty"`
	DateRange  DateRange `json:"date_range,omitempty"`
	Degree     string    `json:"degree,omitempty"`
	Display    string    `json:"display,omitempty"`
	Inferred   bool      `json:"@inferred,omitempty"`
	LastSeen   string    `json:"@last_seen,omitempty"`
	School     string    `json:"school,omitempty"`
	ValidSince string    `json:"@valid_since,omitempty"`
}

// Gender contains a  possible gender of the given person.
// Gender is one of: "male", "female" (There is no default value for this field)
//
// Source: https://docs.pipl.com/reference#gender
type Gender struct {
	Content    string `json:"content,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// Ethnicity contains a possible ethnicity of given person.
//
// Source: https://docs.pipl.com/reference#ethinicity
type Ethnicity struct {
	Content    string `json:"content,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// AllowedEthnicities is all the types
//
// Source: https://docs.pipl.com/reference#ethinicity
//
// Note: all lowercase, case sensitive
var AllowedEthnicities = []string{
	"alaska_native",
	"american_indian",
	"black",
	"chamorro",
	"chinese",
	"filipino",
	"guamanian",
	"japanese",
	"korean",
	"native_hawaiian",
	"other",
	"other_asian",
	"other_pacific_islander",
	"samoan",
	"vietnamese",
	"white",
}

// Language contains information about a possible language known by the given person.
//
// Source: https://docs.pipl.com/reference#language
type Language struct {
	Current    bool   `json:"@current,omitempty"`
	Display    string `json:"display,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Language   string `json:"language,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Region     string `json:"region,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// OriginCountry contains information about a possible origin country of the
// given person.
//
// Source: https://docs.pipl.com/reference#origin-country
type OriginCountry struct {
	Country    string `json:"country,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// Relationship contains information about a person who is closely related to
// the person being searched. This can be family members, spouses, children, etc.
// Type  and Subtype contain information about the nature of the relationship to
// the person being searched. For example, Type = "Family", Subtype = "Father".
// Type can be one of: "work", "family", "friend" (default), "other"
//
// Source: https://docs.pipl.com/reference#relationship
type Relationship struct {
	Addresses       []Address       `json:"addresses,omitempty"`
	Current         bool            `json:"@current,omitempty"`
	DateOfBirth     DateOfBirth     `json:"dob,omitempty"`
	Educations      []Education     `json:"educations,omitempty"`
	Emails          []Email         `json:"emails,omitempty"`
	Ethnicities     []Ethnicity     `json:"ethnicities,omitempty"`
	Gender          Gender          `json:"gender,omitempty"`
	Images          []Image         `json:"images,omitempty"`
	Inferred        bool            `json:"@inferred,omitempty"`
	Jobs            []Job           `json:"jobs,omitempty"`
	Languages       []Language      `json:"languages,omitempty"`
	LastSeen        string          `json:"@last_seen,omitempty"`
	Names           []Name          `json:"names,omitempty"`
	OriginCountries []OriginCountry `json:"origin_countries,omitempty"`
	Phones          []Phone         `json:"phones,omitempty"`
	Relationships   []Relationship  `json:"relationships,omitempty"`
	Subtype         string          `json:"@subtype,omitempty"`
	Type            string          `json:"@type,omitempty"`
	UserIDs         []UserID        `json:"user_ids,omitempty"`
	Usernames       []Username      `json:"usernames,omitempty"`
	ValidSince      string          `json:"@valid_since,omitempty"`
}

// URL contains information about a URL that is closely associated with a given person.
//
// Source: https://docs.pipl.com/reference#url
type URL struct {
	Category   string `json:"@category,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Domain     string `json:"@domain,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Name       string `json:"@name,omitempty"`
	SourceID   string `json:"@source_id,omitempty"`
	URL        string `json:"url,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
}

// Tag contains content classification information
//
// Source: https://docs.pipl.com/reference#tag
type Tag struct {
	Classification string `json:"@classification,omitempty"`
	Content        string `json:"content,omitempty"`
}

// Person contains all the information pertaining to a possible person match,
// including potential multiples of basic fields (names, emails, jobs, etc).
// The Match field represents the confidence of a particular person match, as a
// float: 0 <= Match <= 1. More potential matches returned in a search decreases
// the overall confidence of all matches.
//
// Source: https://docs.pipl.com/reference#person
type Person struct {
	Addresses       []Address       `json:"addresses,omitempty"`
	DateOfBirth     *DateOfBirth    `json:"dob,omitempty"`
	Educations      []Education     `json:"educations,omitempty"`
	Emails          []Email         `json:"emails,omitempty"`
	Ethnicities     []Ethnicity     `json:"ethnicities,omitempty"`
	Gender          *Gender         `json:"gender,omitempty"`
	ID              GUID            `json:"@id,omitempty"`
	Images          []Image         `json:"images,omitempty"`
	Inferred        bool            `json:"@inferred,omitempty"`
	Jobs            []Job           `json:"jobs,omitempty"`
	Languages       []Language      `json:"languages,omitempty"`
	Match           float32         `json:"@match,omitempty"`
	Names           []Name          `json:"names,omitempty"`
	OriginCountries []OriginCountry `json:"origin_countries,omitempty"`
	Phones          []Phone         `json:"phones,omitempty"`
	Relationships   []Relationship  `json:"relationships,omitempty"`
	SearchPointer   string          `json:"@search_pointer,omitempty"`
	URLs            []URL           `json:"urls,omitempty"`
	UserIDs         []UserID        `json:"user_ids,omitempty"`
	Usernames       []Username      `json:"usernames,omitempty"`
}

// Source contains all the information for a given person, gathered from a
// single source. The source structure contains information about the name,
// domain, category, and source URL (amongst other fields).
//
// Source: https://docs.pipl.com/reference#source
type Source struct {
	Addresses       []Address       `json:"addresses"`
	Category        string          `json:"@category"`
	DateOfBirth     DateOfBirth     `json:"dob"`
	Domain          string          `json:"@domain"`
	Educations      []Education     `json:"educations"`
	Emails          []Email         `json:"emails"`
	Ethnicities     []Ethnicity     `json:"ethnicities"`
	Gender          Gender          `json:"gender"`
	ID              string          `json:"@id"`
	Jobs            []Job           `json:"jobs"`
	Languages       []Language      `json:"languages"`
	Match           float32         `json:"@match"`
	Name            string          `json:"@name"`
	Names           []Name          `json:"names"`
	OriginCountries []OriginCountry `json:"origin_countries"`
	OriginURL       string          `json:"@origin_url"`
	PersonID        GUID            `json:"@person_id"`
	Phones          []Phone         `json:"phones"`
	Premium         bool            `json:"@premium"`
	Relationships   []Relationship  `json:"relationships"`
	Sponsored       bool            `json:"@sponsored"`
	Tags            []Tag           `json:"tags"`
	URLs            []URL           `json:"urls"`
	UserIDs         []UserID        `json:"user_ids"`
	Usernames       []Username      `json:"usernames"`
}

// FieldCount contains the count of various attributes returned from a search
//
// Source: https://docs.pipl.com/reference#overview-2
type FieldCount struct {
	Addresses       int `json:"addresses"`
	DOBs            int `json:"dobs"`
	Educations      int `json:"educations"`
	Emails          int `json:"emails"`
	Ethnicities     int `json:"ethnicities"`
	Genders         int `json:"genders"`
	Images          int `json:"images"`
	Jobs            int `json:"jobs"`
	LandlinePhones  int `json:"landline_phones"`
	Languages       int `json:"languages"`
	MobilePhones    int `json:"mobile_phones"`
	Names           int `json:"names"`
	OriginCountries int `json:"origin_countries"`
	Phones          int `json:"phones"`
	Relationships   int `json:"relationships"`
	SocialProfiles  int `json:"social_profiles"`
	UserIDs         int `json:"user_ids"`
	Usernames       int `json:"usernames"`
}

// AvailableData aggregates the counts for found attributes that are relevant to
// your search, divided into free and paid sources.
//
// Source: https://docs.pipl.com/reference#available-data
type AvailableData struct {
	Basic   FieldCount `json:"basic"`
	Premium FieldCount `json:"premium"`
}

// Response holds search results and general request information returned from
// the Pipl API. If an error occurs, the Error field will have more information.
// A search may be successful, but have some warnings. These are held in the
// Warnings field.
//
// Source: https://docs.pipl.com/reference#overview-2
type Response struct {
	AvailableData     AvailableData `json:"available_data"`
	AvailableSources  int           `json:"@available_sources"`
	Error             string        `json:"error"`
	HTTPStatusCode    int           `json:"@http_status_code"`
	MatchRequirements string        `json:"match_requirements"`
	Person            Person        `json:"person"`
	PersonsCount      int           `json:"@persons_count"`
	PossiblePersons   []Person      `json:"possible_persons"`
	Query             Person        `json:"query"`
	SearchID          string        `json:"@search_id"`
	Sources           []Source      `json:"sources"`
	VisibleSources    int           `json:"@visible_sources"`
	Warnings          []string      `json:"warnings"`
}
