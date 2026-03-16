package main

import (
	"fmt"
	"os"
)

func osRootSample() {
	root, err := os.OpenRoot("./safe_zone")
	if err != nil {
		fmt.Println("can't open path:", err)
		return
	}
	defer root.Close()

	f, err := root.Open("../main.go")
	if err != nil {
		fmt.Printf("invalid filepath: %v\n", err)
		return
	}
	f.Close()
}
