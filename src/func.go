package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q
	default:
		panic("unsupported:" + op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	//return
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

//要点
//返回值类型写在最后面
//可返回多个值
//函数作为参数
//没有默认参数，可选参数
//

//函数返回多个值时可以起名字
//仅用于非常简单的函数
//对于调用者而言没有区别
func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(5, 3)

	fmt.Println(q, r)

	//fmt.Println(apply(pow, 3, 4))
	//Calling function main包名.main 方法.func1未知名 with args (3,4)
	//函数式编程--匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 2, 3, 23, 2, 32, 3))
}
