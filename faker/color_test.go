package faker

import (
	"fmt"
	"testing"
)

func ExampleColor() {
	Seed(11)
	fmt.Println(Color("zh_CN"))
	// Output: 黑色
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color("zh_CN")
	}
}
