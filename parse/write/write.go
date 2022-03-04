package write

import (
	"fmt"
	"log"
	"os"
)

func To(file, data string) {

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

