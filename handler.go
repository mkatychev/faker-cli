package gofaker

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

// Not is used to blacklist specific string values
var Not map[string]bool

// Short returns shortform of relevant data.
var Short bool

// Lower returns the lowercased form of relevant data.
var Lower bool

// HandleName handles the boolean map if `faker name` is called
func HandleName(opts map[string]interface{}) string {
	if opts["first"].(bool) {
		return faker.Name().FirstName()
	}
	if opts["last"].(bool) {
		return faker.Name().LastName()
	}
	return faker.Name().String()
}

// HandleSSN handles the boolean map if `faker ssn` is called
func HandleSSN(opts map[string]interface{}) string {
	var ssn string
	if opts["--now"].(bool) {
		ssn = strconv.Itoa(int(time.Now().Unix()))[:9]
	} else {
		ssn = strconv.Itoa(faker.RandomInt(111111111, 999999999))[:9]
	}
	if !Short {
		return fmt.Sprintf("%s-%s-%s", ssn[:3], ssn[3:5], ssn[5:])
	}

	return ssn
}

// HandlePhone handles the boolean map if `faker phone` is called
func HandlePhone(opts map[string]interface{}) string {
	if Short {
		faker.Locale = locales.En_US
		number := faker.PhoneNumber()
		return fmt.Sprintf(
			"%s-%s-%s",
			number.AreaCode(),
			number.ExchangeCode(),
			number.SubscriberNumber(4),
		)
	}
	return faker.PhoneNumber().String()
}

func getInt(from interface{}) (int, bool) {
	str, ok := from.(string)
	if !ok {
		return 0, false
	}
	val, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		panic(err)
	}
	return int(val), true
}

// HandleAddress handles `faker (city|state|zip-code|country)`
func HandleAddress(opts map[string]interface{}) string {
	var keyName string
	if opts["country"].(bool) {
		if Short {
			return faker.Address().CountryCode()
		}
		return faker.Address().Country()
	}
	if opts["city"].(bool) {
		return faker.Address().City()
	}
	if opts["state"].(bool) {
		stateFn := faker.Address().State
		keyName = "state"
		if Short {
			keyName = "state_abbr"
			stateFn = faker.Address().StateAbbr
		}
		if Not != nil {
			// https://stackoverflow.com/questions/15518919/take-address-of-value-inside-an-interface
			faker.Locale["address"].(map[string]interface{})[keyName] = Exclude(
				faker.Locale["address"].(map[string]interface{})[keyName].([]string), Not)
		}
		return stateFn()
	}
	if opts["street"].(bool) {
		return faker.Address().StreetAddress()
	}
	if opts["street2"].(bool) {
		return faker.Address().SecondaryAddress()
	}
	if opts["postal-code"].(bool) || opts["zip"].(bool) {
		if opts["--state"] != nil {
			faker.Address().PostcodeByState(opts["--state"].(string))
		}
		return faker.Address().Postcode()
	}
	return faker.Address().String()
}

// HandleSex handles `faker sex [--short]`
func HandleSex(opts map[string]interface{}) string {
	sexMap := map[bool]string{
		true:  "FEMALE",
		false: "MALE",
	}

	// random seed on a per call basis
	rand.Seed(time.Now().UnixNano())
	// binary gender easier to implement for now
	sex := sexMap[rand.Int()%2 == 0]
	if Short {
		sex = string(sex[0])
	}
	if Lower {
		sex = strings.ToLower(sex)
	}
	return sex
}

// HandleAdult handles `faker adult []` arguments
// TODO Pass in country codes for legal adult age
// current assumption is an adult age is 18 years or older
func HandleAdult(opts map[string]interface{}) string {
	var ok bool
	var min, max int
	// see this for formatting details:
	// https://golang.org/pkg/time/#Time.Format
	dateFormat := "2006-01-02"
	if opts["--min"] != nil {
		min, ok = getInt(opts["--min"])
		if !ok {
			panic("<min> must be an integer")
		}
	}
	if opts["--max"] != nil {
		max, ok = getInt(opts["--max"])
		if !ok {
			panic("<max> must be an integer")
		}
	}
	if opts["--fmt"] != nil {
		dateFormat = opts["--fmt"].(string)
	}
	dob := faker.Date().Birthday(min, max)

	if opts["age"].(bool) {
		return fmt.Sprintf("%d", time.Now().Year()-dob.Year())
	}
	if opts["dob"].(bool) {
		if opts["-Y"].(bool) {
			return fmt.Sprintf("%d", dob.Year())
		}
		if opts["-M"].(bool) {
			return fmt.Sprintf("%d", dob.Month())
		}
		if opts["-D"].(bool) {
			return fmt.Sprintf("%d", dob.Day())
		}

		return dob.Format(dateFormat)
	}
	return "<nil>"
}

// HandleEmail handles `faker email`
func HandleEmail(opts map[string]interface{}) string {
	return faker.Internet().Email()
}

// HandlePassword handles `faker password` generation, allows a max and min length
// default is 8-24
func HandlePassword(opts map[string]interface{}) string {
	var ok bool
	min, max := 8, 24
	if opts["<min>"] != nil {
		min, ok = getInt(opts["<min>"])
		if !ok {
			panic("<min> must be an integer")
		}
	}
	if opts["<max>"] != nil {
		max, ok = getInt(opts["<max>"])
		if !ok {
			panic("<max> must be an integer")
		}
	}
	return faker.Internet().Password(min, max)
}

// HandleGuid handles `faker guid`
func HandleGuid(opts map[string]interface{}) string {
	return uuid.New().String()
}
