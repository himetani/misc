package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			wg.Done()
		})

		future.Execute(setContext("Helloooooooo!"))
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e)
			t.Fail()
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "Error Occured", errors.New("Error")
		})
		wg.Wait()

	})
}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	t.Log("Timeout!")

	t.Fail()
	wg.Done()
}
