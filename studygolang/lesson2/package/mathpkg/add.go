package mathpkg

import "fmt"

// Add public function
func Add(x, y int) int {
	return x + y
}

func init() {
	fmt.Println("add init execute")
}
