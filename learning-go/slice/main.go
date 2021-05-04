package main

import (
	"fmt"
)

func main() {
	/*
	 Unlike arrays, slices are typed only by the elements
	 they contain (not the number of elements).
	 To create an empty slice with non-zero length, use the builtin make.
	 Here we make a slice of strings of length 3
	*/
	s := make([]string, 3)
	fmt.Println("slice :", s)

	// We can use like array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	/*
	 In addition to these basic operations,
	 slices support several more that make them richer than arrays.
	 One is the builtin append, which returns a slice containing one or more new values.
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("after apd:", s)

	//Slices can also be copy’d.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// This gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl2:", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl3:", l)

}
