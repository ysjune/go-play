package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)

	ctx2 := context.WithValue(context.Background(), "number", 9)
	go square(ctx2)
	//cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}

//func square(ctx context.Context) {
//	if v := ctx.Value("number"); v != nil {
//		n := v.(int)
//		fmt.Printf("square:%d", n*n)
//	}
//	wg.Done()
//}
