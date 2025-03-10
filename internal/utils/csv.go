package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"syscall"
	"todoapp/internal/model"
)

func OpenFile(filePath string) (*os.File, error) {
	var f *os.File

	_, err := os.Stat(filePath)
	if err != nil {
		f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to open file for reading")
		}
		if _, err := f.WriteString(model.HeaderString()); err != nil {
			return nil, fmt.Errorf("could not header to new list csv")
		}
	} else {
		f, err = os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("failed to open file for reading")
		}
	}

	// Exclusive lock obtained on file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func ReadCsv(f *os.File) ([][]string, error) {
	reader := csv.NewReader(f)
	return reader.ReadAll()
}