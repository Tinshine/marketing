package util

import (
	"marketing/consts/errs"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

func TestJSON(t *testing.T) {
	type args struct {
		ctx *app.RequestContext
		obj interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantResp string
	}{
		{
			name: "test nil JSON",
			args: args{
				ctx: &app.RequestContext{},
				obj: nil,
			},
			wantCode: 200,
			wantResp: `{"code":0,"data":null,"msg":"ok"}`,
		},
		{
			name: "test empty JSON",
			args: args{
				ctx: &app.RequestContext{},
				obj: "",
			},
			wantCode: 200,
			wantResp: `{"code":0,"data":"","msg":"ok"}`,
		},
		{
			name: "test interface{} JSON",
			args: args{
				ctx: &app.RequestContext{},
				obj: map[string]interface{}{
					"name": "test",
				},
			},
			wantCode: 200,
			wantResp: `{"code":0,"data":{"name":"test"},"msg":"ok"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JSON(tt.args.ctx, tt.args.obj)
			resp := tt.args.ctx.GetResponse()
			if resp.StatusCode() != tt.wantCode {
				t.Errorf("got %d, want %d", resp.StatusCode(), tt.wantCode)
			}
			if string(resp.Body()) != tt.wantResp {
				t.Errorf("got %s, want %s", resp.Body(), tt.wantResp)
			}
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		ctx *app.RequestContext
		err error
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantResp string
	}{
		{
			name: "test error",
			args: args{
				ctx: &app.RequestContext{},
				err: errs.Internal,
			},
			wantCode: 500,
			wantResp: `{"code":-2,"msg":"Bind error"}`,
		},
		{
			name: "test wrap error",
			args: args{
				ctx: &app.RequestContext{},
				err: errors.WithMessage(errs.Bind, "wrap info: "),
			},
			wantCode: 500,
			wantResp: `{"code":-2,"msg":"Bind error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.ctx, tt.args.err)
			resp := tt.args.ctx.GetResponse()
			if resp.StatusCode() != tt.wantCode {
				t.Errorf("got %d, want %d", resp.StatusCode(), tt.wantCode)
			}
			if string(resp.Body()) != tt.wantResp {
				t.Errorf("got %s, want %s", resp.Body(), tt.wantResp)
			}
		})
	}
}
