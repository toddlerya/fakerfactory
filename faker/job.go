package faker

import (
	"math/rand"
)

func Job(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	return getRandValue([]string{"job", lang})
}
