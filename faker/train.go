package faker

import (
	"strconv"
)

func TrainTripis() string {
	number := strconv.Itoa(Number(1, 9999))
	prefix := getRandValue([]string{"train", "prefix"})
	trips := prefix + number
	return trips
}
