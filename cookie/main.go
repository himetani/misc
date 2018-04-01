package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Cookies())
		cookie := &http.Cookie{
			Name:   "hoge",
			Value:  "bar",
			Domain: "localhost",
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Hello world")
	})
	http.ListenAndServe(":8080", nil)
}
