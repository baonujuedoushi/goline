package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func printAlphaMoreTimes() {
	aSignal := make(chan struct{}, 1)
	bSignal := make(chan struct{})
	cSignal := make(chan struct{})
	var wg sync.WaitGroup
	var loopTime int = 5
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < loopTime; i++ {
			<-aSignal
			fmt.Println("A")
			bSignal <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < loopTime; i++ {
			<-bSignal
			fmt.Println("B")
			cSignal <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < loopTime; i++ {
			<-cSignal
			fmt.Println("C")
			aSignal <- struct{}{}
		}
	}()
	aSignal <- struct{}{}
	wg.Wait()
}

func dynamicPrinterA(loop int) {
	//如果有 N 个 协程（N 是动态传入的），要求它们按编号顺序打印。比如 N=5，就要打印 1 2 3 4 5 1 2 3 4 5
	signalChan := make([]chan struct{}, loop)
	for i := 0; i < loop; i++ {
		signalChan[i] = make(chan struct{}, 1)
	}
	printTimes := 5
	//totolCount := printTimes * loop

	var wg sync.WaitGroup
	wg.Add(loop)
	for i := 0; i < loop; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < printTimes; j++ {
				<-signalChan[i]
				printNum := i + 1
				fmt.Println(printNum)
				signalChan[printNum%loop] <- struct{}{}
			}
		}(i)
	}
	signalChan[0] <- struct{}{}
	wg.Wait()
}

func dynamicPrinterB(loop int) {
	//如果有 N 个 协程（N 是动态传入的），要求它们按编号顺序打印。比如 N=5，就要打印 1 2 3 4 5 1 2 3 4 5
	signalChan := make([]chan struct{}, loop)
	for i := 0; i < loop; i++ {
		signalChan[i] = make(chan struct{})
	}
	printTimes := 5
	totolCount := printTimes * loop
	var counter int32

	var wg sync.WaitGroup
	wg.Add(loop)
	for i := 0; i < loop; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < printTimes; j++ {
				<-signalChan[i]
				printNum := i + 1
				fmt.Println(printNum)
				if atomic.AddInt32(&counter, 1) == int32(totolCount) {
					return
				}
				signalChan[printNum%loop] <- struct{}{}
			}
		}(i)
	}
	signalChan[0] <- struct{}{}
	wg.Wait()
}
