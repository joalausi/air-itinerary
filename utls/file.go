package utls

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// ReadFile читает содержимое текстового файла и возвращает его строки.
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

// ReadCSV читает содержимое CSV-файла и возвращает данные как двумерный массив строк.
func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Убираем возможный BOM (если файл в UTF-8 с BOM)
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

	// Вставляем первую строку обратно (если была прочитана)
	return append([][]string{firstLine}, records...), nil
}

// WriteFile записывает данные в выходной текстовый файл.
func WriteFile(filePath string, lines []string) error {
	data := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(filePath, []byte(data), 0644)
}

// ValidateCSV проверяет валидность структуры CSV-файла.
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
