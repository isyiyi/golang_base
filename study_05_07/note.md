1. 函数的类型只与参数列表和返回值列表的顺序和类型有关，与参数名字和返回值名字无关，与函数体也没有关系
```golang
func f1(num1, num2 int, str1 string) (num3 int, str2 string){
	num3 = num1 + num2
	str2 = str1 + "|"
	return
}

func f2(num4, num5 int, str3 string) (int, string) {
	num1 := num4 * num3
	str1 := str3 + str3
	return num1, str1
}

// 虽然这两个函数的参数名字和返回值名字都不相同，并且返回值的方式也不同，函数体的内容更是天差地别，但是它们的参数和返回值列表的类型和数量、顺序都是一样的，则这两个函数的类型就是相同的
var f func(int, int, string) (int, string)
f = f1
f = f2
```
2. 函数多返回值的小窍门
```golang
// 1. 一个多返回值函数可以作为另一个多返回值函数的返回值
func f1() (int, string) {
	xxx
	return 2, "hello, golang"
}

func f2() (int, string) {
	xxxx
	return f1()
}

// 2. 一个多返回值函数可以作为另一个多返回值函数的参数

func f3(int, string) {
	xxxx
	return
}

f3(f1())	// golang会自动解析返回值，将值赋给对应位置

// 3. 如果返回值在函数声明时已经被命名，在return时不需要再带上名字
func f4() (x int, s string) {
	x := 90
	s := "hello, golang"
	return
}

```
3. 对于错误消息的处理（go语言中的错误处理机制与其他语言有很大不同，把返回的错误作为普通的值）
	1. 向上层返回，并在本层返回时添加错误信息，确保main函数收到之后得到一条清晰的因果链
	2. 出错之后间隔一段时间重试，实在不行再向上报错
	3. 使用log库输出日志，并将错误传递给调用者。如果这个错误非常严重，应使用`os.Exit(1)`退出或`log.Fatal()`打印错误并退出
	4. 不是严重错误，可以通过log库将日志打印出来
	5. 有的错误可以直接忽略，即调用时压根不使用变量接收
4. 函数变量不可比较，函数变量、slice、map在未初始化之前的值都为nil
5. 匿名函数作为返回值会形成闭包，也就是在被调用函数运行结束之后，内部的局部变量仍然被保存，闭包不仅是一段代码，它还保存了状态。在匿名函数内部被引用的外部变量，在函数返回之后它们的值仍然会被保留。
6. 闭包会引用外部语法块的变量，所以有时候要避免在闭包中引用被重复使用的变量。
```golang
func f() {
	var funcs = make([]func())
	for i := 0; i < 10; i++ {
		funcs = append(funcs, func(){
			fmt.Println(i)
		})	
	}	
	for _, v := range() {
		v()	// 所有的输出都是9，因为i在闭包内被引用，所以i的值会随着for循环一直改变，最终停留在9
	}
}
```
7. defer语句会在return之后执行，且有多个defer语句的话会逆序执行。因为defer会在return之后且调用者得到返回值之前执行，所以当defer执行时return的返回值已经确定，如果此时对返回值进行修改，那么调用者得到的返回值将会被改变
```golang
func f()(x int){
	defer func() {
		x = 100
	}()
	x = 10
	return
}

// 此时调用f得到的结果将是100，因为虽然x=10已经被返回，但是defer在return之后执行，所以x被修改
```
8. 当发生异常时，程序的执行顺序
	1. 程序终止
	2. goroutine中所有的defer语句和函数会执行(此时使用runtime.Stack()可以看到函数栈信息，因为defer是在栈清理之前调用的)
	3. 程序异常退出并打印出退栈信息
9. 开发者可以手动使用panic("info")来产生异常，但是仅限于那些非常严重的错误。
10. 当产生异常时，程序会直接终止运行，对于大型复杂的系统来说是非常严重的。所以使用recover()可以将程序从宕机状态恢复过来
	+ recover()只有在程序出现异常时才有效，正常的程序使用recover会返回nil且没有任何效果
	+ 使用recover()恢复过来的程序将不会把这个异常主动报出
```golang
func f() {
	s := 90
	defer func() {
		if p := recover(); p != nil {	// 程序正常时，recover不会有任何效果，且返回值为nil
			xxxx			// 程序出现宕机时，p!=nil，程序从宕机的地方转移到defer语句内执行
		}
	}()
	if s > 80 {
		panic("s > 80")
	}
}
```
11. 一般来说，recover函数在defer函数的内部调用，出现异常时，recover函数将程序从宕机状态恢复并返回宕机的值
12. 通常的做法，对于不是自己编写的包中的宕机不要去尝试恢复，因为此时的变量已经处于混乱的状态。如果是自定义的panic(xxx)，recover()的返回值就是panic函数的参数，通过判断宕机值可以知道问题的严重性，不严重的可以直接处理之后返回error，严重的可以再次主动出发panic，使程序结束  <font color="red">重要</font>
```golang
func f() {
	s := 90
	defer func() {
		switch p := recover; p {
		case nil:		// nil说明程序正常运行
		case "x > 80":		// 检测到宕机值是主动宕机，预料到的错误可以直接处理
			fmt.Println("x > 80, recover")
		default:		// 非预料到的错误，继续宕机过程，使程序结束
			panic(p)
		}
	}()
	if s > 80 {
		panic("x > 80")	// 主动触发panic
	}
}
```

---

13. 因为selector可以选择方法和字段，所以结构体的字段和方法名不能一致，否则会冲突。不同类型的字段和方法名可以相同，因为reciver类型不同
14. golang和其他语言不同，可以将方法绑定到任何类型上，对于基础类型和数组等，需要先使用type对其进行命名
15. 如果需要更新变量或者reciver对象太大时，可以使用指针类型作为reciver
16. **如果一个类型的一个方法使用了指针类型作为reciver，那么它的任何方法都应该有一个指针类型作为reciver的版本，即便并不需要这样做**
17. 本身 已经是指针类型的命名类型不允许指定方法
18. reciver为指针类型还是普通类型，在调用的时候都可以使用普通变量或指针类型中的任意一种来调用，因为编译器会进行自动取地址和取值
```golang
type student struct {
	name string
	age int
}

func (stu *student) String() {
	fmt.Println((*stu).name, (*stu).age)
	fmt.Println(stu.name, stu.age)		// 不使用*取地址也可以，编译器会进行隐式转换
}

func main () {
	var stu student = Student {name: "kobe", age: 37}
	(&stu).String()	// 正确
	stu.String() 	// 自动进行隐式转换
}
```

19. 如果定义了一个变量是某个命名类型的指针，应避免复制这个指针的值，因为如果复制之后且使用了这个复制的指针对象的方法，将会修改原来对象的值。（两个指针变量指向同一块内存）
```golang
type newInt int
func (xx *newInt) update(a newInt) {
	*xx = a
}

var tmp newInt = 90
var num = &tmp
var num2 = num
num2.update(10)		// 将会修改num的值

```

20. 即使是临时结构体，通过结构体内嵌，声明的变量仍然能够拿到嵌入结构体的方法
```golang
type student struct {
	name string
}

func (s student) printg() {
	fmt.Println(s.name)
}

var tmp = struct {
	student
	age int
}{age: 90}

tmp.printg()	// 通过定义临时结构体并内嵌student来使用student定义的方法
```

21. 如果子结构体和父结构体的方法名相同，如果直接调用，使用的是父结构体的方法；如果使用完整路径调用，将会使用指定的方法，但是要注意，任意类型的字段和方法名不能相同
```golang
type ast struct {
	name string
}

func (a ast) printg() {
	fmt.Println(a.name)
}

type bst struct {
	ast
	age int
}

func (b bst) printg() {
	fmt.Println(b.age)
}

var b bst = bst{ast{"kobe"}, 37}
b.printg()		// 37，直接调用父结构体的printg方法
b.ast.printg()		// "kobe"，根据匿名类型的路径，调用指定的方法
```

22. 如果结构体内嵌的子结构体为匿名类型，那么不论是普通类型还是指针类型，都可以直接通过子结构体名来访问
```golang
type ast struct {
	name string
}

type bst struct {
	ast
	age int
}

type cst struct {
	*ast
	school string
}

var b = bst{ast{"kobe"}, 37}
var d bst
d.ast = b.ast	// d = bst{ast{"kobe"}, 37}	// 通过ast匿名类型可以访问ast的内容

var c cst = cst{&ast{"kobe"}, 37}
var o cst
o.ast = c.ast	// o = cst{&ast{"kobe"}, 37}	// 将c中初始化的ast{"kobe"}的地址赋值给o
		// 通过ast匿名类型可以直接得到初始化的ast变量的地址
// 总结：无论子结构体匿名类型是普通类型还是指针类型，都可以通过匿名类型的名称直接得到
``` 

23. **方法变量和方法表达式**
	1. 方法变量的类型和函数变量的类型一模一样，区别在于方法变量是为函数指定了接收者；缺点在于，一旦指定了接收器，那么调用方法的对象就固定了，方法表达式可以解决这个问题
	2. 方法表达式相比于方法变量的优点在于，不需要事先指定接收器，声明时直接使用`命名类型.方法名`，它的类型相比于函数变量/方法变量的类型多了一个参数，即接收器（将原来的接收器指定为函数的第一个形参）
```golang
func add(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

var f func(int, int) int
var num1, num2 = 90, 80

f = add
res := f(num1, num2)	// res = 80 + 90 = 170
f = sub
res = f(num1, num2)	// res = 90 - 80 = 10

type ast struct {
	score int
}

func (a ast) printg(num1, num2 int) int {
	return a.score + num1 + num2
}

var a ast = ast{20}
f = a.printg		// 为方法指定接收器，虽然printg是方法，但是使用函数变量同样可以接收，因为类型相同
res = f(num1, num2)	// res = 20 + 80 + 90 = 190

// ==========================================方法表达式


type ast struct {
	name string
	age int
}

func (a ast) printg() {
	fmt.Println(a.name, a.age)
}

var asts = []ast{ast{"kobe", 37}, ast{"james", 38}}
var f func(ast) = ast.printg		// 虽然想要调用的printg方法没有参数，但是使用方法表达式时reciver要作为第一个参数，所以函数变量要将第一个参数命名为指定的类型，且reciver不需要指定为某个固定的对象，直接使用命名类型就好
for _, v := range asts {
	f(v)
}
```

24. golang中结构体中的字段无论是否导出，对于同一个包内的代码都是可见的，golang封装的单元是包而不是类型。（使用首字母大小写来区分是否对外包导出，而不是private、public等）
25. getter和setter函数，getter函数一般以字段的首字符大写命名，setter函数一般就是用Set+字段名字
```golang
type student struct {
	name string
	age int
}

func (s student) Name () string {
	return s.name
}
func (s student) SetName(name string) {
	s.name = name
}
```

26. 封装的优点：
	1. 对外隐藏函数的具体实现细节，使调用者只需要知道参数和调用方式，不需要了解具体实现，对于开发者更新算法内容也有益处
	2. 因为封装之后只有同一个包内的函数才能修改变量的值，所以可以防止调用者肆意修改对象的变量，造成逻辑混乱
	3. 因为调用者不能直接修改对象的值，所以不需要额外去检测值的合法性（开发者只用专心自己开发内容的合逻辑性）
