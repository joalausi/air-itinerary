package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// Flight описывает один маршрут
type Flight struct {
	Origin      string
	Destination string
	Time        string
}

// ParseFile считывает текстовый файл и возвращает список маршрутов
func Parse(filePath string) ([]Flight, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var flights []Flight
	scanner := bufio.NewScanner(file)

	// Регулярное выражение для поиска маршрутов
	flyPattern := regexp.MustCompile(`([A-Z]{3})-([A-Z]{3})\s+(.+)`)

	// Читаем файл построчно
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = cleanText(line) // Очистка спецсимволов

		// Ищем маршрут в строке
		match := flyPattern.FindStringSubmatch(line)
		if match != nil {
			flights = append(flights, Flight{
				Origin:      match[1],
				Destination: match[2],
				Time:        match[3],
			})
		}
	}

	return flights, scanner.Err()
}

// cleanText убирает лишние символы переноса строки и заменяет их на "\n"
func cleanText(input string) string {
	input = strings.ReplaceAll(input, "\v", "\n")
	input = strings.ReplaceAll(input, "\f", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")

	// Убираем лишние пустые строки
	lines := strings.Split(input, "\n")
	var cleaned []string
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 || (len(cleaned) > 0 && cleaned[len(cleaned)-1] != "") {
			cleaned = append(cleaned, line)
		}
	}
	return strings.Join(cleaned, "\n")
}
