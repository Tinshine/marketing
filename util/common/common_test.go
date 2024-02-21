package common

import "testing"

func TestGetRelativePath(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "get relative path",
			want:    "c:\\Users\\acer\\go\\src\\marketing",
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
			if got != tt.want {
				t.Errorf("GetRelativePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
