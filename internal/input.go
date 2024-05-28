package internal

import (
	"errors"
	"fmt"
	"os"
  "io"
  "bufio"
)

func IsInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func GetFile(filePath string) (*os.File, error) {
	if filePath == "" {
		return nil, errors.New("Please input a file")
	}
	if !fileExists(filePath) {
		return nil, errors.New("The file does not exist")
	}
	file, e := os.Open(filePath)
	if e != nil {
		return nil, fmt.Errorf("Unable to read file %s", filePath)
	}
	return file, nil
}

func fileExists(filePath string) bool {
	info, e := os.Stat(filePath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}

func ReadFromFile(r io.Reader, s *string) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		*s += scanner.Text()
		break
	}
	return nil
}

func ReadFromPipe(r io.Reader, s *string) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		*s += scanner.Text()
	}
	return nil
}

