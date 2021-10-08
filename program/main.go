package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/meta", metadata)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page endpoint"))
}

func main() {
	handleRequest()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	fmt.Fprintf(w, "its ok")
}

type Metadata struct {
	Commit      string
	LastUpdated string
}

func metadata(w http.ResponseWriter, r *http.Request) {

	meta := getMetadata()
	b, err := json.Marshal(meta)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(b)
}

func getMetadata() Metadata {

	metadata := Metadata{}

	metadata.Commit = os.Getenv("SHA")
	metadata.LastUpdated = os.Getenv("TIME")

	return metadata
}
