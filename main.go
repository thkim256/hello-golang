package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, thkim256")
	})

	log.Println("start http listening :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
