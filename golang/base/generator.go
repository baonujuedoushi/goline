package main

import (
	"bufio"
	"iter"
	"os"
)

func readFile(filePath string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filePath)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}
