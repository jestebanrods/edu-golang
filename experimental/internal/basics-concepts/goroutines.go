package internal

import (
	"fmt"
	"sync"
	"time"
)

func function(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn", i, time.Now().Format("5"))
	time.Sleep(3 * time.Second)
	fmt.Println("fn", i, time.Now().Format("5"))
}

func Goroutines() {
	var wg sync.WaitGroup

	init := time.Now()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go function(i, &wg)
	}

	wg.Wait()

	fmt.Println("time execution ", time.Since(init))
}
