package main

import (
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		g = "null"
	}

	return g
}

//条件
//if的条件里可以赋值
//if的条件里赋值的变量作用域就在这个if语句里
func main() {
	//const filename = "abc.txt"
	//cc, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println(12312)
	//} else {
	//	fmt.Printf("%s\n", cc)
	//	fmt.Printf("%s\n", 333)
	//}
	//if cc, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//	fmt.Println(12312)
	//} else {
	//	fmt.Printf("%s\n", cc)
	//	fmt.Printf("%s\n", 333)
	//}
	//fmt.Println(cc)//未定义

	//switch
	//自动break 除非是用fallthrough
	//panic 报错 终止执行
	// switch 后可以没有表达式
	fmt.Println(
		grade(59),
		grade(61),
		grade(81),
		grade(91),
		grade(110),
	)
}
