package main

import (
	"fmt"

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
  --fmt=<fmt>       Timestamp formatter, uses the magical reference date of "Mon Jan 2 15:04:05 MST 2006"/"2006-01-02"`
	arguments, _ := docopt.ParseDoc(usage)

	for arg, handler := range argMap {
		if arguments[arg].(bool) {
			fmt.Println(handler(arguments))
		}
	}

}
