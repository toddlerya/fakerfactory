package faker

import (
	"fmt"
	"math/rand"
	"strings"
)

// DomainName will generate a random url domain name
func DomainName() string {
	randomNumber := fmt.Sprintf("%04d", Number(10, 3000))
	domain := getRandValue([]string{"person", "en_US_first"}) + randomNumber
	return domain + "." + DomainSuffix()
}

// DomainSuffix will generate a random domain suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}

func WebSite() string {
	return RandString([]string{"www.", ""}) + DomainName()
}

// URL will generate a random url string
func URL() string {
	url := "http" + RandString([]string{"s", ""}) + "://"
	url += WebSite()

	// Slugs
	num := Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = BS()
	}
	url += "/" + strings.ToLower(strings.Join(slug, "/"))

	return url
}

// HTTPMethod will generate a random http method
func HTTPMethod() string {
	return getRandValue([]string{"internet", "http_method"})
}

// IPv4Address will generate a random version 4 ip address
func IPv4Address() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// IPv6Address will generate a random version 6 ip address
func IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// MacAddress will generate a random mac address
// 根据sep确定分隔符返回, letterType=true返回大写字母, false返回小写字母
func MacAddress(sep string, letterType bool) string {
	// 固定12位，0-9数字+[A-F]或[a-f]
	// 【备注：MAC地址16进制中的第一个字节的第二个数字一定是偶数，即对应到十六进制为0、2、4、6、8、A、C、E中的一个】
	simpleValue := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	specialValue := []string{"0", "2", "4", "6", "8", "A", "C", "E"}
	macSlice := []string{}
	var temp string
	for i := 0; i < 17; i++ {
		switch i {
		case 1:
			temp = RandString(specialValue)
		case 2, 5, 8, 11, 14:
			temp = sep
		default:
			temp = RandString(simpleValue)
		}
		macSlice = append(macSlice, temp)
	}
	macStr := strings.Join(macSlice, "")
	if !letterType {
		macStr = strings.ToLower(macStr)
	}
	return macStr
}

// 随机返回大小写，多种分隔符的mac地址
func RandMacAddress() string {
	sep := RandString([]string{"-", ":"})
	lowUp := RandBool([]bool{true, false})
	return MacAddress(sep, lowUp)

}

// //采集设备ID、固定21位、前9位为安全厂商ID（如FIBERHOME），后12位为采集设备MAC，规则同MAC、所有字母大写
func DeviceID() string {
	lowUp := RandBool([]bool{true, false})
	companySlice := []string{"FIBERHOME", "TAIJIXXXX", "MEIYAXXXXX", "ZHAOWUXXX"}
	com := RandString(companySlice)
	mac := MacAddress("", lowUp)
	return com + mac
}

// Username will genrate a random username based upon picking a random lastname and random numbers at the end
func UserName() string {
	return getRandValue([]string{"person", "en_US_last"}) + replaceWithNumbers("####")
}

// Password will generate a random password
func PassWord(lower bool, upper bool, numeric bool, special bool, space bool, length int) string {
	var passString string
	lowerStr := "abcdefghijklmnopqrstuvwxyz"
	upperStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericStr := "0123456789"
	specialStr := "!@#$%&?-_"
	spaceStr := " "

	if lower {
		passString += lowerStr
	}
	if upper {
		passString += upperStr
	}
	if numeric {
		passString += numericStr
	}
	if special {
		passString += specialStr
	}
	if space {
		passString += spaceStr
	}

	// Set default if empty
	if passString == "" {
		passString = lowerStr + numericStr
	}

	passBytes := []byte(passString)
	finalBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		finalBytes[i] = passBytes[rand.Intn(len(passBytes))]
	}
	return string(finalBytes)
}
