1. 并发：程序由若干个自主的活动单元组成. goroutine和channel都支持通信顺序进程（csp），csp是一个并发的模式，在不同的执行体之间传递值。
2. channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
3. 通道关闭之后将不能再使用发送操作，只能进行接收，直到通道内的数据完全被接收。
4. 无缓冲通道：如果一个协程往chan里写数据，在数据没有被另一个协程接收之前，这个协程将被阻塞。如果一个协程从chan里接收数据，在数据没有从另一个协程发送出来之前，这个协程将被阻塞。（使用无缓冲通道，将导致两个协程的同步）
5. 当使用完channel之后，使用close函数可以关闭channel，但是如果再次接收，将会得到对应类型的零值，当然，golang提供了另一个返回值ok，可以判断当前操作的channel是否被关闭
```golang
func f1(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch chan int) {
	for {
		if v, ok := <-ch; ok { // 当ok为false时，意味着channel被关闭了
			fmt.Println(v)
		}else {
			break
		}
	}
}
```

6. 除了第5种方式判断channel被关闭之外，使用range也能够完成这个功能。当使用range来遍历一个channel时，如果这个channel没有被关闭，那么它会一直阻塞，有数据就处理，**当channel被关闭之后range循环就停止（一定要关闭channel，否则会死锁）**
7. 除了普通的双向通道`chan xxx`外，还有两种单向通道，分别为`chan <- int`单向发送通道和`<- chan int`单向接收通道。**不要试图去关闭一个单向接收通道**
8. 单向通道只是在参数传递时人为规定数据的流向，在创建时只能是双向通道，形参和实参传递时会自动隐式转换（但是，如果已经是单向通道，再传递为双向通道是错误的，单向通道的方向不一样也是错误的）
```golang
func main() {
	var ch = make(chan int)
	f(ch)
}

func f(ch chan <- int) {	// ch形参是单向发送通道，但实参是双向通道，会进行隐式的类型转换
	ch <- 89
	f2(ch)			// error，此时的ch是单向通道，不能再转换为形参的双向通道
}

func f2(ch chan int) {
	fmt.Println(<-ch)
}



```
9. 缓冲通道，通过指定make的第二个参数指定通道的缓冲区大小。如果缓冲区没有满，可以一直发送数据，如果缓冲区已经满了，发送数据将会被阻塞，直到有协程接收了数据，腾出了位置。如果缓冲区没有空，可以一直接收数据，如果缓冲区已经空了，接收数据将会被阻塞，直到有协程发送了数据。
```golang
var ch = make(chan int, 3)
ch <- 90
ch <- 80
ch <- 70
ch <- 60	// 程序将会被阻塞，直到有协程消费了里面的数据
```

10. `cap`可以知道通道缓冲区的容量，`len`可以直到通道缓冲区已经有多少个元素
11. for range xxx的用法（用于不关心内部元素，只关心这个变量的长度） **`for range`可以用来清空一个channel**
```golang
var nums = []int { 1, 2, 3, 4}
for _, v := range nums {
	xxx
}

for range nums {	// 不使用nums的数据，只使用nums的长度，意为循环len(nums)次
	xxx
}

var ch chan int = make(chan int, 3)
ch <- 10
ch <- 20
ch <- 30

// for range的由来
for x := range ch {
	// 对内部元素并不关心，可以将x改为_
}

for _ = range ch {
	// 与其这样写，不如直接省略掉_
}

// 如果对channel里面的元素并不关心，使用for range可以快速清空
for range ch{
}
```

12. select语句也包含一系列的情况和对应的分支，select一直等待，直到一次通信来告知有分支可以执行，执行完该分支后其他分支将不会执行。
13. select的用法
```golang
var ch = make(chan int)
var ch2 = make(chan int, 2)

go func() {
	time.Sleep(2 * time.Second)
	ch <- 90
}()

// 第一个select
select {
	case x := <-ch :
		fmt.Println(x)	// 程序等待两s，输出接收到的x=90
}

// 第二个select
select {
	case x := <-ch :
		fmt.Println(x)
	case ch2 <- 80 :
		xxx		// 执行这一句，执行完之后就结束了
}				// 对于所有的case分支，select是能执行哪个就执行哪个，执行完之后直接结束

// 第三个select
select {
	case x := <-ch :
		fmt.Println(x)
	default :
		fmt.Println("Hello")	// 执行default，因为执行到select时ch没有数据就阻塞了（数据在2s之后才发送），所以直接执行default

}
```

14. 如果有多个case都满足，那么select就随机选择一个，确保每个channel都有机会
15. default分支确保在没有条件满足时select被阻塞，如果没有条件满足则直接执行default分支
16. channel被关闭之后，不能对其发送值，但是可以接收，如果没有元素将得到对应类型的零值（一个通道被关闭且里面所有的值都被接收之后，再进行接收操作将会立即得到响应并返回对应类型的零值）
```golang
var ch = make(chan int)
close(ch)

select {
	case <- ch :  // 虽然ch被关闭了，但是仍然可以取值，并且要快于default（无论怎么取都是0，无需等待和阻塞）
		xxx
	default :
		xxx
}
```

17. 关闭协程，因为golang不允许一个协程直接关闭另一个协程，所以可以通过channel来发送广播信号的方式，让它自己关闭自己。（因为一个channel如果关闭之后select可以迅速get到这个case，所以可以使用select来轮询这个广播信号）
```golang
var done chan int = make(chan int)

func isCancelled() bool {
	select {
		case <- done :
			return true
		default :
			return false
	}
}

// 当需要关闭协程时
close(done)
// 此时所有的协程调用isCancelled函数的结果都是true，因为case<-done这个分支被激活了
```

18. main goroutine要结束时，所有的goroutine都立即停止执行，主函数返回，程序随之退出。此时并不能确定释放了所有的资源。此时可以调用panic，运行时将转储程序中所有goroutine的栈。
