package variables

import "fmt"

func main() {
	var initString string = "initial"
	fmt.Println(initString)

	var numberOne, numberTwo int = 1, 2
	fmt.Println(numberOne, numberTwo)

	var booleanVal bool = true
	fmt.Println(booleanVal)

	var unsignedInt int
	fmt.Println(unsignedInt)

	// should I specify type?
	shortDeclaration := "short"
	fmt.Println(shortDeclaration)
}

/*
	Unsigned values gets a standard value.
	String value gets ""
	Int value get 0
	and so on.
*/