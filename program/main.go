package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/meta", metadata)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page endpoint"))
	//fmt.Fprintf(w, "home page endpoint")
	//io.WriteString(w, "home page")
}

func main() {
	handleRequest()
	//getMetadata()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	fmt.Fprintf(w, "its ok")

}

//getting the pipeline to build this and package
//github actions

type Metadata struct {
	Commit      string
	LastUpdated time.Time
}

//date deployed. Time I deployed a docker image or run the docker image

//pass in the time as variable from the bash script. Time when I do docker build . <- tricky :D
// os.getEnv  pass info to docker container and the app read the info from the container ->hint! pass env. var to container when you build
//when I run the app I can read those variables

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

	//there is a better way of doing it. if there was no git installed? will break.
	//in prod I will have a problem if there is no git.
	//go build - __ and replace a variable there. so that the docker container doesnt have to have git installed. as its baed in
	//

	//cmd, err := exec.Command("git", "log").Output()
	//cmd, err := exec.Command("git", "log", "-1", "--format=%H").Output()
	//have to change that so it works when I dont copy all, but just the go binary.
	//how to execute outside of the app ?
	//should be stored staticly somewhere

	// ideal state of this app in prod. ???
	//aws v1 and v2 . u dont know hwat version is that. when you execute a command on a subshell. Rely on a third party to retrieve data. you should reply on server to retrieve data. Server should handle it internally.

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//str := string(cmd)

	metadata.Commit = os.Getenv("SHA")
	//do the same with the time
	//pass in the version
	return metadata

}

//bash script to build binary and publish in docker image and then to run the image.
