package main

import (
	"fmt"
)

func sum(ch chan int, num []int) {
	result := 0
	for _, v := range num {
		result += v
	}
	ch <- result
}

func main() {
	result := 0
	ch := make(chan int)
	num := []int{
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
		1, 2, 3, 4, 5,
	}
	go sum(ch, num[0:5])
	go sum(ch, num[5:10])
	go sum(ch, num[10:15])
	go sum(ch, num[15:20])
	go sum(ch, num[20:25])
	go sum(ch, num[25:30])

	for range 6 {
		result += <-ch
	}

	fmt.Println(result)
}
