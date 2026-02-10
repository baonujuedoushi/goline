package main

import (
	"fmt"
	"sync"
)

func printAlpha() {
	aFinished := make(chan struct{})
	bFinished := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printC(bFinished)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		printA(aFinished)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		printB(aFinished, bFinished)
	}()
	wg.Wait()
}

func printA(signal chan struct{}) {
	defer close(signal)
	fmt.Println("A")
}

func printB(signalA, signalB chan struct{}) {
	defer close(signalB)
	<-signalA
	fmt.Println("B")
}

func printC(signal chan struct{}) {
	<-signal
	fmt.Println("C")
}
