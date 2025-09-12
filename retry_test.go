package pipl

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Static errors for testing
var (
	ErrMaxCallsReached   = errors.New("max calls reached")
	ErrNetworkFailure    = errors.New("network error")
	ErrPersistentNetwork = errors.New("persistent network error")
)

// mockRetryClient implements HTTPInterface for testing retry behavior
type mockRetryClient struct {
	errors      []error
	callCount   int
	maxCalls    int
	statusCodes []int
}

// Do simulates HTTP requests with configurable responses
func (m *mockRetryClient) Do(_ *http.Request) (*http.Response, error) {
	if m.callCount >= m.maxCalls {
		return nil, ErrMaxCallsReached
	}

	currentCall := m.callCount
	m.callCount++

	// Return error if configured
	if m.errors != nil && currentCall < len(m.errors) && m.errors[currentCall] != nil {
		return nil, m.errors[currentCall]
	}

	// Return response with specific status code
	if m.statusCodes != nil && currentCall < len(m.statusCodes) {
		resp := &http.Response{
			StatusCode: m.statusCodes[currentCall],
			Body:       http.NoBody,
		}
		return resp, nil
	}

	// Return successful response by default
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       http.NoBody,
	}
	return resp, nil
}

// TestRetryableHTTPClient_NoRetry tests client behavior without retries
func TestRetryableHTTPClient_NoRetry(t *testing.T) {
	t.Parallel()

	mock := &mockRetryClient{maxCalls: 1}
	client := &retryableHTTPClient{
		client:     mock, // Mock implements HTTPInterface via Do method
		retryCount: 0,    // No retries
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 1, mock.callCount)
	if resp.Body != nil {
		_ = resp.Body.Close()
	}
}

// TestRetryableHTTPClient_SuccessFirstAttempt tests successful first attempt
func TestRetryableHTTPClient_SuccessFirstAttempt(t *testing.T) {
	t.Parallel()

	mock := &mockRetryClient{maxCalls: 3}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 2,
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 1, mock.callCount)
	if resp.Body != nil {
		_ = resp.Body.Close()
	} // Only one attempt needed
}

// TestRetryableHTTPClient_RetryOn500 tests retry behavior on 500 errors
func TestRetryableHTTPClient_RetryOn500(t *testing.T) {
	t.Parallel()

	mock := &mockRetryClient{
		maxCalls:    3,
		statusCodes: []int{500, 502, 200}, // First two fail, third succeeds
	}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 2,
		backoff: backoffConfig{
			initialTimeout:    1 * time.Millisecond,
			maxTimeout:        10 * time.Millisecond,
			exponentFactor:    2.0,
			maxJitterInterval: 1 * time.Millisecond,
		},
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 3, mock.callCount) // Three attempts total
	if resp.Body != nil {
		_ = resp.Body.Close()
	}
}

// TestRetryableHTTPClient_NoRetryOn400 tests no retry on 4xx errors
func TestRetryableHTTPClient_NoRetryOn400(t *testing.T) {
	t.Parallel()

	mock := &mockRetryClient{
		maxCalls:    3,
		statusCodes: []int{400, 200}, // 400 should not be retried
	}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 2,
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, 1, mock.callCount) // Only one attempt, no retry on 4xx
	if resp.Body != nil {
		_ = resp.Body.Close()
	}
}

// TestRetryableHTTPClient_RetryOnNetworkError tests retry behavior on network errors
func TestRetryableHTTPClient_RetryOnNetworkError(t *testing.T) {
	t.Parallel()

	networkError := ErrNetworkFailure
	mock := &mockRetryClient{
		maxCalls: 3,
		errors:   []error{networkError, networkError, nil}, // First two fail, third succeeds
	}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 2,
		backoff: backoffConfig{
			initialTimeout:    1 * time.Millisecond,
			maxTimeout:        10 * time.Millisecond,
			exponentFactor:    2.0,
			maxJitterInterval: 1 * time.Millisecond,
		},
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 3, mock.callCount)
	if resp.Body != nil {
		_ = resp.Body.Close()
	}
}

// TestRetryableHTTPClient_ExhaustRetries tests behavior when all retries are exhausted
func TestRetryableHTTPClient_ExhaustRetries(t *testing.T) {
	t.Parallel()

	networkError := ErrPersistentNetwork
	mock := &mockRetryClient{
		maxCalls: 3,
		errors:   []error{networkError, networkError, networkError},
	}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 2,
		backoff: backoffConfig{
			initialTimeout:    1 * time.Millisecond,
			maxTimeout:        10 * time.Millisecond,
			exponentFactor:    2.0,
			maxJitterInterval: 1 * time.Millisecond,
		},
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)
	resp, err := client.Do(req) //nolint:bodyclose // Expected to return nil response when retries exhausted

	require.Error(t, err)
	require.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed after 3 attempts")
	assert.Equal(t, 3, mock.callCount) // All attempts used
	// No response to close when all retries are exhausted
}

// TestCalculateBackoff tests the exponential backoff calculation
func TestCalculateBackoff(t *testing.T) {
	t.Parallel()

	client := &retryableHTTPClient{
		backoff: backoffConfig{
			initialTimeout:    2 * time.Millisecond,
			maxTimeout:        100 * time.Millisecond,
			exponentFactor:    2.0,
			maxJitterInterval: 0, // No jitter for predictable testing
		},
	}

	tests := []struct {
		name     string
		attempt  int
		expected time.Duration
	}{
		{"first retry", 0, 2 * time.Millisecond},
		{"second retry", 1, 4 * time.Millisecond},
		{"third retry", 2, 8 * time.Millisecond},
		{"fourth retry", 3, 16 * time.Millisecond},
		{"max timeout", 10, 100 * time.Millisecond}, // Should be capped at max
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delay := client.calculateBackoff(tt.attempt)
			assert.Equal(t, tt.expected, delay)
		})
	}
}

// TestCalculateBackoffWithJitter tests backoff with jitter
func TestCalculateBackoffWithJitter(t *testing.T) {
	t.Parallel()

	client := &retryableHTTPClient{
		backoff: backoffConfig{
			initialTimeout:    2 * time.Millisecond,
			maxTimeout:        100 * time.Millisecond,
			exponentFactor:    2.0,
			maxJitterInterval: 5 * time.Millisecond,
		},
	}

	// Test multiple times to ensure jitter is applied
	delays := make([]time.Duration, 10)
	for i := 0; i < 10; i++ {
		delays[i] = client.calculateBackoff(0)
	}

	// All delays should be >= base delay
	baseDelay := 2 * time.Millisecond
	for _, delay := range delays {
		assert.GreaterOrEqual(t, delay, baseDelay, "delay should be at least base delay")
		assert.LessOrEqual(t, delay, baseDelay+5*time.Millisecond, "delay should not exceed base + max jitter")
	}
}

// TestCreateDefaultHTTPClient_WithRetry tests client creation with retry enabled
func TestCreateDefaultHTTPClient_WithRetry(t *testing.T) {
	t.Parallel()

	opts := DefaultHTTPOptions()
	opts.RequestRetryCount = 3

	client := &Client{
		options: &ClientOptions{
			httpOptions: opts,
		},
	}

	httpClient := createDefaultHTTPClient(client)
	require.NotNil(t, httpClient)

	// Should return retryableHTTPClient
	retryClient, ok := httpClient.(*retryableHTTPClient)
	require.True(t, ok, "should return retryableHTTPClient when retry count > 0")
	assert.Equal(t, 3, retryClient.retryCount)
}

// TestCreateDefaultHTTPClient_NoRetry tests client creation without retry
func TestCreateDefaultHTTPClient_NoRetry(t *testing.T) {
	t.Parallel()

	opts := DefaultHTTPOptions()
	opts.RequestRetryCount = 0

	client := &Client{
		options: &ClientOptions{
			httpOptions: opts,
		},
	}

	httpClient := createDefaultHTTPClient(client)
	require.NotNil(t, httpClient)

	// Should return standard http.Client
	_, ok := httpClient.(*http.Client)
	assert.True(t, ok, "should return http.Client when retry count <= 0")
}

// BenchmarkRetryableHTTPClient_NoRetry benchmarks the client without retries
func BenchmarkRetryableHTTPClient_NoRetry(b *testing.B) {
	mock := &mockRetryClient{maxCalls: b.N}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 0,
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := client.Do(req)
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}
}

// BenchmarkRetryableHTTPClient_WithRetry benchmarks the client with retries
func BenchmarkRetryableHTTPClient_WithRetry(b *testing.B) {
	mock := &mockRetryClient{maxCalls: b.N}
	client := &retryableHTTPClient{
		client:     mock,
		retryCount: 2,
		backoff: backoffConfig{
			initialTimeout:    1 * time.Millisecond,
			maxTimeout:        10 * time.Millisecond,
			exponentFactor:    2.0,
			maxJitterInterval: 1 * time.Millisecond,
		},
	}

	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://example.com", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := client.Do(req)
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}
}
