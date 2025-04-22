package main

import (
	"flag"
	"fmt"
	"itinerary/formatter"
	"itinerary/parser"
	"itinerary/utls"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println(utls.Yellow + "Usage: go run . <input.txt> <output.txt> <airport-lookup.csv>" + utls.Reset)
	}
	// Разбираем аргументы командной строки
	flag.Parse()

	if flag.NArg() < 3 {
		flag.Usage()
		os.Exit(1)
	}

	// // Флаг для списка вещей в поездку
	// packingListFlag := flag.Bool("packing-list", false, "Displays a suggested packing list")
	// flag.Usage = func() {
	// 	fmt.Println(yellow, "Itinerary usage:", reset)
	// 	fmt.Println(green, "go run . input.txt output.txt airport-lookup.csv", reset)
	// 	fmt.Println(yellow, "Optional flag:", reset, green, "-packing-list", reset, yellow, "(Displays a suggested packing list)", reset)
	// }

	// if *packingListFlag { // Если передан флаг -packing-list
	// 	PrintPackingList()
	// 	return
	// }

	// args := flag.Args()
	// if len(args) != 3 { // Ожидаем три аргумента
	// 	flag.Usage()
	// 	os.Exit(1)
	// }

	// Parse command-line arguments
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)
	lookupFile := flag.Arg(2)

	// // Чтение входного файла
	// inputData, err := utls.ReadFile(inputFile)
	// if err != nil {
	// 	fmt.Printf(utls.Red + "Ошибка при чтении входного файла: %v\n", err, utls.Reset)
	// 	os.Exit(1)
	// 	return
	// }

	// // Здесь можно добавить обработку inputData, работу с outputPath и lookupPath

	// Read input file
	inputData, err := utls.ReadFile(inputFile)
	if err != nil {
		fmt.Printf(utls.Red+"Error reading input file: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Process lookup data
	lookupData, err := utls.LoadAirportData(lookupFile)
	if err != nil {
		fmt.Printf(utls.Red+"Error loading airport lookup: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Process input data
	flights, err := parser.Parse(inputData, lookupData)
	if err != nil {
		fmt.Printf(utls.Red+"Error parsing data: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Format the processed data
	formattedData, err := formatter.Format(flights, lookupData)
	if err != nil {
		fmt.Printf(utls.Red+"Error formatting data: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Write processed data to output file
	if err := WriteToFile(outputFile, formattedData); err != nil {
		fmt.Println(utls.Red, "Error: Failed to write to file:", err, utls.Reset)
		os.Exit(1)
	}
	fmt.Println(utls.Blue, "Success:", outputFile, "has been created and written successfully!", utls.Reset)
}

// WriteToFile writes the processed data to the specified output file
func WriteToFile(outputFile string, data []string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// // Получаем аргументы командной строки
// inputFile := os.Args[1]
// outputFile :=

// // Загружаем данные
// dataInput, err := utls.ReadFile(inputFile)
// 	if err != nil {
// 	fmt.Println("Error reading input file:", utls.Red, err, utls.Reset)
// 	os.Exit(1)
// }

// dataInput, err := ReadFile(inputFile)
// if err != nil {
// 	fmt.Println(red, "Error reading input file:", err, reset)
// 	os.Exit(1)
// }

// dataLookup, err := utls.LoadAirportData(lookupFile)
// if err != nil {
// 	fmt.Println(red, "Error loading airport lookup:", err, reset)
// 	os.Exit(1)
// }

// // Обрабатываем файл
// processedLookup := ProcessLookupData(dataLookup) // Предположим, есть такая функция
// parsedData := parser.Parse(dataInput, processedLookup)
// if err != nil {
// 	fmt.Println(red, "Error parsing data:", err, reset)
// 	os.Exit(1)
// }

// 	// Read input file
// 	inputData, err := utls.ReadFile(inputFile)
// 	if err != nil {
// 	fmt.Printf(utls.Red+"Error reading input file: %v\n"+utls.Reset, err)
// 	os.Exit(1)
// }

// // Process lookup data
// lookupData, err := utls.LoadAirportData(lookupFile)
// if err != nil {
// 	fmt.Printf(utls.Red+"Error loading airport lookup: %v\n"+utls.Reset, err)
// 	os.Exit(1)
// }

// // Process input data
// processedData := parser.Parse(inputData, lookupData)

// // Write processed data to output file
// err = WriteToFile(outputFile, processedData)
// if err != nil {
// 	fmt.Println(utls.Red + "Error: Failed to write to file" + utls.Reset)
// 	os.Exit(1)
// } else {
// 	fmt.Println(utls.Blue + "Success: " + outputFile + " has been created and written successfully!" + utls.Reset)
// }
// }

// // WriteToFile writes the processed data to the specified output file
// func WriteToFile(outputFile string, data []string) error {
// file, err := os.Create(outputFile)
// if err != nil {
// 	return err
// }
// defer file.Close()

// for _, line := range data {
// 	_, err := file.WriteString(line + "\n")
// 	if err != nil {
// 		return err
// 	}
// }
// return nil
// }

// 	// Записываем результат в файл
// 	err = WriteToFile(outputFile, parsedData)
// 	if err != nil {
// 		fmt.Println(utls.Red + "Error: Failed to write to file", utls.Reset)
// 	} else {
// 		fmt.Println(utls.Blue + "Success:", outputFile, "has been created and written successfully!", utls.Reset)
// 	}
// }
// 	// // Проверяем количество переданных аргументов
// 	// if len(os.Args) != 4 {

// // Функция записи обработанных данных в выходной файл
// func WriteToFile(output string, data string) error {
// 	if data == "" {
// 		return fmt.Errorf("no data to write")
// 	}

// 	err := os.WriteFile(output, []byte(data), 0664)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// import (
// 	"bufio"
// 	"encoding/csv"
// 	"errors"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strings"
// 	"time"
// )

// var (
// 	airportLookup = make(map[string]string)
// 	iataPattern  = regexp.MustCompile(`#([A-Z]{3})`)
// 	icaoPattern  = regexp.MustCompile(`##([A-Z]{4})`)
// 	datePattern  = regexp.MustCompile(`D\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}([+-]\d{2}:\d{2}|Z))\)`)
// 	timePattern  = regexp.MustCompile(`T(12|24)\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}([+-]\d{2}:\d{2}|Z))\)`)
// )

// func main() {
// 	flag.Usage = func() {
// 		fmt.Println("Usage: go run . <input.txt> <output.txt> <airport-lookup.csv>")
// 	}
// 	flag.Parse()

// 	if flag.NArg() < 3 {
// 		flag.Usage()
// 		os.Exit(1)
// 	}

// 	inputFile, outputFile, lookupFile := flag.Arg(0), flag.Arg(1), flag.Arg(2)

// 	if err := loadAirportLookup(lookupFile); err != nil {
// 		fmt.Println("Error loading airport lookup:", err)
// 		os.Exit(1)
// 	}

// 	lines, err := readLines(inputFile)
// 	if err != nil {
// 		fmt.Println("Error reading input file:", err)
// 		os.Exit(1)
// 	}

// 	processedLines := processLines(lines)

// 	if err := writeLines(outputFile, processedLines); err != nil {
// 		fmt.Println("Error writing output file:", err)
// 		os.Exit(1)
// 	}
// }

// func loadAirportLookup(filePath string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return errors.New("Airport lookup not found")
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return errors.New("Airport lookup malformed")
// 	}

// 	for _, row := range records {
// 		if len(row) < 3 {
// 			return errors.New("Airport lookup malformed")
// 		}
// 		iata, icao, name := row[0], row[1], row[2]
// 		airportLookup[iata] = name
// 		airportLookup[icao] = name
// 	}
// 	return nil
// }

// func readLines(filePath string) ([]string, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, errors.New("Input not found")
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}

// 	return lines, nil
// }

// func processLines(lines []string) []string {
// 	var processed []string
// 	for _, line := range lines {
// 		original := line
// 		line = convertAirportCodes(line)
// 		line = convertDates(line)
// 		line = convertTime(line)
// 		line = normalizeWhitespace(line)
// 		processed = append(processed, line)
// 	}
// 	return processed
// }

// func convertAirportCodes(line string) string {
// 	line = iataPattern.ReplaceAllStringFunc(line, func(match string) string {
// 		code := match[1:]
// 		if name, found := airportLookup[code]; found {
// 			return name
// 		}
// 		return match
// 	})

// 	line = icaoPattern.ReplaceAllStringFunc(line, func(match string) string {
// 		code := match[2:]
// 		if name, found := airportLookup[code]; found {
// 			return name
// 		}
// 		return match
// 	})
// 	return line
// }

// func convertDates(line string) string {
// 	return datePattern.ReplaceAllStringFunc(line, func(match string) string {
// 		parsedTime, _ := time.Parse(time.RFC3339, match[2:len(match)-1])
// 		return parsedTime.Format("02 Jan 2006")
// 	})
// }

// func convertTime(line string) string {
// 	return timePattern.ReplaceAllStringFunc(line, func(match string) string {
// 		parts := timePattern.FindStringSubmatch(match)
// 		if len(parts) < 3 {
// 			return match
// 		}
// 		parsedTime, _ := time.Parse(time.RFC3339, parts[2])
// 		if parts[1] == "12" {
// 			return parsedTime.Format("03:04PM (-07:00)")
// 		}
// 		return parsedTime.Format("15:04 (-07:00)")
// 	})
// }

// func normalizeWhitespace(line string) string {
// 	line = strings.ReplaceAll(line, "\v", "\n")
// 	line = strings.ReplaceAll(line, "\f", "\n")
// 	line = strings.ReplaceAll(line, "\r", "\n")
// 	line = regexp.MustCompile(`\n+`).ReplaceAllString(line, "\n")
// 	return line
// }

// func writeLines(filePath string, lines []string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)
// 	for _, line := range lines {
// 		_, err := writer.WriteString(line + "\n")
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return writer.Flush()
// }
