[![Build Status](https://travis-ci.org/go-jira/jira.svg?branch=master)](https://travis-ci.org/go-jira/jira)
[![GoDoc](https://godoc.org/github.com/go-jira/jira?status.svg)](https://godoc.org/github.com/go-jira/jira)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# faker-cli

Simple command line client for [the wonderful faker pacakge](https://godoc.org/syreclabs.com/go/faker#pkg-constants) allowing one to quickly mock random user data

### Build/Install

You can build and install the official repository with [Go](https://golang.org/dl/):

	go get github.com/mkatychev/faker-cli/cmd/faker

## Usage

Usage:
	faker address
	faker adult [--max-age=<years>] (age|dob [-Y|-M|-D|--fmt=<fmt>])
	faker city
	faker country [--short]
	faker email
	faker name [first|last]
	faker password [<min> <max>]
	faker phone
	faker (postal-code|zip) [--state=<state>]
	faker sex [--short]
	faker state [--short]
	faker street
	faker street2

Options:
  -h --help         Show this screen.
  --version         Show version.
  --short           Return shortform of relevant data.
  --max-age=<years> Upper age limit for fake adult generation [default: 69].
  --fmt=<fmt>       Timestamp formatter, uses the magical reference date of "Mon Jan 2 15:04:05 MST 2006"/"2006-01-02"
```

* see https://golang.org/pkg/time/#Time.Format for help with timestamp formatting
