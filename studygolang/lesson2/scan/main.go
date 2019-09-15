package main

import "fmt"

func main() {
	var (
		name string
		age  int
		male bool
	)
	fmt.Println(name, age, male)
	// fmt.Scan(&name, &age, &male)
	// fmt.Scanln(&name, &age, &male)
	fmt.Scanf("name:%s age:%d male:%t\n", &name, &age, &male)
	fmt.Println(name, age, male)

	// name:babb, age:18, male:true
}
