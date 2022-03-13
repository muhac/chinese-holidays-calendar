package base

import "time"

// Holidays data
type Holidays []Holiday

// Holiday data per day
type Holiday struct {
	Group string
	Date  time.Time
	Name  string
	Type  int
	Nth   int
	Total int
}

func (h Holidays) Where(filter func(Holiday) bool) (result Holidays) {
	for _, item := range h {
		if filter(item) {
			result = append(result, item)
		}
	}
	return
}

func (h Holidays) Len() int           { return len(h) }
func (h Holidays) Less(i, j int) bool { return h[i].Date.Before(h[j].Date) }
func (h Holidays) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
