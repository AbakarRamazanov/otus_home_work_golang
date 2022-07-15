package main

import "testing"

func TestRunCmd(t *testing.T) {
	var envTest = Environment{
		"BAR":   EnvValue{"bar", false},
		"EMPTY": EnvValue{"", false},
		"FOO":   EnvValue{"   foo\nwith new line", false},
		"HELLO": EnvValue{`"hello"`, false},
		"UNSET": EnvValue{"", true},
	}
	var cmd = []string{"echo", "-n", `${BAR}`, "${HELLO}"}
	RunCmd(cmd, envTest)
}
