package faker

// Color will generate a random Chinese color string
func Color(lang string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	return getRandValue([]string{"color", lang})
}
