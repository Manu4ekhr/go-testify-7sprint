package main

import "net/http"

func main() {
	http.HandleFunc(`/cafe`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
