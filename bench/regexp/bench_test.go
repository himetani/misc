package main

import (
	"regexp"
	"testing"
)

var testRegexp string = `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`

func BenchmarkMatchString(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := regexp.MatchString(testRegexp, "jsmith@example.com")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMatchStringCompiled(b *testing.B) {
	r, err := regexp.Compile(testRegexp)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.MatchString("jsmith@example.com")
	}
}
