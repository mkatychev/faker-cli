package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	handler "github.com/mkatychev/faker-cli"
)

var argMap = map[string]func(map[string]interface{}) string{
	"address":     handler.HandleAddress,
	"adult":       handler.HandleAdult,
	"city":        handler.HandleAddress,
	"country":     handler.HandleAddress,
	"email":       handler.HandleEmail,
	"name":        handler.HandleName,
	"password":    handler.HandlePassword,
	"phone":       handler.HandlePhone,
	"postal-code": handler.HandleAddress,
	"sex":         handler.HandleSex,
	"state":       handler.HandleAddress,
	"street":      handler.HandleAddress,
	"street2":     handler.HandleAddress,
	"zip":         handler.HandleAddress,
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
	gofaker state [--short]
	gofaker street
	gofaker street2

Options:
  -h --help         Show this screen.
  --version         Show version.
  --short           Return shortform of relevant data.
  --max-age=<years> Upper age limit for fake adult generation [default: 69].
  --fmt=<fmt>       Timestamp formatter, uses the magical reference date of:
                    "Mon Jan 2 15:04:05 MST 2006"/"2006-01-02".`
	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], "0.1")

	for arg, handler := range argMap {
		if arguments[arg].(bool) {
			fmt.Println(handler(arguments))
		}
	}

}
