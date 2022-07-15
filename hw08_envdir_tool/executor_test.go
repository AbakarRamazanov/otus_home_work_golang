package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	envTest := Environment{
		"BAR":   EnvValue{"bar", false},
		"EMPTY": EnvValue{"", false},
		"FOO":   EnvValue{"   foo\nwith new line", false},
		"HELLO": EnvValue{`"hello"`, false},
		"UNSET": EnvValue{"", true},
	}
	cmd := []string{"pwd"}
	require.Equal(t, 0, RunCmd(cmd, envTest))
	for key, value := range envTest {
		valueEnv := os.Getenv(key)
		require.Equal(t, value.Value, valueEnv)
	}
}
