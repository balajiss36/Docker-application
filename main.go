package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path)) // This will display the route we are on i.e the Slash (/)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi") // If the request is routed to /hi Hi is printed.
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
