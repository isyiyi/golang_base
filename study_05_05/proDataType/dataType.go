package proDataType

import (
	"fmt"
)

func ArrayDemo() {
	var nums = [3] int {1, 2, 3}
	fmt.Println(len(nums))
	for i, v := range nums {
		fmt.Println(i, v)
	}
	
	var nums2 = [...] int {5: 10}
	fmt.Println(len(nums2))
	for i, v := range nums2 {
		fmt.Println(i, v)
	}
	
	var nums3 = [][]int{[]int{1, 2}, []int{2, 3}}
	fmt.Println(nums3)
}

func SliceDemo() {
	var nums = [] int {1, 2, 3}
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	nums = append(nums, 3)	
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	
	var nums2 = [] int {14, 5}
	nums = append(nums, nums2...)
	fmt.Println(nums)

	var nums3 []int
	fmt.Println(nums3 == nil)
	nums3 = append(nums3, 90)
	fmt.Println(nums3)
	
	var nums4 = make([]int, 4, 6)
	nums4[0] = 90
	fmt.Println(nums4)

	var m map[string] int = make(map[string]int)
	m["aa"] = 90
	fmt.Println(m)
}

type student struct {
	name string
	age int
}

type Point struct {
	X int
	Y int
}

type Cricle struct {
	Point
	Radius int
}

func StructDemo() {
	var stu student = student{"wang", 23}
	fmt.Println(stu)
	
	var tmp = struct{
	name string
	age int
	}{"wang", 23}
	fmt.Println(tmp)

	var c = Cricle{
	Point:Point{
	X: 90,
	Y:90,
	},
	Radius:100,
	}
	fmt.Printf("%#v\n", c)
}


type p struct {
	X, Y int
}

type C struct {
	p
	Radius int
}

func StructDemo2() {
	var c = C{
		p:p{23, 23},
		 Radius:10,
}
	fmt.Println(c.p.X)

}
