# CLAUDE.md - go-pipl SDK

## 🎯 Project Overview
**go-pipl** is an unofficial Go SDK wrapper for the [pipl.com API](https://pipl.com/api/), a people search and identity resolution service. The SDK enables programmatic searches for individuals using various identifiers like email, phone, name, address, usernames, and URLs.

**Core Purpose**: Provide a clean, type-safe Go interface to search for people and retrieve comprehensive identity data from Pipl's API.

> 📋 **Project Standards**: See [AGENTS.md](AGENTS.md) for comprehensive development standards and [tech-conventions/](tech-conventions/) for detailed technical guidelines covering Go essentials, testing, commit conventions, and more.

## 🏗️ Architecture Synopsis

### Key Components
- **`client.go`** - HTTP client with retry logic, exponential backoff, configurable timeouts
- **`pipl.go`** - Main search methods: `Search()`, `SearchAllPossiblePeople()`, `SearchByPointer()`
- **`definitions.go`** - Complete data structures for API requests/responses (Person, Address, Phone, Email, etc.)
- **`helpers.go`** - Person construction methods, validation, thumbnail processing
- **`request.go`** - Low-level HTTP request handling with context support
- **`errors.go`** - Comprehensive error definitions for validation and API responses

### Search Flow
1. Build `Person` object with search criteria using `Add*()` methods
2. Client validates minimum criteria (email, phone, name, address, username, userID, or URL required)
3. HTTP request with retry logic and exponential backoff
4. Response parsing and thumbnail URL generation (if enabled)
5. For `SearchAllPossiblePeople()`: automatic follow-up queries using search pointers

## 🔑 Critical Context

### API Constraints
- **Minimum Search Criteria**: Must have one of: full name, email, phone, username, userID, URL, or complete US address
- **Rate Limiting**: Built-in retry logic with exponential backoff (2 attempts by default)
- **API Key Required**: Set via `WithAPIKey()` or environment variable `PIPL_API_KEY`
- **US Address Only**: Addresses forced to `country = "US"` for search reliability

### Dependencies
- **Zero external runtime dependencies** - only uses Go stdlib
- **Testing only**: `github.com/stretchr/testify` for unit tests
- **Go 1.24+** required

### Key Patterns
- **Functional Options**: Client configuration via `WithAPIKey()`, `WithHTTPOptions()`, etc.
- **Context-First**: All search methods accept `context.Context` for cancellation/timeouts
- **Memory-Optimized Structs**: Field ordering optimized with `malign` for memory efficiency
- **Interface-Based Design**: `ClientInterface`, `HTTPInterface` for testability

## ⚡ Quick Commands (MAGE-X)

```bash
# Development setup
magex update:install           # Install/update MAGE-X build tool

# Testing
magex test                     # Run all tests (fast)
magex test:race                # Run tests with race detector (slower)

# Quality checks
magex lint                     # Run linting
magex bench                    # Run benchmarks

# Dependencies
magex deps:update              # Update all dependencies

# Release
magex version:bump bump=patch push  # Bump version and create tag

# Help
magex help                     # View all available commands
```

## 🧪 Testing & Validation

### Test Coverage
- Unit tests for all public methods
- Fuzz tests for critical input validation (`*_fuzz_test.go`)
- Mock tests for HTTP client behavior
- Example programs in `examples/` directory

### Running Tests
```bash
# Standard testing
go test -v ./...

# With coverage
go test -cover ./...

# Race detection
go test -race ./...
```

### Code Quality
- Go Report Card: A+ rating
- OpenSSF Scorecard security analysis
- CodeQL security scanning
- Dependabot automated dependency updates

## 📝 Code Patterns & Best Practices

### Building Search Objects
```go
person := pipl.NewPerson()
person.AddEmail("user@example.com")
person.AddName("John", "", "Doe", "", "")
person.AddPhoneRaw("+1-555-123-4567")
person.AddAddressRaw("123 Main St, Anytown, CA 90210")
```

### Client Configuration
```go
client := pipl.NewClient(
    pipl.WithAPIKey("your-api-key"),
    pipl.WithHTTPOptions(&pipl.HTTPOptions{
        RequestTimeout:    30 * time.Second,
        RequestRetryCount: 3,
    }),
)
```

### Error Handling
```go
response, err := client.Search(ctx, person)
if err != nil {
    if errors.Is(err, pipl.ErrDoesNotMeetMinimumCriteria) {
        // Handle insufficient search criteria
    }
    return fmt.Errorf("search failed: %w", err)
}
```

### Thumbnail Processing
```go
client := pipl.NewClient(
    pipl.WithSearchOptions(&pipl.SearchOptions{
        Thumbnail: &pipl.ThumbnailSettings{
            Enabled:  true,
            Height:   250,
            Width:    250,
            ZoomFace: true,
        },
    }),
)
```

## 🔧 Common Development Tasks

### Adding New Search Fields
1. Add field to relevant struct in `definitions.go` (follow memory optimization comments)
2. Add validation constants/errors in `errors.go`
3. Implement `Add*()` method in `helpers.go` with validation
4. Add corresponding `Has*()` validation method
5. Update `SearchMeetsMinimumCriteria()` if it's a searchable field
6. Add comprehensive tests

### Modifying HTTP Behavior
- Update `HTTPOptions` struct and defaults in `client_options.go`
- Modify retry logic in `retryableHTTPClient` in `client.go`
- Adjust request handling in `request.go`

### Adding New Client Configuration
1. Add field to `ClientOptions` in `client.go`
2. Create functional option in `client_options.go`
3. Update defaults if necessary

### Testing Strategy
- Test all validation logic with edge cases
- Mock HTTP responses for API integration tests
- Use fuzz testing for input validation
- Test error conditions and edge cases
- Verify memory efficiency of struct changes

## 🚨 Important Notes

- **Security**: Never commit API keys - use environment variables
- **Performance**: Struct field order matters (memory-optimized with malign)
- **API Limits**: Respect Pipl API rate limits and cost structure
- **US-Only**: Address searches are limited to US addresses
- **Context**: Always pass context for timeout/cancellation control
- **Validation**: Minimum search criteria validation prevents costly invalid API calls

## 📚 Key Files to Understand

| File             | Purpose           | Key Elements                       |
|------------------|-------------------|------------------------------------|
| `interface.go`   | API contracts     | `ClientInterface`, `SearchService` |
| `definitions.go` | Data structures   | All API request/response types     |
| `helpers.go`     | Person building   | `Add*()` methods, validation       |
| `pipl.go`        | Core search logic | Main search implementations        |
| `client.go`      | HTTP client       | Retry logic, configuration         |
| `examples/`      | Usage patterns    | Real-world usage examples          |

This SDK prioritizes **simplicity, reliability, and zero dependencies** while providing comprehensive access to Pipl's identity resolution capabilities.
