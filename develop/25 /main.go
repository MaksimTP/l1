package main

import (
	"log"
	"time"
)

// This function is an implementation of a sleep function
func sleep(duration time.Duration) {
	// Create a timer that will emit signal after some time (duration)
	tmr := time.NewTimer(duration)
	// blocks function until signal is emmited
	<-tmr.C
}

func main() {
	log.Println("hello")
	sleep(2 * time.Second)
	log.Println("bye")
}
