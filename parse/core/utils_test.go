package core

import (
	"reflect"
	"testing"
	"time"
)

func Test_date(t *testing.T) {
	type args struct {
		year int
		date string
	}
	tests := []struct {
		name       string
		args       args
		wantResult time.Time
	}{
		{"1", args{2001, "1.1"}, time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2", args{2002, "1.11"}, time.Date(2002, 1, 11, 0, 0, 0, 0, time.UTC)},
		{"3", args{2003, "11.1"}, time.Date(2003, 11, 1, 0, 0, 0, 0, time.UTC)},
		{"4", args{2004, "11.11"}, time.Date(2004, 11, 11, 0, 0, 0, 0, time.UTC)},
		{"5", args{2005, "1.41"}, time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := date(tt.args.year, tt.args.date); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("date() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_holidays(t *testing.T) {
	type args struct {
		year    int
		daysRaw string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{"1", args{1, "1.1"}, []string{"1.1"}},
		{"2", args{1, "1.1,2.2"}, []string{"1.1", "2.2"}},
		{"3", args{1, "1.1-1.3"}, []string{"1.1", "1.2", "1.3"}},
		{"4", args{1, "1.1-1.3,2.2,3.3-3.4"}, []string{"1.1", "1.2", "1.3", "2.2", "3.3", "3.4"}},
		{"5", args{1, "1.31-2.2"}, []string{"1.31", "2.1", "2.2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := holidays(tt.args.year, tt.args.daysRaw); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("holidays() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
