package date

import (
	"testing"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func Test_getExpirationDate(t *testing.T) {
	type args struct {
		date                 string
		turnaroundTimeNumber int
	}
	tests := []struct {
		name string
		args args
		want *timestamp.Timestamp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt // Create a new variable tt to avoid the closure issue
		t.Run(tt.name, func(t *testing.T) {
			got := getExpirationDate(tt.args.date, tt.args.turnaroundTimeNumber).Nanos
			if got != tt.want.Nanos {
				t.Errorf("getExpirationDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
