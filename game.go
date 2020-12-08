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
	"os"
)

const (
	CONFIG = "config.js"
	NUM = 4
)

var (
	Games []Game
	Archive map[string][]os.FileInfo
)

// type def for config input
type Game struct {
	Console string `json:console`
	Dir string `json:dir`
}

func (g Game) String() string {
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
	Games = make([]Game, NUM)
	json.Unmarshal(b0, &Games)
	fmt.Println(Games)
	for _, g0 := range Games {
		fmt.Println(g0.String())
	}
}

func parse() {
	Archive = make(map[string][]os.FileInfo)
	for _, g0 := range Games {
		k0 := g0.Console
		d0 := g0.Dir
		f0, err := ioutil.ReadDir(d0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(k0,len(f0))
	}
}

func main() {
	// parse json
	load()
	parse()
}


