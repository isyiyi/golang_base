package gorAndChan

import (
	"fmt"
	"time"
)

func GorUse2() {
	defer func() func() {
		fmt.Println(time.Now().Format(layout))
		return func() {
			fmt.Println(time.Now().Format(layout))
		}
	}()()
	fmt.Println(fib(45))
}

const layout = "Jan 2, 2006 at 3:04:05pm (MST)"

func fib(x int) int {
	if x == 1 || x == 2 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}


func GorUse3() {
	var ch chan int = make(chan int)
	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Println(i, "in")
		}
		close(ch)
	}(ch)	
	go printFib(ch)
	time.Sleep(5 * time.Second)
}

func printFib(ch chan int) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		}else {
			break
		}
	}
}

func GorUse4() {
	var ch chan int = make(chan int)
	go singleCh(ch)
	go fmt.Println(<-ch)
	time.Sleep(time.Second * 2)
}

func singleCh(ch chan <- int) {
	ch <- 90
	f(ch)
}

func f(ch chan <- int) {
	ch <- 10
}

func GorUse6() {
	var nums = []int{1, 2, 3, 4}
	for range nums {
		fmt.Println(0)
	}
}

func GorUse7() {
	var ch = make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 90
	}()
	select {
	case x := <-ch :
	fmt.Println(x)
	default:
		fmt.Println("Hello")
	}
}

func GorUse8() {
	var ch chan int = make(chan int)
	go func() {
		close(ch)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		select {
		case <- ch :
			fmt.Println("Hello")
		default :
			fmt.Println("Default")
		}
	}()
	time.Sleep(2 * time.Second)
}

func GorUse() {
	var ch = make(chan int, 3)
	ch <- 10
	ch <- 20
 	ch <- 30
	close(ch)
	fmt.Println(len(ch))
	for range ch {
	}
	fmt.Println(len(ch))

}
