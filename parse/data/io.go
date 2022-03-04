package data

import "main/parse/base"

type Reader interface {
	Read() Input
}

type Writer interface {
	Write(Output)
}

// Input data
type Input []InputRaw

// InputRaw per year
type InputRaw struct {
	Year int
	Data string
}

type Output string

type Parser interface {
	Parse(Input) base.Holidays
}

type Formatter interface {
	Format(base.Holidays) Output
}
