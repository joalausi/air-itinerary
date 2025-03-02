package main

import (
	"flag"
	"fmt"
	"itinerary/formatter"
	"itinerary/parser"
	"itinerary/utls"
)

func main() {
	flag.Usage = func() {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
	}

	flag.Parse()
	args := flag.Args()

	if len(args) != 3 {
		flag.Usage()
		return
	}

	inputPath := args[0]
	outputPath := args[1]
	airportLookupPath := args[2]

	inputData, err := utls.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Вход не найден")
		return
	}

	airportData, err := utls.LoadAirportData(airportLookupPath)
	if err != nil {
		fmt.Println("Поиск аэропорта не найден")
		return
	}

	parsedData := parser.Parse(inputData, airportData)
	formattedData := formatter.Format(parsedData)

	err = utls.WriteFile(outputPath, formattedData)
	if err != nil {
		fmt.Println("Ошибка записи в файл")
		return
	}
}

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"itinerary/formatter"
// 	"itinerary/parser"
// 	"itinerary/utls"
// )

// func main() {
// 	// Определение флага для вывода справочной информации
// 	flag.Usage = func() {
// 		fmt.Println("itinerary usage:")
// 		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
// 	}

// 	flag.Parse()
// 	args := flag.Args()

// 	// Проверяем количество аргументов
// 	if len(args) != 3 {
// 		flag.Usage()
// 		return
// 	}

// 	inputPath := args[0]
// 	outputPath := args[1]
// 	airportLookupPath := args[2]

// 	// Чтение входных данных
// 	inputData, err := utls.ReadFile(inputPath)
// 	if err != nil {
// 		fmt.Println("Вход не найден")
// 		return
// 	}

// 	// Чтение данных аэропортов
// 	airportData, err := utls.LoadAirportData(airportLookupPath)
// 	if err != nil {
// 		fmt.Println("Поиск аэропорта не найден")
// 		return
// 	}

// 	// Обработка входных данных
// 	parsedData := parser.Parse(inputData, airportData)
// 	formattedData := formatter.Format(parsedData)

// 	// Запись результата в файл
// 	err = utls.WriteFile(outputPath, formattedData)
// 	if err != nil {
// 		fmt.Println("Ошибка записи выходного файла")
// 	}
// }

// COLOR
// func main() {

// 	// Проверяем аргументы
// 	if len(os.Args) != 4 {
// 		color.Red("Usage: go run . <input file> <output file> <airport lookup file>go run . ./input.txt ./output.txt ./airport-lookup.csv")
// 		os.Exit(1)
// 	}

// 	inputFile := os.Args[1]
// 	outputFile := os.Args[2]
// 	lookupFile := os.Args[3]

// 	// Читаем маршруты
// 	flights, err := parser.ParseFile(inputFile)
// 	if err != nil {
// 		color.Red("Error reading input file: %v", err)
// 		os.Exit(1)
// 	}

// 	// Загружаем данные аэропортов
// 	lookup, err := utils.LoadAirportLookup(lookupFile)
// 	if err != nil {
// 		color.Red("Error reading airport lookup file: %v", err)
// 		os.Exit(1)
// 	}

// 	// Форматируем маршрут
// 	formatted := formatter.FormatItinerary(flights, lookup)

// 	// Записываем результат
// 	err = utils.WriteFile(outputFile, formatted)
// 	if err != nil {
// 		color.Red("Error writing output file: %v", err)
// 		os.Exit(1)
// 	}

// 	// Выводим цветной результат
// 	formatter.FormatColored(formatted)

// 	color.Green("✔ Itinerary successfully formatted!")
// }
