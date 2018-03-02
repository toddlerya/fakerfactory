package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
		startTime := time.Now()
		records, count := fakerData(columns, number)
		costTime := time.Now().Sub(startTime)
		fmt.Println("构造数据耗时", costTime)

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
		return "暂未支持"
	case "phone":
		return faker.Phone("zh_CN")
	case "email":
		return faker.Email()
	case "imid":
		return faker.IMID()
	case "nickname":
		return faker.NickName()
	case "website":
		return "暂未支持"
	case "airplane":
		return "暂未支持"
	case "train":
		return "暂未支持"
	case "ipv4":
		return "暂未支持"
	case "ipv6":
		return "暂未支持"
	case "mac":
		return "暂未支持"
	default:
		return "未知字段, 请检查输入参数"
	}
}

func main() {
	StartServer()
	defer Conn.Close()
	//	faker.Seed(11)
	//	fmt.Println(faker.Phone("zh_CN"))
	//	fmt.Println(faker.Email())
	//	fmt.Println(faker.BirthDay())
	//	faker.IdCard()
	//	fmt.Println(faker.AreaCode("zh_CN"))
	//	fmt.Println(faker.Color("zh_CN"))
	//	fmt.Println(faker.Color("en_US", "zh_CN"))
	//	fmt.Println(faker.Name("en_US", "zh_CN"))
	//	fmt.Println(faker.FirstName("en_US", "zh_CN"))
	//	fmt.Println(faker.LastName("en_US", "zh_CN"))
	//	fmt.Println(faker.UserAgent())
	//	var Conn *sql.DB = faker.CreateConn() // 不应该在这里建立连接, 每次请求都会建立连接, 资源消耗比较多
	//	defer Conn.Close()
	//	fmt.Println(faker.Address(Conn))
}
