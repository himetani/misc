package main

import "fmt"

type typedError struct {
	error
}

func (e *typedError) Error() string {
	return e.error.Error()
}

func main() {
	var e interface{} = nil

	if _, ok := e.(*typedError); ok {
		fmt.Println("hoge")
	}
}
