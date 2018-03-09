package faker

import (
	"fmt"
	"math/rand"
	"strings"
)

// DomainName will generate a random url domain name
func DomainName() string {
	randomNumber := fmt.Sprintf("%04d", Number(10, 3000))
	domain := getRandValue([]string{"person", "en_US_first"}) + randomNumber
	return domain + "." + DomainSuffix()
}

// DomainSuffix will generate a random domain suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}

func WebSite() string {
	return RandString([]string{"www.", ""}) + DomainName()
}

// URL will generate a random url string
func URL() string {
	url := "http" + RandString([]string{"s", ""}) + "://"
	url += WebSite()

	// Slugs
	num := Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = BS()
	}
	url += "/" + strings.ToLower(strings.Join(slug, "/"))

	return url
}

// HTTPMethod will generate a random http method
func HTTPMethod() string {
	return getRandValue([]string{"internet", "http_method"})
}

// IPv4Address will generate a random version 4 ip address
func IPv4Address() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// IPv6Address will generate a random version 6 ip address
func IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// Username will genrate a random username based upon picking a random lastname and random numbers at the end
func UserName() string {
	return getRandValue([]string{"person", "en_US_last"}) + replaceWithNumbers("####")
}

// Password will generate a random password
func PassWord(lower bool, upper bool, numeric bool, special bool, space bool, length int) string {
	var passString string
	lowerStr := "abcdefghijklmnopqrstuvwxyz"
	upperStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericStr := "0123456789"
	specialStr := "!@#$%&?-_"
	spaceStr := " "

	if lower {
		passString += lowerStr
	}
	if upper {
		passString += upperStr
	}
	if numeric {
		passString += numericStr
	}
	if special {
		passString += specialStr
	}
	if space {
		passString += spaceStr
	}

	// Set default if empty
	if passString == "" {
		passString = lowerStr + numericStr
	}

	passBytes := []byte(passString)
	finalBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		finalBytes[i] = passBytes[rand.Intn(len(passBytes))]
	}
	return string(finalBytes)
}
