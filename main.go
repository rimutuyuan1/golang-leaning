package main

import (
	"fmt"

	"unicode/utf8"
)

func main() {
	Map()
}

//数组是值类型
//值传递 方法调用会拷贝 引用传递会直接改变值
//指针 引用传递
// go一般不直接使用数组 用切片代替 Slice
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//slice 内部存储结构：视图  不是值类型
func slice() {
	arr := [...]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	s := arr[3:]
	fmt.Println(s)
}

func reSlice() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)
}

func addSlice() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := append(arr[:], 10)
	fmt.Println(s1)
}

func Map() {
	m := map[string]string{
		"name": "viki",
		"sex":  "man",
		"age":  "18",
	}
	m1 := make(map[string]int)

	var m2 map[string]string

	for i, v := range m {
		fmt.Println(i, v)
	}
	fmt.Println("-----------")
	fmt.Println(m1)
	fmt.Println(m)
	fmt.Println(m2)

	fmt.Println("Getting values ")
	name, ok := m["name"]
	fmt.Println(name, ok)
	if name, ok := m["nama"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("empty")
	}

	fmt.Println("delete values")
	delete(m, "age")
	fmt.Println(m)

	fmt.Println("---------")
	fmt.Println(utf8.RuneCountInString("你好 我是王鑫涛"))
	fmt.Println(len("你好 我是王鑫涛"))
	fmt.Println(len([]rune("你好 我是王鑫涛")))

	fmt.Println("---------")
	str := "你好 我是王鑫涛"
	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

}
