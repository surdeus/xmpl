package main

import (
	"net/http"
	"os"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
