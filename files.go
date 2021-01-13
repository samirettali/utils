package utils

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return readlines(file)
}

func readlines(reader io.Reader) ([]string, error) {
	readBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(readBytes), "\n")
	return lines, nil
}
