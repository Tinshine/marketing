package auth

import (
	"marketing/consts/auth"
	"marketing/consts/errs"

	"github.com/pkg/errors"
)

type QueryReq struct {
	AppId     uint   `form:"app_id"`
	Username  string `form:"username"`
	AuthResId uint   `form:"auth_res_id"`
}

func (r *QueryReq) Validate() error {
	if r.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if r.Username == "" {
		return errors.WithMessage(errs.InvalidParams, "username is required")
	}
	if r.AuthResId == 0 {
		return errors.WithMessage(errs.InvalidParams, "auth_res_id is required")
	}
	return nil
}

type QueryResp struct {
	Data  []*RespModel `json:"data"`
	Total int          `json:"total"`
}

type RespModel struct {
	Id        int    `json:"id"`
	AppId     uint   `json:"app_id"`
	Username  string `json:"username"`
	AuthResId uint   `json:"auth_res_id"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
	ExpiredAt int64  `json:"expired_at"`
}

type Auth struct {
	Id        int    `gorm:"column:id"`
	AppId     uint   `gorm:"column:app_id"`
	Username  string `gorm:"column:username"`
	AuthResId uint   `gorm:"column:auth_res_id"`
	CreatedBy string `gorm:"column:created_by"`
	CreatedAt int64  `gorm:"column:created_at"`
	ExpiredAt int64  `gorm:"column:expired_at"`
}

func (a *Auth) TableName() string {
	return "auth"
}

func (a *Auth) ToRespModel() *RespModel {
	return &RespModel{
		Id:        a.Id,
		AppId:     a.AppId,
		Username:  a.Username,
		AuthResId: a.AuthResId,
		CreatedBy: a.CreatedBy,
		CreatedAt: a.CreatedAt,
		ExpiredAt: a.ExpiredAt,
	}
}

type DeleteReq struct {
	Id int `json:"id"`
}

func (r *DeleteReq) Validate() error {
	if r.Id <= 0 {
		return errors.WithMessage(errs.InvalidParams, "id is required")
	}

	return nil
}

type AddReq struct {
	AppId     uint   `json:"app_id"`
	Username  string `json:"username"`
	AuthResId uint   `json:"auth_res_id"`
	ExpiredAt int64  `json:"expired_at"`
}

func (r *AddReq) Validate() error {
	if r.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if r.Username == "" {
		return errors.WithMessage(errs.InvalidParams, "username is required")
	}
	if r.AuthResId == 0 {
		return errors.WithMessage(errs.InvalidParams, "auth_res_id is required")
	}

	return nil
}

type AddResp struct {
	Id int `json:"id"`
}

type CheckReq struct {
	AppId     uint   `json:"app_id"`
	Username  string `json:"username"`
	AuthResId uint   `json:"auth_res_id"`
}

func (r *CheckReq) Validate() error {
	if r.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if r.Username == "" {
		return errors.WithMessage(errs.InvalidParams, "username is required")
	}
	if r.AuthResId == 0 {
		return errors.WithMessage(errs.InvalidParams, "auth_res_id is required")
	}
	return nil
}

type CheckResp struct {
	Pass   bool            `json:"pass"`
	Reason auth.FailReason `json:"reason"`
}
