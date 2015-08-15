package godotviz

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

// DotRender - Method wrapper around DOT tool
func DotRender(inputBytes []byte, outputType string) (graph []byte, err error) {
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	cmd := exec.Command("dot", "-T"+outputType)

	cmd.Stdin = bufio.NewReader(bytes.NewBuffer(inputBytes))
	cmd.Stdout = bufio.NewWriter(&bufferOut)
	cmd.Stderr = bufio.NewWriter(&bufferErr)

	cmd.Start()
	cmdErr := cmd.Wait()
	if cmdErr != nil {
		err = fmt.Errorf("%v\n%v", cmdErr, bufferErr.String())
	}

	return bufferOut.Bytes(), err
}
