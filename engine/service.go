package engine

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/engine/rule"
	"marketing/manager/rule/model"
	"marketing/user"

	"github.com/pkg/errors"
)

func execute(ctx context.Context, req *ExecuteReq, ev consts.Env) (*ExecuteResp, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate")
	}

	user, err := user.GetInfo(ctx, req.UserReq)
	if err != nil {
		return nil, errors.WithMessage(err, "get user info")
	}

	cf, err := getRule(ctx, req, ev)
	if err != nil {
		return nil, errors.WithMessage(err, "get rule")
	}

	res, err := makeRule(cf, user).Execute(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "rule execute")
	}

	return makeExecuteResp(res), nil
}

func getRule(ctx context.Context, req *ExecuteReq, ev consts.Env) (*model.Rule, error) {
	cf, err := model.FindByRuleId(rds.DB(ctx, ev), req.RuleId)
	if err != nil {
		return nil, errors.WithMessage(err, "get rule")
	}
	if cf.ActId != req.ActId || cf.AppId != req.AppId {
		return nil, errs.InvalidParams
	}
	return cf, nil
}

func makeExecuteResp(res *rule.Resp) *ExecuteResp {
	// todo...
	return nil
}

func makeRule(cf *model.Rule, user *user.U) *rule.R {
	// todo...
	return nil
}
