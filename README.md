# go-pipl
**go-pipl** is a simple golang wrapper to help find people using the pipl.com API.

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-pipl.svg?style=flat&p=1) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-pipl?style=flat&p=1)](https://goreportcard.com/report/github.com/mrz1836/go-pipl)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/01708ca3079e4933bafb3b39fe2aaa9d)](https://www.codacy.com/app/mrz1818/go-pipl?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-pipl&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-pipl.svg?branch=master)](https://travis-ci.com/mrz1836/go-pipl)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-pipl.svg?style=flat)](https://github.com/mrz1836/go-pipl/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-pipl?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-pipl) |

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
```bash
$ go get -u github.com/mrz1836/go-pipl
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-pipl).

### Features
- Complete configuration coverage for the pipl.com API
- Pipl client is completely configurable
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
- Thumbnail configuration for `person.Images`
    - Adds `image.ThumbnailURL` with the complete url for a live thumbnail
- Test and example coverage for all methods

## Examples & Tests
All unit tests and [examples](pipl_test.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-pipl) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```bash
$ cd ../go-pipl
$ go test ./... -v
```

Run tests (excluding integration tests)
```bash
$ cd ../go-pipl
$ go test ./... -v -test.short
```

## Benchmarks
Run the Go [benchmarks](pipl_test.go):
```bash
$ cd ../go-pipl
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [examples](pipl_test.go)
- View the [benchmarks](pipl_test.go)
- View the [tests](pipl_test.go)

Basic implementation:
```golang
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/mrz1836/go-pipl"
)

func main() {

    // Create a client with your api key
    client, _ := pipl.NewClient("your-api-key")

    // Create a new person for searching
    person := pipl.NewPerson()
    person.AddUsername("@jeffbezos")

    // Submit the search
    response, _ := client.Search(person)

    // Use the pipl response
    fmt.Println(response.Person.Names[0].Display)
    // Output: Jeff Preston Bezos
}
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836)

## Contributing

This project was based off the original code [go pipl](https://github.com/xpcmdshell/pipl) project by [xpcmdshell](https://github.com/xpcmdshell)

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-pipl)

## License

![License](https://img.shields.io/github/license/mrz1836/go-pipl.svg?style=flat&p=1)
