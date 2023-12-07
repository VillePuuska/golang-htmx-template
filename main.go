package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const port string = "8080"

func handleMissing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request path:")
	fmt.Println(r.URL.Path)
	fmt.Println()
	tmpl := template.Must(template.ParseFiles("pages/missing.html"))
	tmpl.Execute(w, nil)
	//file, err := os.ReadFile("pages/missing.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Fprintln(w, string(file))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request path:")
	fmt.Println(r.URL.Path)
	fmt.Println()
	file, err := os.ReadFile("pages/home.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(file))
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request path:")
	fmt.Println(r.URL.Path)
	fmt.Println()
	fmt.Fprint(w, strings.Split(time.Now().String(), "m")[0])
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Favicon requested. Sending.")
	fmt.Println()
	http.ServeFile(w, r, "imgs/icon.png")
}

func main() {
	fmt.Println("Listening on port", port)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/home", handleHome)
	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/", handleMissing)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
