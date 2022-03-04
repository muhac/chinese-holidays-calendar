package core

import (
	"strings"
)

func (data Raw) Parse() (result Holidays) {
	for _, year := range data {
		days, _ := year.Parse()
		result = append(result, days...)
	}
	return
}

func (data RawInfo) Parse() (result Holidays, err error) {
	for _, holiday := range strings.Split(data.Data, "\n") {
		if holiday[:2] == "//" {
			continue
		}

		info := strings.Split(holiday, ";")

		for i, day := range holidays(data.Year, info[1]) {
			restDay := Holiday{
				Name: info[0],
				Nth:  i + 1,
				Date: date(data.Year, day),
				Type: Rest,
			}
			result = append(result, restDay)
		}

		for i, day := range holidays(data.Year, info[2]) {
			workDay := Holiday{
				Name: info[0],
				Nth:  i + 1,
				Date: date(data.Year, day),
				Type: Work,
			}
			result = append(result, workDay)
		}
	}

	return
}
