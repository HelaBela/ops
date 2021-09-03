package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
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
	//handleRequest()
	getMetadata()
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

//what is stderr and stdout ?

func getMetadata() {

	cmd, err := exec.Command("git", "log").Output()
	if err != nil {
		log.Fatal(err)
	}

	stringOutput := string(cmd)

	strArray := strings.SplitAfter(stringOutput, "\n")

	//strArray := strings.Fields(stringOutput)
	//lastUpdatedString := strArray[len(strArray)-4]
	commitShaString := strArray[len(strArray)-6]
	commitArray := strings.SplitAfter(commitShaString, " ")
	commit := commitArray[1]

	fmt.Printf(string(commit))

}

//git log -> get from there
//when is it done? when do you get the gitlog and when you passit to the app

// func getTime() {
// 	for _, arg := range os.Args[1:] {
// 		fileinfo, err := os.Stat(arg)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		atime := fileinfo.Sys().(*syscall.Stat_t)
// 		fmt.Println(time.Unix((int)atime, (int)atime))
// 	}
// }
