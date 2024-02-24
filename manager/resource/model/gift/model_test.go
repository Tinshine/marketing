package gift

import (
	"marketing/consts"
	"marketing/consts/resource"
	"marketing/util"
	"reflect"
	"testing"
)

func TestQueryReq_Validate(t *testing.T) {
	tests := []struct {
		name    string
		q       *QueryReq
		wantErr bool
	}{
		{
			name: "invalid env",
			q: &QueryReq{
				Env: -1,
			},
			wantErr: true,
		}, {
			name: "normal env",
			q: &QueryReq{
				Env: consts.Test,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("QueryReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGift_ToRespModel(t *testing.T) {
	tests := []struct {
		name string
		g    *Gift
		want *RespModel
	}{
		{
			name: "normal model",
			g: &Gift{
				Id:          1,
				AppId:       100,
				GiftName:    "gift1",
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     1,
				Items: `[
					{
						"item_id": 1,
						"count": 1,
						"role_limit": 1,
						"game_limit": 1
					}
				]`,
				Emails:    `[]`,
				CreatedBy: "tinshine",
			},
			want: &RespModel{
				Id:          1,
				AppId:       100,
				GiftType:    resource.Lottery,
				GiftName:    "gift1",
				LotteryRate: "1.0",
				GroupId:     1,
				Items: []*ItemConfig{
					{
						ItemId:    1,
						Count:     1,
						RoleLimit: 1,
						GameLimit: 1,
					},
				},
				Emails:    EmailConfigs{},
				State:     0,
				CreatedBy: "tinshine",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ToRespModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gift.ToRespModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmailConfigs_Validate(t *testing.T) {
	tests := []struct {
		name    string
		es      EmailConfigs
		wantErr bool
	}{
		{
			name:    "empty emails",
			es:      EmailConfigs{},
			wantErr: false,
		}, {
			name: "empty title",
			es: EmailConfigs{
				{
					Title:     "",
					Language:  []string{"zh_CN"},
					Content:   "hello",
					IsDefault: true,
				},
			},
			wantErr: true,
		}, {
			name: "empty language",
			es: EmailConfigs{
				{
					Title:     "title",
					Content:   "hello",
					IsDefault: true,
				},
			},
			wantErr: true,
		}, {
			name: "duplicate language in single emails",
			es: EmailConfigs{
				{
					Title:     "title",
					Content:   "hello",
					Language:  []string{"zh_CN", "zh_CN"},
					IsDefault: true,
				},
			},
			wantErr: true,
		}, {
			name: "duplicate language in multiple emails",
			es: EmailConfigs{
				{
					Title:     "title1",
					Language:  []string{"zh_CN"},
					Content:   "hello",
					IsDefault: true,
				}, {
					Title:    "title2",
					Language: []string{"zh_CN"},
					Content:  "hello",
				},
			},
			wantErr: true,
		}, {
			name: "no default email",
			es: EmailConfigs{
				{
					Title:    "title1",
					Language: []string{"zh_CN"},
					Content:  "hello",
				},
			},
			wantErr: true,
		}, {
			name: "more than one default email",
			es: EmailConfigs{
				{
					Title:     "title1",
					Language:  []string{"zh_CN"},
					Content:   "hello",
					IsDefault: true,
				}, {
					Title:     "title2",
					Language:  []string{"zh_CN"},
					Content:   "hello",
					IsDefault: true,
				},
			},
			wantErr: true,
		}, {
			name: "normal email",
			es: EmailConfigs{
				{
					Title:     "title1",
					Language:  []string{"zh_CN"},
					Content:   "hello",
					IsDefault: true,
				}, {
					Title:     "title2",
					Language:  []string{"en_US"},
					Content:   "hello",
					IsDefault: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.es.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("EmailConfigs.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddReq_Validate(t *testing.T) {
	tests := []struct {
		name    string
		a       *AddReq
		wantErr bool
	}{
		{
			name: "invalid app_id",
			a: &AddReq{
				AppId: 0,
			},
			wantErr: true,
		}, {
			name: "invalid gift_type",
			a: &AddReq{
				AppId:    1,
				GiftType: 2,
			},
			wantErr: true,
		}, {
			name: "invalid lottery_rate",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "",
			},
			wantErr: true,
		}, {
			name: "invalid group_id",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     0,
			},
			wantErr: true,
		}, {
			name: "invalid giftname length",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     1,
				GiftName:    "",
			},
			wantErr: true,
		}, {
			name: "invalid items",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     1,
				GiftName:    "gift1",
			},
			wantErr: true,
		}, {
			name: "invalid email config",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     1,
				GiftName:    "gift1",
				Items: []*ItemConfig{
					{
						ItemId:    1,
						Count:     1,
						RoleLimit: 1,
						GameLimit: 1,
					},
				},
				Emails: EmailConfigs{
					{
						IsDefault: false,
					},
				},
			},
			wantErr: true,
		}, {
			name: "correct config",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     1,
				GiftName:    "gift1",
				Items: []*ItemConfig{
					{
						ItemId:    1,
						Count:     1,
						RoleLimit: 1,
						GameLimit: 1,
					},
				},
				Emails: EmailConfigs{
					{
						Title:     "title1",
						Language:  []string{"zh_CN"},
						Content:   "hello",
						IsDefault: true,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AddReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddReq_ToModel(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		a       *AddReq
		args    args
		want    *Gift
		wantErr bool
	}{
		{
			name: "normal",
			a: &AddReq{
				AppId:       1,
				GiftType:    resource.Lottery,
				LotteryRate: "1.0",
				GroupId:     1,
				GiftName:    "gift1",
				Items: []*ItemConfig{
					{
						ItemId:    1,
						Count:     1,
						RoleLimit: 1,
						GameLimit: 1,
					},
				},
				Emails: EmailConfigs{
					{
						Title:     "title1",
						Language:  []string{"zh_CN"},
						Content:   "hello",
						IsDefault: true,
					},
				},
			},
			want: &Gift{
				AppId:       1,
				GiftType:    resource.Lottery,
				GroupId:     1,
				LotteryRate: "1.0",
				GiftName:    "gift1",
				Items:       `[{"item_id":1,"count":1,"role_limit":1,"game_limit":1}]`,
				Emails:      `[{"language":["zh_CN"],"title":"title1","content":"hello","is_default":true}]`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.ToModel(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddReq.ToModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddReq.ToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateReq_Validate(t *testing.T) {
	util.SetUnitTestMode()
	tests := []struct {
		name    string
		a       *UpdateReq
		wantErr bool
	}{
		{
			name: "invalid id",
			a: &UpdateReq{
				Id: 0,
			},
			wantErr: true,
		}, {
			name: "invalid lottery rate",
			a: &UpdateReq{
				Id: 1,
				LotteryRate: func() *LotteryRate {
					lr := LotteryRate("1.0")
					return &lr
				}(),
			},
			wantErr: true,
		}, {
			name: "correct request",
			a: &UpdateReq{
				Id: 2,
				LotteryRate: func() *LotteryRate {
					lr := LotteryRate("1.0")
					return &lr
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyncReq_Validate(t *testing.T) {
	tests := []struct {
		name    string
		s       *SyncReq
		wantErr bool
	}{
		{
			name: "invalid id",
			s: &SyncReq{
				Id: 0,
			},
			wantErr: true,
		}, {
			name: "correct request",
			s: &SyncReq{
				Id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("SyncReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteReq_Validate(t *testing.T) {
	tests := []struct {
		name    string
		d       *DeleteReq
		wantErr bool
	}{
		{
			name: "invalid id",
			d: &DeleteReq{
				Id: 0,
			},
			wantErr: true,
		}, {
			name: "invalid env",
			d: &DeleteReq{
				Id:  1,
				Env: 3,
			},
			wantErr: true,
		}, {
			name: "correct",
			d: &DeleteReq{
				Id:  1,
				Env: consts.Test,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReleaseReq_Validate(t *testing.T) {
	tests := []struct {
		name    string
		d       *ReleaseReq
		wantErr bool
	}{
		{
			name: "invalid id",
			d: &ReleaseReq{
				Id: 0,
			},
			wantErr: true,
		}, {
			name: "invalid env",
			d: &ReleaseReq{
				Id:  1,
				Env: 3,
			},
			wantErr: true,
		}, {
			name: "correct",
			d: &ReleaseReq{
				Id:  1,
				Env: consts.Test,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ReleaseReq.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
