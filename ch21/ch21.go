package main

import (
	"fmt"
	"os"
)

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("%T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

type opFunc func(int, int) int

func getOperator2(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

type Writer func(string)

func writeHello(writer Writer) {
	writer("Reconnect World!")
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("fail to create a file")
		return
	}

	defer fmt.Println("must call")
	defer f.Close()
	defer fmt.Println("file close")

	fmt.Println("file write Hello World")
	fmt.Fprintln(f, "Hello World")

	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println("operator result is : ", result)

	result2 := func(a, b int) int {
		return a + b
	}(30, 40)
	fmt.Println("literal function is : ", result2)

	fmt.Println("------------------------")
	CaptureLoop()
	fmt.Println("------------------------")
	CaptureLoop2()
	fmt.Println("------------------------")

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
