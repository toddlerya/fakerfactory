package FakerHub

import (
	"fmt"

	"github.com/toddlerya/FakerHub/faker"
)

func main() {
	faker.Seed(11)
	fmt.Println(faker.SafeColor_zh_CN())
}
