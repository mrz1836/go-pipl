package pipl

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Testing variables
var testImage = "https://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg"
var testThumbnailToken = "AE2861B242686E7BD0CB4D9049298EB7D18FEF66D950E8AB78BCD3F484345CE74536C19A85D0BA3D32DC9E7D1878CD4D341254E7AD129255C6983E6E154C4530A0DAAF665EA325FC0206F8B1D7E0B6B7AD9EBF71FCF610D57D"

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	client, err := NewClient("1234567890", nil)
	require.NoError(t, err)
	require.Equal(t, "1234567890", client.Parameters.Search.apiKey)
	require.Equal(t, float32(MinimumMatch), client.Parameters.Search.MinimumMatch)
	require.Equal(t, float32(MinimumProbability), client.Parameters.Search.MinimumProbability)

	// todo: test changing these values in the SearchParameters
}

// ExampleNewClient example using NewClient()
func ExampleNewClient() {
	client, _ := NewClient("1234567890", nil)
	fmt.Println(client.Parameters.Search.apiKey)
	// Output:1234567890
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewClient("1234567890", nil)
	}
}

// TestDefaultOptions tests setting ClientDefaultOptions()
func TestDefaultOptions(t *testing.T) {

	options := ClientDefaultOptions()
	require.Equal(t, 10, options.TransportMaxIdleConnections)
	require.Equal(t, 10*time.Millisecond, options.BackOffMaxTimeout)
	require.Equal(t, 10*time.Second, options.RequestTimeout)
	require.Equal(t, 2, options.RequestRetryCount)
	require.Equal(t, 2*time.Millisecond, options.BackOffInitialTimeout)
	require.Equal(t, 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	require.Equal(t, 2.0, options.BackOffExponentFactor)
	require.Equal(t, 20*time.Second, options.DialerKeepAlive)
	require.Equal(t, 20*time.Second, options.TransportIdleTimeout)
	require.Equal(t, 3*time.Second, options.TransportExpectContinueTimeout)
	require.Equal(t, 5*time.Second, options.DialerTimeout)
	require.Equal(t, 5*time.Second, options.TransportTLSHandshakeTimeout)
	require.Equal(t, defaultUserAgent, options.UserAgent)
}

// TestSearchMeetsMinimumCriteria test the minimum criteria for a search
//
//	This also tests: HasEmail, HasPhone, HasUserID, HasUsername, HasURL
//	HasName, HasAddress
func TestSearchMeetsMinimumCriteria(t *testing.T) {
	t.Parallel()

	t.Run("missing data", func(t *testing.T) {
		person := new(Person)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("raw name", func(t *testing.T) {
		person := new(Person)
		err := person.AddNameRaw("john smith")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("missing last name", func(t *testing.T) {
		person := new(Person)
		err := person.AddName("john", "", "", "", "")
		require.NoError(t, err)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("missing first name", func(t *testing.T) {
		person := new(Person)
		err := person.AddName("", "", "smith", "", "")
		require.NoError(t, err)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("first and last name", func(t *testing.T) {
		person := new(Person)
		err := person.AddName("john", "", "smith", "", "")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("email address", func(t *testing.T) {
		person := new(Person)
		err := person.AddEmail("clarkkent@gmail.com")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid phone number", func(t *testing.T) {
		person := new(Person)
		err := person.AddPhone(9785550145, 1)
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("missing phone code", func(t *testing.T) {
		person := new(Person)
		err := person.AddPhone(9785550145, 0)
		require.Error(t, err)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid raw phone number", func(t *testing.T) {
		person := new(Person)
		err := person.AddPhoneRaw("19785550145")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid user id", func(t *testing.T) {
		person := new(Person)
		err := person.AddUserID("clarkkent123", "twitter")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid username", func(t *testing.T) {
		person := new(Person)
		err := person.AddUsername("clarkkent", "twitter")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("valid url", func(t *testing.T) {
		person := new(Person)
		err := person.AddURL("https://twitter.com/clarkkent")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})

	t.Run("partial address number", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress("10", "", "", "", "", "", "")
		require.Error(t, err)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("partial address street", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress("10", "Hickory Lane", "", "", "", "", "")
		require.Error(t, err)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("partial address city", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress("10", "Hickory Lane", "", "Smallville", "", "", "")
		require.NoError(t, err)
		require.Equal(t, false, SearchMeetsMinimumCriteria(person))
	})

	t.Run("full address", func(t *testing.T) {
		person := new(Person)
		err := person.AddAddress("10", "Hickory Lane", "", "Smallville", "KS", "", "")
		require.NoError(t, err)
		require.Equal(t, true, SearchMeetsMinimumCriteria(person))
	})
}

// ExampleSearchMeetsMinimumCriteria example using SearchMeetsMinimumCriteria()
func ExampleSearchMeetsMinimumCriteria() {
	person := new(Person)
	if SearchMeetsMinimumCriteria(person) {
		fmt.Println("search meets minimum criteria")
	} else {
		fmt.Println("search does not meet minimum criteria")
	}
	// Output:search does not meet minimum criteria
}

// BenchmarkSearchMeetsMinimumCriteria benchmarks the SearchMeetsMinimumCriteria method
func BenchmarkSearchMeetsMinimumCriteria(b *testing.B) {
	person := new(Person)
	for i := 0; i < b.N; i++ {
		_ = SearchMeetsMinimumCriteria(person)
	}
}

// todo: test Search()

// todo: test SearchByPointer()

// todo: test SearchAllPossiblePeople()

// ======================================================================================================================
// Full Pipl Integration Tests (-test.short to skip)

// todo: make this a mock HTTP request

// TestSearchByPerson tests a live search using a real API key (if set)
// Tests a real pipl request
// Tests the debugging data returned from a request
func TestSearchByPerson(t *testing.T) {

	// Skip this test in short mode (not needed)
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	// Set the testing API key from ENV
	APIKey := os.Getenv("PIPL_CONTACT_API_KEY")

	// Create a new client object to handle your queries (supply an API Key)
	client, err := NewClient(APIKey, nil)
	require.NotNil(t, client)
	require.NoError(t, err)

	// Set your match requirements if you have any. You don't pay for results that
	// don't satisfy your match requirements (but your returned results will be empty)
	client.Parameters.Search.MatchRequirements = "name and phone"

	// Create a blank person to fill out with search terms
	searchObject := NewPerson()

	// Let's find out who this random guy is. We'll search by a username.
	err = searchObject.AddUsername("jeffbezos", "twitter")
	require.NoError(t, err)

	// And we'll add a full name
	err = searchObject.AddName("jeff", "preston", "bezos", "", "")
	require.NoError(t, err)

	// Some field types have a "raw" option that you can use if you're unsure
	// how to break it down. Pipl will attempt to parse it for you.
	// Generally you should use one or the other (AddX() or AddXRaw())
	err = searchObject.AddNameRaw("jeff preston bezos")
	require.NoError(t, err)

	// Launch the search (if you don't meet the minimum search criteria, an error
	// should be returned to you here stating such).
	var results *Response
	results, err = client.Search(context.Background(), searchObject)
	require.NoError(t, err)

	// Test the debugging data
	require.Equal(t, "https://api.pipl.com/search/", client.LastRequest.URL)
	require.Equal(t, "POST", client.LastRequest.Method)
	require.Equal(t, "key="+APIKey+"&match_requirements=name+and+phone&person=%7B%22names%22%3A%5B%7B%22first%22%3A%22jeff%22%2C%22last%22%3A%22bezos%22%2C%22middle%22%3A%22preston%22%7D%2C%7B%22raw%22%3A%22jeff+preston+bezos%22%7D%5D%2C%22usernames%22%3A%5B%7B%22content%22%3A%22jeffbezos%40twitter%22%7D%5D%7D&pretty=false&show_sources=all", client.LastRequest.PostData)
	require.NotNil(t, results)
	require.NotEqual(t, 0, len(results.Person.Names))
	require.Equal(t, "Jeffrey", results.Person.Names[0].First)
}
