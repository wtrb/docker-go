package main

import (
	"fmt"
	"log"
	"net/http"

	"rsc.io/quote"
)

func main() {
	http.HandleFunc("/", homeHanlder)
	http.HandleFunc("/pages", pageViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is a quote for you\n<h1>%s</h1>", quote.Go())
}

func pageViewHandler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "Stop refreshing this page! %v", counter)
}

var counter uint = 0
