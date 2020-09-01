package main

import "fmt"

func main() {
	arr := [3]int{1, 23, 3}

	arr2 := [3][4]int{{1, 2, 3, 4}, {3, 3, 4, 5}, {4, 4, 5, 8}}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[0][0])
}
