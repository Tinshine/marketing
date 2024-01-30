package task

import "context"

type T interface {
	GetId() uint
	Try(context.Context, *Params) (*Resp, error)
	Cancel(context.Context, *Params) (*Resp, error)
	Confirm(context.Context, *Params) (*Resp, error)
}

type Params struct {
}

type Resp struct {
}
