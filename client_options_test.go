package pipl

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDefaultHTTPOptions will test the method DefaultHTTPOptions()
func TestDefaultHTTPOptions(t *testing.T) {
	t.Parallel()

	options := DefaultHTTPOptions()

	assert.Equal(t, 10, options.TransportMaxIdleConnections)
	assert.Equal(t, 2*time.Millisecond, options.BackOffInitialTimeout)
	assert.Equal(t, 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	assert.Equal(t, 2, options.RequestRetryCount)
	assert.InEpsilon(t, 2.0, options.BackOffExponentFactor, 0.001)
	assert.Equal(t, 20*time.Second, options.DialerKeepAlive)
	assert.Equal(t, 20*time.Second, options.TransportIdleTimeout)
	assert.Equal(t, 3*time.Second, options.TransportExpectContinueTimeout)
	assert.Equal(t, 30*time.Second, options.RequestTimeout)
	assert.Equal(t, 5*time.Second, options.DialerTimeout)
	assert.Equal(t, 5*time.Second, options.TransportTLSHandshakeTimeout)
}

// TestDefaultSearchOptions will test the method DefaultSearchOptions()
func TestDefaultSearchOptions(t *testing.T) {
	t.Parallel()

	options := DefaultSearchOptions()
	require.NotNil(t, options.Search)
	require.NotNil(t, options.Thumbnail)

	assert.False(t, options.Search.InferPersons)
	assert.False(t, options.Search.TopMatch)
	assert.Equal(t, MatchRequirementsNone, options.Search.MatchRequirements)
	assert.InDelta(t, float32(MinimumMatch), options.Search.MinimumMatch, 0.001)
	assert.InEpsilon(t, float32(MinimumProbability), options.Search.MinimumProbability, 0.001)
	assert.Equal(t, ShowSourcesAll, options.Search.ShowSources)
	assert.Equal(t, SourceCategoryRequirementsNone, options.Search.SourceCategoryRequirements)
	assert.False(t, options.Search.HideSponsored)
	assert.True(t, options.Search.LiveFeeds)
	assert.False(t, options.Search.Pretty)

	assert.False(t, options.Thumbnail.Enabled)
	assert.False(t, options.Thumbnail.Favicon)
	assert.False(t, options.Thumbnail.ZoomFace)
	assert.Equal(t, ThumbnailHeight, options.Thumbnail.Height)
	assert.Equal(t, thumbnailEndpoint, options.Thumbnail.URL)
	assert.Equal(t, ThumbnailWidth, options.Thumbnail.Width)
}

// TestWithAPIKey will test the method WithAPIKey()
func TestWithAPIKey(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithAPIKey("")
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying empty", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithAPIKey("")
		opt(options)
		assert.Empty(t, options.apiKey)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithAPIKey(testKey)
		opt(options)
		assert.Equal(t, testKey, options.apiKey)
	})
}

// TestWithHTTPClient will test the method WithHTTPClient()
func TestWithHTTPClient(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithHTTPClient(nil)
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying nil", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithHTTPClient(nil)
		opt(options)
		assert.Nil(t, options.httpClient)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		customClient := &http.Client{}
		opt := WithHTTPClient(customClient)
		opt(options)
		assert.Equal(t, customClient, options.httpClient)
	})
}

// TestWithHTTPOptions will test the method WithHTTPOptions()
func TestWithHTTPOptions(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithHTTPOptions(nil)
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying nil", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithHTTPOptions(nil)
		opt(options)
		assert.Nil(t, options.httpOptions)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		customHTTPOpts := DefaultHTTPOptions()
		customHTTPOpts.RequestRetryCount = 3
		opt := WithHTTPOptions(customHTTPOpts)
		opt(options)
		assert.Equal(t, customHTTPOpts, options.httpOptions)
		assert.Equal(t, 3, options.httpOptions.RequestRetryCount)
	})
}

// TestWithSearchOptions will test the method WithSearchOptions()
func TestWithSearchOptions(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithSearchOptions(nil)
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying nil", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithSearchOptions(nil)
		opt(options)
		assert.Nil(t, options.searchOptions)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		customSearchOpts := DefaultSearchOptions()
		customSearchOpts.Search.MinimumProbability = 0.6
		opt := WithSearchOptions(customSearchOpts)
		opt(options)
		assert.Equal(t, customSearchOpts, options.searchOptions)
		assert.InEpsilon(t, float32(0.6), options.searchOptions.Search.MinimumProbability, 0.001)
	})
}

// TestWithUserAgent will test the method WithUserAgent()
func TestWithUserAgent(t *testing.T) {
	t.Parallel()

	t.Run("check type", func(t *testing.T) {
		opt := WithUserAgent("")
		assert.IsType(t, *new(ClientOps), opt)
	})

	t.Run("test applying empty", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithUserAgent("")
		opt(options)
		assert.Empty(t, options.userAgent)
	})

	t.Run("test applying option", func(t *testing.T) {
		options := &ClientOptions{}
		opt := WithUserAgent(testUserAgent)
		opt(options)
		assert.Equal(t, testUserAgent, options.userAgent)
	})
}
