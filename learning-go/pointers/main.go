package main

import "fmt"

func changeValue(a int) {
	a = 0

}

func pchangeValue(ar *int) {
	*ar = 0

}
func main() {
	i := 1
	fmt.Println("initial:", i)
	changeValue(i)
	fmt.Println("after change normal func :", i)
	pchangeValue(&i)
	fmt.Println("after usage of pointer :", i)
	fmt.Println("pointer or adresses :", &i)

}
