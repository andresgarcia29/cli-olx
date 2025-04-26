package helpers

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileAsLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func RemoveIgnoreFileSection(path string) (string, error) {
	lines, err := ReadFileAsLines(path)
	if err != nil {
		return "", err
	}

	newLinesOfFile := []string{}
	ignore := false
	for _, line := range lines {
		if line == "## Ignore" {
			ignore = true
		} else if line == "## /Ignore" {
			ignore = false
		} else {
			if !ignore {
				newLinesOfFile = append(newLinesOfFile, line)
			}
		}
	}

	return strings.Join(newLinesOfFile, "\n"), nil
}

func GetIgnoreFileSection(path string) (string, error) {
	lines, err := ReadFileAsLines(path)
	if err != nil {
		return "", err
	}

	var newLinesOfFile []string
	ignore := true
	for _, line := range lines {
		if line == "## Ignore" {
			ignore = false
			newLinesOfFile = append(newLinesOfFile, "\n\n"+line)
		} else if line == "## /Ignore" {
			ignore = true
			newLinesOfFile = append(newLinesOfFile, line+"\n")
		} else {
			if !ignore {
				newLinesOfFile = append(newLinesOfFile, line)
			}
		}
	}

	return strings.Join(newLinesOfFile, "\n"), nil
}

func SaveTempFile(content string) (string, error) {
	tempFile, err := os.CreateTemp("", "tempfile")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = tempFile.WriteString(content)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}
