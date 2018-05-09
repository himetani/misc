package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4}
	slice2 := make([]int, 2, 2)

	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)
	copy(slice2, slice)

	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)
}
