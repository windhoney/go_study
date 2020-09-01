package main

func main() {

}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	B
	test3()
	test4()
}

type TT struct {
}

func (t TT) test2() {

}

func (t TT) test1() {

}

func (t TT) test3() {

}
