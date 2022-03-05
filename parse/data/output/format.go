package output

import (
	"github.com/google/uuid"
	"main/parse/base"
	"main/parse/data"
)

func NewFormatter(format string) data.Formatter {
	return formatter{}
}

type formatter struct {
	result string
}

func (f formatter) Format(info base.Holidays) (result data.Output) {
	result.Prefix = IcsHead
	result.Suffix = IcsTail

	for _, day := range info {
		outputDay := event{
			Id:    uuid.NewString(),
			Group: day.Group,
			Title: getTitle(day),
			Date:  day.Date,
			Desc:  getDesc(day),
		}
		result.Body = append(result.Body, outputDay.Ics())
	}
	return
}
