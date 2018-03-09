package faker

import (
	"fmt"
	"math/rand"
)

// 航班号，航次，航班
func Voyage() string {
	var randomNumber string
	n := rand.Intn(10)
	if n%2 == 0 {
		randomNumber = fmt.Sprintf("%04d", Number(1000, 9999))
	} else {
		randomNumber = fmt.Sprintf("%03d", Number(100, 999))
	}
	voyage_prefix := getRandValue([]string{"flight", "airline_code"})
	voyage := voyage_prefix + randomNumber
	return voyage
}

// 航空公司名称
func AirlineName() string {
	airline := getRandValue([]string{"flight", "airline_name"})
	return airline
}

// 航空公司代号，名称
func AirlineInfo() string {
	airlineInfo := getRandValue([]string{"flight", "airline_info"})
	return airlineInfo
}
