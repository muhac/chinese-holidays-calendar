package read

import "testing"

func Test_year(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantOk     bool
	}{
		{"2018", args{"2018.txt"}, "2018", true},
		{"2019", args{"2019.avi"}, "", false},
		{"2020", args{"zero.txt"}, "", false},
		{"2021", args{"2021.txt"}, "2021", true},
		{"2022", args{"2022.txt"}, "2022", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotOk := year(tt.args.filename)
			if gotResult != tt.wantResult {
				t.Errorf("year() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotOk != tt.wantOk {
				t.Errorf("year() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
