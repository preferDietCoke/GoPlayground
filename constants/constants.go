package constants

import "fmt"
import "math"
// importing reflect to see which type value is
import "reflect"

// declaring global constant
const randomConstant string = "constant"

func main() {
	fmt.Println(randomConstant)

	// note: cant declare as a int when its actually a long int
	const number = 50000
	const anotherNumber = 3e20/number
	// printing out values
	fmt.Println(anotherNumber)
	// printing as a int64
	fmt.Println(int64(anotherNumber))
	// testing math
	fmt.Println(math.Sin(number))
	// all number types
	fmt.Println("-- Types --")
	fmt.Println("number: ", reflect.TypeOf(number))
	fmt.Println("anotherNumber: ", reflect.TypeOf(anotherNumber))
}