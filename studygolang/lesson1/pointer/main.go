package main

import "fmt"

func modifyArray(arr *[3]int) {
	// go 语言指针不支持修改操作, arr[0] 实际上是 (*arr)[0] = 100
	// arr[0] = 100
	(*arr)[0] = 100
}
func main() {
	var a int
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Printf("type:%T, value:%v\n", b, *b)

	arr := [3]int{1, 2, 3}
	modifyArray(&arr)
	fmt.Println(arr)

	// // aa 是一个 int 类型的空指针
	// var aa *int
	// // panic: runtime error: invalid memory address or nil pointer dereference
	// *aa = 10

	// new 用来初始化值类型指针
	// make 用来初始化引用类型指针
	var ap = new(int) // get int type pointer
	fmt.Println(ap)
	*ap = 10
	fmt.Println(ap)

	var cp = new([3]int)
	cp[0] = 1
	fmt.Println(cp, *cp)
}
