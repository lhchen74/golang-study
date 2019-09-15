package main

import (
	"fmt"
	"time"
)

func printTime(t time.Time) {
	timeStr := t.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)
}

func calcTime() {
	startTime := time.Now()
	start := time.Now().UnixNano() / 1000
	fmt.Println("时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。")
	time.Sleep(time.Millisecond * 30)
	end := time.Now().UnixNano() / 1000
	fmt.Printf("cost %d millisecond\n", end-start)
	fmt.Printf("cost %d millisecond\n", time.Since(startTime))
}

func main() {
	now := time.Now()
	printTime(now)
	calcTime()
}
