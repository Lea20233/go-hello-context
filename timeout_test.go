package go_hello_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CounterTimeOut(ctx context.Context) chan int {

	dest := make(chan int)

	go func() {
		defer close(dest)
		count := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				dest <- count
				count++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return dest
}

func TestTimeOut(t *testing.T) {

	fmt.Println("Total Goroutines = ", runtime.NumGoroutine())
	parent := context.Background()
	//time out to cancel
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	dest := CounterTimeOut(ctx)
	for n := range dest {
		fmt.Println("Counter", n)
	}

	fmt.Println("Total Goroutines = ", runtime.NumGoroutine())
}
