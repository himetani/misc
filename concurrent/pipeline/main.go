package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
					fmt.Println("start!")
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		inStream <-chan int,
		multiplier int,
	) <-chan int {
		outStream := make(chan int)
		go func() {
			defer close(outStream)
			for value := range inStream {
				select {
				case <-done:
					return
				case outStream <- multiplier * value:
				}
			}
		}()
		return outStream
	}

	add := func(
		done <-chan interface{},
		inStream <-chan int,
		additive int,
	) <-chan int {
		outStream := make(chan int)
		go func() {
			defer close(outStream)
			for value := range inStream {
				select {
				case <-done:
					return
				case outStream <- additive + value:
				}
			}
		}()
		return outStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
