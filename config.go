package pipl

import (
	"net"
	"net/http"
	"time"
)

// Package global constants and configuration
const (
	// SearchAPIEndpoint is where we POST queries to
	SearchAPIEndpoint string = "https://api.pipl.com/search/"

	// ThumbnailEndpoint is where the image thumbnails are located
	ThumbnailEndpoint string = "https://thumb.pipl.com/image"

	// ShowSourcesNone specifies that we don't need source info back with search results
	ShowSourcesNone SourceLevel = "false"

	// ShowSourcesAll specifies that we want all source info back with our search results
	ShowSourcesAll SourceLevel = "all"

	// ShowSourcesMatching specifies that we want source info that corresponds to data that satisfies our match requirements
	ShowSourcesMatching SourceLevel = "true"

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

	// todo: finish adding match criteria - also make this flexible and easier to use
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

	// ConnectionExponentFactor backoff exponent factor
	ConnectionExponentFactor float64 = 2.0

	// ConnectionInitialTimeout initial timeout
	ConnectionInitialTimeout time.Duration = 2 * time.Millisecond

	// ConnectionMaximumJitterInterval jitter interval
	ConnectionMaximumJitterInterval time.Duration = 2 * time.Millisecond

	// ConnectionMaxTimeout max timeout
	ConnectionMaxTimeout time.Duration = 1000 * time.Millisecond

	// ConnectionRetryCount retry count
	ConnectionRetryCount int = 3

	// ConnectionWithHTTPTimeout with http timeout
	ConnectionWithHTTPTimeout time.Duration = 1 * time.Second

	// ConnectionTLSHandshakeTimeout tls handshake timeout
	ConnectionTLSHandshakeTimeout time.Duration = 5 * time.Second

	// ConnectionMaxIdleConnections max idle http connections
	ConnectionMaxIdleConnections int = 128

	// ConnectionIdleTimeout idle connection timeout
	ConnectionIdleTimeout time.Duration = 30 * time.Second

	// ConnectionExpectContinueTimeout expect continue timeout
	ConnectionExpectContinueTimeout time.Duration = 3 * time.Second

	// ConnectionDialerTimeout dialer timeout
	ConnectionDialerTimeout time.Duration = 5 * time.Second

	// ConnectionDialerKeepAlive keep alive
	ConnectionDialerKeepAlive time.Duration = 30 * time.Second

	// DefaultUserAgent is the default user agent for all pipl requests
	DefaultUserAgent string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36"
)

// HTTP and Dialer connection variables
var (
	// _Dialer net dialer for ClientDefaultTransport
	_Dialer = &net.Dialer{
		KeepAlive: ConnectionDialerKeepAlive,
		Timeout:   ConnectionDialerTimeout,
	}

	// ClientDefaultTransport is the default transport struct for the HTTP client
	ClientDefaultTransport = &http.Transport{
		DialContext:           _Dialer.DialContext,
		ExpectContinueTimeout: ConnectionExpectContinueTimeout,
		IdleConnTimeout:       ConnectionIdleTimeout,
		MaxIdleConns:          ConnectionMaxIdleConnections,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   ConnectionTLSHandshakeTimeout,
	}
)
