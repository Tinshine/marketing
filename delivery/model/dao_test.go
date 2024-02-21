package model

import (
	"context"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/database/rds"
	"marketing/util/conf"
	"marketing/util/idgen"
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
					TxId:    "1",
					UserId:  "tinshine",
					GroupId: 1,
				},
			},
			want: &Order{
				Id:      2,
				AppId:   100,
				OrderId: "test-order-1",
				TxId:    "1",
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
					TxId:    "2",
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

func TestCreateOrder(t *testing.T) {
	conf.Init()
	log.Init()
	rds.Init()
	txId, err := idgen.Gen(context.Background())
	if err != nil {
		t.Errorf("error generating, err %v", err)
		return
	}
	type args struct {
		c     context.Context
		req   *RewardReq
		state consts.TxState
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create order",
			args: args{
				c: context.Background(),
				req: &RewardReq{
					Ev:      consts.Test,
					TxId:    txId,
					UserId:  "tinshine",
					GroupId: 1,
				},
				state: consts.StateTry,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateOrder(tt.args.c, tt.args.req, tt.args.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("CreateOrder() = %v", got)
			}
		})
	}
}

func TestUpdateOrder(t *testing.T) {
	conf.Init()
	log.Init()
	rds.Init()
	type args struct {
		c    context.Context
		id   uint
		ev   consts.Env
		src  consts.TxState
		dest consts.TxState
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update order",
			args: args{
				c:    context.Background(),
				id:   1,
				ev:   consts.Test,
				src:  consts.StateTry,
				dest: consts.StateCancel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateOrder(tt.args.c, tt.args.id, tt.args.ev, tt.args.src, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
