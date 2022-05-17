package sortPackage

import "math/rand"
import "time"
import "sort"

func BubbleSort(nums [10000]int) [10000]int {
	for i := 0; i < len(nums) - 1; i ++ {
		for j := 0; j < len(nums)-1-i; j ++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func RightSort(nums [10000]int) [10000]int{
	sort.Ints(nums[:])
	return nums
}


func  RandNum()[10000]int{
	newRand := rand.New(rand.NewSource(time.Now().Unix()))
	var nums [10000]int
	for i := 0; i < len(nums); i ++ {
		nums[i] = newRand.Intn(100000)
	}
	return nums
}
