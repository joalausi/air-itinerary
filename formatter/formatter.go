package formatter

import (
	"fmt"
	"itinerary/parser"
	"itinerary/utls"
)

// FormatItinerary форматирует список маршрутов с использованием airport lookup
func Format(flights []parser.Flight, lookup map[string]string) string {
	var result string

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

		// Форматируем время (может быть датой или временем)
		formattedTime := utls.FormatDate(flight.Time)
		if formattedTime == flight.Time { // Если это не дата, пробуем формат времени
			formattedTime = utls.FormatTime(flight.Time)
		}

		// Добавляем отформатированную строку
		result += fmt.Sprintf("%s: Flight from %s (%s) to %s (%s)\n",
			formattedTime, origin, flight.Origin, destination, flight.Destination)
	}

	return result
}
