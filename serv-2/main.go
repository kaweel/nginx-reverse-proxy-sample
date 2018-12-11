package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getMessage())

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//simple method
func getMessage() string {
	return "microservice 2"
}
