package main_test

import "testing"
import sp "study_05_16/sortPackage"

func TestRandNum(t *testing.T) {
	nums := sp.RandNum()
	rightNums := sp.RightSort(nums)
	bubbleNums := sp.BubbleSort(nums)
	for i := 0; i < len(rightNums); i++ {
		if rightNums[i] != bubbleNums[i] {
			t.Errorf("%d != %d\n", rightNums[i], bubbleNums[i])
		}
	}	
}

func BenchmarkRightSort(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		nums := sp.RandNum()
		sp.RightSort(nums)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		nums := sp.RandNum()
		sp.BubbleSort(nums)
	}
}
