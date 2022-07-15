package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	changeEnv(env)
	command := createCommand(cmd)
	if err := command.Run(); err != nil {
		fmt.Println(err)
		return -1
	}
	return command.ProcessState.ExitCode()
}

func createCommand(cmd []string) *exec.Cmd {
	cmd0 := cmd[0]
	command := exec.Command(cmd0, cmd[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	return command
}

func changeEnv(env Environment) {
	for variable, value := range env {
		if value.NeedRemove {
			os.Unsetenv(variable)
		} else {
			os.Setenv(variable, value.Value)
		}
	}
}
