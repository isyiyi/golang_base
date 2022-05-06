package dataType

import "fmt"
import "unsafe"

type Object1 struct {
	s struct {}
	b byte
	a string
}

func AllDataType() {
	var num1 int = 90
	var num2 int32 = 100
	fmt.Println(int32(num1) + num2)	
	var x struct{}
	fmt.Println(unsafe.Sizeof(x))
	var s int
	fmt.Println(unsafe.Sizeof(s))
	o1 := Object1{}
	fmt.Println(unsafe.Sizeof(o1))	
	fmt.Println(o1.b)
	fmt.Println(o1.a)
}

func BitCalc() {
	var a, b int = 3, 5
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	var num int= 0665
	fmt.Printf("%o\n", num)
	fmt.Printf("%#o\n", num)
}

func StrDemo() {
	var s = "Hello world"
	fmt.Println(len(s))
	for i := 0; i < len(s); i ++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	var s2 = "中华人民共和国"
	fmt.Println(len(s2))
	fmt.Println(len([]rune(s2)))
	for i := 0; i < len([]rune(s2)); i ++ {
		fmt.Printf("%c\n", []rune(s2)[i])
	}

	for _, v := range s2 {
		fmt.Printf("%c\n", v)
	}
	var s3 = "\U0000002A"
	fmt.Println(s3)

	const (
		a = iota
		b
		c = 3
		d 
		e = iota
	)
	fmt.Println(a, " ", b, " ", c, " ", d, " ", e)
	// 0 1 3 3 4
	const (
		f = 1 + iota
		g
 		h
	)
	fmt.Println(f, " ", g, " ", h)
	const gg = 90
	fmt.Printf("%T\n", gg)
	var aa = 'c'
	fmt.Printf("%T\n", aa)
}
