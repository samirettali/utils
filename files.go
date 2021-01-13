package utils

import (
	"bufio"
	"os"
)

func readfile(filename string) ([]string, error) {
	file, err := os.Open("sample.txt")
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(text, scanner.Text())
	}

	return lines
}
