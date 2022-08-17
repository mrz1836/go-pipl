package pipl

// Package global constants and configuration
const (
	// searchAPIEndpoint is where we POST queries to
	searchAPIEndpoint string = "https://api.pipl.com/search/"

	// thumbnailEndpoint is where the image thumbnails are located
	thumbnailEndpoint string = "https://thumb.pipl.com/image"

	// ShowSourcesNone specifies that we don't need the source info back with search results
	ShowSourcesNone SourceLevel = valueFalse

	// ShowSourcesAll specifies that we want all source info back with our search results
	ShowSourcesAll SourceLevel = "all"

	// ShowSourcesMatching specifies that we want source info that corresponds to data that satisfies our match requirements
	ShowSourcesMatching SourceLevel = valueTrue

	// MatchRequirementsNone specifies that we don't have any match requirements for this search
	MatchRequirementsNone MatchRequirements = ""

	// MatchRequirementsEmail specifies that we want to match on this field
	MatchRequirementsEmail MatchRequirements = "email"

	// MatchRequirementsPhone specifies that we want to match on this field
	MatchRequirementsPhone MatchRequirements = "phone"

	// MatchRequirementsEmailAndPhone specifies that we want to match on this field
	MatchRequirementsEmailAndPhone MatchRequirements = "email and phone"

	// MatchRequirementsEmailAndName specifies that we want to match on this field
	MatchRequirementsEmailAndName MatchRequirements = "email and name"

	// MatchRequirementsEmailOrPhone specifies that we want to match on this field
	MatchRequirementsEmailOrPhone MatchRequirements = "email or phone"

	// https://docs.pipl.com/reference#match-criteria

	// MinimumProbability is the score for probability
	MinimumProbability = 0.9

	// MinimumMatch is the minimum for a match
	MinimumMatch = 0.0

	// SourceCategoryRequirementsNone specifies that we don't require any specific sources in our results.
	SourceCategoryRequirementsNone SourceCategoryRequirements = ""

	// SourceCategoryRequirementsProfessionalAndBusiness is used for: match_requirements=(emails and jobs)
	SourceCategoryRequirementsProfessionalAndBusiness SourceCategoryRequirements = "professional_and_business"

	// ThumbnailHeight is the default height
	ThumbnailHeight int = 250

	// ThumbnailWidth is the default width
	ThumbnailWidth int = 250

	// DefaultCountry is the default country for address
	DefaultCountry string = "US"

	// DefaultLanguage is the default language
	DefaultLanguage string = "en"

	// DefaultDisplayLanguage is the default display language
	DefaultDisplayLanguage string = "en_US"

	// Internal field for HTTP request
	fieldAPIKey                     = "key"
	fieldHideSponsored              = "hide_sponsored"
	fieldInferPersons               = "infer_persons"
	fieldLiveFeeds                  = "live_feeds"
	fieldMatchRequirements          = "match_requirements"
	fieldMinimumMatch               = "minimum_match"
	fieldPerson                     = "person"
	fieldPretty                     = "pretty"
	fieldSearchPointer              = "search_pointer"
	fieldShowSources                = "show_sources"
	fieldSourceCategoryRequirements = "source_category_requirements"
	fieldTopMatch                   = "top_match"
	valueFalse                      = "false"
	valueTrue                       = "true"
)

// SourceLevel is used internally to represent the possible values
// for show_sources in queries to be submitted: {"all", "matching"/"true", "false"}
type SourceLevel string

// MatchRequirements specifies the conditions for a successful person match in our search.
// This is useful for saving money with the Pipl API, as you only need to pay for the
// data you wanted back. If your search results didn't satisfy the match requirements, then
// no data is returned, and you don't pay.
type MatchRequirements string

// SourceCategoryRequirements specifies the data categories that must be included in
// results for a successful match. If there is no data from the requested categories,
// then the results returned are empty, and you're not charged.
type SourceCategoryRequirements string

// SearchParameters holds options that can affect data returned by a search.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#configuration-parameters
type SearchParameters struct {

	// ShowSources specifies the level of sources info to return with search results, one of {ShowSourcesMatching, ShowSourcesAll, ShowSourcesNone}
	ShowSources SourceLevel

	// MatchRequirements specifies the criteria for a successful Person match.
	// Results that don't fit your match requirements are discarded. If the remaining
	// search results would be empty, you are not charged for the query.
	MatchRequirements MatchRequirements

	// SourceCategoryRequirements specifies the data categories that must be included in
	// results for a successful match. If there is no data from the requested categories,
	// then the results returned are empty, and you're not charged.
	SourceCategoryRequirements SourceCategoryRequirements

	// MinimumProbability is the minimum acceptable probability for inferred data
	MinimumProbability float32

	// MinimumMatch specifies the minimum match confidence for a possible person to be returned in search results
	MinimumMatch float32

	// InferPersons specifies whether the Pipl should return results inferred by statistical analysis
	InferPersons bool

	// HideSponsored specifies whether to omit sponsored data from search results
	HideSponsored bool

	// LiveFeeds specifies whether to use live data sources
	LiveFeeds bool

	// Returns the best high ranking match to your search. API will return either a Person (when high scoring profile is found) or a No Match
	TopMatch bool

	// Returns formatted response in pretty mode for JSON (default is false)
	Pretty bool
}

// ThumbnailSettings is for the thumbnail url settings to be automatically returned
// if any images are found and meet the criteria
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Example: http://thumb.pipl.com/image?height=250&width=250&favicon=true&zoom_face=true&tokens=FIRST_TOKEN,SECOND_TOKEN
type ThumbnailSettings struct {

	// URL is the thumbnail url
	URL string

	// Height of the image
	Height int

	// Width of the image
	Width int

	// Enabled (detects images, automatically adds thumbnail urls)
	Enabled bool

	// Favicon if the icon should be shown or not
	Favicon bool

	// ZoomFace is whether to enable face zoom.
	ZoomFace bool
}

// GUID is a unique format (but is just a string internally, since there's currently
// nothing all that fancy done with GUIDs). Additional guid-handling code may be
// added at a later date if needed.
type GUID string

// Name fields collectively define a possible name for a given person.
// If a search did not return information for a given field, it will be empty.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#name
type Name struct {
	Display    string `json:"display,omitempty"`
	First      string `json:"first,omitempty"`
	Last       string `json:"last,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Middle     string `json:"middle,omitempty"`
	Prefix     string `json:"prefix,omitempty"`
	Raw        string `json:"raw,omitempty"`
	Suffix     string `json:"suffix,omitempty"`
	Type       string `json:"@type,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
	Current    bool   `json:"@current,omitempty"`
}

// Address fields collectively define a possible address for a given person
// If a search did not return information for a given field, it will be empty.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#address
type Address struct {
	Apartment  string `json:"apartment,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	Display    string `json:"display,omitempty"`
	House      string `json:"house,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	POBox      string `json:"po_box,omitempty"`
	Raw        string `json:"raw,omitempty"`
	State      string `json:"state,omitempty"`
	Street     string `json:"street,omitempty"`
	Type       string `json:"@type,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
	ZipCode    string `json:"zip_code,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
}

// Phone fields collectively define a possible phone number for a given person
// If a search did not return information for a given field, it will be empty.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#phone
type Phone struct {
	Display              string `json:"display,omitempty"`
	DisplayInternational string `json:"display_international,omitempty"`
	LastSeen             string `json:"@last_seen,omitempty"`
	Raw                  string `json:"raw,omitempty"`
	Type                 string `json:"@type,omitempty"`
	ValidSince           string `json:"@valid_since,omitempty"`
	CountryCode          int    `json:"country_code,omitempty"`
	Extension            int    `json:"extension,omitempty"`
	Number               int64  `json:"number,omitempty"`
	Current              bool   `json:"@current,omitempty"`
	Inferred             bool   `json:"@inferred,omitempty"`
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
// Note: all lowercase, case-sensitive
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
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#date-range
type DateRange struct {
	End        string `json:"end,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Start      string `json:"start,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
}

// DateOfBirth specifies a possible DOB for a person.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#date-of-birth
type DateOfBirth struct {
	DateRange  DateRange `json:"date_range,omitempty"`
	Display    string    `json:"display,omitempty"`
	LastSeen   string    `json:"@last_seen,omitempty"`
	ValidSince string    `json:"@valid_since,omitempty"`
	Current    bool      `json:"@current,omitempty"`
	Inferred   bool      `json:"@inferred,omitempty"`
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
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#job
type Job struct {
	DateRange    DateRange `json:"date_range,omitempty"`
	Display      string    `json:"display,omitempty"`
	Industry     string    `json:"industry,omitempty"`
	LastSeen     string    `json:"@last_seen,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Title        string    `json:"title,omitempty"`
	ValidSince   string    `json:"@valid_since,omitempty"`
	Current      bool      `json:"@current,omitempty"`
	Inferred     bool      `json:"@inferred,omitempty"`
}

// Education specifies a possible
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#education
type Education struct {
	DateRange  DateRange `json:"date_range,omitempty"`
	Degree     string    `json:"degree,omitempty"`
	Display    string    `json:"display,omitempty"`
	LastSeen   string    `json:"@last_seen,omitempty"`
	School     string    `json:"school,omitempty"`
	ValidSince string    `json:"@valid_since,omitempty"`
	Current    bool      `json:"@current,omitempty"`
	Inferred   bool      `json:"@inferred,omitempty"`
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
// Note: all lowercase, case-sensitive
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
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#language
type Language struct {
	Display    string `json:"display,omitempty"`
	Language   string `json:"language,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Region     string `json:"region,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
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
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#relationship
type Relationship struct {
	DateOfBirth     DateOfBirth     `json:"dob,omitempty"`
	Gender          Gender          `json:"gender,omitempty"`
	Addresses       []Address       `json:"addresses,omitempty"`
	Educations      []Education     `json:"educations,omitempty"`
	Emails          []Email         `json:"emails,omitempty"`
	Ethnicities     []Ethnicity     `json:"ethnicities,omitempty"`
	Images          []Image         `json:"images,omitempty"`
	Jobs            []Job           `json:"jobs,omitempty"`
	Languages       []Language      `json:"languages,omitempty"`
	Names           []Name          `json:"names,omitempty"`
	OriginCountries []OriginCountry `json:"origin_countries,omitempty"`
	Phones          []Phone         `json:"phones,omitempty"`
	Relationships   []Relationship  `json:"relationships,omitempty"`
	UserIDs         []UserID        `json:"user_ids,omitempty"`
	Usernames       []Username      `json:"usernames,omitempty"`
	LastSeen        string          `json:"@last_seen,omitempty"`
	Subtype         string          `json:"@subtype,omitempty"`
	Type            string          `json:"@type,omitempty"`
	ValidSince      string          `json:"@valid_since,omitempty"`
	Current         bool            `json:"@current,omitempty"`
	Inferred        bool            `json:"@inferred,omitempty"`
}

// URL contains information about a URL that is closely associated with a given person.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#url
type URL struct {
	Category   string `json:"@category,omitempty"`
	Domain     string `json:"@domain,omitempty"`
	LastSeen   string `json:"@last_seen,omitempty"`
	Name       string `json:"@name,omitempty"`
	SourceID   string `json:"@source_id,omitempty"`
	URL        string `json:"url,omitempty"`
	ValidSince string `json:"@valid_since,omitempty"`
	Current    bool   `json:"@current,omitempty"`
	Inferred   bool   `json:"@inferred,omitempty"`
}

// Tag contains content classification information
//
// Source: https://docs.pipl.com/reference#tag
type Tag struct {
	Classification string `json:"@classification,omitempty"`
	Content        string `json:"content,omitempty"`
}

// Person contains all the information pertaining to a possible person match,
// including potential multiples of basic fields (names, emails, jobs, etc.).
// The Match field represents the confidence of a particular person match, as a
// float: 0 <= Match <= 1. More potential matches returned in a search decreases
// the overall confidence of all matches.
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#person
type Person struct {
	Addresses       []Address       `json:"addresses,omitempty"`
	Educations      []Education     `json:"educations,omitempty"`
	Emails          []Email         `json:"emails,omitempty"`
	Ethnicities     []Ethnicity     `json:"ethnicities,omitempty"`
	Images          []Image         `json:"images,omitempty"`
	Jobs            []Job           `json:"jobs,omitempty"`
	Languages       []Language      `json:"languages,omitempty"`
	Names           []Name          `json:"names,omitempty"`
	OriginCountries []OriginCountry `json:"origin_countries,omitempty"`
	Phones          []Phone         `json:"phones,omitempty"`
	Relationships   []Relationship  `json:"relationships,omitempty"`
	URLs            []URL           `json:"urls,omitempty"`
	UserIDs         []UserID        `json:"user_ids,omitempty"`
	Usernames       []Username      `json:"usernames,omitempty"`
	ID              GUID            `json:"@id,omitempty"`
	SearchPointer   string          `json:"@search_pointer,omitempty"`
	DateOfBirth     *DateOfBirth    `json:"dob,omitempty"`
	Gender          *Gender         `json:"gender,omitempty"`
	Match           float32         `json:"@match,omitempty"`
	Inferred        bool            `json:"@inferred,omitempty"`
}

// Source contains all the information for a given person, gathered from a
// single source. The source structure contains information about the name,
// domain, category, and source URL (amongst other fields).
//
// DO NOT CHANGE ORDER - Optimized for memory (malign)
//
// Source: https://docs.pipl.com/reference#source
type Source struct {
	DateOfBirth     DateOfBirth     `json:"dob"`
	Gender          Gender          `json:"gender"`
	Addresses       []Address       `json:"addresses"`
	Educations      []Education     `json:"educations"`
	Emails          []Email         `json:"emails"`
	Ethnicities     []Ethnicity     `json:"ethnicities"`
	Jobs            []Job           `json:"jobs"`
	Languages       []Language      `json:"languages"`
	Names           []Name          `json:"names"`
	OriginCountries []OriginCountry `json:"origin_countries"`
	Phones          []Phone         `json:"phones"`
	Relationships   []Relationship  `json:"relationships"`
	Tags            []Tag           `json:"tags"`
	URLs            []URL           `json:"urls"`
	UserIDs         []UserID        `json:"user_ids"`
	Usernames       []Username      `json:"usernames"`
	Category        string          `json:"@category"`
	Domain          string          `json:"@domain"`
	ID              string          `json:"@id"`
	Name            string          `json:"@name"`
	OriginURL       string          `json:"@origin_url"`
	PersonID        GUID            `json:"@person_id"`
	Match           float32         `json:"@match"`
	Premium         bool            `json:"@premium"`
	Sponsored       bool            `json:"@sponsored"`
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
	TopMatch          bool          `json:"top_match"`
	VisibleSources    int           `json:"@visible_sources"`
	Warnings          []string      `json:"warnings"`
}
