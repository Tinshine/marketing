package item

import "marketing/consts/resource"

type QueryReq struct {
	Id    *int `form:"id"`
	AppId uint `form:"app_id"`
	Name  *int `form:"id"`
}

type QueryResp struct {
	Items []ItemResp `json:"items"`
	Total int        `json:"total"`
}

type ItemResp struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Item struct {
	Id        int               `gorm:"column:id"`
	AppId     uint              `gorm:"column:app_id"`
	ItemType  resource.ItemType `gorm:"column:item_type"`
	ItemId    int               `gorm:"column:item_id"`
	ItemName  string            `gorm:"column:item_name"`
	Descr     string            `gorm:"column:descr"`
	CreatedBy string            `gorm:"column:created_by"`
	UpdatedBy string            `gorm:"column:updated_by"`
	CreatedAt int64             `gorm:"column:created_at"`
	UpdatedAt int64             `gorm:"column:updated_at"`
}

func (i *Item) TableName() string {
	return "item"
}
