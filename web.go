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
	GAMES = "cares.js"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web serve on localhost:8080")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,GAMES)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", GamesHandler)
	http.ListenAndServe(":8080", nil)
}


