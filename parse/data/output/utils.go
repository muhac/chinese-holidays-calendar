package output

import (
	"fmt"
	"main/parse/base"
	"time"
)

const (
	IcsHead  = "BEGIN:VCALENDAR\nVERSION:2.0\nPRODID:-//Rank Technology//Chinese Holidays//EN"
	IcsEvent = "BEGIN:VEVENT\nDTSTART:%s\nDTEND:%s\nSUMMARY:%s\nDESCRIPTION:%s\nEND:VEVENT"
	IcsTail  = "END:VCALENDAR"
)

// event data
type event struct {
	Id    string
	Group string
	Title string
	Date  time.Time
	Desc  string
}

func (d event) Ics() string {
	return fmt.Sprintf(
		IcsEvent,
		d.Date.Format("20060102T150405"),
		d.Date.Add(time.Hour*24).Format("20060102T150405"),
		d.Title,
		d.Desc,
	)
}

func getStatusName(status int) string {
	name := map[int]string{
		base.Rest: "假期",
		base.Work: "补班",
	}
	return name[status]
}

func getTitle(item base.Holiday) string {
	return fmt.Sprintf("%s%s", item.Name, getStatusName(item.Type))
}

func getDesc(item base.Holiday) string {
	return fmt.Sprintf("%s 第%d天/共%d天", getStatusName(item.Type), item.Nth, item.Total)
}