package main

import (
	"fmt"
	"time"
)

// passing information between routines
// enable thread safety
// listen for data updates in channels and prevent it
func deadlock() {
	c := make(chan int)
	// assign value
	c <- 1   // c : [1] // deadlock here
	i := <-c // pop value from c. now c : []
	fmt.Println(i)

	// this will result in a deadlock error. the channel waits for a routine to read from it
}

func chanel() {
	// c := make(chan int)
	c := make(chan int, 5) // buffer channel that can hold 5 values
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func process(c chan int) {
	defer close(c) // close the channel right before the function exits
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}
