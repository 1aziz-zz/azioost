package main

import (
	"fmt"
)

func main() {

	fmt.Println(exampleMap())

}

// A method with two args and two return types
func getFullNameTextA(firstName, lastName string) (string, string) {
	return "The person:", firstName + " " + lastName
}

// A method with two args and two NAMED return types (better only used in short functions)
func getFullNameTextB(firstName, lastName string) (x string, y string) {
	x = "The person:"
	y = firstName + " " + lastName
	return
}

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Example usages: https://blog.golang.org/defer-panic-and-recover
func aDeferExample() {
	defer fmt.Println("This is before return!")
	fmt.Println("This is the return value!")
}

// Slice doesn't store any data. It defines a range or slice
// length of a slice is number of elements it contains, capacity is the size of the array inside

func sliceExample() {
	names := [5]string{
		"Aziz", "Tom", "Freddy", "Max", "Ben",
	}
	fmt.Println(names)

	slice := names[0:3]

	fmt.Println(slice)
	names[1] = "Mot"
	fmt.Println(slice)

}

// Maps

func exampleMap() map[int]string {
	m := map[int]string {
		1: "Aziz",
		2: "Jan",
		3: "Tom",
	}

	delete(m, 1)
	m[2] = "Peter"
	return m
}
