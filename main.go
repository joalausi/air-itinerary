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
