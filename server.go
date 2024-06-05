package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"

	http.ListenAndServe(":8080", nil)
	log.Printf("Listening on %s ...\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}

}

// handles the hello world request
func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// if method is GET
	writeResponse(writer, "Hello, World!")
}

// handles the health request
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// if method is GET
	writeResponse(writer, "OK!")
}

func writeResponse(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString)
	_, err := writer.Write(response)
	if err != nil {
		log.Printf("Error writing response: %s\n", err)
	}

}
