package faker

import (
	"math/rand"
)

// AreaCode will generate a random areacode string
func AreaCode(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	return getRandValue([]string{"areacode", lang})
}
