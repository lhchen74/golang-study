package main

import "fmt"

func main() {
	var a = [3]string{"ruby", "python", "php"}
	d := [...]int{1, 2, 3}
	fmt.Printf("a:%v, a:%T; d:%v, d:%T\n", a, a, d, d)

	// 根据索引初始化
	var e [10]int
	e = [10]int{9: 1}
	fmt.Println(e)

	for _, v := range a {
		fmt.Println(v)
	}

	var f = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	var g = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(f, g)

	// 数组是值类型

	h := [2]int{1, 2}
	i := h
	i[0] = 66
	fmt.Println(h, i)

}
