package gift

import (
	"os"
	"reflect"
	"testing"
)

func TestInitDAO(t *testing.T) {
	tests := []struct {
		name func() string
		want DAO
	}{
		{
			name: func() string {
				os.Setenv("unit_test", "0")
				return "prod env"
			},
			want: &rdsDAO{},
		}, {
			name: func() string {
				os.Setenv("unit_test", "1")
				return "test env"
			},
			want: &mockDAO{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name(), func(t *testing.T) {
			if got := InitDAO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDAO() = %v, want %v", got, tt.want)
			}
		})
	}
}
