package parser

import (
	"regexp"
	"strings"
)

type Flight struct {
	Origin      string
	Destination string
	Departure   string
	Arrival     string
	Date        string
	RawLines    []string
}

// Parse transforms input strings into a Flights slice according to the given rules.
func Parse(lines []string, lookup map[string]string) ([]Flight, error) {
	var flights []Flight
	currentIndex := -1

	// Preprocessing strings: split by ^t or ^f markers (insert line feed)
	var procLines []string
	for _, line := range lines {
		if strings.Contains(line, "^") {
			var buf strings.Builder
			for i := 0; i < len(line); i++ {
				if line[i] == '^' && i+1 < len(line) && (line[i+1] == 't' || line[i+1] == 'f') {
					// Break line at marker
					procLines = append(procLines, buf.String())
					buf.Reset()
					if line[i+1] == 't' {
						i++
					} else if line[i+1] == 'f' {
						// Skip all consecutive 'f''
						j := i + 1
						for j < len(line) && line[j] == 'f' {
							j++
						}
						i = j - 1
					}
					continue
				}
				buf.WriteByte(line[i])
			}
			procLines = append(procLines, buf.String())
		} else {
			procLines = append(procLines, line)
		}
	}

	// Regular expression for strings like "X to Y on D(...)"
	routePattern := regexp.MustCompile(`[#*()]*([A-Za-z0-9]+)[^A-Za-z0-9]+to[^A-Za-z0-9]*[#*()]*([A-Za-z0-9]+)[^A-Za-z0-9]*on\s*D\(([^)]+)\)`)

	for _, raw := range procLines {
		line := strings.TrimSpace(raw)
		if line == "" {
			// Empty string - close current flight block
			currentIndex = -1
			continue
		}
		// Processing label "Departure:"
		if strings.HasPrefix(line, "Departure:") {
			dep := strings.TrimSpace(strings.TrimPrefix(line, "Departure:"))
			if currentIndex >= 0 {
				flights[currentIndex].Departure = dep
			} else {
				flights = append(flights, Flight{Departure: dep})
				currentIndex = len(flights) - 1
			}
			continue
		}
		// Processing the label "Arrival:"
		if strings.HasPrefix(line, "Arrival:") {
			arr := strings.TrimSpace(strings.TrimPrefix(line, "Arrival:"))
			if currentIndex >= 0 {
				flights[currentIndex].Arrival = arr
			} else {
				flights = append(flights, Flight{Arrival: arr})
				currentIndex = len(flights) - 1
			}
			continue
		}
		// Separate libes T12(...), T24(...), D(...)
		if strings.HasPrefix(line, "T12(") {
			flights = append(flights, Flight{Departure: line})
			currentIndex = len(flights) - 1
			continue
		}
		if strings.HasPrefix(line, "T24(") {
			flights = append(flights, Flight{Arrival: line})
			currentIndex = len(flights) - 1
			continue
		}
		if strings.HasPrefix(line, "D(") {
			flights = append(flights, Flight{Date: line})
			currentIndex = len(flights) - 1
			continue
		}
		// Route lines of the type "X to Y on D(...)"
		if len(line) > 0 {
			first := line[0]
			if first == '#' || first == '(' || first == '*' {
				if match := routePattern.FindStringSubmatch(line); match != nil {
					origin := match[1]
					dest := match[2]
					datePart := match[3]
					flights = append(flights, Flight{
						Origin:      origin,
						Destination: dest,
						Date:        "D(" + datePart + ")",
					})
					currentIndex = len(flights) - 1
					continue
				}
			}
		}
		// Other lines - raw text
		if currentIndex >= 0 {
			flights[currentIndex].RawLines = append(flights[currentIndex].RawLines, line)
		} else {
			flights = append(flights, Flight{RawLines: []string{line}})
			currentIndex = len(flights) - 1
		}
	}
	return flights, nil
}
