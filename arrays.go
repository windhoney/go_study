package main

import "fmt"

//数组
//遍历
//rang 美观 意义明确
//C++没有类似能力
//java/python for each value 不能同时获取i v
//数组是值类型 非引用
//调用函数 func(arr [10]int) 拷贝数组 值传递不会改变数组本身
//go语言一般不直接使用数组

//func printArray(arr [5]int) {
func printArray(arr *[5]int) { //指针 会改变数组
	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int //[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]] 4行5列 4个 长度为5的int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//for i, v := range arr3 {
	//for _, v := range arr3 {
	//	fmt.Println(v)
	//}
	//
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	printArray(&arr3)
	printArray(&arr1)
	fmt.Println(arr1, arr2, arr3)
}
