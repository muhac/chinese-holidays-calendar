package input

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func date(year int, date string) (result time.Time) {
	input := fmt.Sprintf("%04d-%s", year, date)
	result, _ = time.Parse("2006-1.2", input)

	if date[0] == '0' { // => 0001-1.1
		delta, _ := strconv.Atoi(date[2:]) // days before
		result = result.AddDate(year-1, 0, -delta)
	}
	return
}

func holidays(year int, days string) (result []time.Time) {
	if days == "" {
		return
	}

	for _, day := range strings.Split(days, ",") {
		if strings.Contains(day, "-") {
			period := strings.Split(day, "-")
			d := date(year, period[0])
			for !d.After(date(year, period[1])) {
				result = append(result, d)
				d = d.AddDate(0, 0, 1)
			}
		} else {
			result = append(result, date(year, day))
		}
	}
	return result
}
