package common

import (
	"strings"
	"testing"
)

func TestGetRelativePath(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "get relative path",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRelativePath()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRelativePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.HasSuffix(got, "marketing") {
				t.Errorf("GetRelativePath() = %v, HasSuffix %v", got, strings.HasSuffix(got, "marketing"))
			}
		})
	}
}
