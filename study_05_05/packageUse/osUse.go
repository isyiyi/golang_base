package packageUse

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io/ioutil"
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

func DupFromFile2() {
	var files = [3]string{"packageUse/data1.txt", "packageUse/data2.txt", "packageUse/data3.txt"}
	counts := make(map[string] int)
	for _, v := range files {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println(err)
		}
		strs := strings.Split(string(data), "\n")
		for _, s := range strs {
			counts[s] ++
		}
	}	
	
	for k, v := range counts {
		fmt.Println(k, " ", v)
	}
}

func DupHomeWork() {
	var files = [3]string{"packageUse/data1.txt", "packageUse/data2.txt", "packageUse/data3.txt"}
	
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		
		// 当一个文件出现重复语句时，立即返回不再往下计数
		counts := make(map[string] int)
		strs := strings.Split(string(data), "\n")
		for _, s := range strs {
			counts[s] ++
			if counts[s] == 2 {
				fmt.Println(file)
				break
			}
		}
	}
}
