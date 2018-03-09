package faker

import (
	"strconv"
)

func Age() string {
	age := strconv.Itoa(Number(0, 105))
	return age
}
