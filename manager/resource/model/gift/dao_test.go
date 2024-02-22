package gift

import (
	"marketing/util"
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
				util.UnsetUnitTestMode()
				return "prod env"
			},
			want: &rdsDAO{},
		}, {
			name: func() string {
				util.SetUnitTestMode()
				return "test env"
			},
			want: &mockDAO{
				gifts: map[uint]*Gift{},
			},
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
