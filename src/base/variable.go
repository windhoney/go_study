package main

import "fmt"

func main() {
	name,sex := "zhang","nan"

	name,sex = "123","nan"
	fmt.Println(name,sex)
	fmt.Println("a\\nb")
}
