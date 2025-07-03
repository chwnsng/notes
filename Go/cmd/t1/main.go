package main

import (
	"fmt"
	"unicode/utf8"
) // in Go, imported packages Must be used

func main() {
	// ch1()
	// ch2()
	// grt()
	// bestDeal()
	generic()

}

func ch1() { // in the main package, the main() function must be defined
	// -------------------------------- NUMBERS --------------------------------
	var intNum int = 300000 // "int" is 32-bit on a 32 bit system and 64-bit on a 64 bit system
	var unsigned uint8      // (0, 255)
	// var intNum int8 -- for 8-bit integer
	// var intNum int16 -- max 32767
	// var intNum int32
	// var intNum int64

	// var myfloat float32 = 12345678.9 // becomes 12345679.0 instead
	var myfloat float64 = 12345678.9 // 1234567878.9 --more accurate

	var myInt8 int8 = 1
	var myfloat32 float32 = 1.5
	var result float32 = myfloat32 + float32(myInt8)

	fmt.Println(intNum)
	fmt.Println(unsigned)
	fmt.Println(myfloat)
	fmt.Println(result)

	var a int = 3
	var b int = 2
	fmt.Println(a / b) // round down to 1
	fmt.Println(a % b) // 1

	// -------------------------------- TEXT --------------------------------
	var myString = "Hello" + "\n world"
	var s = `Hello
	World`
	var ss string // default value = ""
	fmt.Println(ss)

	fmt.Println(myString + s)
	fmt.Println(len("Test")) // gets NUMBER of BYTES, not chars
	// chars outside the vanilla ASCII are more than 1 byte.
	// If your text will contain those chars, do this instead:
	fmt.Println(utf8.RuneCountInString("α"))

	// runes are not that useful YET
	var myRune rune = 'a' // defaultvalue = 0
	fmt.Println(myRune)

	// -------------------------------- MORE --------------------------------

	var myfalse bool = false // default value is false
	fmt.Print(myfalse)
	fmt.Println(true)

	// inferring types
	var myvar = "text"
	fmt.Println(myvar)

	coolvar := "text"
	fmt.Println(coolvar)

	var1, var2 := 1, 2
	fmt.Println(var1 + var2)

	// -------------------------------- CONST --------------------------------

	const myConst string = "hey"
	fmt.Println(myConst)

}
