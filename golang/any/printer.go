package main

import "fmt"

func printAny[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
