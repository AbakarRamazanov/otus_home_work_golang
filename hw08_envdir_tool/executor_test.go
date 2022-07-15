package main

import (
	"os"
	"testing"
)

func TestRunCmd(t *testing.T) {
	os.Clearenv()
	var envTest = Environment{
		"BAR":   EnvValue{"bar", false},
		"EMPTY": EnvValue{"", false},
		"FOO":   EnvValue{"   foo\nwith new line", false},
		"HELLO": EnvValue{`"hello"`, false},
		"UNSET": EnvValue{"", true},
	}
	var cmd = []string{"printenv"}
	RunCmd(cmd, envTest)
}
