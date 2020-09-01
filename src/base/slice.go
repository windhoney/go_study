package main

import "fmt"

//slice切片  视图view
//半开半闭区间 [2:6] 2345(0123456)
//Slice本身没有数据,是对底层array的一个view

//[]没写长度就是切片 写长度就是数组

//修改 会改变数值本身
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	//fmt.Println("a[2:6]", s)
	//fmt.Println("a[:6]", arr[:6])
	//s1 := arr[2:]
	//s2 := arr[:]
	//updateSlice(s1)
	//fmt.Println("s1", s1)
	//fmt.Println("a[2:]", arr[2:])
	//fmt.Println("a[:]", arr[:])
	//fmt.Println("--reslice")
	//fmt.Println(s2)
	//s2 = s2[:5]
	//fmt.Println(s2)
	//s2 = s2[2:]
	//fmt.Println()
	//fmt.Println(s2)
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//slice 扩展
	//slice实现 ptr起点 len cap到最后
	//slice 可以向后扩展 不可以向前扩展
	//s[i] 不可以超越len(s) 向后扩展不可以超越底层数组cap(s)
	fmt.Println("extending slice")
	s1 := arr[2:6] //2 3 4 5
	s2 := s1[3:5]  //5 6 最大扩展到6[3:6]
	fmt.Println(s1, s2)
	fmt.Println(s1, "\n", len(s1), cap(s1)) //4 6
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s3, 12)
	//s4 s5重新分配了一个数组
	//垃圾回收，原来数组没人用就回收
	//向slice添加元素
	//添加元素时如果超越cap 系统会重新分配更大的底层数组
	//由于值传递的关系，必须接收append的返回值
	//s=append(s,val)
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)

	s8 := make([]int,2,10)
	fmt.Println(s8)

	s10 := []int{1,2,3,4}
	s11 := s10[3]
	s10 = append(s10, 5,6,7,8)
	s10[3] = 100
	fmt.Println(s10,s11)

}
