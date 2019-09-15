package main

import (
	"fmt"
	"time"
)

// Sayer 接口
type Sayer interface {
	say()
}

type dog struct{}

type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func showType(a interface{}) {
	fmt.Printf("%T\n", a)
}

type Flyer interface {
	fly()
}

type bird struct{}

type chicken struct{}

// 值接收者
func (d bird) fly() {
	fmt.Println("鸟会飞")
}

// 指针接收者
func (d *chicken) fly() {
	fmt.Println("鸡也会飞, 😄")
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪

	// 定义一个空接口
	var i interface{}
	i = 100
	i = "hello"
	i = false
	i = struct{ name string }{name: "babb"}
	fmt.Println(i)
	showType(i)
	showType(time.Second)
	showType(time.Now())

	// map 的值为空接口 interface {}
	var stuInfo = make(map[string]interface{}, 100)
	stuInfo["java"] = 100
	stuInfo["ruby"] = true
	stuInfo["python"] = "哈哈"
	fmt.Println(stuInfo)

	// 类型断言
	var x1 interface{}
	x1 = 100
	v, ok := x1.(int)
	if !ok {
		fmt.Printf("not int\n")
	} else {
		fmt.Printf("%d is int\n", v)
	}

	justifyType(v)

	var f Flyer
	var tom = bird{}
	f = tom

	var tom2 = &bird{} // tom2 是 *bird 类型
	f = tom2           // f 可以接收 *dog 类型
	f.fly()

	// var k = chicken{}
	// f = k             // f 不可以接收 chicken
	var k = &chicken{} // f 可以接收 *chicken 类型
	f = k

	// 接口是由 类型和值 两部分组成
	var ii interface{}
	var c int = 1
	ii = c
	value, ok := ii.(int)
	if ok {
		fmt.Printf("%d is int", value)
	}
}
