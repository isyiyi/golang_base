package packageUse

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func OsUse() {
	// 在程序执行时，获取执行时输入的命令行参数，第一个为执行该程序的命令
	// echo命令：将echo命令后的所有数据原封不动重复一遍
	var s string
	for k, v := range os.Args {
		if k == 0 {
			continue
		}
		s += v + " "
	}
	fmt.Println(s)
	s = ""

	// 直接使用切片，遍历[1:]
	for _, v := range os.Args[1:] {
		s += v + " "
	}
	fmt.Println(s)
	s = ""

	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)

	fmt.Println(os.Args[0])

	for k, v := range os.Args[1:] {
		fmt.Println(k, " ", v)
	}
}

func Dup1() {
	var counts = make(map[string] int)
	var input = bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] ++
	}
	
	for line, count := range counts {
		if count > 1 {
			fmt.Println(line, " ", count)
		}
	}
}

func DupFromFile() {
	var files = [3]string{"packageUse/data1.txt", "packageUse/data2.txt", "packageUse/data3.txt"}
	counts := make(map[string]int)
	for _, v := range files {
		f, err := os.Open(v)
		defer f.Close()
		if err != nil {
			fmt.Println("open file error")
			fmt.Println(err)
			return
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()] ++
		}
	}
	for k, v := range counts {
		fmt.Println(k, " ", v)
	}
}
