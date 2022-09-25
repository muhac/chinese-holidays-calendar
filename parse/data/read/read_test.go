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
			gotResult, err := year(tt.args.filename, `^\d{4}\.txt$`)
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
		{"1", args{"// none"}, []string{}},
		{"2", args{";1.1;2.2"}, []string{}},
		{"3", args{"3;1.1;2.2"}, []string{"3;1.1;2.2"}},
		{"4", args{"4;1.1;"}, []string{"4;1.1;"}},
		{"5", args{"5;1.1;2.2,3.3"}, []string{"5;1.1;2.2,3.3"}},
		{"6", args{"6;1.1,2.2;3.3,4.4"}, []string{"6;1.1,2.2;3.3,4.4"}},
		{"7", args{"7;1.1,2.2;3.3,4.4-5.5"}, []string{"7;1.1,2.2;3.3,4.4-5.5"}},
		{"8", args{"8;1.1;2.2;"}, []string{}},
		{"9", args{"9;,1.1;2.2"}, []string{}},
		{"10", args{"10;1.1"}, []string{}},
		{"11", args{"11;1.1;2.2,"}, []string{}},
		{"12", args{"// 13;1.1;2.2     "}, []string{}},
		{"13", args{"13;1.1;2.2 // none"}, []string{"13;1.1;2.2"}},
		{"14", args{"14;1.1;2.2        "}, []string{"14;1.1;2.2"}},
		{"15", args{"        15;1.1;2.2"}, []string{"15;1.1;2.2"}},
		{"16", args{"   16;1.1;2.2     "}, []string{"16;1.1;2.2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := lines(tt.args.data); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("lines() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
