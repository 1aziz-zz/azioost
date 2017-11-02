package main

import "fmt"

func main() {
	var mySchool *school
	mySchool = &school{"TOP!"}
	fmt.Println(mySchool.show())

	fmt.Println(mySchool.name)
}

type school struct {
	name string
}

// if it's not a pointer (without *), assigning s.name will not affect the injected object above.
// so if the injected object is supposed to get changed in a method, * must be added before the
// struct type.
func (s *school) show() string {
	s.name = "POT!"
	return s.name
}
