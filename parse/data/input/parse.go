package input

import (
	"fmt"
	"strings"

	"main/parse/base"
	"main/parse/data"
)

func NewParser() data.Parser {
	return parser{}
}

type parser struct{}

func (p parser) Parse(raw data.Input) (result base.Holidays) {
	for _, year := range raw {
		days, _ := parse(year)
		result = append(result, days...)
	}
	return
}

func parse(raw data.InputRaw) (result base.Holidays, err error) {
	dayCount := make(map[string]map[int]int)

	for group, holiday := range raw.Data {
		groupName := fmt.Sprintf("%04d%02d", raw.Year, group+1)
		dayCount[groupName] = make(map[int]int)
		info := strings.Split(holiday, ";")

		for i, day := range holidays(raw.Year, info[1]) {
			restDay := base.Holiday{
				Group: groupName,
				Name:  info[0],
				Nth:   i + 1,
				Date:  day,
				Type:  base.Rest,
			}
			result = append(result, restDay)
			dayCount[restDay.Group][restDay.Type]++
		}

		for i, day := range holidays(raw.Year, info[2]) {
			workDay := base.Holiday{
				Group: groupName,
				Name:  info[0],
				Nth:   i + 1,
				Date:  day,
				Type:  base.Work,
			}
			result = append(result, workDay)
			dayCount[workDay.Group][workDay.Type]++
		}
	}

	for i, holiday := range result {
		result[i].Total = dayCount[holiday.Group][holiday.Type]
	}

	return
}
