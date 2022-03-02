package main

import (
	"fmt"
	"log"
	"main/parse/read"
	"os"
)

func main() {
	data := read.InDir("./data/").Data()
	output := fmt.Sprintf("%+v", data)
	fmt.Println(output)
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

	fmt.Println(n, "done")
}
