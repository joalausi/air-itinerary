package utls

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

// LoadAirportData loads airport data from a CSV file and returns a map of airport codes and names.
func LoadAirportData(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.TrimLeadingSpace = true

	// 1) Read first line (header) and validate
	header, err := r.Read()
	if err != nil {
		return nil, errors.New("airport lookup malformed")
	}
	for i := range header {
		header[i] = strings.TrimSpace(header[i])
	}

	expected := []string{
		"name",
		"iso_country",
		"municipality",
		"icao_code",
		"iata_code",
		"coordinates",
	}
	// Find the index of each expected column in the header
	idx := make(map[string]int, len(expected))
	for _, col := range expected {
		found := false
		for j, h := range header {
			if h == col {
				idx[col] = j
				found = true
				break
			}
		}
		if !found {
			return nil, errors.New("airport lookup malformed")
		}
	}

	lookup := make(map[string]string)

	// 2) Read the rest of the file
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("airport lookup malformed")
		}
		// Check if all expected columns are present
		for _, col := range expected {
			cell := strings.TrimSpace(rec[idx[col]])
			if cell == "" {
				return nil, errors.New("airport lookup malformed")
			}
		}
		name := rec[idx["name"]]
		icao := rec[idx["icao_code"]]
		iata := rec[idx["iata_code"]]

		// Map for both codes
		lookup[icao] = name
		lookup[iata] = name
	}

	return lookup, nil
}
