package engine

import (
	"marketing/user"
)

type ExecuteReq struct {
	AppId   uint         `json:"app_id"`
	UserReq user.InfoReq `json:"user"`
	RuleId  string       `json:"rule_id"`
	ActId   string       `json:"act_id"`
}

func (r *ExecuteReq) Validate() error {
	// todo...
	return nil
}

type ExecuteResp struct{}
