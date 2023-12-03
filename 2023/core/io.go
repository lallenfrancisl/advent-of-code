package core

import (
	"os"
	"strings"
)

// Read a file and split them into lines
func ReadLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
