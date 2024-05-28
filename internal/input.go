package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"unicode"
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


// cleanString removes all quotes and whitespace from the input string
func cleanString(str string) string {
    var builder strings.Builder
    for _, ch := range str {
        if ch != '"' && ch != '\'' && ch != '`' && !unicode.IsSpace(ch) {
            builder.WriteRune(ch)
        }
    }
    return builder.String()
}

// isValidURL checks if the given string is a valid URL
func isValidURL(str string) bool {
    u, err := url.Parse(str)
    if err != nil {
        return false
    }

    // Check if the URL has a valid scheme and host
    return u.Scheme != "" && u.Host != ""
}
