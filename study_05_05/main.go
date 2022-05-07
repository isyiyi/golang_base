package main
// 包声明，指定该文件属于哪个包，一个包可以有多个.go文件组成
// 一个包中只能有一个main函数，且main包只能有一个.go文件

import(
//	"study_05_05/packageUse"
	"fmt"
//	"study_05_05/dataType"
	pdt "study_05_05/proDataType"
) 
// 导入其他的包，多个时可使用import()

// import之后就是 包级别的常量、变量、类型、函数的声明
var Count int = 90

const Page int = 100

type Student struct {
	name string
	age  int
}

func main() {
	// packageUse.OsUse()
	// packageUse.Dup1()
	// packageUse.DupFromFile()
	// packageUse.DupFromFile2()
	// packageUse.DupHomeWork()
	// packageUse.MainProcess()
	var x int = 0
	switch x {
	case 1, 2, 3 :
		fmt.Println("1, 2, 3")	
	case 0, 4, 5 :
		fmt.Println("0, 4, 5")
	default :
		fmt.Println("< 5")
	}

	s := &Student{
		name: "kobe",
		age: 37,
	}
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println("**************************************")
	// dataType.DataTypeUse()
	dataType.UseArea()
	dataType.UseArea2()
	// dataType.UseArea()
	// dataType.UseArea2()
	// dataType.AllDataType()
	// dataType.BitCalc()
	// dataType.StrDemo()
	
	// pdt.ArrayDemo()	
	// pdt.SliceDemo()	
	pdt.MainProcess2()
	// pdt.StructDemo2()
	
}
