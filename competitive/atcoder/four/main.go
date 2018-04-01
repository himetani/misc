package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	nn := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nn[i])
	}

	cnt := 0
	stop := false

	for {
		for i := 0; i < n; i++ {
			if nn[i]%2 != 0 {
				stop = true
				break
			}
			nn[i] = nn[i] / 2
		}
		if stop {
			break
		}
		cnt++
	}

	fmt.Println(cnt)
}
