package main

import "fmt"

func recv(ch chan bool) {
	ret := <-ch // block
	fmt.Println(ret)
}

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	var ch1 chan int        // 声明 int 类型 channel
	ch1 = make(chan int, 1) // make 函数初始化 channel 容量为 1, 不能连续多次发送消息, 需要发送一次接收一次, 之后才可以再发送
	ch1 <- 10               // send
	// ch1 <- 10            // fatal error: all goroutines are asleep - deadlock! 容量不足, 一直等待
	ret := <-ch1 // accept
	fmt.Println(ret)
	ch1 <- 20
	ret = <-ch1
	fmt.Println(ret)
	close(ch1) // close
	ret = <-ch1
	fmt.Println(ret) // 0 close 之后再接收会得到类型的 0 值
	// close(ch1) // panic: close of closed channel 关闭已经关闭的 channel 会引发 panic

	// 1 表示有缓冲通道, 每次连续接收一个值
	ch := make(chan bool, 1)
	ch <- true
	fmt.Println(len(ch), cap(ch))
	go recv(ch)
	ch <- true // ch 有值, recv 函数 release
	fmt.Println("Main end")

	var ch2 = make(chan int, 100)
	send(ch2)

	// for {
	// 	ret, ok := <-ch2
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(ret)
	// }

	for ret := range ch2 {
		fmt.Println(ret)
	}
}
