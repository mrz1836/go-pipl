# go-pipl
> The unofficial golang wrapper for the [pipl.com API](https://pipl.com/api/).

[![Release](https://img.shields.io/github/release-pre/mrz1836/go-pipl.svg?logo=github&style=flat)](https://github.com/mrz1836/go-pipl/releases)
[![Build Status](https://img.shields.io/github/workflow/status/mrz1836/go-pipl/run-go-tests?logo=github&v=3)](https://github.com/mrz1836/go-pipl/actions)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-pipl?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-pipl)
[![codecov](https://codecov.io/gh/mrz1836/go-pipl/branch/master/graph/badge.svg)](https://codecov.io/gh/mrz1836/go-pipl)
[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-pipl)](https://golang.org/)
[![Sponsor](https://img.shields.io/badge/sponsor-MrZ-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/mrz1836)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat)](https://mrz1818.com/?tab=tips&af=go-pipl)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-pipl** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-pipl
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/mrz1836/go-pipl?tab=doc)

[![GoDoc](https://godoc.org/github.com/mrz1836/go-pipl?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-pipl?tab=doc)

### Features
- Complete coverage for the [pipl.com API](https://pipl.com/api/)
- [Client](client.go) is completely configurable
- Using [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Minimum criteria detection before submitting a pipl query
- Search by pipl pointer reference
- Search for a single person via any of the following:
    - Full Name
    - Full Street Address
    - Email
    - Phone
    - Username or UserID or URL
- Search **all possible people**
    - Returns the original full person record
    - Searches all possible persons and gets full details
    - Combines all persons into one single response
- Thumbnail configuration setting for `person.Images`
    - Adds `image.ThumbnailURL` with the complete url for a live thumbnail
- Test and example coverage for all methods


<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                  Runs lint, test-short and vet
clean                Remove previous builds and any test cache data
clean-mods           Remove all the Go mod cache
coverage             Shows the test coverage
godocs               Sync the latest tag with GoDocs
help                 Show this help message
install              Install the application
install-go           Install the application (Using Native Go)
lint                 Run the golangci-lint application (install if not found)
release              Full production release (creates release in Github)
release              Runs common.release then runs godocs
release-snap         Test the full release (build binaries)
release-test         Full production test release (everything except deploy)
replace-version      Replaces the version in HTML/JS (pre-deploy)
tag                  Generate a new tag and push (tag version=0.0.0)
tag-remove           Remove a tag if found (tag-remove version=0.0.0)
tag-update           Update an existing tag to current commit (tag-update version=0.0.0)
test                 Runs vet, lint and ALL tests
test-ci              Runs all tests via CI (exports coverage)
test-ci-no-race      Runs all tests via CI (no race) (exports coverage)
test-ci-short        Runs unit tests via CI (exports coverage)
test-short           Runs vet, lint and tests (excludes integration tests)
uninstall            Uninstall the application (and remove files)
update-linter        Update the golangci-lint package (macOS only)
vet                  Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](pipl_test.go) run via [Github Actions](https://github.com/mrz1836/go-pipl/actions) and
uses [Go version 1.15.x](https://golang.org/doc/go1.15). View the [configuration file](.github/workflows/run-tests.yml).

- [helper examples & tests](helpers_test.go)
- [pipl examples &  tests](pipl_test.go)
- [response tests](response_test.go)

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go [benchmarks](pipl_test.go):
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](.github/CODE_STANDARDS.md).

<br/>

## Usage
- View the [pipl examples & benchmarks](pipl_test.go)
- View the [helper examples & benchmarks](helpers_test.go)
- View the [response tests](response_test.go)

Basic implementation:
```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/mrz1836/go-pipl"
)

func main() {

    // Create a client with your api key
    client, _ := pipl.NewClient("your-api-key", nil)

    // Create a new person for searching
    person := pipl.NewPerson()
    _ = person.AddUsername("jeffbezos", "twitter")

    // Submit the search
    response, _ := client.Search(person)

    // Use the pipl response
    fmt.Println(response.Person.Names[0].Display)
    // Output: Jeff Preston Bezos
}
```
 
<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:------------------------------------------------------------------------------------------------:|
|                                [MrZ](https://github.com/mrz1836)                                 |
 
<br/>

## Contributing
View the [contributing guidelines](.github/CONTRIBUTING.md) and follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://mrz1818.com/?tab=tips&af=go-pipl) to ensure this journey continues indefinitely! :rocket:

### Credits
This project was based off the original code [go pipl](https://github.com/xpcmdshell/pipl) project by [xpcmdshell](https://github.com/xpcmdshell)

<br/>

## License

[![License](https://img.shields.io/github/license/mrz1836/go-pipl.svg?style=flat)](LICENSE)
