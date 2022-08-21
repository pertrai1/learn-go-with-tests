// package comment
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Buffer for something or another
type Buffer interface{}

// Greet with a print
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreetHandler does something
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
