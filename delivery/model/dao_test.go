package model

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
				return "init rds dao"
			},
			want: &rdsDAO{},
		}, {
			name: func() string {
				util.SetUnitTestMode()
				return "init mock dao"
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
