package main

import "fmt"
import "github.com/mkatychev/farker-cli"

var argMap = map[string]func(map[string]interface{}) string{
	"address":     faker_cli.HandleAddress,
	"adult":       faker_cli.HandleAdult,
	"city":        faker_cli.HandleAddress,
	"country":     faker_cli.HandleAddress,
	"email":       faker_cli.HandleEmail,
	"name":        faker_cli.HandleName,
	"password":    faker_cli.HandlePassword,
	"phone":       faker_cli.HandlePhone,
	"postal-code": faker_cli.HandleAddress,
	"sex":         faker_cli.HandleSex,
	"state":       faker_cli.HandleAddress,
	"street":      faker_cli.HandleAddress,
	"zip":         faker_cli.HandleAddress,
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
		if arguments[arg].(bool) == true {
			fmt.Println(handler(arguments))
		}
	}

}
