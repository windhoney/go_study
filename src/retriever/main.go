package main

import (
	"fmt"
	"retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"12312"}
	fmt.Println(download(r))
}
