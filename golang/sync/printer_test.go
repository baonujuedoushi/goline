package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func aomicVersion(loop, printTimes int) {
	signalChan := make([]chan struct{}, loop)
	for i := 0; i < loop; i++ {
		signalChan[i] = make(chan struct{})
	}
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

func bufferVersion(loop, printTimes int) {
	signalChan := make([]chan struct{}, loop)
	for i := 0; i < loop; i++ {
		signalChan[i] = make(chan struct{}, 1)
	}
	//totolCount := printTimes * loop

	var wg sync.WaitGroup
	wg.Add(loop)
	for i := 0; i < loop; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 0; j < printTimes; j++ {
				<-signalChan[i]
				printNum := i + 1
				signalChan[printNum%loop] <- struct{}{}
			}
		}(i)
	}
	signalChan[0] <- struct{}{}
	wg.Wait()
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aomicVersion(1000, 1000)
	}
}
func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bufferVersion(10000, 1000)
	}
}
