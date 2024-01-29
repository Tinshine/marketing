package middleware

import (
	"context"
	"marketing/consts/auth"
	"marketing/consts/errs"
	"marketing/util"
	"marketing/util/log"

	authModel "marketing/manager/auth/model/auth"
	resModel "marketing/manager/auth/model/resource"
	authService "marketing/manager/auth/service/auth"
	resService "marketing/manager/auth/service/resource"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthCheck(resType auth.ResType, authType auth.AuthType) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		username, err := util.GetUsername(ctx)
		if err != nil {
			log.Error("AuthCheck.GetUsername", errs.EmptyUsername)
			util.Error(c, errs.EmptyUsername)
			c.Abort()
			return
		}

		resp, err := resService.Query(ctx, &resModel.QueryReq{
			ResType:   &resType,
			AuthTypes: &[]auth.AuthType{authType, auth.Admin},
		})
		if err != nil {
			log.Error("AuthCheck.Query", err)
			util.Error(c, err)
			c.Abort()
			return
		}
		if len(resp.Data) <= 0 {
			log.Error("AuthCheck.Query", errs.NoAuthRes)
			util.Error(c, errs.NoAuthRes)
			c.Abort()
			return
		}

		reason := ""
		for _, d := range resp.Data {
			resp, err := authService.Check(ctx, &authModel.CheckReq{
				AppId:     d.AppId,
				Username:  username,
				AuthResId: uint(d.Id),
			})
			if err != nil {
				log.Error("AuthCheck.Check", err)
				util.Error(c, err)
				c.Abort()
				return
			}
			if resp.Pass {
				c.Next(ctx)
				return
			}
			reason += string(resp.Reason) + ","
		}

		log.Error("AuthCheck.Check", errs.NoAuth, "reason", reason)
		util.Error(c, errs.NoAuth)
		c.Abort()

	}

}
