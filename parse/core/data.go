package core

import "time"

// Holidays data
type Holidays []Holiday

// Holiday data per day
type Holiday struct {
	Date time.Time
	Name string
	Type int
	Nth	int
}

// Raw data
type Raw []RawInfo

// RawInfo per year
type RawInfo struct {
	Year    int
	Data string
}
