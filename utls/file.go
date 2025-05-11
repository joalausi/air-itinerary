package utls

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// ReadFile reads contents of a text file and returns lines as a slice of strings.
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// ReadCSV reads a CSV file and returns its contents as a slice of string slices.
func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Remove possible BOM (Byte Order Mark) from the first line(if the file is in UTF-8 with BOM)
	firstLine, err := reader.Read()
	if err != nil {
		return nil, err
	}
	if len(firstLine) > 0 {
		firstLine[0] = strings.TrimPrefix(firstLine[0], "\uFEFF")
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Insert the first line back (if it was read)
	return append([][]string{firstLine}, records...), nil
}

// WriteFile writes a slice of strings to output file.
func WriteFile(filePath string, lines []string) error {
	data := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(filePath, []byte(data), 0644)
}

// ValidateCSV checks the structure of a CSV file.
func ValidateCSV(data [][]string) error {
	if len(data) == 0 {
		return errors.New("CSV file is empty")
	}

	columnCount := len(data[0])
	for i, row := range data {
		if len(row) != columnCount {
			return errors.New("CSV structure is inconsistent (row " + string(i+1) + ")")
		}
	}

	return nil
}
