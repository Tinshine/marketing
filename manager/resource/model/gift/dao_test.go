package gift

import (
	"context"
	"marketing/consts"
	"marketing/consts/resource"
	"marketing/database/rds"
	"marketing/util"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestFindById(t *testing.T) {
	util.TestInit()
	type args struct {
		db *gorm.DB
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "id not found",
			args: args{
				db: rds.DB(context.Background(), consts.Test),
				id: 0,
			},
			want:    false,
			wantErr: true,
		}, {
			name: "id found",
			args: args{
				db: rds.DB(context.Background(), consts.Test),
				id: 1,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindById(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) != tt.want {
				t.Errorf("FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindGiftTypeById(t *testing.T) {
	util.TestInit()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    resource.GiftType
		wantErr bool
	}{
		{
			name: "id not found",
			args: args{
				id: 0,
			},
			want:    -1,
			wantErr: true,
		}, {
			name: "find id=2 lottery",
			args: args{
				id: 2,
			},
			want:    resource.Lottery,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindGiftTypeById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindGiftTypeById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindGiftTypeById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindByAppId(t *testing.T) {
	util.TestInit()
	type args struct {
		db    *gorm.DB
		appId uint
	}
	tests := []struct {
		name    string
		args    args
		want    []*Gift
		wantErr bool
	}{
		{
			name: "app_id 100 have records",
			args: args{
				db:    rds.DB(context.Background(), consts.Test),
				appId: 100,
			},
			want: []*Gift{{
				Id:          2,
				Items:       "[]",
				Emails:      "[]",
				GiftType:    resource.Lottery,
				AppId:       100,
				GiftName:    "lottery",
				LotteryRate: "1.0",
				GroupId:     1,
			}},
			wantErr: false,
		}, {
			name: "app_id 0 have no records",
			args: args{
				db:    rds.DB(context.Background(), consts.Test),
				appId: 0,
			},
			want:    []*Gift{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindByAppId(tt.args.db, tt.args.appId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByAppId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByAppId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindByGroupId(t *testing.T) {
	util.TestInit()
	type args struct {
		db      *gorm.DB
		appId   uint
		groupId uint
	}
	tests := []struct {
		name    string
		args    args
		want    []*Gift
		wantErr bool
	}{
		{
			name: "find by group id 1",
			args: args{
				db:      rds.DB(context.Background(), consts.Test),
				appId:   100,
				groupId: 1,
			},
			want: []*Gift{{
				Id:          2,
				Items:       "[]",
				Emails:      "[]",
				GiftType:    resource.Lottery,
				AppId:       100,
				GiftName:    "lottery",
				LotteryRate: "1.0",
				GroupId:     1,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindByGroupId(tt.args.db, tt.args.appId, tt.args.groupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByGroupId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByGroupId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateById(t *testing.T) {
	util.TestInit()
	type args struct {
		db     *gorm.DB
		id     int
		fields map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "id not found",
			args: args{
				db:     rds.DB(context.Background(), consts.Test),
				id:     0,
				fields: map[string]interface{}{},
			},
			wantErr: true,
		}, {
			name: "id found",
			args: args{
				db:     rds.DB(context.Background(), consts.Test),
				id:     1,
				fields: map[string]interface{}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateById(tt.args.db, tt.args.id, tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("UpdateById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
