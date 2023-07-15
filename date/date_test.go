package date

import (
	"reflect"
	"testing"
	"time"
)

func Test_GetExpirationDate(t *testing.T) {
	type args struct {
		date                 string
		turnaroundTimeNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Valid Business Day, within working hours", args{"2023-07-13T09:00:00.000Z", 48}, "2023-07-17T09:00:00.000Z", false},
		{"Weekend to Next Business Day", args{"2023-07-15T14:00:00.000Z", 24}, "2023-07-17T14:00:00.000Z", false},
		{"Before Working Hours to Start of Working Hours", args{"2023-07-14T07:30:00.000Z", 12}, "2023-07-17T09:00:00.000Z", false},
		{"After Working Hours to Start of Next Working Day", args{"2023-07-14T19:30:00.000Z", 8}, "2023-07-17T09:00:00.000Z", false},
		{"Valid Business Day, within working hours, large turnaround time", args{"2023-07-14T09:00:00.000Z", 168}, "2023-07-21T09:00:00.000Z", false},
		{"Invalid Timestamp", args{"invalid_timestamp", 24}, "", true},
	}
	for _, tt := range tests {
		tt := tt // Create a new variable tt to avoid the closure issue
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetExpirationDate(tt.args.date, tt.args.turnaroundTimeNumber)
			if err != nil && tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("getExpirationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateExpirationDate(t *testing.T) {
	type args struct {
		t                    time.Time
		turnaroundTimeNumber int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Normal Business Day, within working hours",
			args: args{
				t:                    time.Date(2023, 7, 13, 10, 0, 0, 0, time.UTC),
				turnaroundTimeNumber: 72,
			},
			want: time.Date(2023, 7, 17, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "Weekend to Next Business Day",
			args: args{
				t:                    time.Date(2023, 7, 15, 14, 0, 0, 0, time.UTC),
				turnaroundTimeNumber: 24,
			},
			want: time.Date(2023, 7, 17, 14, 0, 0, 0, time.UTC),
		},
		{
			name: "Before Working Hours to Start of Working Hours",
			args: args{
				t:                    time.Date(2023, 7, 14, 7, 0, 0, 0, time.UTC),
				turnaroundTimeNumber: 12,
			},
			want: time.Date(2023, 7, 17, 9, 0, 0, 0, time.UTC),
		},
		{
			name: "After Working Hours to Start of Next Working Day",
			args: args{
				t:                    time.Date(2023, 7, 14, 19, 0, 0, 0, time.UTC),
				turnaroundTimeNumber: 8,
			},
			want: time.Date(2023, 7, 17, 9, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateExpirationDate(tt.args.t, tt.args.turnaroundTimeNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateExpirationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTimestamp(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "Valid Timestamp",
			args: args{
				date: "2023-07-13T23:15:26.371Z",
			},
			want:    time.Date(2023, 7, 13, 23, 15, 26, 371000000, time.UTC),
			wantErr: false,
		},
		{
			name: "Invalid Timestamp",
			args: args{
				date: "invalid_timestamp",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimestamp(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
