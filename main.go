package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//超时控制
func handle() {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan struct{}, 1)
	go func() {
		// Rpc(ctx, ...) // TODO
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("nice")
		// nice
	case <-ctx.Done():
		// timeout
	}
}

func proc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("procDone")
			return
		default:
			fmt.Println("procDefault")
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

	wait()
	channel()
	// handle()
}
