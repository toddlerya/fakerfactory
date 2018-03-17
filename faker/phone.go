package faker

import (
	"fmt"
	"math/rand"
)

func MobilePhone(langs ...string) string {
	// lang: zh_CN --> 中文
	// lang: en_US --> 英文
	lang := langs[rand.Intn(len(langs))]
	phonePrefix := getRandValue([]string{"phone", lang})
	phone8Num := fmt.Sprintf("%08d", Number(1, 99999999))
	phoneNumber := phonePrefix + phone8Num
	return phoneNumber
}

// IMSI是15位的十进制数, 结构为: MCC + MNC + MSIN
// 国内为460开头、<=15、纯数字（国内一般为15位，国际规范为不超过15位）
// 暂时只做国内的
func Imsi() string {
	mcc := "460" // 中国为460
	mncSlice := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "20"}
	mnc := RandString(mncSlice)
	msin := fmt.Sprintf("%010d", Number(1, 9999999999))
	return mcc + mnc + msin
}

// IMEI（GSM网络）即EQUIREMENT_ID，15-17位、纯数字
func Imei() string {
	// IMEI = TAC [+FAC] + SNR + CD [+SVN]
	// TAC: 8位数字（早起是6位），中国的前两位是86
	// FAC: 2位数字，仅在早起TAC为6位的手机中存在
	// SNR: 由第9位数字开始的6位数字组成
	// CD:  验证码，由前14位数字通过Luhn算法得出
	// SVN: 软件版本号，仅在部分机型中存在
}

// Luhn算法
func Luhn(preNumArr []int) int {
	total := 0
	temp := 0
	preNumArr = append(preNumArr, 0) // 补充校验数字占位
	fmt.Println(preNumArr)
	length := len(preNumArr)

	for i := length - 1; i > -1; i-- {
		if i%2 != 0 { // 原始数组的第奇数个乘以二，若新数字大于9则减去9

		}
	}
}
