package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envTest := Environment{
		"BAR":   EnvValue{"bar", false},
		"EMPTY": EnvValue{"", false},
		"FOO":   EnvValue{"   foo\nwith new line", false},
		"HELLO": EnvValue{`"hello"`, false},
		"UNSET": EnvValue{"", true},
	}
	env, err := ReadDir("testdata/env")
	require.NoError(t, err)
	require.Equal(t, len(env), len(envTest))
	for key := range env {
		_, ok := envTest[key]
		require.True(t, ok)
		require.Equal(t, env[key], envTest[key])
	}
}

func TestGenerateAndReadDir(t *testing.T) {
	err := os.Mkdir("testdir", 0o777)
	require.NoError(t, err)
	defer os.RemoveAll("testdir")
	files := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < 10; i++ {
		file, err := os.Create("testdir/" + files[i])
		require.NoError(t, err)
		file.WriteString(files[i])
		file.Close()
	}
	env, err := ReadDir("testdir")
	require.NoError(t, err)
	for i := 0; i < 10; i++ {
		value, ok := env[files[i]]
		require.True(t, ok)
		require.False(t, value.NeedRemove)
		require.Equal(t, value.Value, files[i])
	}
}

// TODO func check files with =
