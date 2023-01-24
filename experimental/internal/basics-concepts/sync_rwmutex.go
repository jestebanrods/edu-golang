package internal

import (
	"fmt"
	"sync"
	"time"
)

func SyncRWMutex() {
	var m = map[int]int{}
	var wg sync.WaitGroup
	var mu sync.RWMutex

	for i := 1; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(10 * time.Millisecond)
			mu.RLock()
			x := m[i]
			mu.RUnlock()
			_ = x * x
			wg.Done()
		}(i)
	}

	for j := 0; j < 1000; j++ {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		m[j] = j * 2
		mu.Unlock()
	}

	wg.Wait()

	fmt.Println(m[100])
}
