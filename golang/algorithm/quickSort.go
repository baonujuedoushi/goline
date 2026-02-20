package main

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	return quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high)

		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
