package faker

import (
	"fmt"
	"math/rand"
)

func MobilePhone(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	phonePrefix := getRandValue([]string{"phone", lang})
	phone8Num := fmt.Sprintf("%08d", Number(1, 99999999))
	phoneNumber := phonePrefix + phone8Num
	return phoneNumber
}
