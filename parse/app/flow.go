package app

import "main/parse/core"

func Data(optional ...core.Holidays) Handler {
	return newHandler(optional...)
}

type Handler interface {
	Read(filename string) setDirIn
	Write(filename string) setDirOut
}

type setDirIn interface {
	From(directory string) readData
}

type readData interface {
	Parse() getData
}

type getData interface {
	Sort() getData
	Get() core.Holidays
}

type setDirOut interface {
	To(directory string) setTitle
}

type setTitle interface {
	Title(name string) writeData
}

type writeData interface {
	Set()
}
