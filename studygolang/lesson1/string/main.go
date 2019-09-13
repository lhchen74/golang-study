package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "babb"
	fmt.Println(len(s1))

	s2 := "goland"
	s3 := "python"
	fmt.Println(s2 + ": " + s3)
	s4 := fmt.Sprintf("%s: %s", s2, s3)
	fmt.Println(s4)

	ret := strings.Split(s4, ": ")
	fmt.Println(ret)

	ret2 := strings.Contains(s4, ":")
	fmt.Println(ret2)

	ret3 := strings.HasPrefix(s4, "g")
	ret4 := strings.HasSuffix(s4, "python")
	fmt.Println(ret3, ret4)

	s5 := "apple"
	fmt.Println(strings.Index(s5, "p"), strings.LastIndex(s5, "p"))

	a6 := []string{"python", "ruby", "php"}
	fmt.Println(strings.Join(a6, "*"))
}
