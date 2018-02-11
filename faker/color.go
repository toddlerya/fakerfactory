package faker

// Color will generate a random Chinese color string
func SafeColor_zh_CN() string {
	return getRandValue([]string{"color", "safe_zh_CN"})
}

// Color will generate a random English color string
func Color_en_US() string {
	return getRandValue([]string{"color", "safe_en_US"})
}

// SafeColor will generate a random safe en_US color string
func SafeColor_en_US() string {
	return getRandValue([]string{"color", "full_en_US"})
}
