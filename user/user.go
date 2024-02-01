package user

import (
	"context"
	"marketing/consts"
)

type InfoReq struct {
	LoginType consts.LoginType `json:"login_type"`
	Token     string           `json:"token"`
}

func GetInfo(ctx context.Context, rq InfoReq) (*U, error) {
	// todo...
	return nil, nil
}

type U interface {
	GetId() string
}
