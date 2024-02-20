package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/util/conf"
	"marketing/util/log"
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	conf.Init()
	log.Init()
	rds.Init()
	type args struct {
		c context.Context
		r *RewardReq
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr error
	}{
		{
			name: "find order",
			args: args{
				c: context.Background(),
				r: &RewardReq{
					Ev:      consts.Test,
					TxId:    1,
					UserId:  "tinshine",
					GroupId: 1,
				},
			},
			want: &Order{
				Id:      2,
				AppId:   100,
				OrderId: "test-order-1",
				TxId:    1,
				GroupId: 1,
				UserId:  "tinshine",
				TxState: consts.StateTry,
			},
			wantErr: nil,
		},
		{
			name: "not found order",
			args: args{
				c: context.Background(),
				r: &RewardReq{
					Ev:      consts.Test,
					TxId:    2,
					UserId:  "tinshine",
					GroupId: 3,
				},
			},
			want:    nil,
			wantErr: errs.OrderNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindOrder(tt.args.c, tt.args.r)
			if err != tt.wantErr {
				t.Errorf("FindOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
