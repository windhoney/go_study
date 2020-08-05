package main

import (
	"fmt"
)

//fmt 标准库 https://studygolang.com/pkgdoc
//rune 相当于GO的char
//string函数
//fields Split Join
//Contains Index
//ToLower ToUpper
//Trim TrimRight TrimLeft

func main() {
	s := "Yes我爱慕课网！"
	fmt.Println(len(s))
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
}
