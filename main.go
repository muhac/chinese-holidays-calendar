package main

import (
	"fmt"
	"log"
	"main/parse/read"
	"os"
)

func main() {
	data := read.From("./data/").Read().Parse()
	for _, d := range data {
		fmt.Println(fmt.Sprintf("%+v", d))
	}

	output := fmt.Sprintf("%+v", data)
	write("./docs/index.html", output)
	write("./docs/holiday.ics", output)
}

func write(file, data string) {

	f, err := os.Create(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	n, err := f.WriteString(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file, n, "done")
}
