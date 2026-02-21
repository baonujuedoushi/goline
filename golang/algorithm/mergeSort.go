package main

func mergeSort(s []int) []int {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}
	mid := sLen / 2
	lSlice := mergeSort(s[0:mid])
	rSlice := mergeSort(s[mid:])
	return merge(lSlice, rSlice)
}

func merge(ls, rs []int) []int {
	var lPoint, rPoint int
	lsLen := len(ls)
	rsLen := len(rs)
	newSlice := make([]int, 0, lsLen+rsLen)
	for lPoint < lsLen && rPoint < rsLen {
		if ls[lPoint] < rs[rPoint] {
			newSlice = append(newSlice, ls[lPoint])
			lPoint++
		} else {
			newSlice = append(newSlice, rs[rPoint])
			rPoint++
		}
	}
	newSlice = append(newSlice, ls[lPoint:]...)
	newSlice = append(newSlice, rs[rPoint:]...)
	return newSlice
}
