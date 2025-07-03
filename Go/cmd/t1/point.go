package main

import "fmt"

func pointers() {
	// var p *int32
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("p points to %v", *p) // where the pointer is pointing to
	fmt.Printf("\ni is equal to %v\n", i)
	*p = 10 // set the value at location p to 10
	p = &i  // set p to point to the memory location of i
	fmt.Printf("p points to %v", *p)
	fmt.Printf("\ni is equal to %v\n", i)
	*p = 1 // i will also change to 1 since it's pointing to the same int32
	fmt.Printf("p points to %v", *p)
	fmt.Printf("\ni is equal to %v\n", i)

	// different for regular vars
	var k int32 = 2
	i = k // i is actually a copy of k, pointing to a different location

	// but slices point to the same location
	slice := []int32{1, 2, 3}
	sliceCopy := slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}

// using pointers with functions
func square(nums [5]float64) [5]float64 {
	fmt.Printf("Location of nums array: %p\n", &nums) // not the same array as the one inside the function!
	// this arrays is copied from the one passed in the argument. So it will take more  memory!
	// use pointers to optimize this

	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	return nums
}

// when parameters are large, passing in a pointer to their value instead might be a good idea!
func square2(nums *[5]float64) [5]float64 { // takes the value of at pointer that points to nums
	fmt.Printf("Location of nums array: %p\n", nums) // same array
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	return *nums
}

func callSquare() {
	n1 := [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Location of n1 array: %p\n", &n1)

	result := square2(&n1) // pass in a pointer that points to n1
	fmt.Printf("result: %v\n", result)
}
