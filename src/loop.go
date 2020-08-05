package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
//for if 后面的条件没有括号
//if 条件里也定义变量
//没有while
//switch 不需要break 也可以直接switch多个条件


func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	for  {
		fmt.Println("abc")
	}

}

/**
for 条件里不需要括号
for的条件里可以省略初始条件，结束条件，递增表达式
*/
func main() {
	fmt.Println(
		convertToBin(5), //101
		//对2取模 奇数第一位1
		//13/2=6/2=0
		//6/2=3/2=1
		//3/2=1 1
		//1011---》1101
		convertToBin(13),
		convertToBin(2342342),
		convertToBin(0),
	)
	printFile("abccc.txt")
	//forever()
}
