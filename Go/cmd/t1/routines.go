package main

import (
	"fmt"
	"sync"
	"time"
)

// "go"
// goroutines: get concurrency
// concurrency can either be parallel execution or constantly switching between tasks
// if we have a multicore cpu, we typically get parallel execution in Go

// var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func grt() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // basically a counter
		go dbCall(i) // "go" continue without waiting for the next step
	}
	wg.Wait() // wait for the counter to become 0 before proceeding. Else, the program might exit before tasks finish
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nResults are: %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond) // simulating database communication delay
	// fmt.Println("db response:", dbData[i])
	save(dbData[i])
	log()
	wg.Done() // decrements the counter
}

func save(result string) {
	m.Lock() // thread lock -- let goroutines wait for each other
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() // multiple routines can hold a readlock at once, but if a readlock is held, no one can hold the (write) lock
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}

// note:
// in this example --> basically contant time for the dbcall function
// more computationally intensive tasks --> how much execution time improves will depend on how many cpu cores we have
