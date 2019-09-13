package main

import "fmt"

// %v 以默认方式打印变量值
// %T 打印变量类型
// %% 字面上的百分号

// %+d 带符号的整型

// %5d 右对齐
// %-5d 左对齐
// %05d 左边补 0
func main() {
	a := 100
	b := "hello"
	c := false
	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)
	fmt.Printf("a:%T, b:%T, c:%T\n", a, b, c)

	d := -10
	fmt.Printf("d=%+d\n", d)

	fmt.Printf("|%d|\n", 6)
	fmt.Printf("|%8d|\n", 6)
	fmt.Printf("|%-8d|\n", 6)
	fmt.Printf("|%08d|\n", 6)

}
