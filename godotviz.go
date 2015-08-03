package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// main - implements simple processing DOT language from stdin to stdout
// Example:
// echo 'graph {a -- b a -- b  b -- c }' | go run godotviz.go png > ~/graph.png
func main() {
	fileFormat := os.Args[1]
	stdinContent := stdinToStr()

	// result := DotRender("graph {a -- b a -- b  b -- a }", "png")
	result := DotRender(stdinContent, fileFormat)

	os.Stdout.Write(result.Bytes())
}

func stdinToStr() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(data)
}

// DotRender - Method wrapper around DOT tool
func DotRender(inputString string, outputType string) bytes.Buffer {
	bufferIn := bytes.NewBufferString(inputString)
	var bufferOut bytes.Buffer

	cmd := exec.Command("dot", "-T"+outputType)

	cmd.Stdin = bufferIn
	cmd.Stdout = bufio.NewWriter(&bufferOut)
	cmd.Stderr = os.Stderr

	cmd.Run()

	return bufferOut
}
