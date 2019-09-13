package main

import "fmt"

func changeString() {
	s1 := "big"
	bs1 := []byte(s1)
	bs1[0] = 'p'
	fmt.Println(string(bs1))

	s2 := "白萝卜"
	rs2 := []rune(s2)
	rs2[0] = '红'
	fmt.Println(string(rs2))
}

func main() {
	s := "hello, 中国"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	// rune 代笔一个 unicode 字符
	// for range 按照 rune 类型遍历
	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
	}

	changeString()
}
