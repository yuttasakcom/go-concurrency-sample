package main

import (
	"fmt"
	"time"
)

func main() {
	defer timer()()

	// syncSlowly()
	// syncSlowly()
	// syncSlowly()
	// syncSlowly()
	// syncSlowly()

	c := make(chan bool, 5)

	go asyncSlowly(c)
	go asyncSlowly(c)
	go asyncSlowly(c)
	go asyncSlowly(c)
	go asyncSlowly(c)

	<-c
	<-c
	<-c
	<-c
	<-c
}

func syncSlowly() {
	time.Sleep(1000 * time.Microsecond)
}

func asyncSlowly(c chan bool) {
	// defer timer()()
	time.Sleep(1000 * time.Microsecond)
	c <- true
}

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("timer: %s\n", time.Since(start))
	}
}
