package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sc0rp1us/godotviz"
)

// main - implements simple processing DOT language from stdin to stdout
// Example:
// echo 'graph {a -- b a -- b  b -- c }' | go run godotviz.go png > ~/graph.png
func main() {
	fileFormat := os.Args[1]
	stdinContent := stdinToStr()

	// result := DotRender("graph {a -- b a -- b  b -- a }", "png")
	result := godotviz.DotRender(stdinContent, fileFormat)

	os.Stdout.Write(result)
}

func stdinToStr() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(data)
}
