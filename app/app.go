package main

import (
	"fmt"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customer", getAllCustomer)

	http.ListenAndServe("localhost:8000", nil)
	fmt.Println("Server is running at port :8000")
}
