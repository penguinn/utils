package ptime_test

import (
	"github.com/penguinn/utils/ptime"
	"testing"
)

func TestFormatTimeToDay(t *testing.T) {
	type args struct {
		timeStamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				timeStamp: 1514390401,
			},
			want: "2017-12-28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ptime.FormatTimeToDay(tt.args.timeStamp); got != tt.want {
				t.Errorf("FormatTimeToDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeToSecond(t *testing.T) {
	type args struct {
		timeStamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				timeStamp: 1514390401,
			},
			want: "2017-12-28 00:00:01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ptime.FormatTimeToSecond(tt.args.timeStamp); got != tt.want {
				t.Errorf("FormatTimeToSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDayToTime(t *testing.T) {
	type args struct {
		timeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				timeStr: "2017-12-28",
			},
			want:    1514390400,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ptime.ParseDayToTime(tt.args.timeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDayToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDayToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSecondToTime(t *testing.T) {
	type args struct {
		timeStr string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				timeStr: "2017-12-28 00:00:01",
			},
			want:    1514390401,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ptime.ParseSecondToTime(tt.args.timeStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSecondToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseSecondToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
