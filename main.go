package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Цвета для терминала
	yellow := "\033[93m"
	green := "\033[32m"
	red := "\033[31m"
	reset := "\033[0m"
	blue := "\033[44m"

	// Флаг для списка вещей в поездку
	packingListFlag := flag.Bool("packing-list", false, "Displays a suggested packing list")
	flag.Usage = func() {
		fmt.Println(yellow, "Itinerary usage:", reset)
		fmt.Println(green, "go run . input.txt output.txt airport-lookup.csv", reset)
		fmt.Println(yellow, "Optional flag:", reset, green, "-packing-list", reset, yellow, "(Displays a suggested packing list)", reset)
	}

	// Разбираем аргументы командной строки
	flag.Parse()

	if *packingListFlag { // Если передан флаг -packing-list
		PrintPackingList()
		return
	}

	args := flag.Args()
	if len(args) != 3 { // Ожидаем три аргумента
		flag.Usage()
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := args[1]
	lookupFile := args[2]

	// Загружаем данные
	dataInput, err := ReadInputFile(inputFile)
	if err != nil {
		fmt.Println(red, "Error reading input file:", err, reset)
		os.Exit(1)
	}

	dataLookup, err := ReadAirportLookupFile(lookupFile)
	if err != nil {
		fmt.Println(red, "Error loading airport lookup:", err, reset)
		os.Exit(1)
	}

	// Обрабатываем файл
	parsedData := ParseInputFile(dataInput, dataLookup)

	// Записываем результат в файл
	err = WriteToFile(outputFile, parsedData)
	if err != nil {
		fmt.Println(red, "Error: Failed to write to file", reset)
	} else {
		fmt.Println(blue, "Success:", outputFile, "has been created and written successfully!", reset)
	}
}

// Функция записи обработанных данных в выходной файл
func WriteToFile(output string, data string) error {
	if data == "" {
		return fmt.Errorf("no data to write")
	}

	err := os.WriteFile(output, []byte(data), 0664)
	if err != nil {
		return err
	}
	return nil
}

// Функция вывода списка рекомендуемых вещей в поездку
func PrintPackingList() {
	green := "\033[32m"
	brightGreen := "\033[92m"
	red := "\033[91m"
	reset := "\033[0m"

	fmt.Println(green, "\nWelcome to the travel itinerary tool!", reset)
	fmt.Println(red, "\nHere is a suggested packing list for your trip:", reset)

	packingList := []string{
		"Passport", "Boarding Pass", "Phone", "Chargers",
		"Laptop", "Headphones", "Travel Pillow", "Snacks",
		"Water Bottle", "Sunglasses", "Notebook & Pen",
	}

	for _, item := range packingList {
		fmt.Println(brightGreen, "-", item, reset)
	}
}
