package internal

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var d = map[int]int{}

func SyncMutex() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := rand.Intn(10)
			mu.Lock()
			d[n]++
			mu.Unlock()
		}()
	}

	wg.Wait()

	var total int

	for k, v := range d {
		fmt.Printf("%d = %d\n", k, v)
		total += v
	}

	fmt.Printf("total = %d", total)
}

func SyncWithChannels() {
	var out = make(chan map[int]int, 1000)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(ch chan<- map[int]int) {
			defer wg.Done()
			n := rand.Intn(10)
			d := make(map[int]int)
			d[n]++
			ch <- d
		}(out)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for c := range out {
		for k, v := range c {
			d[k] += v
		}
	}

	var total int

	for k, v := range d {
		fmt.Printf("%d = %d\n", k, v)
		total += v
	}

	fmt.Printf("total = %d", total)
}
