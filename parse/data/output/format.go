package output

import (
	"fmt"
	"main/parse/base"
	"main/parse/data"
)

func NewFormatter(format string) data.Formatter {
	return formatter{}
}

type formatter struct{}

func (f formatter) Format(info base.Holidays) data.Output {
	return data.Output(fmt.Sprintf("%+v", info))
}
