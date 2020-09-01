package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	input, err := inputReader.ReadString('\n')
	fmt.Println(input, err)
}
