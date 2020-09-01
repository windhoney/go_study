package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 23, 3, 5, 4, 45, 6546, 4, 678, 8}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
