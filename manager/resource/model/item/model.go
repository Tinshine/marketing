package item

import (
	"marketing/consts/errs"
	"marketing/consts/resource"

	"github.com/pkg/errors"
)

type QueryReq struct {
	Id     *int `form:"id"`
	AppId  uint `form:"app_id"`
	ItemId *int `form:"item_id"`
}

type QueryResp struct {
	Data  []*RespModel `json:"data"`
	Total int          `json:"total"`
}

type RespModel struct {
	Id           int                   `json:"id"`
	AppId        uint                  `json:"app_id"`
	ItemType     resource.ItemType     `json:"item_type"`
	ItemId       int                   `json:"item_id"`
	ItemName     string                `json:"item_name"`
	DeliveryType resource.DeliveryType `json:"delivery_type"`
	Descr        string                `json:"descr"`
	CreatedBy    string                `json:"created_by"`
	UpdatedBy    string                `json:"updated_by"`
	CreatedAt    int64                 `json:"created_at"`
	UpdatedAt    int64                 `json:"updated_at"`
}

type Item struct {
	Id           int                   `gorm:"column:id"`
	AppId        uint                  `gorm:"column:app_id"`
	ItemType     resource.ItemType     `gorm:"column:item_type"`
	ItemId       int                   `gorm:"column:item_id"`
	ItemName     string                `gorm:"column:item_name"`
	DeliveryType resource.DeliveryType `gorm:"column:delivery_type"`
	Descr        string                `gorm:"column:descr"`
	CreatedBy    string                `gorm:"column:created_by"`
	UpdatedBy    string                `gorm:"column:updated_by"`
	CreatedAt    int64                 `gorm:"column:created_at"`
	UpdatedAt    int64                 `gorm:"column:updated_at"`
}

func (i *Item) TableName() string {
	return "item"
}

func (i *Item) ToRespModel() *RespModel {
	return &RespModel{
		Id:           i.Id,
		AppId:        i.AppId,
		ItemType:     i.ItemType,
		ItemId:       i.ItemId,
		ItemName:     i.ItemName,
		DeliveryType: i.DeliveryType,
		Descr:        i.Descr,
		CreatedBy:    i.CreatedBy,
		UpdatedBy:    i.UpdatedBy,
		CreatedAt:    i.CreatedAt,
		UpdatedAt:    i.UpdatedAt,
	}
}

type AddReq struct {
	AppId        uint                  `json:"app_id"`
	ItemType     resource.ItemType     `json:"item_type"`
	ItemId       int                   `json:"item_id"`
	ItemName     string                `json:"item_name"`
	DeliveryType resource.DeliveryType `json:"delivery_type"`
	Descr        string                `json:"descr"`
	CreatedBy    string                `json:"created_by"`
}

func (a *AddReq) Validate() error {
	if a.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if err := a.ItemType.Validate(); err != nil {
		return errors.WithMessage(err, "item_type is invalid")
	}
	if len(a.ItemName) == 0 || len(a.ItemName) > 50 {
		return errors.WithMessage(errs.InvalidParams, "item_name length is invalid")
	}
	if len(a.Descr) > 255 {
		return errors.WithMessage(errs.InvalidParams, "descr length is invalid")
	}
	if a.CreatedBy == "" {
		return errors.WithMessage(errs.InvalidParams, "created_by is required")
	}
	if err := a.DeliveryType.Validate(); err != nil {
		return errors.WithMessage(err, "delivery_type is invalid")
	}
	return nil
}

type AddResp struct {
	Id int `json:"id"`
}

type DeleteReq struct {
	Id int `json:"id"`
}

func (d *DeleteReq) Validate() error {
	if d.Id == 0 {
		return errors.WithMessage(errs.InvalidParams, "id is required")
	}
	return nil
}

type UpdateReq struct {
	Id    int    `json:"id"`
	Descr string `json:"descr"`
}

func (u *UpdateReq) Validate() error {
	if u.Id == 0 {
		return errors.WithMessage(errs.InvalidParams, "id is required")
	}
	if len(u.Descr) > 255 {
		return errors.WithMessage(errs.InvalidParams, "descr length is invalid")
	}
	return nil
}
