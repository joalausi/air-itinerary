package utls

import (
	"encoding/csv"
	"errors"
	"os"
)

// LoadAirportData загружает данные аэропортов из CSV-файла и возвращает карту кодов аэропортов и их названий
func LoadAirportData(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("поиск аэропорта не найден")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("поиск аэропорта неверный")
	}

	if len(records) < 2 {
		return nil, errors.New("поиск аэропорта неверный")
	}

	columnIndex := make(map[string]int)
	headers := records[0]
	for i, header := range headers {
		columnIndex[header] = i
	}

	requiredColumns := []string{"name", "iata_code", "icao_code"}
	for _, col := range requiredColumns {
		if _, exists := columnIndex[col]; !exists {
			return nil, errors.New("поиск аэропорта неверный")
		}
	}

	airportMap := make(map[string]string)
	for _, record := range records[1:] {
		if len(record) < len(headers) {
			return nil, errors.New("поиск аэропорта неверный")
		}
		name := record[columnIndex["name"]]
		iata := record[columnIndex["iata_code"]]
		icao := record[columnIndex["icao_code"]]

		if name == "" || (iata == "" && icao == "") {
			return nil, errors.New("поиск аэропорта неверный")
		}

		if iata != "" {
			airportMap[iata] = name
		}
		if icao != "" {
			airportMap[icao] = name
		}
	}

	return airportMap, nil
}

// package utls

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// )

// // LoadAirportLookup загружает airport lookup из CSV-файла
// func LoadAirportLookup(filePath string) (map[string]string, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	lookup := make(map[string]string)

// 	for _, record := range records[1:] { // Пропускаем заголовок
// 		if len(record) < 5 || record[4] == "" || record[0] == "" {
// 			return nil, fmt.Errorf("поиск аэропорта неверный")
// 		}
// 		lookup[record[4]] = record[0] // IATA → Название аэропорта
// 	}

// 	return lookup, nil
// }
