package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 5, 7, 8}
	v := 8
	for key, val := range arr {
		for i := key + 1; i < len(arr); i++ {
			if val+arr[i] == v {
				fmt.Printf("和为%d的两个元素的下标为（%d,%d）\n", v, key, i)
			}
		}
	}
}
