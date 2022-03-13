package main

import (
	"main/parse/base"
	"main/parse/core"
)

func main() {
	holidays := core.Data().Read(`^20\d\d`).From("data").Parse().Sort().Get()

	holidays.Print("==== HOLIDAYS ====")

	core.Data(holidays).Write("index.html").To("docs").Title("节假日").Set()
	core.Data(holidays).Write("holiday.ics").To("docs").Title("节假日").Set()

	core.Data(holidays.Select(base.Rest)).Write("rest.ics").To("docs").Title("节假日（假期）").Set()
	core.Data(holidays.Select(base.Work)).Write("work.ics").To("docs").Title("节假日（补班）").Set()
}
