package main

import "fmt"

//map[K]V,map[K1]map[K2]v
//创建 make(map[string]int)
//获取元素 m[key]
//key不存在时 获得value类型的初始值
//用value,ok :=m[key] 来判断时候存在key
//用delete删除一个key
//遍历
//rang 遍历key 或者key value
//不保证遍历顺序 如需排序 需手动对key排序 加到slice能排序
//len 元素个数
//key类型
//map使用哈希表 必须可以比较相等
//除了slice map function 的内建类型都可以作为key
//struct类型不包含上述字段 也可以作为key  编译时检查
func main() {
	m := map[string]string{
		"name":    "zz", //无序的
		"course":  "golang",
		"site":    "github",
		"quality": "good",
	}
	m2 := make(map[string]int)    //==empty
	var m3 = make(map[string]int) //==nil
	fmt.Println(m, m2, m3)
	//遍历
	for k, v := range m { //_,v
		fmt.Println(k, v) //
	}
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	courseName1, ok := m["course1"] //不存在打印空值  string-空字符 zero value
	fmt.Println(courseName1, ok)
	if courseName, ok := m["course1"]; ok {
		fmt.Println("exist", courseName)
	} else {
		fmt.Println("not exist")
	}
	//删除
	delete(m, "name")
	fmt.Println(m)
}
