package formatter

import (
	"fmt"
	"itinerary/parser"
	"itinerary/utls"
	"regexp"
	"strings"
)

// Format принимает список Flight и карту lookup (код → название аэропорта),
// и возвращает готовые строки для записи в output.txt.
func Format(flights []parser.Flight, lookup map[string]string) ([]string, error) {
	var out []string

	// Регэксп для замены кодов аэропортов: #CODE
	codePattern := regexp.MustCompile(`#([A-Za-z0-9]+)`)
	// Регэксп для поиска дат D(...)
	datePattern := regexp.MustCompile(`D\([^)]*\)`)

	for _, f := range flights {
		// 1) Маршрут: Origin → Destination [on DATE]
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
				// Обработка суффикса Z) → +00:00)
				tok := f.Date
				if strings.HasSuffix(tok, "Z)") {
					tok = strings.Replace(tok, "Z)", "+00:00)", 1)
				}
				formattedDate := utls.FormatDateTime(tok)
				line = fmt.Sprintf("%s on %s", line, formattedDate)
			}
			out = append(out, line)
		}

		// 2) Сырые строки (RawLines): код/дата/время внутри произвольного текста
		if len(f.RawLines) > 0 {
			for _, raw := range f.RawLines {
				// 2.1) Сначала заменяем коды аэропортов
				line := codePattern.ReplaceAllStringFunc(raw, func(match string) string {
					code := match[1:]
					if name, ok := lookup[code]; ok {
						return name
					}
					return match
				})

				// 2.2) Заменяем даты D(...)
				line = datePattern.ReplaceAllStringFunc(line, func(match string) string {
					tok := match
					if strings.HasSuffix(tok, "Z)") {
						tok = strings.Replace(tok, "Z)", "+00:00)", 1)
					}
					return utls.FormatDateTime(tok)
				})

				// 2.3) Обрабатываем временные метки T12(...) и T24(...)
				parts := strings.Fields(line)
				for i, token := range parts {
					if strings.HasPrefix(token, "T12(") || strings.HasPrefix(token, "T24(") {
						tkn := token
						// Преобразуем Z) в +00:00)
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

		// 3) Блоки только с датой (без маршрута и без RawLines)
		if f.Date != "" && f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 {
			tok := f.Date
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			out = append(out, utls.FormatDateTime(tok))
		}

		// 4) Поле Departure
		if f.Departure != "" {
			tok := f.Departure
			if strings.HasSuffix(tok, "Z)") {
				tok = strings.Replace(tok, "Z)", "+00:00)", 1)
			}
			formatted := utls.FormatDateTime(tok)
			if f.Origin == "" && f.Destination == "" && len(f.RawLines) == 0 && f.Date == "" {
				// если это единственная метка в блоке
				out = append(out, formatted)
			} else {
				out = append(out, "Departure: "+formatted)
			}
		}

		// 5) Поле Arrival
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

		// Пустая строка-разделитель между блоками
		out = append(out, "")
	}

	// Удаляем финальную пустую строку, если она есть
	if len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}

	return out, nil
}
