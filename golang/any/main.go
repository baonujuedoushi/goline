package main

type anyNode[T any] struct {
	val  T
	next *anyNode[T]
}

func main() {

}

func printIt() {
	sSlice := []string{"a1", "b2"}
	intSlice := []int{1, 2}
	printAny(sSlice)
	printAny(intSlice)
}
