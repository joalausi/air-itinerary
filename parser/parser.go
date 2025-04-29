package parser

import (
	"bufio"
	"fmt"
	"itinerary/utls"
	"os"
	"regexp"
	"strings"
)

// Flight describes one route
type Flight struct {
	Origin      string
	Destination string
	DateTime    string
}

// Parse reads a text file and returns list of routes
func Parse(filePath string) ([]Flight, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("%sError: Failed to open file %s: %w%s", utls.Red, filePath, err, utls.Reset)
	}
	defer file.Close()

	var flights []Flight
	scanner := bufio.NewScanner(file)

	// Regular expression for finding routes
	flyPattern := regexp.MustCompile(`^([A-Z]{3})-([A-Z]{3})\s+D?\(?(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(?:Z|[-+]\d{2}:\d{2}))\)?$`)

	// read file line by line
	for scanner.Scan() {
		line := cleanText(scanner.Text())

		// looking for aroute in line
		match := flyPattern.FindStringSubmatch(line)
		if match != nil {
			flights = append(flights, Flight{
				Origin:      match[1],
				Destination: match[2],
				DateTime:    match[3],
			})
		} else {
			fmt.Printf("%sWarning: string does not match route format: %s\n%s", utls.Yellow, line, utls.Reset)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("%sError reading file: %w%s", utls.Red, err, utls.Reset)
	}

	return flights, nil
}

// cleanText убирает лишние символы переноса строки и заменяет их на "\n"
func cleanText(input string) string {
	replacer := strings.NewReplacer("\v", "\n", "\f", "\n", "\r", "\n")
	input = replacer.Replace(input)

	// Убираем лишние пустые строки
	lines := strings.Split(input, "\n")
	var cleaned []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" || (len(cleaned) > 0 && cleaned[len(cleaned)-1] != "") {
			cleaned = append(cleaned, trimmed)
		}
	}
	return strings.Join(cleaned, "\n")
}
