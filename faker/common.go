package faker

import (
	"math/rand"

	"github.com/toddlerya/FakerHub/faker/data"
)

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	_, checkOk = data.Data[dataVal[0]]
	if len(dataVal) == 2 && checkOk {
		_, checkOk = data.Data[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Get Random Value
func getRandValue(dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return data.Data[dataVal[0]][dataVal[1]][rand.Intn(len(data.Data[dataVal[0]][dataVal[1]]))]
}
