package main

import (
	"fmt"
	"io/ioutil"
)

//条件
//if的条件里可以赋值
//if的条件里赋值的变量作用域就在这个if语句里
func main() {
	const filename = "abc.txt"
	//cc, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println(12312)
	//} else {
	//	fmt.Printf("%s\n", cc)
	//	fmt.Printf("%s\n", 333)
	//}
	if cc, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
		fmt.Println(12312)
	} else {
		fmt.Printf("%s\n", cc)
		fmt.Printf("%s\n", 333)
	}
	//fmt.Println(cc)//未定义
}
