package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Printf("Hello World %d\n", i)
}

func a() {
	defer wg.Done() // 计数器 -1
	for i := 1; i < 100; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go hello(i)
	// }
	// fmt.Println("Main")
	// // time.Sleep(time.Second)
	// wg.Wait()

	// 设置当前程序并发时占用的CPU逻辑核心数
	runtime.GOMAXPROCS(2)
	// 计数器计数
	wg.Add(2)
	go a()
	go b()
	// 等待
	wg.Wait()
}
