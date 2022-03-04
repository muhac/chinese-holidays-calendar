package input

import (
	"main/parse/base"
	"main/parse/data"
	"strings"
)

func NewParser() data.Parser {
	return parser{}
}

type parser struct {}

func (p parser) Parse(raw data.Input) (result base.Holidays) {
	for _, year := range raw {
		days, _ := parse(year)
		result = append(result, days...)
	}
	return
}

func parse(raw data.InputRaw) (result base.Holidays, err error) {
	for _, holiday := range strings.Split(raw.Data, "\n") {
		if holiday[:2] == "//" {
			continue
		}

		info := strings.Split(holiday, ";")

		for i, day := range holidays(raw.Year, info[1]) {
			restDay := base.Holiday{
				Name: info[0],
				Nth:  i + 1,
				Date: date(raw.Year, day),
				Type: base.Rest,
			}
			result = append(result, restDay)
		}

		for i, day := range holidays(raw.Year, info[2]) {
			workDay := base.Holiday{
				Name: info[0],
				Nth:  i + 1,
				Date: date(raw.Year, day),
				Type: base.Work,
			}
			result = append(result, workDay)
		}
	}

	return
}
