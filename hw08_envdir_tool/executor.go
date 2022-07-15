package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here.
	for variable, value := range env {
		if value.NeedRemove {
			os.Unsetenv(variable)
		} else {
			os.Setenv(variable, value.Value)
		}
	}
	var out bytes.Buffer
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		return -1
	}
	fmt.Println(out.String())
	return command.ProcessState.ExitCode()
}
