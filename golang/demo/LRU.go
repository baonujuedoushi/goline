package main

import "fmt"

func lru() {
	//维护一个切片获取某个key或者更新某个key时，将该key移动到
	aSlice := []int{1, 2, 3, 4}
	aSlice = append(aSlice[1:], 5)
	fmt.Printf("%v", aSlice[0])
	aSlice = append(aSlice[0:1], aSlice[2:]...)
	aSlice = append(aSlice, aSlice[1])
}
