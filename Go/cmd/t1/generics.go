package main

import "fmt"

func generic() {
	intSlice := []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	float32Slice := []float32{1.5, 2.5, -4.5}
	fmt.Println(sumSlice[float32](float32Slice))
}

// allows different types to use the same function. T for Type
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// the 'any' type can be used when the every operations used is possible on every data type
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// generics can also be used with structs
// type gasEngine struct {
// 	gallons float32
// 	mpg     float32
// }

// type electricEngine struct {
// 	kwh   float32
// 	mpkwh float32
// }

type car[T gasEngine | electricEngine] struct {
	manufacturer string
	model        string
	engine       T
}
