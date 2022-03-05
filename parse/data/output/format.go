package output

import (
	"main/parse/base"
	"main/parse/data"
)

func NewFormatter(format string) data.Formatter {
	return formatterICS{}
}

type formatterICS struct{}

func (f formatterICS) Format(info base.Holidays) (result data.Output) {
	result.Prefix = IcsHead
	result.Suffix = IcsTail

	for _, day := range info {
		outputDay := event{
			Group: day.Group,
			Title: getTitle(day),
			Date:  day.Date,
			Desc:  getDesc(day),
		}
		result.Body = append(result.Body, outputDay.Ics())
	}
	return
}
