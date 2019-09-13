package main

import (
	"errors"
	"fmt"
	"strings"
)

func intSum(x, y int) int {
	return x + y
}

func intSum2(x ...int) int {
	fmt.Println(x) // x 是一个切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc2(x, y int) (sum, sub int) {
	// no new variables on left side of :=
	// sum := x + y
	// sub := x - y
	sum = x + y
	sub = x - y
	return
}

// 定义函数类型
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("Unsupport operation")
		return nil, err
	}
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc4(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	s := intSum2(1, 2, 3)
	fmt.Println(s)

	var c calculation
	c = add
	fmt.Println(c(1, 2))

	r := calc3(1, 2, add)
	fmt.Println(r)

	f := adder(10)
	fmt.Println(f(20))
	fmt.Println(f(30))

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	f1, f2 := calc4(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7

}
