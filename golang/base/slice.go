package main

func sliceT() {
	s := make([]int, 0, 2)
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
}
