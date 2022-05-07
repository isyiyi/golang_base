package funcs

import "fmt"
import "time"

func f1(num1, num2 int, str1 string) (num3 int, str2 string) {
	num3 = num1 + num2 
	str2 = str1 + "|"
	return
}

func f2(num3, num4 int, str3 string) (int, string) {
	num1 := num3 * num4
	str1 := str3 + str3
	return num1, str1
}

func FuncDemo() {
	var f func(int, int, string)(int, string)
	f = f1
	num, str := f(2, 3, "hello")
	fmt.Println(num, str)
	
	f= f2
	num, str = f(2, 3, "hello")
	fmt.Println(num, str)
}

func BDemo() {
	var funcss []func()
	for i := 0; i < 10; i++ {
		funcss = append(funcss, func(){
			fmt.Println(i)
		})
	}

	for _, v := range funcss {
		v()
	}
}

func CalcTime() {
	defer trace()()
	time.Sleep(time.Second * 5)
}

func trace() func() {
	start := time.Now()
	return func() {
		end := time.Now()
		fmt.Println(start)
		fmt.Println(end)
	}
}
