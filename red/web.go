// golang web server for google cloud
// okaq.funka
// Tue Jan 21, 2014
package web

import (
    "io/ioutil"
    "log"
    "net/http"
    "appengine"
)

const (
    PATH = "index.html"
)

var (
    Data []byte
    err error
)

func init() {
    // cache single page html webapp
    Data, err = ioutil.ReadFile(PATH)
    if err != nil {
        log.Fatal(err)
    }
    // http handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write(Data)
        // logging
        c := appengine.NewContext(r)
        c.Infof("okaq.funka: hello %s", r.RemoteAddr)
    })
}


