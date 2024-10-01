package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// All goroutines stop their work if a main goroutine finishes its work
/*
func main() {
	go func() {
		fmt.Println("Hello world")
	}()
}
*/

// -----------------------------------------

// Exit with a channel emmiting a signal
/*
 func main() {
 	quit := make(chan struct{})

 	go func() {
 		for {
 			select {
 			case <-quit:
 				return
 			default:
 				fmt.Println("Hello")
 			}
 		}
 	}()

 	quit <- struct{}{}
}
*/

// -----------------------------------------

// Exiting a goroutine with a contexts with a cancel function
func main() {
	forever := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				forever <- 1
				return
			default:
				fmt.Println("hello")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(1500 * time.Millisecond)
		cancel()
	}()

	<-forever
	fmt.Println("finish")
}
