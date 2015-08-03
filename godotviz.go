package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fileFormat := os.Args[1]
	fmt.Println(fileFormat)
	os.Exit(0)

	result := DotRender("graph {a -- b a -- b  b -- a }", "png")
	fmt.Print(result.String())
}

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
