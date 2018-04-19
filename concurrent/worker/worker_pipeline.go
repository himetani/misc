package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	bufferSize := 3
	var dispatcher Dispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkLauncher = &PrefixSuffixWorker{
			prefix: fmt.Sprintf("WorkerID: %d -> ", i),
			suffix: " World",
			id:     i,
		}
		dispatcher.LaunchWorker(w)
	}

	request := 10
	var wg sync.WaitGroup
	wg.Add(request)

	for i := 0; i < request; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data: "Hello", Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
	return myRequest
}
