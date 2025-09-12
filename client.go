package pipl

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"time"
)

const (

	// version is the current version
	version = "v0.5.1"

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

// retryableHTTPClient implements HTTPInterface with retry logic using native Go
type retryableHTTPClient struct {
	client     HTTPInterface
	retryCount int
	backoff    backoffConfig
}

// backoffConfig holds the exponential backoff configuration
type backoffConfig struct {
	initialTimeout    time.Duration
	maxTimeout        time.Duration
	exponentFactor    float64
	maxJitterInterval time.Duration
}

// Do executes the HTTP request with retry logic
func (r *retryableHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if r.retryCount <= 0 {
		return r.client.Do(req)
	}

	var lastErr error
	for attempt := 0; attempt <= r.retryCount; attempt++ {
		resp, err := r.client.Do(req)
		if err == nil && resp != nil {
			// Success - check if we should retry based on status code
			if resp.StatusCode < 500 {
				return resp, nil
			}
			// Server error - close body and retry
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
		}

		lastErr = err

		// Don't sleep after the last attempt
		if attempt < r.retryCount {
			delay := r.calculateBackoff(attempt)
			time.Sleep(delay)
		}
	}

	return nil, fmt.Errorf("request failed after %d attempts: %w", r.retryCount+1, lastErr)
}

// calculateBackoff calculates the backoff delay with exponential backoff and jitter
func (r *retryableHTTPClient) calculateBackoff(attempt int) time.Duration {
	// Calculate exponential backoff
	delay := float64(r.backoff.initialTimeout) * math.Pow(r.backoff.exponentFactor, float64(attempt))

	// Apply maximum timeout limit
	if delay > float64(r.backoff.maxTimeout) {
		delay = float64(r.backoff.maxTimeout)
	}

	// Add jitter to prevent thundering herd
	if r.backoff.maxJitterInterval > 0 {
		jitterMax := big.NewInt(int64(r.backoff.maxJitterInterval))
		jitterVal, _ := rand.Int(rand.Reader, jitterMax)
		delay += float64(jitterVal.Int64())
	}

	return time.Duration(delay)
}

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

	// Create base HTTP client
	baseClient := &http.Client{
		Transport: clientDefaultTransport,
		Timeout:   c.options.httpOptions.RequestTimeout,
	}

	// Return client with or without retry logic
	if c.options.httpOptions.RequestRetryCount <= 0 {
		return baseClient
	}

	return &retryableHTTPClient{
		client:     baseClient, // baseClient implements HTTPInterface
		retryCount: c.options.httpOptions.RequestRetryCount,
		backoff: backoffConfig{
			initialTimeout:    c.options.httpOptions.BackOffInitialTimeout,
			maxTimeout:        c.options.httpOptions.BackOffMaxTimeout,
			exponentFactor:    c.options.httpOptions.BackOffExponentFactor,
			maxJitterInterval: c.options.httpOptions.BackOffMaximumJitterInterval,
		},
	}
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
