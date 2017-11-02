package main

import "fmt"

func main() {

	a := "TestA"
	var b = &a      // b gets the memory address of a

	fmt.Println(b)
	fmt.Println(*b) // prints value of the memory address of a

	*b = "TestB"  // assign new value to the memory address of a

	fmt.Println(a)
}



