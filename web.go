// minimal web server
// test f1-micro capabilities
// logging, ssl
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "index.html"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,INDEX)
}

func main() {
	fmt.Println("web server started on localhost:80")
	fmt.Printf("start time: %s\n.", time.Now().String())
	http.HandleFunc("/", WebHandler)
	http.ListenAndServe(":80", nil)
}
