package data

import "main/parse/core"

type Reader interface {
	Read() Input
}

type Parser interface {
	Parse(Input) core.Holidays
}

type Formatter interface {
	Format(core.Holidays) Output
}

type Writer interface {
	Write(Output)
}

// Input data
type Input []InputRaw

// InputRaw per year
type InputRaw struct {
	Year int
	Data []string
}

type Output struct {
	Prefix string
	Body   []string
	Suffix string
}
