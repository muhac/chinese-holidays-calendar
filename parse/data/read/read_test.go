package read

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_year(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    error
	}{
		{"2018", args{"2018.txt"}, 2018, nil},
		{"2019", args{"2019.avi"}, 0, fmt.Errorf("%s", "invalid year")},
		{"2020", args{"zero.txt"}, 0, fmt.Errorf("%s", "invalid year")},
		{"2021", args{"2021.txt"}, 2021, nil},
		{"2022", args{"2022.txt"}, 2022, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := year(tt.args.filename)
			if err == nil && err != tt.wantErr || err != nil && tt.wantErr == nil {
				t.Errorf("year() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("year() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_lines(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{"1", args{"// none"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := lines(tt.args.data); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("lines() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
