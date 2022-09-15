package input

import (
	"fmt"
	"strings"

	"main/parse/core"
	"main/parse/data"
)

func NewParser() data.Parser {
	return parser{}
}

type parser struct{}

func (p parser) Parse(raw data.Input) (result core.Holidays) {
	for _, year := range raw {
		days, _ := parse(year)
		result = append(result, days...)
	}
	return
}

func parse(raw data.InputRaw) (result core.Holidays, err error) {
	dayCount := make(map[string]map[core.Status]int)

	for group, holiday := range raw.Data {
		groupName := fmt.Sprintf("%04d%02d", raw.Year, group+1)
		dayCount[groupName] = make(map[core.Status]int)
		info := strings.Split(holiday, ";")

		for i, day := range holidays(raw.Year, info[1]) {
			restDay := core.Holiday{
				Group: groupName,
				Name:  info[0],
				Nth:   i + 1,
				Date:  day,
				Type:  core.Rest,
			}
			result = append(result, restDay)
			dayCount[restDay.Group][restDay.Type]++
		}

		for i, day := range holidays(raw.Year, info[2]) {
			workDay := core.Holiday{
				Group: groupName,
				Name:  info[0],
				Nth:   i + 1,
				Date:  day,
				Type:  core.Work,
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
