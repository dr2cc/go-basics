package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	go genrateNumbers(1000, numbers)
	for n := range numbers {
		fmt.Println(n)
	}
}

func genrateNumbers(n int, res chan int) {
	for i := 0; i <= n; i++ {
		res <- i * 2
	}
}
