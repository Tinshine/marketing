package engine

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/util"
	"marketing/util/log"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

func Execute(ctx context.Context, c *app.RequestContext) {
	req := new(ExecuteReq)
	if err := c.Bind(req); err != nil {
		log.Error("Execute.Bind", err)
		util.Error(c, errors.WithMessage(errs.Bind, err.Error()))
		return
	}

	ev, _ := strconv.Atoi(c.Param("env"))
	resp, err := execute(ctx, req, consts.Env(ev))
	if err != nil {
		log.Error("Execute.Execute", err)
		util.Error(c, errors.WithMessage(errs.Internal, err.Error()))
		return
	}

	util.JSON(c, resp)
}
