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
