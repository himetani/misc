package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var x int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	cnt := 0
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if i*500+j*100+k*50 == x {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
