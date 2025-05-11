package formatter

import (
	"fmt"
	"itinerary/parser"
	"itinerary/utls"
	"regexp"
	"strings"
)

// Format takes a Flight list and lookup map (code → airport name),
// and returns ready strings for writing to output.txt.
func Format(flights []parser.Flight, lookup map[string]string) ([]string, error) {
	var out []string

	// Regexp for replacing airport codes: #CODE
	codePattern := regexp.MustCompile(`#([A-Za-z0-9]+)`)
	// Regexp for searching dates D(...)
	datePattern := regexp.MustCompile(`D\([^)]*\)`)

	for _, f := range flights {
		// 1) route(itinerary): Origin → Destination [on DATE]
		if f.Origin != "" && f.Destination != "" {
			originName := lookup[f.Origin]
			if originName == "" {
				originName = f.Origin
			}
			destName := lookup[f.Destination]
			if destName == "" {
				destName = f.Destination
			}

			line := fmt.Sprintf("%s to %s", originName, destName)
			if f.Date != "" {
				// Processing suffix Z) → +00:00)
				tok := f.Date
				if strings.HasSuffix(tok, "Z)") {
					tok = strings.Replace(tok, "Z)", "+00:00)", 1)
				}
				formattedDate := utls.FormatDateTime(tok)
				line = fmt.Sprintf("%s on %s", line, formattedDate)
			}
			out = append(out, line)
		}

		// 2) (RawLines): code/date/time inside arbitrary text
		if len(f.RawLines) > 0 {
			for _, raw := range f.RawLines {
				// 2.1) First, replace the airport codes
				line := codePattern.ReplaceAllStringFunc(raw, func(match string) string {
					code := match[1:]
					if name, ok := lookup[code]; ok {
						return name
					}
					return match
				})

				// 2.2) replace dates D(...)
				line = datePattern.ReplaceAllStringFunc(line, func(match string) string {
					tok := match
					if strings.HasSuffix(tok, "Z)") {
						tok = strings.Replace(tok, "Z)", "+00:00)", 1)
					}
					return utls.FormatDateTime(tok)
				})

				// 2.3) Processing timestamps T12(...) и T24(...)
				parts := strings.Fields(line)
				for i, token := range parts {
					if strings.HasPrefix(token, "T12(") || strings.HasPrefix(token, "T24(") {
						tkn := token
						// transform Z) t +00:00)
						if strings.HasSuffix(tkn, "Z)") {
							tkn = strings.Replace(tkn, "Z)", "+00:00)", 1)
						}
						parts[i] = utls.FormatDateTime(tkn)
					}
				}
				line = strings.Join(parts, " ")
				out = append(out, line)
			}
		}

		// 3) Date-only blocks (no route and no RawLines)
		if f.Date != "" && f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 {
			tok := f.Date
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			out = append(out, utls.FormatDateTime(tok))
		}

		// 4) block Departure
		if f.Departure != "" {
			tok := f.Departure
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			formatted := utls.FormatDateTime(tok)
			if f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 && f.Date == "" {
				// if its only one label in block
				out = append(out, formatted)
			} else {
				out = append(out, "Departure: "+formatted)
			}
		}

		// 5) block Arrival
		if f.Arrival != "" {
			tok := f.Arrival
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			formatted := utls.FormatDateTime(tok)
			if f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 && f.Date == "" {
				out = append(out, formatted)
			} else {
				out = append(out, "Arrival: "+formatted)
			}
		}

		// Empty line separator between blocks
		out = append(out, "")
	}

	// Remove the final blank line if there is one
	if len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}

	return out, nil
}
