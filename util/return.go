package util

import (
	"errors"
	"marketing/const/errs"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func JSON(ctx *app.RequestContext, obj interface{}) {
	ctx.JSON(http.StatusOK, utils.H{
		"code": 0,
		"msg":  "ok",
		"data": obj,
	})
}

func Error(ctx *app.RequestContext, err error) {
	if err == nil {
		panic("err is nil")
	}

	code := -1
	msg := err.Error()

	var e *errs.Error
	if errors.As(err, &e) {
		code = e.GetCode()
	}

	ctx.JSON(http.StatusInternalServerError, utils.H{
		"code": code, "msg": msg,
	})
}
