package main

import (
	"github.com/toddlerya/FakerHub/server"
)

func main() {
	//fmt.Println(faker.Color("zh_CN"))
	//fmt.Println(faker.Color("en_US", "zh_CN"))
	//	fmt.Println(faker.Name("en_US", "zh_CN"))
	//	fmt.Println(faker.FirstName("en_US", "zh_CN"))
	//	fmt.Println(faker.LastName("en_US", "zh_CN"))
	//fmt.Println(faker.UserAgent())
	server.StartServer()
}
