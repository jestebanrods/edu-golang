package internal

import (
	"fmt"
	"time"
)

func BufferedChannels() {
	// is a buffer when the cap of channel is different to zero
	// if is a buffer no need to wait to send a channel
	ch := make(chan int, 1)
	fmt.Println(len(ch), cap(ch))

	go func() {
		ch <- 5
		fmt.Println("send ok")
		fmt.Println(len(ch), cap(ch))
	}()

	time.Sleep(3 * time.Second)

	fmt.Printf("receive ok %d\n", <-ch)
	fmt.Println(len(ch), cap(ch))
}
