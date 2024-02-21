package model

import (
	"context"
	"marketing/consts"
	"marketing/user"
)

type T interface {
	GetTxId() string
	Try(context.Context, *Params) (*Resp, error)
	Cancel(context.Context, *Params) (*Resp, error)
	Confirm(context.Context, *Params) (*Resp, error)
}

type Params struct {
	Input map[string]interface{}
	User  user.U
	Ev    consts.Env
	AppId uint
}

type Resp struct {
}
