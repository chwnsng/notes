package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var MAX_PIZZA_PRIZE float32 = 5
var MAX_SALMON_PRIZE float32 = 3

func bestDeal() {
	pizzaChannel := make(chan string)
	salmonChannel := make(chan string)
	websites := []string{"grab.com", "lineman.com", "foodpanda.com"}
	for i := range websites {
		go checkPizzaPrices(websites[i], pizzaChannel)
		go checkSalmonPrices(websites[i], salmonChannel)
	}
	sendMessage(pizzaChannel, salmonChannel)
}

func checkPizzaPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var pizzaPrice = rand.Float32() * 20
		if pizzaPrice <= MAX_PIZZA_PRIZE {
			c <- website
			break
		}
	}
}

func checkSalmonPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var salmonPrice = rand.Float32() * 20
		if salmonPrice <= MAX_SALMON_PRIZE {
			c <- website
			break
		}
	}
}

// this function will be called many times from multiple routines
// the value of both channels will be empty (just a memory location) until a good deal has been found
// when that happens, the value will be popped into to a channel in its checkPrice function
// then, this function will be called and the 1 obtained channel data will be passed into the select statement
func sendMessage(pizzaChannel chan string, salmonChannel chan string) {
	// fmt.Printf("Best deal for pizza is at %s\n", <- channel)
	select {
	case website := <-pizzaChannel: //
		fmt.Printf("\nBest deal for pizza is at %v", website)
	case website := <-salmonChannel:
		fmt.Printf("\nBest deal for salmon is at %v", website)
	}
}
