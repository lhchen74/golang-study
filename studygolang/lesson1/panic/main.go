package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {

	// defer一定要在可能引发panic的语句之前定义。
	// panic可以在任何地方引发，但recover只有在defer调用的函数中有效

	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// ### defer语句
// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
// 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
// 也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
// 比如：资源清理、文件关闭、解锁及记录时间等。
// ### defer执行时机
// 在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
// 而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
func deferTest() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// deferTest()

	// ???
	// fmt.Println(f1())
	// fmt.Println(f2())
	// fmt.Println(f3())
	// fmt.Println(f4())

	// ???
	// a := 1
	// b := 2
	// defer calc("1", a, calc("10", a, b))
	// a = 0
	// defer calc("2", a, calc("20", a, b))
	// b = 1
	// fmt.Println(a, b)

	// funcA()
	// funcB()
	// funcC()

	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	var a1 []int
	a1[0] = 100

}
