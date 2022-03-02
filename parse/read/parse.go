package read

import (
	"fmt"
	"main/parse/base"
	"regexp"
)

func parse(name, data string) (result base.Holiday, err error) {
	fmt.Println(name)
	fmt.Println(data)
	return
}

func year(filename string) (result string, ok bool) {
	regex := regexp.MustCompile(`^\d{4}\.txt$`)
	if !regex.MatchString(filename) {
		return
	}
	return filename[:4], true
}
