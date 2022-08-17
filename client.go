package pipl

import (
	"net"
	"net/http"
	"time"

	"github.com/gojektech/heimdall/v6"
	"github.com/gojektech/heimdall/v6/httpclient"
)

const (

	// version is the current version
	version = "v0.5.0"

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-pipl: " + version
)

type (
	// Client is the client configuration and options
	Client struct {
		options *ClientOptions // Options are all the default settings / configuration
	}

	// ClientOptions holds all the configuration for client requests and default resources
	ClientOptions struct {
		apiKey        string         // The user's API key for NOWNode API
		httpClient    HTTPInterface  // HTTP client interface
		httpOptions   *HTTPOptions   // Options for the HTTP client
		searchOptions *SearchOptions // contains search options
		userAgent     string         // User agent for all outgoing requests
	}

	// HTTPOptions holds all the configuration for the HTTP client
	HTTPOptions struct {
		BackOffExponentFactor          float64       `json:"back_off_exponent_factor"`
		BackOffInitialTimeout          time.Duration `json:"back_off_initial_timeout"`
		BackOffMaximumJitterInterval   time.Duration `json:"back_off_maximum_jitter_interval"`
		BackOffMaxTimeout              time.Duration `json:"back_off_max_timeout"`
		DialerKeepAlive                time.Duration `json:"dialer_keep_alive"`
		DialerTimeout                  time.Duration `json:"dialer_timeout"`
		RequestRetryCount              int           `json:"request_retry_count"`
		RequestTimeout                 time.Duration `json:"request_timeout"`
		TransportExpectContinueTimeout time.Duration `json:"transport_expect_continue_timeout"`
		TransportIdleTimeout           time.Duration `json:"transport_idle_timeout"`
		TransportMaxIdleConnections    int           `json:"transport_max_idle_connections"`
		TransportTLSHandshakeTimeout   time.Duration `json:"transport_tls_handshake_timeout"`
	}

	// HTTPInterface is used for the HTTP client
	HTTPInterface interface {
		Do(req *http.Request) (*http.Response, error)
	}

	// SearchOptions are custom search options for conducting searches
	SearchOptions struct {
		Search    *SearchParameters  // contains the search parameters that are submitted with your query, which may affect the data returned
		Thumbnail *ThumbnailSettings // is for the thumbnail url settings
	}
)

// createDefaultHTTPClient will create a default HTTP client interface
func createDefaultHTTPClient(c *Client) HTTPInterface {

	// dial is the net dialer for clientDefaultTransport
	dial := &net.Dialer{
		KeepAlive: c.options.httpOptions.DialerKeepAlive,
		Timeout:   c.options.httpOptions.DialerTimeout,
	}

	// clientDefaultTransport is the default transport struct for the HTTP client
	clientDefaultTransport := &http.Transport{
		DialContext:           dial.DialContext,
		ExpectContinueTimeout: c.options.httpOptions.TransportExpectContinueTimeout,
		IdleConnTimeout:       c.options.httpOptions.TransportIdleTimeout,
		MaxIdleConns:          c.options.httpOptions.TransportMaxIdleConnections,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   c.options.httpOptions.TransportTLSHandshakeTimeout,
	}

	// Determine the strategy for the http client (no retry enabled)
	if c.options.httpOptions.RequestRetryCount <= 0 {
		return httpclient.NewClient(
			httpclient.WithHTTPTimeout(c.options.httpOptions.RequestTimeout),
			httpclient.WithHTTPClient(&http.Client{
				Transport: clientDefaultTransport,
				Timeout:   c.options.httpOptions.RequestTimeout,
			}),
		)
	}

	// Create exponential back-off
	backOff := heimdall.NewExponentialBackoff(
		c.options.httpOptions.BackOffInitialTimeout,
		c.options.httpOptions.BackOffMaxTimeout,
		c.options.httpOptions.BackOffExponentFactor,
		c.options.httpOptions.BackOffMaximumJitterInterval,
	)

	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(c.options.httpOptions.RequestTimeout),
		httpclient.WithRetrier(heimdall.NewRetrier(backOff)),
		httpclient.WithRetryCount(c.options.httpOptions.RequestRetryCount),
		httpclient.WithHTTPClient(&http.Client{
			Transport: clientDefaultTransport,
			Timeout:   c.options.httpOptions.RequestTimeout,
		}),
	)
}

// NewClient will make a new client with the provided options
func NewClient(opts ...ClientOps) ClientInterface {

	// Create a client with defaults
	c := &Client{
		options: &ClientOptions{
			httpOptions:   DefaultHTTPOptions(),
			searchOptions: DefaultSearchOptions(),
			userAgent:     defaultUserAgent,
		},
	}

	// Overwrite defaults with any set by user
	for _, opt := range opts {
		opt(c.options)
	}

	// Set a default http client if one does not exist
	if c.options.httpClient == nil {
		c.options.httpClient = createDefaultHTTPClient(c)
	}

	return c
}

// HTTPClient will return the current HTTP client
func (c *Client) HTTPClient() HTTPInterface {
	return c.options.httpClient
}

// UserAgent will return the current user agent
func (c *Client) UserAgent() string {
	return c.options.userAgent
}
