package unilog

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		logPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "logger creation with correct file path",
			args:    args{logPath: "./log.txt"},
			wantErr: false,
		},
		{
			name:    "logger creation with incorrect file path",
			args:    args{logPath: "/home/tmp/floder/log.txt"},
			wantErr: true,
		},
		{
			name:    "logger creation without file path",
			args:    args{logPath: ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLogger(tt.args.logPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got.Error("test error")
			}
		})
	}
}
