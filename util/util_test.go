package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeKey(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty slice",
			args: args{
				items: []interface{}{},
			},
			want: "",
		},
		{
			name: "one element slice",
			args: args{
				items: []interface{}{"test"},
			},
			want: "test",
		},
		{
			name: "multiple elements slice",
			args: args{
				items: []interface{}{"test", 1, true},
			},
			want: "test_1_true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MakeKey(tt.args.items...))
		})
	}
}
