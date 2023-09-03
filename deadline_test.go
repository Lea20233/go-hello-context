package go_hello_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CounterDeadline(ctx context.Context) chan int {

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

func TestDeadline(t *testing.T) {

	fmt.Println("Total Goroutines = ", runtime.NumGoroutine())
	parent := context.Background()
	//deadline to cancel
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	dest := CounterDeadline(ctx)
	for n := range dest {
		fmt.Println("Counter", n)
	}

	fmt.Println("Total Goroutines = ", runtime.NumGoroutine())
}
