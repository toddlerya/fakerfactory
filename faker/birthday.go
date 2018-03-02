// 生成生日日期

package faker

import (
	"fmt"
	"strconv"
	"time"
)

func BirthDay() string {
	var day string
	yearStart := time.Now().Year() - 105 // 最大年龄105岁
	year := Number(yearStart, time.Now().Year())
	month := fmt.Sprintf("%02d", Number(1, 12))
	// 判断是否为闰年
	term01 := year % 4
	term02 := year % 100
	term20 := year % 400
	if (term01 == 0 && term02 != 0) || term20 == 0 { // 闰年
		switch month {
		case "01", "03", "05", "07", "08", "10", "12":
			day = fmt.Sprintf("%02d", Number(1, 31))
		case "02":
			day = fmt.Sprintf("%02d", Number(1, 28))
		case "04", "06", "09", "11":
			day = fmt.Sprintf("%02d", Number(1, 30))
		default:
			panic("月份错误")
		}
	} else { // 非闰年
		switch month {
		case "01", "03", "05", "07", "08", "10", "12":
			day = fmt.Sprintf("%02d", Number(1, 31))
		case "02":
			day = fmt.Sprintf("%02d", Number(1, 29))
		case "04", "06", "09", "11":
			day = fmt.Sprintf("%02d", Number(1, 30))
		default:
			panic("月份错误")
		}
	}
	var birthday string
	birthday = strconv.Itoa(year) + month + day
	return birthday
}
