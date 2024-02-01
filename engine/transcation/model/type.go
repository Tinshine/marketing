package model

import (
	"context"
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
	User user.U
}

type Resp struct {
}
