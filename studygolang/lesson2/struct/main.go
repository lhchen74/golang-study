package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name, city string
	age        int8
}

type test struct {
	a, b, c, d int8
	// b int8
	// c int8
	// d int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 函数指定接收者之后就是方法
// 需要使用接收者调用
// func (p person) dream() {
// 	p.age = 100
// 	fmt.Printf("%s dream is eat & sleep\n", p.name)
// }

func (p *person) dream() {
	p.age = 100
	fmt.Printf("%s dream is eat & sleep\n", p.name)
}

type myInt int // 可以为当前包任意类型添加方法

func (m *myInt) sayHi() {
	fmt.Println("Hello myInt")
}

// 结构体模拟继承
type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Println("move...")
}

type dog struct {
	feet int
	*animal
}

func (d *dog) bark() {
	fmt.Printf("%s wangwang...\n", d.name)
}

// // Student is public type
// type Student struct {
// 	// 大写表示可以公开访问, 如果小写, json 序列化结果 为 {}
// 	// 因为小写在 json 包中无法访问
// 	ID     int
// 	Gender string
// 	Name   string
// }

// Student is public type
type Student struct {
	ID     int    `json:"id"` // json 序列化结果为小写 id
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	var babb = person{
		name: "babb",
		city: "suzhou",
		age:  26,
	}
	fmt.Println(babb)

	var julian = person{}
	julian.name = "julian"
	fmt.Println(julian)

	// 得到对应结构体的指针
	var owen = new(person)
	owen.name = "owen" // <=> (*owen).name = "owen"
	fmt.Println(owen, *owen)

	var yoyo = &person{} // <=> var owen = new(person)
	yoyo.name = "yoyo"
	fmt.Println(yoyo, *yoyo)

	var sunny = person{
		"sunny",
		"suzhou",
		18,
	}

	var lucy = &person{
		name: "lucy",
	}

	fmt.Println(sunny, lucy, *lucy)

	n := test{
		1, 2, 3, 4,
	}
	// 结构体占用一块连续的内存。
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	// 结构体是值类型
	p1 := person{
		name: "p1",
	}
	p2 := p1
	p2.name = "p2"
	fmt.Println(p1, p2)
	p3 := &p1
	p3.name = "p3"
	fmt.Println(p1, p3, *p3)

	p4 := newPerson("np", "suzhou", 18)
	fmt.Printf("%v\n", p4)

	// var p5 = person{
	// 	name: "babb",
	// 	age:  26,
	// 	city: "suzhou",
	// }
	// p5.dream() // <=> (&p5).dream()
	// fmt.Println(p5.age)

	var p5 = &person{
		name: "babb",
		age:  26,
		city: "suzhou",
	}
	p5.dream()
	fmt.Println(p5.age)

	var mi myInt
	fmt.Println(mi)
	mi.sayHi()

	var d = dog{
		feet: 4,
		animal: &animal{
			name: "ali",
		},
	}
	d.move()
	d.bark()

	var stu1 = Student{
		ID:     1,
		Gender: "male",
		Name:   "babb",
	}
	// 结构体序列化为 JSON
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON serialized failed!")
		fmt.Println(err)
	}
	fmt.Println(v)
	fmt.Println(string(v))

	// 反序列化
	str := "{\"ID\":1,\"Gender\":\"male\",\"Name\":\"babb\"}"
	var stu2 = &Student{}
	json.Unmarshal([]byte(str), stu2)
	fmt.Printf("%T\n", stu2)

}
