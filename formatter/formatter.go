package formatter

import (
	"fmt"
	"itinerary/parser"
	"itinerary/utls"
	"strings"
)

// Format принимает список маршрутов и карту кодов аэропортов и возвращает отформатированную строку
func Format(flights []parser.Flight, lookup map[string]string) string {
	var result strings.Builder

	// Проходим по всем маршрутам
	for _, flight := range flights {
		// Получаем названия аэропортов
		origin := lookup[flight.Origin]
		if origin == "" {
			origin = flight.Origin // Если код не найден, оставляем как есть
		}
		destination := lookup[flight.Destination]
		if destination == "" {
			destination = flight.Destination
		}

		// Форматируем дату и время
		formattedDateTime := utls.FormatDateTime(flight.DateTime)

		// Добавляем отформатированную строку в буфер
		result.WriteString(fmt.Sprintf("%s: Flight from %s (%s) to %s (%s)\n",
			formattedDateTime, origin, flight.Origin, destination, flight.Destination))
	}

	return result.String()
}
