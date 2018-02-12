package faker

import (
	"math/rand"
	"strings"
)

// Name will generate a random First and Last Name
func Name(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	first := strings.Join([]string{lang, "first"}, "_")
	last := strings.Join([]string{lang, "last"}, "_")
	switch lang {
	case "en_US":
		return getRandValue([]string{"person", first}) + " " + getRandValue([]string{"person", last})
	case "zh_CN":
		return getRandValue([]string{"person", last}) + getRandValue([]string{"person", first})
	default:
		return "未知的语言类型[姓名报错]"
	}
}

// FirstName will generate a random first name
func FirstName(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	first := strings.Join([]string{lang, "first"}, "_")
	switch lang {
	case "en_US", "zh_CN":
		return getRandValue([]string{"person", first})
	default:
		return "未知的语言类型[姓氏报错]"
	}
}

// LastName will generate a random last name
func LastName(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	last := strings.Join([]string{lang, "last"}, "_")
	switch lang {
	case "en_US", "zh_CN":
		return getRandValue([]string{"person", last})
	default:
		return "未知的语言类型[名字报错]"
	}
}

// NamePrefix will generate a random name prefix
//func NamePrefix() string {
//	return getRandValue([]string{"person", "prefix"})
//}

// NameSuffix will generate a random name suffix
//func NameSuffix() string {
//	return getRandValue([]string{"person", "suffix"})
//}
