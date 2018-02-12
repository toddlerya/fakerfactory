package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/toddlerya/FakerHub/faker"
)

func StartServer() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/fakerhub", GetFaker)
	}

	router.Run(":8000")
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
		records, count := fakerData(columns, number)

		startTime := time.Now()
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
	itemCols := strings.Split(columns, ",")
	fakerNumber, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}

	var results []map[string]string
	for i := 0; i < fakerNumber; i++ {
		resultMap := make(map[string]string)
		for _, col := range itemCols {
			resultMap[col] = matchFaker(strings.ToLower(col))
		}
		results = append(results, resultMap)
	}
	count := len(results)
	return results, count
}

func matchFaker(col string) string {
	switch col {
	case "color":
		return faker.Color("zh_CN")
	case "name":
		return faker.Name("zh_CN")
	case "age":
		return "暂未支持"
	case "sex":
		return "暂未支持"
	case "job":
		return "暂未支持"
	default:
		return "未知字段, 请检查输入参数"
	}
}
