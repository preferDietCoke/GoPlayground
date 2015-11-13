package main

import "fmt"

func main() {
	simpleSlice := make([]string, 3)
	fmt.Println("empty slice: ", simpleSlice)
	// adds values to slice
	simpleSlice[0] = "a"
	simpleSlice[1] = "b"
	simpleSlice[2] = "c"
	fmt.Println("set:", simpleSlice)
	fmt.Println("get:", simpleSlice[2])
	fmt.Println("length:", len(simpleSlice))
	// adds new values to simpleSlice even if the length is
	// out of bounds
	simpleSlice = append(simpleSlice, "d")
	// can be added multiple values
	simpleSlice = append(simpleSlice, "e", "f")
	// make another slice with the same length as simpleSlice
	anotherSlice := make([]string, len(simpleSlice))
	// copy the values of slimpleSlice
	copy(anotherSlice, simpleSlice)
	fmt.Println("anotherSlice: ", anotherSlice)
	// creating another slice with values from 2 to 4
	// the last value in position 5 is not included
	slice3 := simpleSlice[2:5]
	fmt.Println("slice3:", slice3)
	// declaring a two dimensional slice
	twoDimensionalSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDimensionalSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimensionalSlice[i][j] = i + j
		}
	}
	fmt.Println("2D slice:", twoDimensionalSlice)
}