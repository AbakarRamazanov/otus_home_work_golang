package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0o755)
	if err != nil {
		return err
	}
	defer fromFile.Close()
	toFile, err := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		return err
	}
	defer toFile.Close()
	fromFileStat, _ := fromFile.Stat()
	toFileStat, _ := toFile.Stat()
	err = checkErrors(fromFileStat, toFileStat, offset)
	if err != nil {
		return err
	}
	if limit+offset > fromFileStat.Size() || limit == 0 {
		limit = fromFileStat.Size() - offset
	}
	fromFile.Seek(offset, 0)
	byteset := make([]byte, 32)
	allWriteBytes := int64(0)
	var countReadBytes int
	for !errors.Is(err, io.EOF) && limit > allWriteBytes {
		countReadBytes, err = fromFile.Read(byteset)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
		if allWriteBytes+int64(countReadBytes) > limit {
			countReadBytes = int(limit - allWriteBytes)
		}
		toFile.Write(byteset[:countReadBytes])
		allWriteBytes += int64(countReadBytes)
		fmt.Printf("\rТекущий процент: %d (%d / %d)", allWriteBytes*100/limit, allWriteBytes, limit)
	}
	fmt.Println("\nКопирование заверешено")
	return nil
}

func checkErrors(fromFileStat, toFileStat os.FileInfo, offset int64) error {
	if fromFileStat.IsDir() || !fromFileStat.Mode().IsRegular() || toFileStat.IsDir() || !toFileStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	if fromFileStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}
	return nil
}
