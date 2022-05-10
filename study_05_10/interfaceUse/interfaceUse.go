package interfaceUse

import (
	"fmt"
	"flag"
	"strconv"
	"io"
	"bytes"
	"sort"
)

type student struct {
	name string
	age int
}

func (stu *student) string() {
	fmt.Println(stu.name, stu.age)
}

func (stu student) printg() {
	fmt.Println(stu.name, stu.age)
}

func Func1() {
	var stu = student {
		name : "kobe",
		age : 37,
	}
	var stu1 = student{}
	fmt.Println(stu)
	fmt.Printf("%p\n", &stu1)
	fmt.Printf("%p\n", &student{})

	stu1.string()
	// student{}.string()
	var stuPtr = &student{}
	stuPtr.string()
	student{}.printg()	
}

func Func2() {
	// testInterface(20, 30, 40)
	var args = []interface{}{20, true, "hello"}
	testInterface(args...)
}

func testInterface(nums ...interface{}) {
	fmt.Println(nums)
}

type Age int

func (a *Age) String()string {
 	return 	strconv.Itoa(int(*a))		// 没有什么意义，单纯为了实现Value接口
}

func (a *Age) Set(s string) error {
	// fmt.Println(s)
	*a = 33380
	return nil
}

func Func4() {
	var age Age = 90
	flag.CommandLine.Var(&age, "age", "set age")
	flag.Parse()
	fmt.Printf("%v\n", age)
	var w io.Writer
	fmt.Printf("%T, %[1]v\n", w)
}

func Func5() {
	var w io.Writer
	var b *bytes.Buffer
	w = b
	fmt.Printf("%T, %[1]v\n", w)
	fmt.Println(w == nil)
	fmt.Println(b == nil)
}

type player struct {
	name string
	score int
}

type players []player

func (p players) Len() int {
	return len(p)
}

func (p players) Less(i, j int) bool {
	return p[i].score > p[j].score
}

func (p players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Func6() {
	var ps players = []player{
		player{"kobe", 81},
		player{"james", 61},
		player{"wade", 76},
	}
	sort.Sort(sort.Reverse(ps))
	
	for k, v := range ps {
		fmt.Println(k, v)
	}
}


func Func7() {
	var nums = []int{1, 9, 8, 2, 4}
	newSlice := sort.IntSlice(nums)
	sort.Sort(newSlice)
	for _, v := range newSlice{
		fmt.Println(v)
	}	
}

func Func8() {
	var x interface{} = 90
	// y := x.(int)
	fmt.Printf("%T\n", x)
	// fmt.Println(y + 90)
	fmt.Printf("%T\n", x)
}

type Wow interface {
	printg()
}

type Wow2 interface {
	Wow
	printgg()
}

type newString string

func (n newString) printg() {
	fmt.Println(n)
}

type newBool bool

func (b newBool) printg() {
	fmt.Println(b)
}

func (b newBool) printgg() {
	fmt.Println(b, 20)
}

func Func9() {
	var n newBool = true
	var w Wow2 = n
	w.printg()
	// res := w.(Wow)
	var res Wow = w
	fmt.Printf("%T, %[1]v\n", res)
}

func Func3() {
	var w Wow = newBool(true)
	w = w.(newBool)
	fmt.Println(w)
}
