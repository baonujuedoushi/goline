package main

func binSearch(searchList []int, searchNum int) int {
	return doSearch(searchList, 0, len(searchList)-1, searchNum)
}

func doSearch(searchList []int, startIdx, endIdx, searchNum int) int {
	//s := []int{1, 4, 7, 8, 11, 15, 17, 20} len = 8
	//           0  1  2  3   4   5   6   7
	//i := 4
	//
	if startIdx > endIdx {
		return -1
	}
	midlePoint := (startIdx + endIdx) / 2
	midleNum := searchList[midlePoint]
	if midleNum == searchNum {
		return midlePoint
	}

	if midleNum > searchNum {
		endIdx = midlePoint - 1
	} else {
		startIdx = midlePoint + 1
	}

	return doSearch(searchList, startIdx, endIdx, searchNum)
}
