package arrays

import "fmt"

func main() {
	// basic declaration of an array
	var array [5]int
	fmt.Println("emp:", array)
	// adding values to array
	array[4] = 100
	fmt.Println("set:", array)
	fmt.Println("get:", array[4])
	// length of an array
	fmt.Println("length:", len(array))
	// another array declaration with values
	anotherArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("anotherArray: ", anotherArray)
	// two dimensional array
	var twoDimensionalArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensionalArray[i][j] = i + j
		}
	}
	fmt.Println("2D array: ", twoDimensionalArray)
}