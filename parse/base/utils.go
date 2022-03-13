package base

import "fmt"

func selectType(t int) func(Holiday) bool {
	return func(holidays Holiday) bool {
		return holidays.Type == t
	}
}

func (h Holidays) Select(t int) Holidays {
	return h.Where(selectType(t))
}

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
