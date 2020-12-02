package main

import (
	"context"
	"fmt"
	"sync"
)

func proc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("proc")
			return
		default:
		}
	}
}

func channel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go proc(ctx)
	go proc(ctx)
	go proc(ctx)

	cancel()
}

func wait() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		fmt.Println(1)
		wg.Done()
	}()
	go func() {
		fmt.Println(2)
		wg.Done()
	}()
	go func() {
		fmt.Println(3)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	channel()
	wait()

}
