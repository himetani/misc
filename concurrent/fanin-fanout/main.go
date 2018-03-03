package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	rand := func() interface{} { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Println(numFinders)
	finders := make([]<-chan int, numFinders)

	for i := 0; i < len(finders); i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	fmt.Println("Primes:")
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\t\n", prime)
	}

	fmt.Printf("Search took: %v\n", time.Since(start))
}

func primeFinder(done chan interface{}, valueStream <-chan int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for {
			select {
			case <-done:
			case v := <-valueStream:
				if isPrime(v) {
					intStream <- v
				}
			}
		}
	}()

	return intStream
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan int, num int) <-chan int {
	takeStream := make(chan int)
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)

	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
			case stringStream <- v.(string):
			}
		}
	}()

	return stringStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
			case intStream <- v.(int):
			}
		}
	}()

	return intStream
}

func fanIn(
	done <-chan interface{},
	channels ...<-chan int,
) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()

		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}
