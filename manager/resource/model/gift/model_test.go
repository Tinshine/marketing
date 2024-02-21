package gift

import (
	"marketing/consts/resource"
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
