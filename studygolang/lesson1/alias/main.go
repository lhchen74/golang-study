package main

import "fmt"

// NewInt 是一个新类型
type NewInt int

// 类型别名: 代码编译后不存在
// 可以提高代码的可读性
type aliasInt = int

func main() {
	var a NewInt
	var b aliasInt
	fmt.Printf("a:%T, b:%T", a, b)
}
