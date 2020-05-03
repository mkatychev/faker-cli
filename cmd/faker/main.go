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
	faker (postal-code|zip)
	faker sex [--short]
	faker state
	faker street
	`
	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)

	for arg, handler := range argMap {
		if arguments[arg].(bool) {
			fmt.Println(handler(arguments))
		}
	}

}
