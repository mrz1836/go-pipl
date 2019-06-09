# go-pipl
**go-pipl** Simple golang wrapper to help find people using the Pipl API.  This project follows Go best practices and you can view the standards and specifications at the [end of this readme](https://github.com/mrz1836/go-pipl#adheres-to-effective-go-standards).

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-pipl.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-pipl?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-pipl)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/b11a08d5619849a0ae911d91e3bb47c7)](https://www.codacy.com/app/mrz1818/go-pipl?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-pipl&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-pipl.svg?branch=master)](https://travis-ci.com/mrz1836/go-pipl)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-pipl.svg?style=flat)](https://github.com/mrz1836/go-pipl/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-pipl?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-pipl) |


## Table of Contents
- [Installation](https://github.com/mrz1836/go-pipl#installation)
- [Documentation](https://github.com/mrz1836/go-pipl#documentation)
- [Examples & Tests](https://github.com/mrz1836/go-pipl#examples--tests)
- [Benchmarks](https://github.com/mrz1836/go-pipl#benchmarks)
- [Code Standards](https://github.com/mrz1836/go-pipl#code-standards)
- [Usage](https://github.com/mrz1836/go-pipl#usage)
- [Maintainers](https://github.com/mrz1836/go-pipl#maintainers)
- [Contributing](https://github.com/mrz1836/go-pipl#contributing)
- [License](https://github.com/mrz1836/go-pipl#license)

## Installation

**go-pipl** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-pipl
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-pipl).

## Examples & Tests
All unit tests and [examples](https://github.com/mrz1836/go-pipl/blob/master/pipl_test.go) run via [Travis CI](https://travis-ci.com/mrz1836/go-pipl) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](https://github.com/mrz1836/go-pipl/blob/master/.travis.yml).

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
Run the Go [benchmarks](https://github.com/mrz1836/go-pipl/blob/master/pipl_test.go):
```bash
$ cd ../go-pipl
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](https://github.com/mrz1836/go-pipl/blob/master/CODE_STANDARDS.md).

## Usage
- View the [examples](https://github.com/mrz1836/go-pipl/blob/master/pipl_test.go)
- View the [benchmarks](https://github.com/mrz1836/go-pipl/blob/master/pipl_test.go)
- View the [tests](https://github.com/mrz1836/go-pipl/blob/master/pipl_test.go)

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
    client := pipl.NewClient("your-api-key")

    // Set your match requirements
    client.SearchParameters.MatchRequirements = "name and phone"

    // Create a new person
    search := pipl.NewPerson()

    // Search by username
    search.AddUsername("@jeffbezos")

    // Submit the search
    response, err := client.SearchByPerson(search)
    if err != nil {
        log.Fatalln(err)
    }

    // Show the response
    fmt.Println(response.Person.Names[0].Display)
    // Output: Jeff Preston Bezos
}
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836)

## Contributing

This project was based off the original code [go pipl](https://github.com/xpcmdshell/pipl) project by [xpcmdshell](https://github.com/xpcmdshell)

View the [contributing guidelines](https://github.com/mrz1836/go-pipl/blob/master/CONTRIBUTING.md) and follow the [code of conduct](https://github.com/mrz1836/go-pipl/blob/master/CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-pipl)

## License

![License](https://img.shields.io/github/license/mrz1836/go-pipl.svg?style=flat)
