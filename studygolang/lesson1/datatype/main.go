package main

import "fmt"

const pi = 3.14

const (
	n1 = iota
	n2
	_
	n4
)

// iota 在 const 关键字出现时被重置为 0
// const 中每增加一行变量声明将使 iota 计数一次 (可以理解为 const 语句块的行索引)
// 如果等号右边不写和上一行一样
const (
	a1 = iota
	a2 = 100
	a3 = iota
	a4 = iota
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << 10
	MB
	GB
	TB
	PB
)

const (
	aa, bb = iota + 1, iota + 2
	cc, dd
	ee, ff
)

func main() {
	var a string = "a"

	var (
		b string
		c int
		d bool
	)

	var e = 'e'

	f := "f"

	fmt.Println(a, b, c, d, e, f)

	fmt.Println(n1, n2, n4)

	fmt.Println(a1, a2, a3, a4)

	fmt.Println(MB, GB)

	fmt.Println(aa, bb, cc, dd, ee, ff)
}
