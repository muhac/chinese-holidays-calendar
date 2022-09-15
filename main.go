package main

import (
	"main/parse/app"
	"main/parse/core"
)

func main() {
	holidays := app.Data().Read(`^20\d\d`).From("data").Parse().Sort().Get().Print("==== HOLIDAYS ====")

	app.Data(holidays).Write("index.html").To("docs").Title("节假日").Set()
	app.Data(holidays).Write("holiday.ics").To("docs").Title("节假日").Set()

	app.Data(holidays.Select(core.Rest)).Write("rest.ics").To("docs").Title("节假日（假期）").Set()
	app.Data(holidays.Select(core.Work)).Write("work.ics").To("docs").Title("节假日（补班）").Set()
}
