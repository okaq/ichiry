// okaq.com home site web server
// google cloud app engine
// golang version 1.14
// aq@okaq.com
// 2020-12-20
package main

import (
	"log"
	"net/http"
	"os"
)

const (
	INDEX = "index.html"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	// log req
	http.ServeFile(w,r,INDEX)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", WebHandler)
	log.Printf("listening on localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}


