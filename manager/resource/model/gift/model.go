package gift

import (
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/consts/resource"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

type QueryReq struct {
	Id      *int  `json:"id"`
	AppId   uint  `json:"app_id"`
	GroupId *uint `json:"group_id"`
}

type QueryResp struct {
	Data  []*RespModel `json:"data"`
	Total int          `json:"total"`
}

type RespModel struct {
	Gift
}

type Gift struct {
	Id          int                 `json:"id" gorm:"column:id"`
	AppId       uint                `json:"app_id" gorm:"column:app_id"`
	GiftType    resource.GiftType   `json:"gift_type" gorm:"column:gift_type"`
	GiftName    string              `json:"gift_name" gorm:"column:gift_name"`
	LotteryRate LotteryRate         `json:"lottery_rate" gorm:"column:lottery_rate"`
	GroupId     int64               `json:"group_id" gorm:"column:group_id"`
	Items       GiftItems           `json:"items" gorm:"column:items"`
	Emails      GiftEmails          `json:"emails" gorm:"column:emails"`
	State       consts.ReleaseState `json:"state" gorm:"column:state"`
	FilterId    int                 `json:"filter_id" gorm:"column:filter_id"`
	CreatedBy   string              `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   string              `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   int64               `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   int64               `json:"updated_at" gorm:"column:updated_at"`
}

func (g *Gift) TableName() string {
	return "gift"
}

func (g *Gift) ToRespModel() *RespModel {
	return &RespModel{*g}
}

type LotteryRate string

type GiftItems string

type ItemConfig struct {
	Count     int `json:"count"`
	RoleLimit int `json:"role_limit"`
	GameLimit int `json:"game_limit"`
}

func (g GiftItems) Decode() []*ItemConfig {
	var cfs []*ItemConfig
	if err := sonic.UnmarshalString(string(g), &cfs); err != nil {
		return nil
	}
	return cfs
}

type GiftEmails string

type EmailConfigs []*EmailConfig

type EmailConfig struct {
	Language  []string `json:"language"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	IsDefault bool     `json:"is_default"`
}

func (g GiftEmails) Decode() []*EmailConfig {
	var cfs []*EmailConfig
	if err := sonic.UnmarshalString(string(g), &cfs); err != nil {
		return nil
	}
	return cfs
}

func (es EmailConfigs) Validate() error {
	if len(es) == 0 {
		return nil
	}
	hasDefault := false
	visited := map[string]struct{}{}
	for _, c := range es {
		if len(c.Language) == 0 || len(c.Title) == 0 || len(c.Content) == 0 {
			return errors.WithMessage(errs.InvalidParams, "all email fields must offered")
		}
		if c.IsDefault {
			if hasDefault {
				return errors.WithMessage(errs.InvalidParams, "only one default config needed")
			}
			hasDefault = true
		}
		for _, l := range c.Language {
			if _, ok := visited[l]; ok {
				return errors.WithMessage(errs.InvalidParams, "duplicate language")
			}
			visited[l] = struct{}{}
		}
	}
	return nil
}

type AddReq struct {
	AppId       uint              `json:"app_id"`
	GiftType    resource.GiftType `json:"gift_type"`
	GiftName    string            `json:"gift_name"`
	LotteryRate LotteryRate       `json:"lottery_rate"`
	GroupId     int64             `json:"group_id"`
	Items       []*ItemConfig     `json:"items"`
	Emails      EmailConfigs      `json:"emails"`
}

func (a *AddReq) Validate() error {
	if a.AppId == 0 {
		return errors.WithMessage(errs.InvalidParams, "app_id is required")
	}
	if a.GiftType != resource.Normal && a.GiftType != resource.Lottery {
		return errors.WithMessage(errs.InvalidParams, "gift_type is invalid")
	}
	if a.GiftType == resource.Normal && a.LotteryRate != "" {
		return errors.WithMessage(errs.InvalidParams, "normal gift's lottery_rate should be empty")
	}
	if a.GiftType == resource.Lottery && a.LotteryRate == "" {
		return errors.WithMessage(errs.InvalidParams, "lottery gift's lottery_rate is required")
	}
	if len(a.GiftName) == 0 || len(a.GiftName) > 50 {
		return errors.WithMessage(errs.InvalidParams, "gift_name length invalid")
	}
	if len(a.Items) == 0 {
		return errors.WithMessage(errs.InvalidParams, "items is required")
	}
	if err := a.Emails.Validate(); err != nil {
		return errors.WithMessage(errs.InvalidParams, err.Error())
	}
	return nil
}

func (a *AddReq) ToModel(username string) (*Gift, error) {
	items, err := sonic.MarshalString(a.Items)
	if err != nil {
		return nil, errors.WithMessage(errs.Internal, "marshal items")
	}
	emails, err := sonic.MarshalString(a.Emails)
	if err != nil {
		return nil, errors.WithMessage(errs.Internal, "marshal emails")
	}

	gift := &Gift{
		AppId:       a.AppId,
		GiftType:    a.GiftType,
		GiftName:    a.GiftName,
		LotteryRate: a.LotteryRate,
		GroupId:     a.GroupId,
		Items:       GiftItems(items),
		Emails:      GiftEmails(emails),
		State:       consts.StateCreated,
		CreatedBy:   username,
	}

	return gift, nil
}

type AddResp struct {
	Id int `json:"id"`
}
