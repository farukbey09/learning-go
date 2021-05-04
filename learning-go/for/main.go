package main

import (
	"fmt"
)

func main() {

	i := 1
	fmt.Println("\nFirst For")
	for {
		if i == 5 {
			break
		}
		i++
		fmt.Println(i)
	}
	j := 1
	fmt.Println("\nSecond For")
	for j <= 3 {

		fmt.Println(j)
		j = j + 1
	}
	fmt.Println("\nThird For")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("\nForth For")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
