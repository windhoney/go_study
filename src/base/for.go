package main

import "fmt"

func main() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
	sum := 0
	for i = 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
