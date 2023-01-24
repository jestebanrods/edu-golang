package internal

import (
	"fmt"
	"time"
)

func double(n int) int {
	return n * 2
}

func calculate(ch chan<- int) {
	for i := 0; i < 6; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- double(i)
	}

	// if the channel is the type receive
	// only the function that receive the channel can close it
	close(ch)
}

func Channels() {
	ch := make(chan int)

	go calculate(ch)

	// blocks the executions until exists a data
	// read manual channel n, ok := <-ch
	for n := range ch {
		fmt.Printf("%d ", n)
	}

	fmt.Println("finish execution")
}
