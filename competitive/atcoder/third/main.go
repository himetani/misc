package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ss string

	fmt.Scan(&ss)

	cnt := 0
	for _, s := range ss {
		n, _ := strconv.Atoi(string(s))
		if n == 1 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
