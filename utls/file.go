package utls

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ReadCSV читает содержимое CSV-файла и возвращает данные как двумерный массив строк.
func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// WriteFile записывает данные в выходной текстовый файл.
func WriteFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// ValidateCSV проверяет валидность структуры CSV-файла.
func ValidateCSV(data [][]string) error {
	if len(data) == 0 {
		return errors.New("CSV файл пуст")
	}

	// Проверка, что все строки имеют одинаковое количество колонок
	columnCount := len(data[0])
	for _, row := range data {
		if len(row) != columnCount {
			return errors.New("Некорректная структура CSV")
		}
	}

	// Проверка на пустые данные
	for _, row := range data {
		for _, cell := range row {
			if cell == "" {
				return errors.New("CSV файл содержит пустые ячейки")
			}
		}
	}

	return nil
}
