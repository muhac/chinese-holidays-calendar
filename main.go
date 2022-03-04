package main

import (
	"fmt"
	"main/parse/read"
	"main/parse/write"
)

func main() {
	data := read.From("./data/").Read().Parse()
	for _, d := range data {
		fmt.Printf("%+v\n", d)
	}

	output := fmt.Sprintf("%+v", data)
	write.To("./docs/index.html", output)
	write.To("./docs/holiday.ics", output)
}
