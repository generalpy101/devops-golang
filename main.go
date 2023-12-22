package main

import (
	"fmt"
	"log"
	"net/http"
)

var logger *log.Logger = log.Default()

func main() {

	fmt.Println("Go server starting up....")

	http.HandleFunc("/health", handleHealthCheck)
	http.HandleFunc("/hello", handleHelloWorld)

	logger.Println("Go server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Println("Health OK")
	w.Write([]byte("OK"))

	if _, err := w.Write([]byte("Hello World!")); err != nil {
		logger.Panic(err)
	}
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.Println("Invalid request method for /hello")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	logger.Println("Handling /hello request...")

	if _, err := w.Write([]byte("Hello World!")); err != nil {
		logger.Panic(err)
	}
}
