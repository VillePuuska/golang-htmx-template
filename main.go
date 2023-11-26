package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const port string = "8080"

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request path:")
	fmt.Println(r.URL.Path)
	fmt.Println()
	file, err := os.ReadFile("pages/missing.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(file))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Favicon requested. Sending.")
	http.ServeFile(w, r, "imgs/icon.png")
}

func main() {
	fmt.Println("Listening on port", port)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
