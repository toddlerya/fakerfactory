package main

import (
	"fmt"

	"github.com/toddlerya/FakerHub/faker"
//	"github.com/toddlerya/FakerHub/server"
)

func main() {
	faker.Seed(11)
	fmt.Println(faker.Color("zh_CN"))
	fmt.Println(faker.Color("en_US", "zh_CN"))
	//fmt.Println(faker.UserAgent())
	//server.StartServer()
}
