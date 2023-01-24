package internal

import (
	"fmt"
	"time"
)

func Semaphores() {
	var maxCh int8 = 2
	var s = make(chan struct{}, maxCh)
	var ch = make(chan int)

	for i := 0; i < 1000; i++ {
		go func(i int) {
			s <- struct{}{}
			time.Sleep(1 * time.Second)
			ch <- i
			<-s
		}(i)
	}

	for i := 0; i < 1000; i++ {
		fmt.Printf("%d ", <-ch)
	}
}
