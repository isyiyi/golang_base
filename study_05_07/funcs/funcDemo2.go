package funcs

import "fmt"

type student struct {
	name string
	age int
}

type Path []int

func (p Path) app (x int)Path {
	p = append(p, x)
	return p
}

func (stu *student) string() {
	fmt.Println((*stu).name, (*stu).age)
}

func FuncsDemo1() {
	var p Path
	p = p.app(20)
	fmt.Println(p)

	var stu = student{name: "kobe", age: 37}
	stu.string()
}

type newInt int

func (n *newInt) update(x newInt){
	*n = x
}

func FuncsDemo2() {
	var num newInt = 90
	var numPtr = &num
	var numPtr2 = numPtr
	numPtr2.update(10)
	fmt.Println(num)
}

type ast struct {
	name string
}

func (a ast) printg() {
	fmt.Println(a.name)
}

type bst struct {
	*ast
	age int
}

func (b bst) printg() {
	fmt.Println(b.age)
}

func FuncsDemo3() {
	var b bst = bst{&ast{"kobe"}, 37}
	// b.ast.printg()
	var c bst
	c.ast = b.ast
	fmt.Println(*(b.ast))
	fmt.Println(b)
	fmt.Println(c)
}

func FuncsDemo5() {
	var num1, num2 = 90, 80
	var f func(int, int) int
	f = add
	fmt.Println(f(num1, num2))
	f = sub
	fmt.Println(f(num1, num2))
	var d dst = dst{"kobe", 20}
	f = d.addAge
	fmt.Println(f(num1, num2))
}

func add(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

type dst struct{
	name string
	age int
}

func (d dst) addAge(num1, num2 int) int {
	return d.age + num1 + num2
}

func (d dst) printg() {
	fmt.Println(d.name, d.age)
}	

func FuncsDemo4() {
	var d = dst{"kobe", 37}
	// f := d.printg
	var f func() = d.printg
	f()
	fmt.Printf("%#T\n", f)
}

func FuncsDemo() {
	var dsts = []dst{dst{"kobe", 37}, dst{"james", 20}}
	var f func(dst) = dst.printg
	for _, v := range dsts {
		f(v)
	}
>>>>>>> dev
}
