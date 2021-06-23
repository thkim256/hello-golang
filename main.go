package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, webserver")
	})

	log.Println("시작 http listening :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
