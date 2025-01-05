package main

import (
	"fmt"
	"net/http"
)

func motdHandler(w http.ResponseWriter, r *http.Request) {
	motd := "Welcome to the Go MOTD Server! Have a great day!"
	fmt.Fprintf(w, motd)
}

func main() {
	http.HandleFunc("/motd", motdHandler)
	port := ":10002"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

