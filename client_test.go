package pipl

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewClient will test the method NewClient()
func TestNewClient(t *testing.T) {
	t.Parallel()

	t.Run("basic client", func(t *testing.T) {
		c := NewClient()
		require.NotNil(t, c)
		client := c.HTTPClient()
		require.NotNil(t, client)
		ua := c.UserAgent()
		assert.Equal(t, defaultUserAgent, ua)
	})

	t.Run("custom user agent", func(t *testing.T) {
		c := NewClient(WithUserAgent("custom-agent"))
		require.NotNil(t, c)
		ua := c.UserAgent()
		assert.Equal(t, "custom-agent", ua)
	})

	t.Run("custom http client", func(t *testing.T) {
		hc := &http.Client{}
		c := NewClient(WithHTTPClient(hc))
		require.NotNil(t, c)
		assert.Equal(t, hc, c.HTTPClient())
	})

	t.Run("custom http options, no retry", func(t *testing.T) {
		opts := DefaultHTTPOptions()
		opts.RequestRetryCount = 0
		c := NewClient(WithHTTPOptions(opts))
		require.NotNil(t, c)
		require.NotNil(t, c.HTTPClient())

		// Should return standard http.Client
		_, ok := c.HTTPClient().(*http.Client)
		assert.True(t, ok, "should return http.Client when retry count is 0")
	})

	t.Run("custom http options, with retry", func(t *testing.T) {
		opts := DefaultHTTPOptions()
		opts.RequestRetryCount = 3
		opts.BackOffInitialTimeout = 5 * time.Millisecond
		opts.BackOffMaxTimeout = 50 * time.Millisecond
		opts.BackOffExponentFactor = 2.5
		opts.BackOffMaximumJitterInterval = 10 * time.Millisecond

		c := NewClient(WithHTTPOptions(opts))
		require.NotNil(t, c)
		require.NotNil(t, c.HTTPClient())

		// Should return retryableHTTPClient
		retryClient, ok := c.HTTPClient().(*retryableHTTPClient)
		require.True(t, ok, "should return retryableHTTPClient when retry count > 0")
		assert.Equal(t, 3, retryClient.retryCount)
		assert.Equal(t, opts.BackOffInitialTimeout, retryClient.backoff.initialTimeout)
		assert.Equal(t, opts.BackOffMaxTimeout, retryClient.backoff.maxTimeout)
		assert.InDelta(t, opts.BackOffExponentFactor, retryClient.backoff.exponentFactor, 0.001)
		assert.Equal(t, opts.BackOffMaximumJitterInterval, retryClient.backoff.maxJitterInterval)
	})
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client := NewClient(WithAPIKey(testKey))
	fmt.Println(client.UserAgent())
	// Output:go-pipl: v0.5.1
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClient(WithAPIKey(testKey))
	}
}
