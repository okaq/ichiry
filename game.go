// file coordination
// <aq@okaq.com>
// 2020-12-06
package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	// "io"
	"io/ioutil"
	// "os"
)

const (
	CONFIG = "config.js"
)

// type def for config input
type Games struct {
	Console string `json:console`
	Dir string `json:dir`
}

func (g Games) String() string {
	b0, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}
	return string(b0)
}

func load() {
	b0, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b0))
	g0 := make([]Games, 4)
	json.Unmarshal(b0, &g0)
	fmt.Println(g0)
	for _, Game := range g0 {
		fmt.Println(Game.String())
	}
}

func parse() {

}

func main() {
	// parse json
	load()
	parse()
}


