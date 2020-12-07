// file coordination
// <aq@okaq.com>
// 2020-12-06
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/util"
	"os"
)

const (
	CONFIG = "config.js"
)

// type def for config input
type Games struct {
	console string `json:console`
	dir string `json:dir`
}

func (g Games) String() string {
	b0, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}
	return string(b0)
}

func load() {
	ioutil.ReadAll(CONFIG)
}

func parse() {

}

func main() {
	// parse json
	load()
	parse()
}


