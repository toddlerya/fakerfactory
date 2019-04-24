package faker

import (
	"math/rand"
)

//获取所有车辆品牌的名称
func CarBrand(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	return getRandValue([]string{"carbrand", lang})
}