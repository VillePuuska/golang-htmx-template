package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = "8080"

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request path:")
	fmt.Println(r.URL.Path)
	fmt.Println()
}

func main() {
	fmt.Println("Listening on port", port)
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
