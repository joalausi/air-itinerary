package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Flight описывает один маршрут
type Flight struct {
	Origin      string
	Destination string
	DateTime    string
}

// ParseFile считывает текстовый файл и возвращает список маршрутов
func Parse(filePath string) ([]Flight, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Ошибка: не удалось открыть файл %s: %w", filePath, err)
	}
	defer file.Close()

	var flights []Flight
	scanner := bufio.NewScanner(file)

	// Регулярное выражение для поиска маршрутов
	flyPattern := regexp.MustCompile(`^([A-Z]{3})-([A-Z]{3})\s+D?\(?(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(?:Z|[-+]\d{2}:\d{2}))\)?$`)

	// Читаем файл построчно
	for scanner.Scan() {
		line := cleanText(scanner.Text())

		// Ищем маршрут в строке
		match := flyPattern.FindStringSubmatch(line)
		if match != nil {
			flights = append(flights, Flight{
				Origin:      match[1],
				Destination: match[2],
				DateTime:    match[3],
			})
		} else {
			fmt.Printf("Предупреждение: строка не соответствует формату маршрута: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при чтении файла: %w", err)
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
