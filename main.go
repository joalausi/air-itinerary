package main

import (
	"flag"
	"fmt"
	"itinerary/formatter"
	"itinerary/parser"
	"itinerary/utls"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Println(utls.Yellow + "Itinerary usage:" + utls.Reset)
		fmt.Println(utls.Yellow + "Usage: go run . ./input.txt ./output.txt ./airport-lookup.csv" + utls.Reset)
	}

	// Parse command-line arguments
	flag.Parse()
	if flag.NArg() < 3 {
		flag.Usage()
		os.Exit(1)
	}

	// Parse command-line arguments
	inputFile := flag.Arg(0)  // Path to input text file
	outputFile := flag.Arg(1) // Path to output text file (to be created)
	lookupFile := flag.Arg(2) // Path to airport lookup CSV file

	// Read input file
	inputData, err := utls.ReadFile(inputFile)
	if err != nil {
		fmt.Printf(utls.Red+"Error reading input file: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Clean input data
	for i, line := range inputData {
		inputData[i] = cleanText(line)
	}

	//Process lookup data
	lookupName, lookupCity, err := utls.LoadAirportData(lookupFile)
	if err != nil {
		fmt.Printf(utls.Red+"Error loading airport lookup: %v\n"+utls.Reset, err)
		os.Exit(1)
		// lookupData = map[string]string{
		// "LAX":  "Los Angeles International Airport",
		// "EGLL": "London Heathrow Airport",
		// "HAJ":  "Hannover Airport",
		// "EDDW": "Bremen Airport",
		// "CEG":  "Hawarden Airport",
		// "WWK":  "Wewak International Airport",
		// "BIH":  "Eastern Sierra Regional Airport",
		// }
	}

	// parse input data
	flights, err := parser.Parse(inputData, lookupName)
	if err != nil {
		fmt.Printf(utls.Red+"Parse error: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Format the processed data
	formattedData, err := formatter.Format(flights, lookupName, lookupCity)
	if err != nil {
		fmt.Printf(utls.Red+"Format error: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Write the formatted output lines to the output file
	err = WriteToFile(outputFile, formattedData)
	if err != nil {
		fmt.Printf(utls.Red+"Error: Failed to write to file: %v\n"+utls.Reset, err)
		os.Exit(1)
	}
	fmt.Println(utls.Blue, "Success:", outputFile, "has been created and written successfully!", utls.Reset)
}

// WriteToFile writes the processed data to the specified output file
func WriteToFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// cleanText cleans the input string by replacing certain characters with newlines
func cleanText(s string) string {
	replacer := strings.NewReplacer("\v", "\n", "\f", "\n", "\r", "\n")
	return replacer.Replace(s)
}
