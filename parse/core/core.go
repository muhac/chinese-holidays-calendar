package core

import (
	"main/parse/base"
)

func Data(optional ...base.Holidays) Handler {
	return newHandler(optional...)
}

type Handler interface {
	ReadFrom(directory string) Parser
	WriteTo(file string) Formatter
}

type Parser interface {
	Parse() Sorter
}

type Sorter interface {
	Sort() Getter
}

type Formatter interface {
	Format(format string) Setter
}

type Getter interface {
	Get() base.Holidays
}

type Setter interface {
	Set()
}
