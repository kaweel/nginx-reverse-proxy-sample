package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getMessage())
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getMessageHello())
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handlerHello)
	http.ListenAndServe(":8080", nil)
}

//simple method
func getMessage() string {
	return "microservice 1"
}

//simple method
func getMessageHello() string {
	return "Hello!!!!!!!!"
}
