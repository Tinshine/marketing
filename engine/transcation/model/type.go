package model

import (
	"context"
	"marketing/consts"
	"marketing/user"
)

type T interface {
	SetTrId(uint)
	GetTrId() uint
	Try(context.Context, *Params) (*Resp, error)
	Cancel(context.Context, *Params) (*Resp, error)
	Confirm(context.Context, *Params) (*Resp, error)
}

type Params struct {
	Reward map[string]interface{}
	User   user.U
	Ev     consts.Env
}

type Resp struct {
}
