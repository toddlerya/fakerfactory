package faker

import (
	"fmt"
	"math/rand"
	"strconv"
)

func IMID() string {
	n := Number(10000, 99999999999)
	id := strconv.Itoa(n)
	return id
}

func NickName() string {
	splitString := []string{"_", ".", "-", "*", "!", "~", "丶", "氵", "灬"}
	randomNumber := fmt.Sprintf("%04d", Number(10, 3000))
	nickname := Name("en_US", "zh_CN") + getRandValue([]string{"person", "en_US_last"}) + splitString[rand.Intn(len(splitString))] + randomNumber
	return nickname
}
