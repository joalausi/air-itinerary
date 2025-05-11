package utls

import (
	"fmt"
	"regexp"
	"time"
)

// monthNames maps month numbers to their abbreviated names
var monthNames = map[string]string{
	"01": "Jan", "02": "Feb", "03": "Mar", "04": "Apr", "05": "May", "06": "Jun",
	"07": "Jul", "08": "Aug", "09": "Sep", "10": "Oct", "11": "Nov", "12": "Dec",
}

// FormatDateTime formats date and time into a more readable format
func FormatDateTime(input string) string {
	// Regex for D(...) and T12(...)/T24(...)
	re := regexp.MustCompile(`(D|T12|T24)\((\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2})([+-]\d{2}:\d{2}|Z)\)`)
	match := re.FindStringSubmatch(input)

	if match == nil {
		return input // If format is incorrect, return as is
	}

	formatType, year, month, day, hour, minute, offset := match[1], match[2], match[3], match[4], match[5], match[6], match[7]

	// If it's a date format (D), return formatted date
	if formatType == "D" {
		monthName := monthNames[month]
		return fmt.Sprintf("%s %s %s", day, monthName, year)
	}

	// Construct time string
	timeStr := fmt.Sprintf("%s-%s-%sT%s:%s%s", year, month, day, hour, minute, offset)

	// Parse into time.Time object
	parsedTime, err := time.Parse("2006-01-02T15:04-07:00", timeStr)
	if err != nil {
		return input
	}

	// Determine time format
	var formattedTime string
	if formatType == "T12" {
		formattedTime = parsedTime.Format("03:04PM")
	} else {
		formattedTime = parsedTime.Format("15:04")
	}

	// Handle offset formatting
	formattedOffset := formatOffset(offset)

	return fmt.Sprintf("%s %s", formattedTime, formattedOffset)
}

// formatOffset formats time zone offset
func formatOffset(offset string) string {
	if offset == "Z" {
		return "(+00:00)"
	}
	return fmt.Sprintf("(%s)", offset)
}
