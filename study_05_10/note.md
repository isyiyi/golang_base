1. 先前介绍的基本类型或大多数引用类型都是具体类型，具体类型指定了它所含数据的精确布局，还暴漏了这个精确布局的内部操作。（只要知道了这个具体类型的数据，就知道了它是什么和它能够干什么）
2. 接口类型是一种抽象类型，它并没有暴漏它所含数据的布局和内部结构，当然也没有这些数据的具体操作。它仅仅提供以下方法。（知道了这个接口，仅仅能够知道它包含哪些方法，并且这些方法因为类型的不同要做的事情也不同）
3. golang中的接口并不要求显式实现，只要某个类型的方法与接口中定义的所有方法类型和名称一致，那么这个类型就实现了这个接口
4. 接口同样可以和结构体一样进行嵌入，嵌入之后的接口会得到被嵌入接口的所有方法
```golang
type Writer interface {
	Write(p []byte)(n int, err error)
}

type ReadWriter interface {
	Writer
	Read(p []byte)(n int, err error)
}
```

5. 对接口类型赋值时，只要=右边的类型或接口实现了左边接口变量的所有方法，就可以进行赋值
```golang
var w io.Writer
w = os.Stdout	// os.Stdout是*os.File类型，且*os.File类型实现了io.Writer、io.Reader、io.Closer接口
w = new(bytes.Buffer)	// *bytes.Buffer类型实现了io.Writer、io.Reader接口
w = time.Second	// 错误，因为time.Second并没有Write方法

var wrc io.ReadWriteCloser
var w io.Writer
w = wrc		// 正确，=右边的接口类型实现了io.Writer接口，即拥有Write方法
wrc = w		// 失败，=右边的接口类型只有Write方法，并不拥有Read和Close方法
```

6. 对于命名类型，无论它的方法接收器是指针类型还是普通类型，通过普通类型或指针类型的变量都可以调用
```golang
type student struct {
	name string
}

func (stu *student) string() {
	fmt.Println(stu.name)
}

func (stu student) printg() {
	fmt.Println(stu.name)
}

var stu = student{}	// 不显式赋值时，所有的字段都被赋值为对应类型的零值
stu.string()		// 尽管reciver是指针类型，编译器会自动取地址

var stuPtr = student{}
(&stuPtr).printg()	// 尽管reciver是普通类型，编译器会自动解引用

var stuPtr2 = &student{}
stuPtr2.printg()	

student{}.string()	// 错误error，编译器无法对一个临时变量取地址并调用对应的方法
student{}.printg()	// 正确，直接通过临时变量来调用，编译器无需对其取地址（printg的接收器是普通类型）
```

7. 由6可知道，调用方法时无需特别注意接收器的类型是golang提供的语法糖，但**使用临时变量调用时就会出现问题，无法对一个临时变量隐式取地址，从而调用revicer为指针的方法**，所以，golang建议如果一个命名类型的方法存在指针类型为接收器的方法，那么所有的方法都应该有指针类型为接收器的版本
8. 只有接口暴漏的方法才可以通过接口变量来调用，如果一个类型实现了io.Writer方法，但是它自身存在Write、Read、Close等方法，io.Writer接口变量只能调用Write方法
9. 一个接口中定义的方法数越多，实现它的门槛就越高。空接口`interface{}`意为没有任何门槛的接口，所以任何值都可以赋值给它
```golang
var i interface{}
i = true
i = 90
i = "hello golang"
// 任意类型的值都可以赋值给空接口
```

9.1. `x.(T)`是类型断言的语法，x是一个接口类型的表达式，T是一个类型。类型断言会检查x的动态类型是否满足断言类型T
	1. 如果断言类型是一个具体类型，类型断言就会检查x的动态类型是否满足T，如果满足: x的动态类型就转为T，值不变(确定x的具体类型，因为空接口不能进行任何操作且其他接口类型能够使用的动态类型太过宽泛，如：`io.Writer接口能够保存的值就包括*os.File, *bytes.Buffer等多种具体类型，只有通过断言才能确定具体的类型`)；检查失败就直接崩溃

**如果是x是接口类型，T是具体类型，也就是第一种情况，那么断言之后的结果可以直接使用x接收**

	2. 如果断言类型是一个接口类型，类型断言就会检查x的动态类型是否满足T，如果满足(也就是说x包含T接口的所有方法):那么结果的动态类型和动态值仍然是x，只不过方法数量相对于原先的接口类型变多了。如果失败(也就是说x不包含T接口的全部方法)：直接崩溃**(理解断言类型是接口类型时，动态类型没有改变这个特点)**
```golang
var w io.Writer
w = os.Stdout	// os.Stdout是*os.File类型，并且包含Write和Read方法
res := w.(io.ReadWriter)	// 因为os.Stdout本身包含Write和Read方法，所以实现了io.ReadWriter接口；而结果res的动态类型没改变，仍然是*os.File，但相比于上一行只实现io.Writer方法来说，能够多调用Read方法
```
	3. 如果断言类型是一个接口类型，且T的类型相较于x本来的接口类型方法更多（即：T类型的方法包含x类型的方法，x类型方法是T类型方法的子集），这个时候一般不使用类型断言，因为直接赋值即可
```golang
var rw io.ReadWriter
res := rw.(io.Writer)	// 从一个更大方法集的接口类型向更小方法集的接口类型断言
// 等价于
res = io.Writer(rw)
var res io.Writer = rw
```
**类型断言一般会返回两个值，一个是断言之后的结果，一个是是否断言成功，如果使用一个返回值的断言，断言失败将会崩溃**


10. 通过一个可变的空接口参数，可以传递任意的值，但是我们不能进行任何操作，因为空接口不包含任何方法，需要使用类型断言确定它的具体类型再操作
```golang
func Func1() {
	f2(20, "hello world", true)	// 可以传递任意类型
	var args = []interface{}{20, "hello world", true}
	f2(args...)
}

func f2(args ...interface{}) {
	// args可以接收任意类型的参数
}

```
 
11. 接口值：一个具体类型（接口的动态类型）和该类型的一个值（动态值）
接口的零值就是动态类型和值都是nil，一个接口值是否是nil取决于它的动态类型，每一次对接口值的赋值都可能会改变接口值的动态类型，当将接口值赋为nil时，它的动态类型和值都变成nil
```golang
var w io.Writer	// 类型为nil，值为nil
w = os.Stdout	// 类型变成*os.File，值为os.Stdout
w = new(bytes.Buffer)	// 类型变成*bytes.Buffer，值为新分配的缓冲区的指针
w = nil		// 类型和值都变成nil
```

11.1. 对于接口值的可比较性，如果两个接口值都是nil或两个接口值的动态类型和动态值(前提是动态值具有可比较性，如果动态值是slice类型，将会崩溃)都相等，则这两个接口相等，使用==来比较，所以接口值可以作为map的键或switch的操作数
（但是请注意，接口值是非平凡的，不像基础类型和指针类型是完全可比较的，亦或是map、slice、函数类型完全不可比较，接口值只有动态类型和动态值都可比较时才是可比较的，因此结构体或数组中有接口值元素时比较要谨慎）

**调用一个nil接口（动态类型和动态值都为nil）的任何方法都会出错，对于一个命名类型的方法，有的方法允许接收器为nil，有的则不允许**

**注意：一个接口值要想与nil比较为true，那么它的动态类型和动态值都必须为nil，仅仅只有动态值为nil，比较结果为false（因为接口值存在动态类型不为nil，而动态值为nil的情况）**
```golang
var w io.Writer
var b *bytes.Buffer
fmt.Println("%T, %[1]v", b)	// *bytes.Buffer，nil
fmt.Println("%T, %[1]v", w)	// nil，nil
fmt.Println(w == nil)		// true
fmt.Println(b == nil)		// true，普通类型没有接口值那么麻烦，只要值为nil就为true1
w = b
fmt.Println("%T, %[1]v", w)	// *bytes.Buffer, nil
fmt.Println(w == nil)		// false
```

12. **如果一个命名类型的方法接收器包括普通类型和指针类型，并且无论是普通类型还是指针类型都没有完全实现接口的所有方法，那么这个类型就不能称为实现了这个接口**
```golang
type inter interface{
	f1()
	f2()
}

type newInt int

func (n newInt) f1() {}
func (n *newInt) f2() {}

var i inter
var n newInt
i = n	// 错误，newInt类型并没有实现这个接口，要么使用普通类型完全实现接口，要么使用指针类型完全实现接口
```

13. error是一个内置类型(是一个接口类型)，自主创建error的途径只有两种，errors.New()和fmt.Errorf()，其中第二个可提供格式化的字符串，因为error是一个接口类型，所以任何一种命名类型实现了Error方法，都可以被返回
```golang
// error类型的源代码
type error interface {
	Error() string
}
```

14. 通常类型断言用来存储一些任何类型都可以传递的值，而不是去调用断言之后确定类型的方法，所以使用if-else来断言类型的确切类型是臃肿的，使用switch分支可以实现更高效的操作（这里的switch与分支语句中的switch不同点在于，这里不允许使用fallthrough条件击穿（因为不可能有一个接口满足多种精确类型），deafult在所有的类型都不满足时运行，当第一个case被满足之后，switch将会执行case块语句并退出）
```golang
func f(x interface{}) {
	if x, ok := x.(int); ok {
		xxxx		// 断言成功，此时的x是int类型
	}else if x, ok := x.(string); ok {
		xxxx 		// 断言成功，此时的x是string类型
	}
}

// 更简洁的switch语句（每一个分支块内，变量x的类型都与该分支断言的类型一致，deafult分支中x就是原始类型）

func f(x interface{}) {
	switch x := x.(type) {
		case int:
			xxxx	// 此时的x是int类型
		case string:
			xxxx	// 此时的x是string类型
		deafult:
			xxxx	// 所有类型都不满足时执行	
	}
}
```
