package write

import (
	"fmt"
	"log"
	"os"
	"strings"

	"main/parse/data"
)

func NewWriter(dir, file string) data.Writer {
	return dataWriter{File: "./" + dir + "/" + file}
}

type dataWriter struct {
	File string
}

func (dw dataWriter) Write(data data.Output) {
	output := strings.Join(
		[]string{
			data.Prefix,
			strings.Join(data.Body, "\n\n"),
			data.Suffix,
		},
		"\n\n\n",
	)

	f, err := os.Create(dw.File)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = f.Close() }()

	n, err := f.WriteString(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("write", n, "bytes to", dw.File)
}
