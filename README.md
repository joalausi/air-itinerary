# Flight Itinerary Parser

A program that reads a text file with an itinerary and converts airport codes, dates and times into a more readable format into a output file.

## Features

- Converts IATA codes (#XXX) to airport names (works midword with brackets)
- Converts ICAO codes (##XXXX) to airport names (works midword with brackets)
- Supports short form (*) for city names
- Processes codes within brackets: (), [], {}
- Converts dates from D(YYYY-MM-DDThh:mm±hh:mm) to "DD Mon YYYY"
- Converts times in both 12h (T12) and 24h (T24) formats
- Preserves punctuation and surrounding text
- Removes extra blank lines from the output ensuring that the output is clean and readable


## Usage

```go
go run . ./input.txt ./output.txt ./airports_lookup.csv
```

- input: Textfile with the itinerary that needs to be prettified.
- output: Output file, where the prettified itinerary will be written.
- airport-lookup: The path to the CSV file that contains the information for airport code lookup.