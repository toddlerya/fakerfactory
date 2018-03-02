package faker

import (
	"fmt"
	"math/rand"
)

// Email will generate a random email string
func Email() string {
	var email string
	splitString := []string{"_", ".", "-"}
	randomNumber := fmt.Sprintf("%04d", Number(10, 3000))
	flag := rand.Intn(10)
	if i := flag % 2; i == 0 {
		email = getRandValue([]string{"person", "en_US_first"}) + getRandValue([]string{"person", "en_US_last"}) + randomNumber
	} else {
		email = getRandValue([]string{"person", "en_US_first"}) + splitString[rand.Intn(len(splitString))] + getRandValue([]string{"person", "en_US_last"}) + randomNumber
	}

	email += getRandValue([]string{"email", "postfix"})

	return email
}
