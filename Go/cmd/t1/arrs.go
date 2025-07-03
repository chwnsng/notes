package main

import (
	"fmt"
)

func arr1() {
	var intArr [3]int32 // can't change length and type // contiguus memory
	fmt.Println(intArr)
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// prints the memory location!
	fmt.Println(&intArr[0]) // 1st 4 bytes
	fmt.Println(&intArr[1]) // adjacent 4 bytes
	fmt.Println(&intArr[2]) // adjacent 4 bytes

	var myArr [3]int32 = [3]int32{1, 2, 3}
	myA := [3]int32{1, 2, 3}
	ma := [...]int32{1, 2, 3} // compiler will know the size is 3
	fmt.Println(myArr, myA, ma)

}

func sl() {
	// using slice: able to append to the array
	mySlice := []int8{4, 5, 6} // size not specified
	fmt.Print(mySlice)
	fmt.Printf(" length: %v, capacity: %v\n", len(mySlice), cap(mySlice))
	mySlice = append(mySlice, 7)
	fmt.Print(mySlice)
	fmt.Printf(" length: %v, capacity: %v\n", len(mySlice), cap(mySlice))

	// append multiple values
	mySlice2 := []int8{8, 9, 0}
	mySlice2 = append(mySlice, mySlice2...)
	fmt.Println(mySlice2)

	// specify slice capacity
	mySlice3 := make([]int32, 3, 6) // this runs faster than init without cap specs
	fmt.Print(mySlice3)
	fmt.Printf(" length: %v, capacity: %v\n", len(mySlice3), cap(mySlice3))

}

func map1() {
	myMap := make(map[string]uint8) // string keys, uint8 values
	myMap2 := map[string]uint8{"John": 1, "Taobin": 2}
	fmt.Println(myMap, myMap2)
	fmt.Println(myMap2["John"])
	fmt.Println(myMap2["Nonexistent key"]) // get 0 <-- default value of uint8

	// map will always return something so we may have to manually check if the value actually exists.
	val, exist := myMap2["Key"]
	fmt.Println(val, exist)

	// delete (no return value)
	delete(myMap2, "Taobin")
	fmt.Println(myMap, myMap2)

}

func loop() {
	// order is not preserved when iterating over a map!
	myMap := map[string]uint8{"a": 1, "b": 2, "c": 3, "d": 4}
	myArr := [5]int32{1, 2, 3, 4, 5}
	// for name := range myMap {
	// 	fmt.Printf("%v\n", name)
	// }
	for name, num := range myMap {
		fmt.Printf("%v %v\n", name, num)
	}

	for idx, val := range myArr {
		fmt.Printf("%v: %v\n", idx, val)
	}

	// note: Go doesn't have a while loop. custom a for loop for it.

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

}
