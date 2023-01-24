package internal

import (
	"fmt"
	"time"
)

func doubleNumber(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 2
}

func tripleNumber(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 3
}

func Select() {
	d := make(chan int)
	t := make(chan int)

	var f chan struct{}

	go func() {
		time.Sleep(3 * time.Second)
		f = make(chan struct{})
		f <- struct{}{}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			d <- doubleNumber(i)
			t <- tripleNumber(i)
		}

		close(d)
		close(t)
	}()

	var breaker = true

	for breaker {
		select {
		case <-f:
			fmt.Println("cancel")
			close(d)
			close(t)
			close(f)
			breaker = false
		case c, ok := <-d:
			if !ok {
				breaker = false
			}

			fmt.Printf("%d ", c)
		case c, ok := <-t:
			if !ok {
				breaker = false
			}

			fmt.Printf("%d ", c)
		}
	}

	fmt.Println("finish")
}
