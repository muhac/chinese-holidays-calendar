package base

import (
	"fmt"
	"strings"
	"time"
)

func Date(year int, date string) (result time.Time) {
	input := fmt.Sprintf("%d-%s", year, date)
	result, _ = time.Parse(TimeLayout, input)
	return
}

func DayRange(daysRaw string) (result []string) {
	if daysRaw == "" {
		return
	}

	days := strings.Split(daysRaw, ",")
	for _, day := range days {
		if strings.Contains(day, "-") {
			r := strings.Split(day, "-")
			start, _ := time.Parse("1.2", fmt.Sprintf("%s", r[0]))
			end, _ := time.Parse("1.2", fmt.Sprintf("%s", r[1]))
			for d := start; d.Before(end.Add(time.Hour)); d = d.AddDate(0, 0, 1) {
				result = append(result, d.Format("1.2"))
			}
		} else {
			result = append(result, day)
		}
	}
	return result
}
