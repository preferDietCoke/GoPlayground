package main

import "fmt"

func main() {
	// works like a while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Classic for loop:")
	// classic for loop
	for j := 0; j < 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("Always true then breaks:")
	// for loop without arguments same as a while (true)
	for {
		fmt.Println("loop")
		break
	}
}