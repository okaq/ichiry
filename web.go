// browser emulator
// <aq@okaq.com>
// 2020-12-10
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "web.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web serve on localhost:8080")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.ListenAndServe(":8080", nil)
}


