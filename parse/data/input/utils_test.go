package input

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
		{"6", args{2006, "0.1"}, time.Date(2005, 12, 31, 0, 0, 0, 0, time.UTC)},
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
		{"1", args{1, "1.1"}, []string{"01.1.1"}},
		{"2", args{1, "1.1,2.2"}, []string{"01.1.1", "01.2.2"}},
		{"3", args{1, "1.1-1.3"}, []string{"01.1.1", "01.1.2", "01.1.3"}},
		{"4", args{1, "1.1-1.3,2.2,3.3-3.4"}, []string{"01.1.1", "01.1.2", "01.1.3", "01.2.2", "01.3.3", "01.3.4"}},
		{"5", args{1, "1.31-2.2"}, []string{"01.1.31", "01.2.1", "01.2.2"}},
		{"6", args{1, "0.2"}, []string{"00.12.30"}},
		{"7", args{1, "0.4-0.1"}, []string{"00.12.28", "00.12.29", "00.12.30", "00.12.31"}},
		{"8", args{1, "0.2-1.2"}, []string{"00.12.30", "00.12.31", "01.1.1", "01.1.2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := holidays(tt.args.year, tt.args.daysRaw)
			for idx, result := range gotResult{
				if !reflect.DeepEqual(result.Format("06.1.2"), tt.wantResult[idx]) {
					t.Errorf("holidays() = %v, want %v", result, tt.wantResult[idx])
				}
			}
		})
	}
}
