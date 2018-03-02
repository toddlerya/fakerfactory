package faker

import (
	"math/rand"
)

// Color will generate a random color string
func Gender(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	return getRandValue([]string{"gender", lang})
}
