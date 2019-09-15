package main

import (
	"fmt"
	"time"

	// m 包的别名
	m "code.jbn.com/studygolang/lesson2/package/mathpkg"
	// 只执行 init() 初始化函数
	// _ "code.jbn.com/studygolang/lesson2/package/mathpkg"
)

var a int = 10

func init() {
	fmt.Println(a)
	fmt.Println("init exeute before main")
}

// call 顺序 main -> A -> B
// init 顺序 B(init) -> A(init) -> main(init)

func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func main() {
	// sum := mathpkg.Add(100, 200)
	// fmt.Println(sum)
	sum := m.Add(100, 200)
	fmt.Println(sum)

	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	timestampDemo()

	// tickDemo()

	formatDemo()

}
