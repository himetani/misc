package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGSTOP)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		cancel()
	}()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	defer fmt.Println("End of doSomething func")
	for {
		select {
		case <-time.Tick(1 * time.Second):
			fmt.Println("hello")
		case <-ctx.Done():
			return
		}
	}
}
