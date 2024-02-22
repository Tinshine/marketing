package log

// import (
// 	"marketing/consts/errs"
// 	"marketing/util/conf"
// 	"testing"
// )

// func TestInfo(t *testing.T) {
// 	conf.Init()
// 	Init()
// 	type args struct {
// 		event string
// 		kvs   []interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "test info",
// 			args: args{
// 				event: "Test.Info",
// 				kvs: []interface{}{
// 					"foo", map[string]interface{}{
// 						"foo": "bar",
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Info(tt.args.event, tt.args.kvs...)
// 		})
// 	}
// }

// func TestError(t *testing.T) {
// 	conf.Init()
// 	Init()
// 	type args struct {
// 		event string
// 		err   error
// 		kvs   []interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "test Error",
// 			args: args{
// 				event: "Test.Error",
// 				err:   errs.Internal,
// 				kvs: []interface{}{
// 					"foo", map[string]interface{}{
// 						"foo": "bar",
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Error(tt.args.event, tt.args.err, tt.args.kvs...)
// 		})
// 	}
// }

// func TestWarn(t *testing.T) {
// 	conf.Init()
// 	Init()
// 	type args struct {
// 		event string
// 		kvs   []interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "test Warn",
// 			args: args{
// 				event: "Test.Warn",
// 				kvs: []interface{}{
// 					"foo", map[string]interface{}{
// 						"foo": "bar",
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Warn(tt.args.event, tt.args.kvs...)
// 		})
// 	}
// }
