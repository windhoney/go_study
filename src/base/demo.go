package main

import (
	"fmt"
	"time"
)

type Student struct {
	name   string
	number string
}

var studentList []Student
var isLogin = true

func main() {
	login()
}

//进入系统
func login() {
	var operate string
	var studentNum string
	var studentName string
	var num int
	fmt.Println("欢迎光临学生管理系统")
	fmt.Println("1.查看所有学生")
	fmt.Println("2.新增学生")
	fmt.Println("3.删除学生")
	fmt.Println("4.编辑学生")
	fmt.Println("5.退出系统")
	fmt.Println("请输入要执行的操作：")
	fmt.Scanln(&operate)
	switch operate {
	case "1":
		allStudent(&studentList)
	case "2":
		fmt.Println("请输入学生学号：")
		fmt.Scanln(&studentNum)
		fmt.Println("请输入学生姓名：")
		fmt.Scanln(&studentName)
		addStudent(studentNum, studentName)
	case "3":
		fmt.Println("请输入要删除的学生序号：")
		fmt.Scanln(&num)
		delStudent(num)
	case "4":
		fmt.Println("请输入要编辑的学生序号：")
		fmt.Scanln(&num)
		fmt.Println("请输入学生学号：")
		fmt.Scanln(&studentNum)
		fmt.Println("请输入学生姓名：")
		fmt.Scanln(&studentName)
		editStudent(num, studentNum, studentName)
	default:
		logout()
	}
	time.Sleep(1 * time.Second)
	if isLogin {
		login()
	}
}

//查看所有学生
func allStudent(studentList *[]Student) {
	if len(*studentList) == 0 {
		fmt.Println("没有学生")
	} else {
		for k, v := range *studentList {
			fmt.Printf("%d.姓名:%s，学号:%s \n", k+1, v.name, v.number)
		}
	}
}

func addStudent(studentNum, studentName string) {
	//studentInfo = Student{studentName, studentNum}
	studentList = append(studentList, Student{studentName, studentNum})
	fmt.Println("学生添加成功")
}

//删除学生
func delStudent(num int) {
	studentList = append(studentList[:num-1], studentList[num:]...)
	fmt.Println("删除成功")
}

//编辑学生
func editStudent(num int, studentNum, studentName string) {
	studentList[num-1] = Student{studentName, studentNum}
	fmt.Println("学生修改成功")
}

//退出系统
func logout() {
	fmt.Println("退出系统成功")
	isLogin = false
}
