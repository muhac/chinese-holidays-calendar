package main

import (
	"main/parse/base"
	"main/parse/core"
)

func main() {
	holidays := core.Data().ReadFrom(base.SourceDir).Parse().Sort().Get()

	holidays.Print("==== HOLIDAYS ====")

	core.Data(holidays).WriteTo(base.IndexPage).Format(base.ICS).Set()
	core.Data(holidays).WriteTo(base.HolidayICS).Format(base.ICS).Set()
}
