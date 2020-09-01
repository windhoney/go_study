package main

import "fmt"

func main() {
	abc := 8
	if num := abc; num > 1 {
		fmt.Println("大于1")
	} else if num > 6 {
		fmt.Println("大于6")
	} else {
		fmt.Println("不大于")
	}
}
