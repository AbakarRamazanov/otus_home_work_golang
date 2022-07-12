package main

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var ErrFilesAreDiffent = errors.New("files are different")

func TestCopyCantWrite(t *testing.T) {
	err := Copy(`unavailable.otus_home_work_golang`, "greateble.otus_home_work_golang", 0, 0)
	require.Error(t, err)
	err = Copy(`testdata/input.txt`, ".otus_home_work_golang/greateble", 0, 0)
	require.Error(t, err)
}

func TestCopyOffsetExceedsFileSize(t *testing.T) {
	err := os.Mkdir("testdata/testdir", 0o755)
	defer os.RemoveAll("testdata/testdir")
	require.NoError(t, err)
	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset7000_limit0.txt", 7000, 0)
	require.EqualError(t, err, ErrOffsetExceedsFileSize.Error())
}

func TestCopyDir(t *testing.T) {
	err := Copy(`testdata`, "testdataCopy", 0, 0)
	require.EqualError(t, err, ErrUnsupportedFile.Error())
}

func TestCopy(t *testing.T) {
	err := os.Mkdir("testdata/testdir", 0o755)
	defer os.RemoveAll("testdata/testdir")
	require.NoError(t, err)
	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset0_limit0.txt", 0, 0)
	require.NoError(t, err)
	err = compare("testdata/out_offset0_limit0.txt", "testdata/testdir/out_offset0_limit0.txt")
	require.NoError(t, err)

	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset0_limit10.txt", 0, 10)
	require.NoError(t, err)
	compare("testdata/out_offset0_limit10.txt", "testdata/testdir/out_offset0_limit10.txt")
	require.NoError(t, err)

	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset0_limit1000.txt", 0, 1000)
	require.NoError(t, err)
	err = compare("testdata/out_offset0_limit1000.txt", "testdata/testdir/out_offset0_limit1000.txt")
	require.NoError(t, err)

	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset0_limit10000.txt", 0, 10000)
	require.NoError(t, err)
	err = compare("testdata/out_offset0_limit10000.txt", "testdata/testdir/out_offset0_limit10000.txt")
	require.NoError(t, err)

	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset100_limit1000.txt", 100, 1000)
	require.NoError(t, err)
	err = compare("testdata/out_offset100_limit1000.txt", "testdata/testdir/out_offset100_limit1000.txt")
	require.NoError(t, err)

	err = Copy(`testdata/input.txt`, "testdata/testdir/out_offset6000_limit1000.txt", 6000, 1000)
	require.NoError(t, err)
	err = compare("testdata/out_offset6000_limit1000.txt", "testdata/testdir/out_offset6000_limit1000.txt")
	require.NoError(t, err)
}

func compare(filepath1 string, filepath2 string) error {
	f1, err := os.ReadFile(filepath1)
	if err != nil {
		return err
	}
	f2, err := os.ReadFile(filepath2)
	if err != nil {
		return err
	}
	if bytes.Equal(f1, f2) {
		return nil
	}
	return ErrFilesAreDiffent
}
