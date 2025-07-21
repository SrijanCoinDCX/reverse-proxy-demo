package main

import (
	"fmt"
	"net/http"
)

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go payment service!")
}

func main() {
	http.HandleFunc("/payments/", paymentsHandler)
	fmt.Println("Go server running at http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}