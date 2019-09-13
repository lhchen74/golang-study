package main

import "fmt"

// 切片是一个拥有相同类型元素的可变长度序列, 它是基于数组类型的一层封装,
// 切片是引用类型, 颞部包含 `地址`、 `大小`、`容量`
// 切片一般用于快速的操作一块数据集合
func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Printf("a:%T, b:%T\n", a, b)

	var c []int
	// c = a[0:2]
	c = a[:]
	fmt.Println(c)
	fmt.Println(len(c), cap(c))

	// Array
	x := [...]string{"python", "ruby", "javascript", "php", "nodejs", "kotlin"}
	// Slice
	y := x[1:4]
	fmt.Println(len(y), cap(y))

	// 切片每次数量扩容为原来的两倍
	a1 := []int{}
	a1 = append(a1, 1)
	fmt.Printf("a:%v  len:%d  cap:%d  ptr:%p\n", a1, len(a1), cap(a1), a1)
	a1 = append(a1, 1)
	fmt.Printf("a:%v  len:%d  cap:%d  ptr:%p\n", a1, len(a1), cap(a1), a1)
	a1 = append(a1, 1)
	fmt.Printf("a:%v  len:%d  cap:%d  ptr:%p\n", a1, len(a1), cap(a1), a1)
	a1 = append(a1, 1)
	fmt.Printf("a:%v  len:%d  cap:%d  ptr:%p\n", a1, len(a1), cap(a1), a1)
	a1 = append(a1, 1)
	fmt.Printf("a:%v  len:%d  cap:%d  ptr:%p\n", a1, len(a1), cap(a1), a1)
	a1 = append(a1, 1)
	fmt.Printf("a:%v  len:%d  cap:%d  ptr:%p\n", a1, len(a1), cap(a1), a1)
	// a:[1]  len:1  cap:1  ptr:0xc4200161e8
	// a:[1 1]  len:2  cap:2  ptr:0xc420016210
	// a:[1 1 1]  len:3  cap:4  ptr:0xc42001c420
	// a:[1 1 1 1]  len:4  cap:4  ptr:0xc42001c420
	// a:[1 1 1 1 1]  len:5  cap:8  ptr:0xc420012180
	// a:[1 1 1 1 1 1]  len:6  cap:8  ptr:0xc420012180

	// 切片是引用类型
	b1 := []int{1, 2, 3}
	c1 := b1
	c1[0] = 100
	fmt.Println(b1, c1)

	// var d1 []int // 未申请内存
	// copy(d1, c1)
	// fmt.Println(d1) // []

	// var d1 []int
	// d1 = make([]int, 3, 3)
	// copy(d1, c1)
	// fmt.Println(d1)

	d1 := []int{0, 0, 0}
	copy(d1, c1)
	fmt.Println(d1)

	e1 := []string{"python", "ruby", "javascript", "php", "nodejs", "kotlin"}
	e1 = append(e1[:1], e1[2:]...)
	e1 = append(e1, []string{"scala", "haskell"}...)
	fmt.Println(e1)
}
