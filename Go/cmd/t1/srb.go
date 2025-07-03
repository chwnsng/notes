package main

import (
	"fmt"
	"strings"
)

func str() {
	// s := "résumé"
	// index := s[0]
	// fmt.Printf("%v, %T\n", index, index)
	// for i, v := range s { // <-- leaads to unexpected outcomes due to non ASCII character
	// 	fmt.Println(i, v)
	// }

	// cast to runes first <-- UTF8
	ss := []rune("résumé")
	for i, v := range ss { // <-- leaads to unexpected outcomes due to non ASCII character
		fmt.Println(i, v)
	}

	// concatenating strings
	strSlice := []string{"s", "a", "y", "h", "i"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // use string builder to concatenate strings is faster than "+" due to strings being immutable <-- a new one is created everytime
	}
	fmt.Printf("\n%v\n", strBuilder.String())

}
