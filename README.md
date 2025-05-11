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

```
/itinerary
│
├── main.go                # Entry
├── parser/                # Route parsing module
│   ├── parser.go
├── formatter/             # Data formatting
│   ├── formatter.go
├── utils/                 # Utilities (reading files, transform, error handling)
│   ├── file.go
│   ├── airport_lookup.go
│   ├── time_parser.go
├──input.txt
├──output.txt
├──airports_lookup.csv 
```

## Usage

```go
go run . ./input.txt ./output.txt ./airports_lookup.csv
```

- input: Textfile with the itinerary that needs to be prettified.
- output: Output file, where the prettified itinerary will be written.
- airport-lookup: The path to the CSV file that contains the information for airport code lookup.

### Help
To display the help message, run the program with the -h flag.
```go
go run . -h
```

### IATA Codes
- Format: `#XXX`
- Example: `#HEL` → "Helsinki-Vantaa Airport"
- Short form: `*#HEL` → "Helsinki"

### ICAO Codes
- Format: `##XXXX`
- Example: `##EGLL` → "London Heathrow Airport"
- Short form: `*##EGLL` → "London"

### Brackets and Punctuation
- Supports brackets: `(#HEL)`, `[#LAX]`, `{##EGLL}`
- Preserves punctuation: `#HEL.`, `##EGLL,`
- Works mid-word with brackets: `text(#LAX)more`

### Dates and Times
- Date: `D(2024-01-15T08:30Z)` → "15 Jan 2024"
- 24h time: `T24(2042-09-01T21:43Z)` → "21:43 (00:00)"
- 12h time: `T12(2024-07-23T15:29-11:00)` → "03:29 PM (-11:00)"

## CSV File Structure

The airport lookup file should contain the following columns(can be in any order):
- name
- iso_country
- municipality
- icao_code
- iata_code
- coordinates


## Project Files 
