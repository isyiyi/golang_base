package dataType

import "fmt"

func DataTypeUse() {
	// var name type = expression
	// type or expression 可以丢弃一个，但不能全部丢弃
	var s string = "Hello Golang"
	fmt.Println(s)
	
	// golang声明时默认初始化为对应类型的零值
	var s1 string
	s1 = "Hello Golang"
	fmt.Println(s1)

	var s2 = "Hello Golang"
	fmt.Println(s2)

	// 短变量声明，不能用在包级别变量的声明中
	s3 := "Hello Golang"
	fmt.Println(s3)

	// 多变量声明
	// 如果在var后跟数据类型，则多个变量必须是同一类型
	var s4, s5, s6 string = "Hello", "Golang", "World"
	fmt.Println(s4, s5, s6)

	// 如果在var后没有跟数据类型，则每个变量的类型可以不同
	var s7, t1, n1 = "Hello Golang", true, 100
	fmt.Println(s7, t1, n1)

	// 使用短变量声明多个变量时，每个变量的类型可以不同
	s8, t2, n2 := "Hello Golang", false, 120
	fmt.Println(s8, t2, n2)
	
}
