package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Sending HTML response to the client
	fmt.Fprintln(w, `<h1 style="text-align:center; color:blue;">"Embrace every challenge as a chance to grow, for it is through overcoming obstacles that you unlock your true potential."</h1>`)
	// Printing to the server console
	fmt.Println("Someone hit me!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on 80")
	http.ListenAndServe(":80", nil)
}
