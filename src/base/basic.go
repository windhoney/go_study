package main

import (
	"fmt"
	"math/cmplx"
)

func variable() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func aaa() {
	var a int = 1
	var b string = "2"
	fmt.Println(a, b)
}

func bbb() {
	var abc = 123
	var cc = "213"
	fmt.Println(abc, cc)
}

var dd = 123 //函数外必须有var关键字  包内部变量  不存在全局变量的说法
var (
	name = 1
	aaa1 = 3
)

func ccc() {
	abc := 12 //第一次用：
	fmt.Println(abc)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

//func ttt()  {
//	a,b :=3,4
//	var c int
//	c= math.Sqrt();
//
//}

//内建变量类型
//bool string
//u无符号整数
//不加u有符号整数 int8规定长度，
//不规定长度跟操作系统来32 32 64-64
//ptr指针
//byte rune(go语言的char)32位--4字节int32
//float32 64---complex64--float32位 128==float64位 负数
//支持复数类型
//强制类型转换
//强制的
//--------
//常量
//常量名字 不用全部大写，普通变量命名  go语言大小写有含义
//可以作为各种类型使用
//枚举类型
// const(
//  php =1//iota 自增值//  b = 1<<10*iota
//  java = 2
//)
func con()  {
	const  sss string = "12312"
	const a = 23  //常量不用转类型 不确定类型可以是任何
}

func main() {
	//fmt.Printf("hello world1")
	euler()
}
