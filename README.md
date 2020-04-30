# go-pipl
**go-pipl** is a simple golang wrapper to help find people using the [pipl.com API](https://pipl.com/api/).

[![Go](https://img.shields.io/github/go-mod/go-version/mrz1836/go-pipl)](https://golang.org/)
[![Build Status](https://travis-ci.com/mrz1836/go-pipl.svg?branch=master)](https://travis-ci.com/mrz1836/go-pipl)
[![Report](https://goreportcard.com/badge/github.com/mrz1836/go-pipl?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-pipl)
[![codecov](https://codecov.io/gh/mrz1836/go-pipl/branch/master/graph/badge.svg)](https://codecov.io/gh/mrz1836/go-pipl)
[![Release](https://img.shields.io/github/release-pre/mrz1836/go-pipl.svg?style=flat)](https://github.com/mrz1836/go-pipl/releases)
[![GoDoc](https://godoc.org/github.com/mrz1836/go-pipl?status.svg&style=flat)](https://pkg.go.dev/github.com/mrz1836/go-pipl?tab=doc)

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

## Installation

**go-pipl** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/mrz1836/go-pipl
```

## Documentation
You can view the generated [documentation here](https://pkg.go.dev/github.com/mrz1836/go-pipl?tab=doc).

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

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
all                            Runs test, install, clean, docs
bench                          Run all benchmarks in the Go application
clean                          Remove previous builds and any test cache data
clean-mods                     Remove all the Go mod cache
coverage                       Shows the test coverage
godocs                         Sync the latest tag with GoDocs
help                           Show all make commands available
lint                           Run the Go lint application
release                        Full production release (creates release in Github)
release-test                   Full production test release (everything except deploy)
release-snap                   Test the full release (build binaries)
tag                            Generate a new tag and push (IE: make tag version=0.0.0)
tag-remove                     Remove a tag if found (IE: make tag-remove version=0.0.0)
tag-update                     Update an existing tag to current commit (IE: make tag-update version=0.0.0)
test                           Runs vet, lint and ALL tests
test-short                     Runs vet, lint and tests (excludes integration tests)
update                         Update all project dependencies
update-releaser                Update the goreleaser application
vet                            Run the Go vet application
```
</details>

## Examples & Tests
All unit tests and [examples](pipl_test.go) run via [Travis CI](https://travis-ci.org/mrz1836/go-pipl) and uses [Go version 1.14.x](https://golang.org/doc/go1.14). View the [deployment configuration file](.travis.yml).

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

## Benchmarks
Run the Go [benchmarks](pipl_test.go):
```shell script
make bench
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

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

## Maintainers

| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |

## Contributing

This project was based off the original code [go pipl](https://github.com/xpcmdshell/pipl) project by [xpcmdshell](https://github.com/xpcmdshell)

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-pipl)

## License

![License](https://img.shields.io/github/license/mrz1836/go-pipl.svg?style=flat)
