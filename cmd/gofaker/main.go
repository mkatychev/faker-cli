package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	gofaker "github.com/mkatychev/faker-cli"
)

var argMap = map[string]func(map[string]interface{}) string{
	"address":     gofaker.HandleAddress,
	"adult":       gofaker.HandleAdult,
	"city":        gofaker.HandleAddress,
	"country":     gofaker.HandleAddress,
	"email":       gofaker.HandleEmail,
	"guid":        gofaker.HandleGUID,
	"name":        gofaker.HandleName,
	"now":         gofaker.HandleNow,
	"password":    gofaker.HandlePassword,
	"phone":       gofaker.HandlePhone,
	"postal-code": gofaker.HandleAddress,
	"sex":         gofaker.HandleSex,
	"state":       gofaker.HandleAddress,
	"street":      gofaker.HandleAddress,
	"street2":     gofaker.HandleAddress,
	"ssn":         gofaker.HandleSSN,
	"zip":         gofaker.HandleAddress,
}

const gofakerVersion = "0.4.1"
const usage string = `
Usage:
	gofaker address
	gofaker adult [--min=<years>] [--max=<years>] (age|dob [-Y|-M|-D|--fmt=<fmt>])
	gofaker city
	gofaker country [--short]
	gofaker email
	gofaker (guid|uuid)
	gofaker name [first|last]
	gofaker now [--fmt=<fmt>]
	gofaker password [<min> <max>]
	gofaker phone [--short]
	gofaker (postal-code|zip) [--short] [--state=<state>]
	gofaker sex [--short] [--lower]
	gofaker state [--short] [-n <val,>]
	gofaker street
	gofaker street2
	gofaker ssn [--short] [--now]

Options:
  -h --help                    Show this screen.
  --version                    Show version.
  --short                      Return shortform of relevant data.
  --min=<years>                Lower age limit for fake adult generation [default: 18].
  --max=<years>                Upper age limit for fake adult generation [default: 69].
  --fmt=<fmt>                  Timestamp formatter, uses the magical reference date of:
                               "Mon Jan 2 15:04:05 MST 2006"/"2006-01-02" [default: 2006-01-02].
  --not <val,>, -n <val,>      Blacklist specific string values, comma separated.
  --now                        Creates an SSN from the first 9 characters of the current timestamp.
`

func ifErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	opts, err := docopt.ParseArgs(usage, os.Args[1:], gofakerVersion)
	ifErr(err)

	// convert --not into a string slice using commas as separators
	if opts["--not"] != nil {
		gofaker.Not = make(map[string]bool)
		for _, v := range strings.Split(opts["--not"].(string), ",") {
			gofaker.Not[v] = true
		}
	}
	gofaker.DateFormat = opts["--fmt"].(string)
	gofaker.Short = opts["--short"].(bool)
	gofaker.Lower = opts["--lower"].(bool)
	for arg, handler := range argMap {
		if opts[arg].(bool) {
			fmt.Fprintln(os.Stdout, handler(opts))
		}
	}

}
