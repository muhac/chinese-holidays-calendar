package output

import (
	"fmt"
	"time"

	"main/parse/core"
)

const (
	icsHead  = "BEGIN:VCALENDAR\nVERSION:2.0\nPRODID:-//Rank Technology//Chinese Holidays//EN\nX-WR-CALNAME:%s"
	icsEvent = "BEGIN:VEVENT\nUID:%s\nDTSTART;VALUE=DATE:%s\nSUMMARY:%s\nDESCRIPTION:%s\nEND:VEVENT"
	icsTail  = "END:VCALENDAR"
)

// event data
type event struct {
	id    string
	group string
	title string
	date  time.Time
	desc  string
}

func (d event) Ics() string {
	return fmt.Sprintf(
		icsEvent,
		d.id,
		d.date.Format("20060102"),
		d.title,
		d.desc,
	)
}

func getStatusName(status core.Status) string {
	name := map[core.Status]string{
		core.Rest: "假期",
		core.Work: "补班",
	}
	return name[status]
}

func getTitle(item core.Holiday) string {
	return fmt.Sprintf("%s%s", item.Name, getStatusName(item.Type))
}

func getDesc(item core.Holiday) string {
	return fmt.Sprintf("%s 第%d天/共%d天", getStatusName(item.Type), item.Nth, item.Total)
}
