package write

import (
	"fmt"
	"log"
	"main/parse/data"
	"os"
)

func NewWriter(filename string) data.Writer {
	return dataWriter{File: filename}
}

type dataWriter struct {
	File string
}

func (dw dataWriter) Write(data data.Output) {

	f, err := os.Create(dw.File)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	n, err := f.WriteString(string(data))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dw.File, n, "done")
}
