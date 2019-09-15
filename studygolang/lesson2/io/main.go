package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readByByte(filename string) {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	for {

		// var tmp [128]byte
		// n, err := file.Read(tmp[:])
		var tmp = make([]byte, 128, 128)
		n, err := file.Read(tmp)

		if err == io.EOF {
			fmt.Println(string(tmp))
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("read %d byte\n", n)
		fmt.Println(string(tmp))

	}
}

func readByLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(str)
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(str)
	}
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func basicWrite() {
	// os.O_WRONLY|os.O_CREATE|os.O_APPEND 写入 | 没有就创建 | 追加写
	// 0755 属主 读（04）写（02）执行（01), 其他用户 读 & 执行
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write([]byte("hello"))
	file.WriteString(" world\n")
}

func bufioWrite() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("Life is short, I use python\n")
	writer.Flush()
}

func ioutilWrite() {
	err := ioutil.WriteFile("test.txt", []byte("PHP is the best language of the world! \n"), 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// normal read data
	// readByByte("./main.go")

	// buffio read data
	// readByLine("./main.go")

	// ioutil read data
	// readFile("./main.go")

	basicWrite()

	bufioWrite()

	// ioutilWrite()

}
