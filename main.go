package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/health", healthCheck)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page endpoint"))
	//fmt.Fprintf(w, "home page endpoint")
	//io.WriteString(w, "home page")
}

func main() {
	handleRequest()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	fmt.Fprintf(w, "its ok")

}

//getting the pipeline to build this and package
//github actions

type Metadata struct {
	commit      string
	lastUpdated time.Time
}

//tracablity
