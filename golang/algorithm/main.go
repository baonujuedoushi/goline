package main

import "fmt"

func main() {
	sorted := mergeSort([]int{5, 1, 2, 8, 20, 15, 17, 60, 33, 6})
	fmt.Println(sorted)
}
func missingNumber(nums []int) int {
	//	s := []int{3, 0, 1}
	//fmt.Println(missingNumber(s))
	mMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		mMap[num] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := mMap[i]; !ok {
			return i
		}
	}
	return 0
}

func missingNumberB(nums []int) int {
	var listSum, shouldSum int
	nLen := len(nums)
	shouldSum = nLen * (nLen + 1) / 2
	for _, num := range nums {
		listSum += num
	}
	return shouldSum - listSum
}
