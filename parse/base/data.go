package base

type Holiday struct {
	Year string
	Days []Day
}

type Day struct {
	Name string
	Date string
	Type int
}
