package pipl

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testApartment               = "1"
	testCity                    = "Smallville"
	testEmail                   = "clark.kent@example.com"
	testEmailSecondary          = "clarkkent@gmail.com"
	testEthnicity               = "white"
	testFirstName               = "clark"
	testHouseNumber             = "10"
	testImage                   = "https://vignette3.wikia.nocookie.net/smallville/images/5/55/S10E18-Booster21.jpg"
	testKey                     = "test-key-1234567"
	testLastName                = "kent"
	testMiddleName              = "ryan"
	testPhone                   = int64(9785550145)
	testPhoneCountryCode        = 1
	testPhoneRaw                = "19785550145"
	testPOBox                   = "123"
	testSearchPointer           = "1906090343183157724859920008073008866"
	testState                   = "KS"
	testStreet                  = "Hickory Lane"
	testThumbnailToken          = "AE2861B242686E7BD0CB4D9049298EB7D18FEF66D950E8AB78BCD3F484345CE74536C19A85D0BA3D32DC9E7D1878CD4D341254E7AD129255C6983E6E154C4530A0DAAF665EA325FC0206F8B1D7E0B6B7AD9EBF71FCF610D57D" //nolint:gosec // Test token
	testURL                     = "https://twitter.com/clarkkent"
	testUserAgent               = "test-user-agent"
	testUserName                = "clarkkent"
	testUserNameServiceProvider = "twitter"
)

// TestClient_Search will test the method Search()
func TestClient_Search(t *testing.T) {
	t.Parallel()

	t.Run("error - does not meet criteria", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()

		response, err := c.Search(ctx, searchObject)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrDoesNotMeetMinimumCriteria)
		require.Nil(t, response)
	})

	t.Run("error - bad http response", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&errorHTTPResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()
		err := searchObject.AddUsername("superman", "facebook")
		require.NoError(t, err)

		var response *Response
		response, err = c.Search(ctx, searchObject)
		require.Error(t, err)
		require.Nil(t, response)
	})

	t.Run("error - bad JSON response", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&errorBadJSONResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()
		err := searchObject.AddUsername("superman", "facebook")
		require.NoError(t, err)

		var response *Response
		response, err = c.Search(ctx, searchObject)
		require.Error(t, err)
		require.Nil(t, response)
	})

	t.Run("error - missing api key response", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&errorMissingAPIKeyResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()
		err := searchObject.AddUsername("superman", "facebook")
		require.NoError(t, err)

		var response *Response
		response, err = c.Search(ctx, searchObject)
		require.Error(t, err)
		require.Nil(t, response)
	})

	t.Run("valid response - username", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()
		err := searchObject.AddUsername("superman", "facebook")
		require.NoError(t, err)

		var response *Response
		response, err = c.Search(ctx, searchObject)
		require.NoError(t, err)
		require.NotNil(t, response)

		require.Equal(t, "0", response.SearchID)
		require.Equal(t, 1, response.PersonsCount)
		require.Equal(t, 3, response.VisibleSources)
		require.Equal(t, 4, response.AvailableSources)
		require.Equal(t, http.StatusOK, response.HTTPStatusCode)

		require.Equal(t, testEmail, response.Query.Emails[0].Address)
	})

	t.Run("valid response - username - custom options", func(t *testing.T) {
		c := NewClient(
			WithAPIKey(testKey),
			WithHTTPClient(&validResponse{}),
			WithSearchOptions(&SearchOptions{
				Search: &SearchParameters{
					ShowSources:                ShowSourcesAll,
					MatchRequirements:          MatchRequirementsEmail,
					SourceCategoryRequirements: SourceCategoryRequirementsProfessionalAndBusiness,
					MinimumProbability:         MinimumProbability,
					MinimumMatch:               0.1,
					InferPersons:               true,
					HideSponsored:              true,
					LiveFeeds:                  false,
					TopMatch:                   true,
					Pretty:                     true,
				},
				Thumbnail: &ThumbnailSettings{},
			}),
		)
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()
		err := searchObject.AddUsername("superman", "facebook")
		require.NoError(t, err)

		var response *Response
		response, err = c.Search(ctx, searchObject)
		require.NoError(t, err)
		require.NotNil(t, response)

		require.Equal(t, "0", response.SearchID)
		require.Equal(t, 1, response.PersonsCount)
		require.Equal(t, 3, response.VisibleSources)
		require.Equal(t, 4, response.AvailableSources)
		require.Equal(t, http.StatusOK, response.HTTPStatusCode)

		require.Equal(t, testEmail, response.Query.Emails[0].Address)
	})
}

// TestClient_SearchByPointer will test the method SearchByPointer()
func TestClient_SearchByPointer(t *testing.T) {
	t.Parallel()

	t.Run("error - does not meet criteria", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()

		response, err := c.SearchByPointer(ctx, "")
		require.Error(t, err)
		require.ErrorIs(t, err, ErrInvalidSearchPointer)
		require.Nil(t, response)
	})

	t.Run("valid response - search_pointer", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()

		response, err := c.SearchByPointer(ctx, testSearchPointer)
		require.NoError(t, err)
		require.NotNil(t, response)

		require.Equal(t, "0", response.SearchID)
		require.Equal(t, 1, response.PersonsCount)
		require.Equal(t, 3, response.VisibleSources)
		require.Equal(t, 4, response.AvailableSources)
		require.Equal(t, http.StatusOK, response.HTTPStatusCode)
		require.Equal(t, testSearchPointer, response.Person.SearchPointer)

		require.Equal(t, testEmail, response.Query.Emails[0].Address)
	})

	t.Run("valid response - search_pointer - with options", func(t *testing.T) {
		c := NewClient(
			WithAPIKey(testKey),
			WithHTTPClient(&validResponse{}),
			WithSearchOptions(&SearchOptions{
				Search: &SearchParameters{
					ShowSources:                ShowSourcesAll,
					MatchRequirements:          MatchRequirementsEmail,
					SourceCategoryRequirements: SourceCategoryRequirementsProfessionalAndBusiness,
					MinimumProbability:         MinimumProbability,
					MinimumMatch:               0.1,
					InferPersons:               true,
					HideSponsored:              true,
					LiveFeeds:                  false,
					TopMatch:                   true,
					Pretty:                     true,
				},
				Thumbnail: &ThumbnailSettings{},
			}),
		)
		require.NotNil(t, c)

		ctx := context.Background()

		response, err := c.SearchByPointer(ctx, testSearchPointer)
		require.NoError(t, err)
		require.NotNil(t, response)

		require.Equal(t, "0", response.SearchID)
		require.Equal(t, 1, response.PersonsCount)
		require.Equal(t, 3, response.VisibleSources)
		require.Equal(t, 4, response.AvailableSources)
		require.Equal(t, http.StatusOK, response.HTTPStatusCode)
		require.Equal(t, testSearchPointer, response.Person.SearchPointer)

		require.Equal(t, testEmail, response.Query.Emails[0].Address)
	})
}

// TestClient_SearchAllPossiblePeople will test the method SearchAllPossiblePeople()
func TestClient_SearchAllPossiblePeople(t *testing.T) {
	t.Parallel()

	t.Run("error - does not meet criteria", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()

		response, err := c.SearchAllPossiblePeople(ctx, searchObject)
		require.Error(t, err)
		require.ErrorIs(t, err, ErrDoesNotMeetMinimumCriteria)
		require.Nil(t, response)
	})

	t.Run("error - bad http response", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&errorHTTPResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		response, err := c.SearchByPointer(ctx, testSearchPointer)
		require.Error(t, err)
		require.Nil(t, response)
	})

	t.Run("error - bad JSON response", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&errorBadJSONResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		response, err := c.SearchByPointer(ctx, testSearchPointer)
		require.Error(t, err)
		require.Nil(t, response)
	})

	t.Run("error - missing api key response", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&errorMissingAPIKeyResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		response, err := c.SearchByPointer(ctx, testSearchPointer)
		require.Error(t, err)
		require.Nil(t, response)
	})

	t.Run("basic search, no possible persons", func(t *testing.T) {
		c := NewClient(WithAPIKey(testKey), WithHTTPClient(&validResponse{}))
		require.NotNil(t, c)

		ctx := context.Background()
		searchObject := NewPerson()
		err := searchObject.AddUsername("superman", "facebook")
		require.NoError(t, err)

		var response *Response
		response, err = c.SearchAllPossiblePeople(ctx, searchObject)
		require.NoError(t, err)
		require.NotNil(t, response)

		require.Equal(t, "0", response.SearchID)
		require.Equal(t, 1, response.PersonsCount)
		require.Equal(t, 3, response.VisibleSources)
		require.Equal(t, 4, response.AvailableSources)
		require.Equal(t, http.StatusOK, response.HTTPStatusCode)

		require.Equal(t, testEmail, response.Query.Emails[0].Address)
	})

	t.Run("basic search, multiple possible persons", func(_ *testing.T) {
		// todo: add new JSON with possible persons

		// todo: test thumbnail generation of persons
	})
}
