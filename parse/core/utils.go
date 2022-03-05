package core

import (
	"main/parse/base"
	"main/parse/data"
	"main/parse/data/input"
	"main/parse/data/output"
	"main/parse/data/read"
	"main/parse/data/write"
	"sort"
)

func newHandler(optional ...base.Holidays) Handler {
	if len(optional) == 0 {
		return handler{}
	}
	return handler{data: optional[0]}
}

type handler struct {
	data base.Holidays

	reader data.Reader
	writer data.Writer

	input  data.Input
	output data.Output
}

func (h handler) ReadFrom(directory string) Parser {
	h.reader = read.NewReader(directory)
	h.input = h.reader.Read()
	return h
}

func (h handler) Parse() Sorter {
	h.data = input.NewParser().Parse(h.input)
	return h
}

func (h handler) Sort() Getter {
	sort.Sort(h.data)
	return h
}

func (h handler) WriteTo(file string) Formatter {
	h.writer = write.NewWriter(file)
	return h
}

func (h handler) Format(format string) Setter {
	h.output = output.NewFormatter(format).Format(h.data)
	return h
}

func (h handler) Get() base.Holidays {
	return h.data
}

func (h handler) Set() {
	h.writer.Write(h.output)
}
