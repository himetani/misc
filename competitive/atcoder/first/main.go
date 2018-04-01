package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var s string

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&s)

	fmt.Println(a+b+c, s)
}
