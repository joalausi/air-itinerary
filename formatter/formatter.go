package formatter

import (
	"fmt"
	"itinerary/parser"
	"itinerary/utls"
	"regexp"
	"strings"
)

// Format takes a Flight list and lookup maps (code → airport, code→city)
// and returns ready strings for writing to output.txt.
func Format(flights []parser.Flight, lookupName, lookupCity map[string]string) ([]string, error) {
	var out []string

	// Regexp for replacing airport codes: optional "*" then "#" or "##", then code
	codePattern := regexp.MustCompile(`(\*)?(?:#{1,2})([A-Za-z0-9]+)`) // groups[1]=="*" or "", groups[2]==code
	// Regexp for searching dates D(...)
	datePattern := regexp.MustCompile(`D\([^)]*\)`)

	for _, f := range flights {
		// 1) route(itinerary): Origin → Destination [on DATE]
		if f.Origin != "" && f.Destination != "" {
			// choose city or full name based on flags
			var originName string
			if f.OriginCityOnly {
				if city, ok := lookupCity[f.Origin]; ok {
					originName = city
				}
			}
			if originName == "" {
				originName = lookupName[f.Origin]
				if originName == "" {
					originName = f.Origin
				}
			}

			var destName string
			if f.DestCityOnly {
				if city, ok := lookupCity[f.Destination]; ok {
					destName = city
				}
			}
			if destName == "" {
				destName = lookupName[f.Destination]
				if destName == "" {
					destName = f.Destination
				}
			}

			line := fmt.Sprintf("%s to %s", originName, destName)
			if f.Date != "" {
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
				line := codePattern.ReplaceAllStringFunc(raw, func(match string) string {
					groups := codePattern.FindStringSubmatch(match)
					star, code := groups[1], groups[2]
					if star == "*" {
						if city, ok := lookupCity[code]; ok {
							return city
						}
					}
					if name, ok := lookupName[code]; ok {
						return name
					}
					return match
				})

				line = datePattern.ReplaceAllStringFunc(line, func(match string) string {
					tok := match
					if strings.HasSuffix(tok, "Z)") {
						tok = strings.Replace(tok, "Z)", "+00:00)", 1)
					}
					return utls.FormatDateTime(tok)
				})

				// timestamps T12(), T24()
				parts := strings.Fields(line)
				for i, token := range parts {
					if strings.HasPrefix(token, "T12(") || strings.HasPrefix(token, "T24(") {
						tkn := token
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

		// 3) Date-only
		if f.Date != "" && f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 {
			tok := f.Date
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			out = append(out, utls.FormatDateTime(tok))
		}

		// 4) Departure
		if f.Departure != "" {
			tok := f.Departure
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			formatted := utls.FormatDateTime(tok)
			if f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 && f.Date == "" {
				out = append(out, formatted)
			} else {
				out = append(out, "Departure: "+formatted)
			}
		}

		// 5) Arrival
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
	}
	// Remove the final blank line if there is one
	if len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}

	return out, nil
}
