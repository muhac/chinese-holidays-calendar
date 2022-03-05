package base

import "fmt"

func (h Holidays) Print(titles ...string) (result string) {
	for _, title := range titles {
		fmt.Println(title)
	}

	for _, day := range h {
		result += fmt.Sprintf("%+v\n", day)
	}

	fmt.Println(result)
	return
}
