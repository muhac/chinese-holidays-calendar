package input

import (
	"fmt"
	"strings"
	"time"
)

func date(year int, date string) (result time.Time) {
	input := fmt.Sprintf("%04d-%s", year, date)
	result, _ = time.Parse("2006-1.2", input)
	return
}

func holidays(year int, daysRaw string) (result []string) {
	if daysRaw == "" {
		return
	}

	days := strings.Split(daysRaw, ",")
	for _, day := range days {
		if strings.Contains(day, "-") {
			period := strings.Split(day, "-")
			for d := date(year, period[0]); !d.After(date(year, period[1])); d = d.AddDate(0, 0, 1) {
				result = append(result, d.Format("1.2"))
			}
		} else {
			result = append(result, day)
		}
	}
	return result
}
