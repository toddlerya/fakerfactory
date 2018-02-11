package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func StartServer() {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/fakehub", GetFaker)
	}

	router.Run(":8000")
}

func GetFaker(c *gin.Context) {
	var users = []Users{
		Users{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
		Users{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
	}
	columns := c.Query("columns")
	number := c.Query("number")
	fmt.Println("columns==>", columns)
	fmt.Println("number==>", number)

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"status": "ok",
			"code":   "0"},
		"data": users,
	})
}
