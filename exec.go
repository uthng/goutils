package utils

import (
	"bytes"
	"io"
	"os/exec"
)

// ExecCmd executes command shell. It redirects output and errors to standard and then stores in a bytes.
// It returns command's stdout  and stderr in bytes and error of command execution
func ExecCmd(command string, args []string, o io.Writer, e io.Writer) ([]byte, []byte, error) {
	var outBuf, errBuf bytes.Buffer
	//var errStdout, errStderr error

	cmd := exec.Command(command, args...)

	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	outWriter := io.MultiWriter(o, &outBuf)
	errWriter := io.MultiWriter(e, &errBuf)

	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	go func() {
		//_, errStdout = io.Copy(stdout, stdoutIn)
		io.Copy(outWriter, stdoutIn)
	}()

	go func() {
		//_, errStderr = io.Copy(stderr, stderrIn)
		io.Copy(errWriter, stderrIn)
	}()

	err = cmd.Wait()

	//output := append(outBuf.Bytes(), errBuf.Bytes()...)
	return outBuf.Bytes(), errBuf.Bytes(), err
}
