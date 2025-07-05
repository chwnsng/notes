package main

import "fmt"

// closest equivalent to  classes but doesn't have inheritance
type gasEngine struct {
	mpg     uint8 // default 0
	gallons uint8 // default 0
	owner
	// int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

// methods
func (e gasEngine) milesleft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesleft() uint8 {
	return e.kwh * e.mpkwh
}

// interface
type engine interface {
	milesleft() uint8
}

// passing interface to function instead of struct --> more flexibility
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesleft() {
		fmt.Println("Can make it!")
	} else {
		fmt.Println("Not enough fuel!")
	}
}

func strc() {
	var myEngine gasEngine = gasEngine{25, 10, owner{"John"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	canMakeIt(myEngine, 50)

	// creating an anonymous struct
	myEngine2 := struct {
		mpg     uint8
		gallons uint8
	}{30, 25}

	fmt.Println(myEngine2)

}

// using structs with JSON encoders/decoders
// add struct tags to define each json field name
type LoginRequest struct {
	Username string `json:"username"` // will look for field name "username" when decoding from json
	Password string `json:"password"`
}
