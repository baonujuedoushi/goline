package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// 同时请求多个广告商，先返回结果的广告商为胜出者
func adSearch() {
	adProviders := []string{"google", "microsoft"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1200)
	defer cancel()
	result := make(chan string, 1)
	var wg sync.WaitGroup
	timer := time.NewTimer(time.Microsecond * time.Duration(rand.IntN(50)))
	defer timer.Stop()
	for _, provider := range adProviders {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			i := 1
			gapTime := 100
			for {
				if i > 3 {
					fmt.Printf("[%s] can't sent\n", p)
					return
				}
				select {
				case result <- p:
					fmt.Printf("[%s] sent\n", p)
					return
				case <-ctx.Done():
					fmt.Printf("[%s] timeout\n", p)
					return

				default:
					fmt.Printf("[%s] buffer is full,prepare %d retry...\n", p, i)
					timer.Reset(time.Duration(gapTime*i) * time.Microsecond)
					select {
					case <-timer.C:
						i++
					case <-ctx.Done():
						return
					}
				}
			}

		}(provider)
	}
	select {
	case first := <-result:
		fmt.Printf("fastest is %s", first)

	case <-ctx.Done():

		fmt.Println("all providers timeout!")
	}
	wg.Wait()
}
