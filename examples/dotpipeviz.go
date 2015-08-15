package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/sc0rp1us/godotviz"
)

// main - implements simple processing DOT language from stdin to stdout
// Example:
// echo 'graph {a -- b a -- b  b -- c }' | go run dotpipeviz.go png > ~/graph.png
func main() {
	fileFormat := os.Args[1]

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	result, err := godotviz.DotRender(data, fileFormat)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(result)
}
