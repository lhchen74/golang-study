package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["java"] = 90
	fmt.Println(m1)

	m2 := map[string]string{
		"java":       "good",
		"python":     "normal",
		"javascript": "bad",
	}
	fmt.Println(m2)

	v, ok := m2["ruby"]
	if !ok {
		fmt.Println("Not Exists")
	} else {
		fmt.Println(v)
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	delete(m2, "java")
	fmt.Println(m2)

	rand.Seed(time.Now().UnixNano())
	var sourceMap = make(map[string]int, 10)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		sourceMap[key] = value
	}

	// 取出 map 的 key, 存入 slice
	var keys = make([]string, 0, 10)
	for key := range sourceMap {
		keys = append(keys, key)
	}

	// 对 slice 排序
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, sourceMap[key])
	}

	// 切片中的元素为map类型
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	// syntax error: unexpected comma, expecting type
	// mapSlice[0] = make(map[string][string], 10)
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "babb"
	mapSlice[0]["age"] = "18"
	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["hobby"] = "coding"
	for index, value := range mapSlice {
		fmt.Printf("index:%d, value:%v\n", index, value)
	}

	// map中值为切片类型
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
