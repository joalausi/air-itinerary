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
	
	// Parse command-line arguments
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


	// Parse command-line arguments
	inputFile := flag.Arg(0)    // Path to input text file
    outputFile := flag.Arg(1)   // Path to output text file (to be created)
    lookupFile := flag.Arg(2)   // Path to airport lookup CSV file

	// // Здесь можно добавить обработку inputData, работу с outputPath и lookupPath

	// Read input file
	inputData, err := utls.ReadFile(inputFile)
	if err != nil {
		fmt.Printf(utls.Red+"Error reading input file: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	//Process lookup data
	lookupData, err := utls.LoadAirportData(lookupFile)
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
	flights, err := parser.Parse(inputData, lookupData)
	if err != nil {
		fmt.Printf(utls.Red+"Parse error: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Format the processed data
	formattedData, err := formatter.Format(flights, lookupData)
	if err != nil {
		fmt.Printf(utls.Red+"Format error: %v\n"+utls.Reset, err)
		os.Exit(1)
	}

	// Write processed data to output file
	if err := WriteToFile(outputFile, formattedData); err != nil {
		fmt.Printf(utls.Red+"Error: Failed to write to file: %v\n"+utls.Reset, err)
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

    // Read input file (each line as an element in the slice)
    inputData, err := utls.ReadFile(inputFile)
    if err != nil {
        fmt.Printf(utls.Red+"Error reading input file %s: %v\n"+utls.Reset, inputFile, err)
        os.Exit(1)
    }

    // Load airport data from CSV file into a map (code -> full name)
    lookupData, err := utls.LoadAirportData(lookupFile)
    if err != nil {
        // If loading from CSV fails (function not implemented or file missing), use a fallback map
        fmt.Println(utls.Yellow + "Warning: using default airport data map (limited) due to error:", err, utls.Reset)
    }

    // Parse the input lines into a list of Flight structures
    flights, err := parser.Parse(inputData, lookupData)
    if err != nil {
        fmt.Printf(utls.Red+"Error parsing data: %v\n"+utls.Reset, err)
        os.Exit(1)
    }

    // Format the list of Flight structures into output lines
    formattedData, err := formatter.Format(flights, lookupData)
    if err != nil {
        fmt.Printf(utls.Red+"Error formatting data: %v\n"+utls.Reset, err)
        os.Exit(1)
    }

    // Write the formatted output lines to the output file
    err = WriteToFile(outputFile, formattedData)
    if err != nil {
        fmt.Printf(utls.Red+"Error: Failed to write to file %s: %v\n"+utls.Reset, outputFile, err)
        os.Exit(1)
    }

    // If we reached here, everything succeeded
    fmt.Println(utls.Blue + "Success: output written to " + outputFile + utls.Reset)
}

// WriteToFile writes the slice of strings to the specified output file, one line at a time.
func WriteToFile(filePath string, lines []string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // Write each line followed by a newline
    for _, line := range lines {
        _, err := file.WriteString(line + "\n")
        if err != nil {
            return err
        }
    }
    return nil
}
