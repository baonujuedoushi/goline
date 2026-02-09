package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func safeQuit() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	jChan := make(chan int, 3)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		defer wg.Done()
		defer close(jChan)
		jobGeneretor(jChan, ctx, randIntFactory)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		jobReciver(jChan)
	}()

	<-sigChan
	cancel()

	wg.Wait()
}
func randIntFactory() int {
	return rand.IntN(100)
}
func jobGeneretor[T any](jChan chan T, ctx context.Context, factory func() T) {
	//generetor job, add to queue
	//in channel is open case, try add job, when can't add case, queue is full. drop that
	//in channel is close case , stop generate
	jobGap := time.Microsecond * 1000
	tiker := time.NewTicker(jobGap)
	defer tiker.Stop()
	for {
		//nolint:gosimple
		select {
		case <-tiker.C:
			select {
			case <-ctx.Done():
				fmt.Println("stop the generetor")
				return
			default:
				//we are fine
			}
			select {
			case jChan <- factory():
				fmt.Println("new job added")
				//double check
			case <-ctx.Done():
				fmt.Println("stop the generetor")
				return
			default:
				fmt.Println("job is full, drop")
				//drop
			}
		}
	}
}

func jobReciver[T any](jChan chan T) {
	jobCostsTime := time.Microsecond * 2000

	for value := range jChan {

		fmt.Printf("current job content: %v \n", value)
		time.Sleep(jobCostsTime)
	}
}
