package system

import (
	"os"
	"testing"
)

func TestGetIntEnvVar(t *testing.T) {
	type args struct {
		envVar string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Env var is int",
			args: args{
				envVar: "VAR_1",
			},
			want:    23,
			wantErr: false,
		},
		// Env var is not int
		{
			name: "Env var is NOT int",
			args: args{
				envVar: "VAR_2",
			},
			want:    0,
			wantErr: true,
		},
	}
	os.Setenv("VAR_1", "23")
	os.Setenv("VAR_2", "SOME STRING")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIntEnvVar(tt.args.envVar)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIntEnvVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIntEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
