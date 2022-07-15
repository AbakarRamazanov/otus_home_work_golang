package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	bytesNL = []byte{'\n'}
	bytes00 = []byte{0x00}
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	environment := make(Environment)
	for _, f := range files {
		if isInvalidFile(f) {
			continue
		}
		if f.Size() == 0 {
			environment[f.Name()] = EnvValue{Value: "", NeedRemove: true}
			continue
		}
		s, err := getValue(filepath.Join(dir, f.Name()))
		if err != nil {
			return nil, err
		}
		environment[f.Name()] = EnvValue{Value: s, NeedRemove: false}
	}
	return environment, nil
}

func isInvalidFile(f fs.FileInfo) bool {
	return f.IsDir() || strings.Contains(f.Name(), "=")
}

func getValue(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	s, err := r.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	s = string(bytes.ReplaceAll([]byte(s), bytes00, bytesNL))
	s = strings.TrimRight(s, " \t\r\n") // TODO delete \r
	return s, nil
}
