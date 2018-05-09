package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4}
	slice2 := slice

	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)

	slice[4] = 5
	slice2[2] = 0

	fmt.Println("slice: ", slice, "cap: ", cap(slice))
	fmt.Println("slice2: ", slice2, "cap: ", cap(slice2))

	fmt.Println("address: ", &slice[0])
	fmt.Println("address: ", &slice2[0])

	slice = append(slice, 0)

	fmt.Println("slice: ", slice, "cap: ", cap(slice))
	fmt.Println("slice2: ", slice2, "cap: ", cap(slice2))

	fmt.Println("address: ", &slice[0])
	fmt.Println("address: ", &slice2[0])

	slice[0] = 55
	slice2[2] = 55

	fmt.Println("slice: ", slice, "cap: ", cap(slice))
	fmt.Println("slice2: ", slice2, "cap: ", cap(slice2))

}
