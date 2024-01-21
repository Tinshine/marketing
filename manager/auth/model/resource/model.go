package resource

import (
	"marketing/const/auth"
	"marketing/const/errs"

	"github.com/pkg/errors"
)

type QueryReq struct {
	AppId   uint          `form:"app_id"`
	ResType *auth.ResType `form:"res_type"`
	ResId   string        `form:"res_id"`
}

func (r *QueryReq) Validate() error {
	if r.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if r.ResType == nil {
		return nil
	}
	if r.ResId == "" {
		return errors.WithMessage(errs.InvalidParams, "res_id is required")
	}

	return nil
}

type QueryResp struct {
	Data  []*RespModel `json:"data"`
	Total int          `json:"total"`
}

type RespModel struct {
	Id        int          `json:"id"`
	AppId     uint         `json:"app_id"`
	ResType   auth.ResType `json:"res_type"`
	ResId     string       `json:"res_id"`
	CreatedBy string       `json:"created_by"`
	CreatedAt int64        `json:"created_at"`
	UpdateBy  string       `json:"update_by"`
	UpdatedAt int64        `json:"updated_at"`
}

type AuthRes struct {
	Id        int           `gorm:"column:id"`
	AppId     uint          `gorm:"column:app_id"`
	ResType   auth.ResType  `gorm:"column:res_type"`
	ResId     string        `gorm:"column:res_id"`
	AuthType  auth.AuthType `gorm:"column:auth_type"`
	CreatedBy string        `gorm:"column:created_by"`
	CreatedAt int64         `gorm:"column:created_at"`
}

func (a *AuthRes) TableName() string {
	return "auth_res"
}

func (a *AuthRes) ToRespModel() *RespModel {
	return &RespModel{
		Id:        a.Id,
		AppId:     a.AppId,
		ResType:   a.ResType,
		ResId:     a.ResId,
		CreatedBy: a.CreatedBy,
		CreatedAt: a.CreatedAt,
	}
}

type DeleteReq struct {
	Id int `json:"id"`
}

func (r *DeleteReq) Validate() error {
	if r.Id == 0 {
		return errors.WithMessage(errs.InvalidParams, "id is required")
	}

	return nil
}

type AddReq struct {
	AppId    uint          `json:"app_id"`
	ResType  auth.ResType  `json:"res_type"`
	ResId    string        `json:"res_id"`
	AuthType auth.AuthType `json:"auth_type"`
}

func (r *AddReq) Validate() error {
	if r.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if err := r.ResType.Validate(); err != nil {
		return errors.WithMessage(errs.InvalidParams, "res_type is invalid")
	}
	if r.ResId == "" {
		return errors.WithMessage(errs.InvalidParams, "res_id is required")
	}
	if err := r.AuthType.Validate(); err != nil {
		return errors.WithMessage(errs.InvalidParams, "auth_type is invalid")
	}
	return nil
}

type AddResp struct {
	Id int `json:"id"`
}
