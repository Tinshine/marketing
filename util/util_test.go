package util

import (
	"context"
	"marketing/consts"
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

func TestGetUsername(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "nil context",
			args:    args{ctx: nil},
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty context",
			args:    args{ctx: context.Background()},
			want:    "",
			wantErr: true,
		},
		{
			name:    "valid context",
			args:    args{ctx: context.WithValue(context.Background(), consts.CtxUsernameKey, "test")},
			want:    "test",
			wantErr: false,
		},
		{
			name:    "invalid context",
			args:    args{ctx: context.WithValue(context.Background(), consts.CtxUsernameKey, nil)},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsername(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}
