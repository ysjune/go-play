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

}
