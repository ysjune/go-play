package main

import "fmt"

//func main() {
//	var wg sync.WaitGroup
//	ch := make(chan int)
//	//quit := make(chan bool)
//
//	wg.Add(1)
//	//go square(&wg, ch, quit)
//	go square(&wg, ch)
//
//	for i := 0; i < 10; i++ {
//		ch <- i * 2
//	}
//
//	//quit <- true
//	wg.Wait()
//}

//func square(wg *sync.WaitGroup, ch chan int) {
//	for n := range ch {
//		fmt.Printf("square: %d\n", n*n)
//		time.Sleep(time.Second)
//	}
//	wg.Done()
//}

//func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
//	for {
//		select {
//		case n := <-ch:
//			fmt.Printf("square :%d\n", n*n)
//			time.Sleep(time.Second)
//		case <-quit:
//			wg.Done()
//			return
//		}
//	}
//}

//func square(wg *sync.WaitGroup, ch chan int) {
//	tick := time.Tick(time.Second)
//	terminate := time.After(10 * time.Second)
//
//	for {
//		select {
//		case <-tick:
//			fmt.Println("tick")
//		case <-terminate:
//			fmt.Println("terminated!")
//			wg.Done()
//			return
//		case n := <-ch:
//			fmt.Printf("square: %d\n", n*n)
//			time.Sleep(time.Second)
//		}
//	}
//
//	for n := range ch {
//		fmt.Printf("square: %d\n", n*n)
//		time.Sleep(time.Second)
//	}
//	wg.Done()
//}

func main() {
	ch := make(chan int, 2) // 크기가 2인 채널 생성
	ch <- 1
	ch <- 2
	ch <- 3 //deadlock 발생

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
