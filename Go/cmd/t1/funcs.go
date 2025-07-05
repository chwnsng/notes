package main

import (
	"errors"
	"fmt"
)

// private funcs start with lowercase (scope=package)
func ch2() {
	fmt.Println("Hey there")
}

// public funcs start with uppercase (scope=any package)
// other packages can import this function
func Printme(text string) {
	fmt.Println(text)
}

func divide(num int, denom int) int {
	return num / denom
}

// returning multiple values
func multDiv(num int, denom int) (int, int, error) {
	var err error // default = nil
	if denom == 0 {
		err = errors.New("Can't divide by zero you fool!\n")
		return 0, 0, err
	}
	return num / denom, num * denom, err
}

func checkMD() {
	// getting multiple returned values
	var ans1, ans2, err = multDiv(10, 0)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Div = %v, Mult = %v\n", ans1, ans2)
	}
}

func vibeCheck(a int, b int) {
	if a == b && b == a {
		fmt.Println("Passed the vibe check")
	}
	if a == b || b == a {
		fmt.Println("Checked the vibe pass")
	}

}

func switchCheck(a int) {
	// in Go, the 'break' for each case is implied
	switch {
	case a == 1:
		fmt.Println("y")
	case a == 2:
		fmt.Println("y")
	default:
		fmt.Println("n")
	}

	// can also write conditionals like this
	switch a {
	case 0, 1, 2:
		fmt.Println("yy")
	default:
		fmt.Println("nn")
	}
}
