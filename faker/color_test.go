package faker

import (
	"fmt"
)

func ExampleColor() {
	Seed(11)
	fmt.Println(SafeColor_zh_CN())
	// Output: MediumOrchid
}

func main() {
	ExampleColor()
}
