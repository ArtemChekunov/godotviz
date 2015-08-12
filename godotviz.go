package godotviz

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

// DotRender - Method wrapper around DOT tool
func DotRender(inputString string, outputType string) (graph []byte, err error) {
	bufferIn := bytes.NewBufferString(inputString)
	var bufferOut bytes.Buffer
	var bufferErr bytes.Buffer

	cmd := exec.Command("dot", "-T"+outputType)

	cmd.Stdin = bufferIn
	cmd.Stdout = bufio.NewWriter(&bufferOut)
	cmd.Stderr = bufio.NewWriter(&bufferErr)

	cmd.Start()
	cmdErr := cmd.Wait()
	if cmdErr != nil {
		err = fmt.Errorf("%v\n%v", cmdErr, bufferErr.String())
	}

	return bufferOut.Bytes(), err
}
