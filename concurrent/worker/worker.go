package main

import (
	"fmt"
	"strings"
)

type WorkLauncher interface {
	LaunchWorker(in chan Request)
}

type PrefixSuffixWorker struct {
	id     int
	prefix string
	suffix string
}

func (w *PrefixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefixs(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = strings.ToUpper(s)

			out <- msg
		}

		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			uppercaseString, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = fmt.Sprintf("%s%s", uppercaseString, w.suffix)

			out <- msg
		}

		close(out)
	}()

	return out
}

func (w *PrefixSuffixWorker) prefixs(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			uppercaseStringWithSuffix, ok := msg.Data.(string)

			if !ok {
				msg.Handler(nil)
				continue
			}

			msg.Data = fmt.Sprintf("%s%s", w.prefix, uppercaseStringWithSuffix)

			out <- msg
		}

		close(out)
	}()

	return out
}
