package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/toddlerya/FakerHub/faker"
)

var dbPath string = `./faker/data/data.db`
var Conn *sql.DB = faker.CreateConn(dbPath) // 不应该在这里建立连接, 每次请求都会建立连接, 资源消耗比较多, 后续改进

func StartServer() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/fakerhub", GetFaker)
	}
	router.Run(":8001")
}

func GetFaker(c *gin.Context) {
	// todo: 需要对Query参数进行bind，先粗暴的判断下长度
	columns := c.Query("columns")
	number := c.Query("number")
	//	fmt.Println("columns==>", columns, len(columns))
	//	fmt.Println("number==>", number, len(number))
	if len(columns) == 0 || len(number) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status": "error",
				"code":   "100"},
			"data": gin.H{
				"number":  nil,
				"records": "请输入有效的参数"},
		})
	} else {
		//		startTime := time.Now()
		records, count := fakerData(columns, number)
		//		costTime := time.Now().Sub(startTime)
		//		fmt.Println("构造数据耗时", costTime)

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status": "ok",
				"code":   "0"},
			"data": gin.H{
				"count":   count,
				"records": records},
		})
	}

}

func fakerData(columns, number string) ([]map[string]string, int) {
	//	dbPath := `./faker/data/data.db`
	//	var Conn *sql.DB = faker.CreateConn(dbPath) // 不应该在这里建立连接, 每次请求都会建立连接, 资源消耗比较多, 后续改进
	//	defer Conn.Close()

	itemCols := strings.Split(columns, ",")
	fakerNumber, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	if fakerNumber >= 10000 {
		fakerNumber = 10000
	}
	var results []map[string]string
	for i := 0; i < fakerNumber; i++ {
		resultMap := make(map[string]string)
		for _, col := range itemCols {
			resultMap[col] = matchFaker(strings.ToLower(col), Conn)
		}
		results = append(results, resultMap)
	}
	count := len(results)
	return results, count
}

func matchFaker(col string, c *sql.DB) string {
	switch col {
	case "color":
		return faker.Color("zh_CN")
	case "job":
		return faker.Job("zh_CN")
	case "name":
		return faker.Name("zh_CN")
	case "sex":
		return faker.Gender("zh_CN")
	case "address":
		return faker.Address(c)
	case "idcard":
		return faker.IdCard()
	case "age":
		return faker.Age()
	case "mobilephone": // 移动电话
		return faker.MobilePhone("zh_CN")
	case "telphone": // 固定电话
		return "暂未支持"
	case "email":
		return faker.Email()
	case "imid":
		return faker.IMID()
	case "nickname":
		return faker.NickName()
	case "username":
		return faker.UserName()
	case "password":
		return faker.PassWord(true, true, true, true, true, 10)
	case "website":
		return faker.WebSite()
	case "url":
		return faker.URL()
	case "airport":
		return faker.AirPortInfo()
	case "voyage": // 航班号
		return faker.Voyage()
	case "airlineinfo": // 航空公司信息(代号+名称)
		return faker.AirlineInfo()
	case "traintrips":
		return faker.TrainTripis()
	case "trainseat":
		return faker.SeatOfTrain()
	case "flightseat":
		return faker.SeatOfFlight()
	case "ipv4":
		return faker.IPv4Address()
	case "ipv6":
		return faker.IPv6Address()
	case "mac": // 暂时随机返回各种类型的MAC地址
		return faker.RandMacAddress()
	case "imsi": // 暂时只提供460开头的中国imsi
		return faker.Imsi()
	case "imei":
		return faker.Imei()
	case "meid":
		return faker.RandMeid()
	case "deviceid": //采集设备ID、固定21位、前9位为安全厂商ID（如FIBERHOME），后12位为采集设备MAC，规则同MAC、所有字母大写
		return "暂未支持"
	case "date": // 时间字段,两种，10位绝对秒(当天数据)，数据库日期格式{YYYYMMDD,hh:mm:ss}
		return "暂未支持"
	case "useragent":
		return faker.UserAgent()
	case "gapassport":
		return "暂未支持"
	case "twpassport":
		return "暂未支持"

	default:
		return "未知字段, 请检查输入参数"
	}
}

func main() {
	StartServer()
	defer Conn.Close()
}
