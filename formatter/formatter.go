package formatter

import (
	"fmt"
	"itinerary/parser"
	"itinerary/utls"
	"regexp"
)

// Format форматирует блоки Flight в текстовый вид с заменой кодов и форматированием времени.
func Format(flights []parser.Flight, lookup map[string]string) ([]string, error) {
	var out []string

	codePattern := regexp.MustCompile(`#([A-Za-z0-9]+)`)
	datePattern := regexp.MustCompile(`D\([^)]*\)`)

	for _, f := range flights {
		// Форматирование маршрута (если указан Origin и Destination)
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
				formattedDate := utls.FormatDateTime(f.Date)
				line = fmt.Sprintf("%s on %s", line, formattedDate)
			}
			out = append(out, line)
		}
		// Печать необработанных строк (RawLines) с заменой кодов и дат
		if len(f.RawLines) > 0 {
			for _, raw := range f.RawLines {
				line := raw
				line = codePattern.ReplaceAllStringFunc(line, func(match string) string {
					code := match[1:]
					if name, ok := lookup[code]; ok {
						return name
					}
					return match
				})
				line = datePattern.ReplaceAllStringFunc(line, func(match string) string {
					return utls.FormatDateTime(match)
				})
				out = append(out, line)
			}
		}
		// Отдельное поле Date (без Origin/Destination)
		if f.Date != "" && f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 {
			formattedDate := utls.FormatDateTime(f.Date)
			out = append(out, formattedDate)
		}
		// Поле Departure
		if f.Departure != "" {
			formattedDep := utls.FormatDateTime(f.Departure)
			if f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 {
				out = append(out, formattedDep)
			} else {
				out = append(out, "Departure: "+formattedDep)
			}
		}
		// Поле Arrival
		if f.Arrival != "" {
			formattedArr := utls.FormatDateTime(f.Arrival)
			if f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 {
				out = append(out, formattedArr)
			} else {
				out = append(out, "Arrival: "+formattedArr)
			}
		}
		// Разрыв строки между блоками полётов
		out = append(out, "")
	}
	// Удаление последней пустой строки
	if len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return out, nil
}
