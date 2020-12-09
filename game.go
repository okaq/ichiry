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
	Flat map[string]string
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
		Archive[k0] = f0
	}
}

func output() {
	// new json game data
	// format for html
	// write to disk
	for k0, v0 := range Archive {
		fmt.Println(k0, len(v0))
		for i0, f0 := range v0 {
			b0, err := json.Marshal(f0)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(i0)
			fmt.Println(f0.Name())
			fmt.Printf("%d\n", f0.Size())
			fmt.Printf("%v\n", f0.ModTime())
			fmt.Println(len(b0))
		}
	}
	// flatten to map[string]string
	// resolve full filepath = dir + name
	// include size on disk
	// encode to json array of strings
}

func output2() {
	// flat key value
	// get file paths
	f0 := paths()
	fmt.Println(f0)
}

func paths() []string {
	p0 := make([]string, 1)
	for _, g0 := range Games {
		d0 := g0.Dir
		k0 := g0.Console
		p1 := Archive[k0]
		fmt.Println(d0,k0,len(p1))
		for _, p2 := range p1 {
			if p2.IsDir() == true {
				continue
			}
			f0 := p2.Name()
			f1 := fmt.Sprintf("%s/%s", d0, f0)
			f2 := fmt.Sprintf("%v|%s", p2.Size(), p2.ModTime().String())
			f3 := fmt.Sprintf("%s|%s", f1, f2)
			p0 = append(p0,f3)
		}
	}
	fmt.Println(p0)
	return p0
}

func main() {
	// parse json
	load()
	parse()
	output()
	output2()
}


