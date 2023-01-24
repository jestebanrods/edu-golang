package internal

import (
	"fmt"
	"time"
)

func sum(a, b int64) int64 {
	time.Sleep(1 * time.Microsecond)
	return a + b
}

func parallel(top int64) {
	out := make(chan int64)
	values := make(chan int64)

	// productor
	go func() {
		defer close(values)

		for i := 0; i < int(top); i++ {
			values <- int64(i)
		}

		if top%2 != 0 {
			values <- 0
		}
	}()

	// consumer
	go func() {
		defer close(out)

		for i := 0; i < int(top); i++ {
			a, ok := <-values
			if !ok {
				break
			}

			b, ok := <-values
			if !ok {
				break
			}

			out <- sum(a, b)
		}
	}()

	var suma int64
	breaker := true

	for breaker {
		select {
		case a, ok := <-out:
			if !ok {
				breaker = false
			}

			suma += a
		}
	}

	fmt.Println(suma)
}

func main() {
	now := time.Now()
	parallel(1000)
	fmt.Println(time.Since(now))
}
