package go_utils

import (
	"bufio"
	"bytes"
	"context"
	"github.com/pkg/errors"
	"io"
	"os/exec"
	"strings"
	"time"
)

// 执行Shell脚本，返回行解析对象数组
func ExecuteBashLiner(shellScripts string, liner func(line string) bool) error {
	cmd := exec.Command("bash", "-c", shellScripts)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	var eout bytes.Buffer
	cmd.Stderr = &eout

	cmd.Start()
	defer cmd.Process.Kill()
	defer cmd.Wait()

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line != "" {
			if !liner(line) {
				return nil
			}
		}
	}

	eoutput := eout.String()
	if eoutput != "" {
		return errors.New(eoutput)
	}

	return nil
}

func ExecuteBash(shellScripts string) (string, error) {
	stdout := ""

	err := ExecuteBashLiner(shellScripts, func(line string) bool {
		stdout += line
		return true
	})

	return stdout, err
}

func BashTimeout(cmds string, timeout time.Duration) (string, string) {
	return ExecuteCommandTimeout(timeout, "bash", "-c", cmds)
}

// https://medium.com/@vCabbage/go-timeout-commands-with-os-exec-commandcontext-ba0c861ed738
func ExecuteCommandTimeout(timeout time.Duration, name string, args ...string) (string, string) {
	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	cmd := exec.CommandContext(ctx, name, args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	var eout bytes.Buffer
	cmd.Stderr = &eout

	err := cmd.Run()
	if err != nil {
		return "", err.Error()
	}

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		return out.String(), "timed out"
	}

	output := out.String()
	eoutput := eout.String()

	return output, eoutput
}
