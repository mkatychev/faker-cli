package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/mkatychev/faker-cli"
)

var argMap = map[string]func(map[string]interface{}) string{
	"address":     gofaker.HandleAddress,
	"adult":       gofaker.HandleAdult,
	"city":        gofaker.HandleAddress,
	"country":     gofaker.HandleAddress,
	"email":       gofaker.HandleEmail,
	"name":        gofaker.HandleName,
	"password":    gofaker.HandlePassword,
	"phone":       gofaker.HandlePhone,
	"postal-code": gofaker.HandleAddress,
	"sex":         gofaker.HandleSex,
	"state":       gofaker.HandleAddress,
	"street":      gofaker.HandleAddress,
	"street2":     gofaker.HandleAddress,
	"zip":         gofaker.HandleAddress,
}

func main() {
	usage := `Usage:
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
  --not <val,>, -n <val,>      Blacklist specific string values, comma separated.`
	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], "0.1.1")

	// convert --not into a string slice using commas as separators
	if arguments["--not"] != nil {
		gofaker.Not = make(map[string]bool)
		for _, v := range strings.Split(arguments["--not"].(string), ",") {
			gofaker.Not[v] = true
		}
	}
	gofaker.Short = arguments["--short"].(bool)
	for arg, handler := range argMap {
		if arguments[arg].(bool) {
			fmt.Println(handler(arguments))
		}
	}

}
