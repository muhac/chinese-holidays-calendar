package core

import (
	"log"
	"time"

	"github.com/samber/lo"
)

type Status string

const (
	Rest Status = "rest" // 假日
	Work Status = "work" // 补班
)

// Holidays data
type Holidays []Holiday

// Holiday data per day
type Holiday struct {
	Group string
	Date  time.Time
	Name  string
	Type  Status
	Nth   int
	Total int
}

func (h Holidays) Select(t Status) Holidays {
	return lo.Filter(h, func(d Holiday, _ int) bool { return d.Type == t })
}

func (h Holidays) Print(titles ...string) Holidays {
	lo.ForEach(titles, func(title string, _ int) { log.Println(title) })
	lo.ForEach(h, func(day Holiday, _ int) { log.Printf("%+v\n", day) })
	return h
}

func (h Holidays) Len() int           { return len(h) }
func (h Holidays) Less(i, j int) bool { return h[i].Date.Before(h[j].Date) }
func (h Holidays) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
