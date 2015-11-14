package helloworld

import "fmt"

func main() {
	var sayThis string = "Go or Google"
	const sayThisOnly string = "Inside constant"
	fmt.Println(sayThis)
	fmt.Println(sayThisOnly)
}