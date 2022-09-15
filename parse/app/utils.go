package app

import (
	"sort"

	"main/parse/core"
	"main/parse/data"
	"main/parse/data/input"
	"main/parse/data/output"
	"main/parse/data/read"
	"main/parse/data/write"
)

func newHandler(optional ...core.Holidays) Handler {
	if len(optional) == 0 {
		return handler{}
	}
	return handler{data: optional[0]}
}

type handler struct {
	data core.Holidays

	reader   data.Reader
	writer   data.Writer
	filename string

	input  data.Input
	output data.Output
}

func (h handler) Read(filename string) setDirIn {
	h.filename = filename
	return h
}

func (h handler) From(directory string) readData {
	h.reader = read.NewReader(directory, h.filename)
	return h
}

func (h handler) Parse() getData {
	h.input = h.reader.Read()
	h.data = input.NewParser().Parse(h.input)
	return h
}

func (h handler) Sort() getData {
	sort.Sort(h.data)
	return h
}

func (h handler) Write(filename string) setDirOut {
	h.filename = filename
	return h
}

func (h handler) To(directory string) setTitle {
	h.writer = write.NewWriter(directory, h.filename)
	return h
}

func (h handler) Title(name string) writeData {
	h.output = output.NewFormatter(name).Format(h.data)
	return h
}

func (h handler) Get() core.Holidays {
	return h.data
}

func (h handler) Set() {
	h.writer.Write(h.output)
}
