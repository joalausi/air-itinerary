package utls

import (
	"fmt"
	"regexp"
	"time"
)

// monthNames - названия месяцев для удобного формата даты
var monthNames = map[string]string{
	"01": "Jan", "02": "Feb", "03": "Mar", "04": "Apr", "05": "May", "06": "Jun",
	"07": "Jul", "08": "Aug", "09": "Sep", "10": "Oct", "11": "Nov", "12": "Dec",
}

// FormatDate преобразует дату D(2007-04-05T12:30−02:00) → "05 апр. 2007"
func FormatDate(input string) string {
	// Регулярное выражение для даты формата D(2007-04-05T12:30−02:00)
	re := regexp.MustCompile(`D\((\d{4})-(\d{2})-(\d{2})T\d{2}:\d{2}([+-]\d{2}:\d{2}|Z)\)`)
	match := re.FindStringSubmatch(input)

	if match == nil {
		return input // Если не найдено, вернуть как есть
	}

	year, month, day := match[1], match[2], match[3]
	monthName := monthNames[month]
	return fmt.Sprintf("%s %s %s", day, monthName, year)
}

// FormatTime преобразует время
func FormatTime(input string) string {
	// Регулярное выражение для времени (12-часовой и 24-часовой формат)
	re := regexp.MustCompile(`T(12|24)\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2})([+-]\d{2}:\d{2}|Z)\)`)
	match := re.FindStringSubmatch(input)

	if match == nil {
		return input
	}

	timeFormat := match[1] // 12 или 24
	timestamp := match[2]  // 2007-04-05T12:30
	offset := match[3]     // -02:00 или Z

	// Парсим время в формат Go
	parsedTime, err := time.Parse("2006-01-02T15:04", timestamp)
	if err != nil {
		return input
	}

	// Определяем формат вывода
	if timeFormat == "12" {
		return fmt.Sprintf("%s %s", parsedTime.Format("03:04PM"), formatOffset(offset))
	} else {
		return fmt.Sprintf("%s %s", parsedTime.Format("15:04"), formatOffset(offset))
	}
}

// formatOffset преобразует смещение часового пояса
func formatOffset(offset string) string {
	if offset == "Z" {
		return "(+00:00)"
	}
	return fmt.Sprintf("(%s)", offset)
}
