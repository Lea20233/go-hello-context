package go_hello_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//example leak in goroutines

func CreateCounterLeak() chan int {

	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1
		//leak happen in here, destination still get data from counter
		//thats why there still goroutine run in background
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

func TestCounterLeak(t *testing.T) {

	//in here goroutines 2
	fmt.Println("Total Goroutines", runtime.NumGoroutine())

	destination := CreateCounterLeak()

	//leak from 'for' in Counter
	for number := range destination {
		fmt.Println("Counter", number)
		if number == 10 {
			break
		}
	}

	//in here goroutines 3
	fmt.Println("Total Goroutines", runtime.NumGoroutine())
}

// this is after fixed the leak using context.WithCancel
func CreateCounterCancel(ctx context.Context) chan int {

	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		//if the number==10, then it go in case <- ctx.Done()
		//// then it will return nil
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func TestCounterCancel(t *testing.T) {

	fmt.Println("Total Goroutines = ", runtime.NumGoroutine())

	//using context withcancel to cancel after return nil
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterCancel(ctx)

	for number := range destination {
		fmt.Println("Counter", number)
		if number == 10 {
			break
		}
	}
	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutines = ", runtime.NumGoroutine())

}
