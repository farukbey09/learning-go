package main

import "fmt"

func main() {
	var arrayDefine [5]int
	fmt.Println(arrayDefine)
	arrayDefine[2] = 500

	fmt.Println("get 2. element :", arrayDefine[2])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
