[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

# faker-cli

Simple command line client for [the wonderful faker pacakge](https://godoc.org/syreclabs.com/go/faker#pkg-constants) allowing one to quickly mock random user data

### Build/Install

You can build and install the official repository with [Go](https://golang.org/dl/):

	go get github.com/mkatychev/faker-cli/cmd/gofaker

## Usage

```
Usage:
	gofaker address
	gofaker adult [--max-age=<years>] (age|dob [-Y|-M|-D|--fmt=<fmt>])
	gofaker city
	gofaker country [--short]
	gofaker email
	gofaker name [first|last]
	gofaker password [<min> <max>]
	gofaker phone
	gofaker (postal-code|zip) [--state=<state>]
	gofaker sex [--short]
	gofaker state [--short] [-n <val,>]
	gofaker street
	gofaker street2

Options:
  -h --help                    Show this screen.
  --version                    Show version.
  --short                      Return shortform of relevant data.
  --max-age=<years>            Upper age limit for fake adult generation [default: 69].
  --fmt=<fmt>                  Timestamp formatter, uses the magical reference date of:
                               "Mon Jan 2 15:04:05 MST 2006"/"2006-01-02".
  --not <val,>, -n <val,>      Blacklist specific string values, comma separated.
```

* see https://golang.org/pkg/time/#Time.Format for help with timestamp formatting
