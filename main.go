package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, World!")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", echoString)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/static/", static)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
