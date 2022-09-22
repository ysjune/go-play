package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d 부터 %d까지의 합계는 %d 입니다.\n", a, b, sum)
	wg.Done()
}

type Account struct {
	Balance int
}

var mutex sync.Mutex

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {

	//for i := 0; i < 10; i++ {
	//	go SumAtoB(1, 1000000)
	//}
	//
	//wg.Wait()
	//fmt.Println("모든 계산이 완료되었습니다.")
	//-----------------------------------------
	//account := &Account{0}
	//wg.Add(10)
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for {
	//			DepositAndWithdraw(account)
	//		}
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//------------------------------------------
	//rand.Seed(time.Now().UnixNano())
	//
	//wg.Add(2)
	//fork := &sync.Mutex{}
	//spoon := &sync.Mutex{}
	//
	//go diningProblem("a", fork, spoon, "포크", "수저")
	//go diningProblem("b", spoon, fork, "수저", "포크")
	//wg.Wait()
	//-------------------------------------------
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value : %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다. \n", name)
		first.Lock()
		fmt.Printf("%s %s Get\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s Get\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다. \n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
