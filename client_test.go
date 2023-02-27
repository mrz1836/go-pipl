package pipl

import (
	"fmt"
	"net/http"
	"testing"

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
