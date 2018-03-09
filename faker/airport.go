package faker

// 机场的IATA三位编码
func AirPort3Code() string {
	code := getRandValue([]string{"airport", "airport_iata_code"})
	return code
}

// 机场的ICAO四位编码
func AirPort4Code() string {
	code := getRandValue([]string{"airport", "airport_icao_code"})
	return code
}

// 机场所在的城市
func AirPortCity() string {
	city := getRandValue([]string{"airport", "airport_city"})
	return city
}

// 机场所在城市的拼音
func AirPortCityPinyin() string {
	cityPinyin := getRandValue([]string{"airport", "airport_city_pinyin"})
	return cityPinyin
}

// 机场名字
func AirPortName() string {
	name := getRandValue([]string{"airport", "airport_name"})
	return name
}

// 机场信息
func AirPortInfo() string {
	info := getRandValue([]string{"airport", "airport_info"})
	return info
}
