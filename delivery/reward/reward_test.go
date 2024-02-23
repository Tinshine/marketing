package reward

import (
	"context"
	"marketing/consts"
	trM "marketing/engine/transcation/model"
	"marketing/user"
	"marketing/util"
	"reflect"
	"testing"
)

func Test_reward_Try(t *testing.T) {
	util.SetUnitTestMode()
	var n1 uint = 1
	var n3 uint = 3
	type args struct {
		ctx    context.Context
		params *trM.Params
	}
	tests := []struct {
		name    string
		r       *reward
		args    args
		want    *trM.Resp
		wantErr bool
	}{
		{
			name: "normal try",
			r:    &reward{TxId: "3"},
			args: args{
				ctx: context.Background(),
				params: &trM.Params{
					Input: map[string]interface{}{
						"quota_id": n3,
						"group_id": n3,
					},
					User:  &user.SdkUser{Id: "3"},
					Ev:    consts.Test,
					AppId: 3,
				},
			},
			want:    &trM.Resp{},
			wantErr: false,
		}, {
			name: "try without quota_id",
			r:    &reward{TxId: "3"},
			args: args{
				ctx: context.Background(),
				params: &trM.Params{
					Input: map[string]interface{}{
						"group_id": n3,
					},
					User:  &user.SdkUser{Id: "3"},
					Ev:    consts.Test,
					AppId: 3,
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "try without group_id",
			r:    &reward{TxId: "3"},
			args: args{
				ctx: context.Background(),
				params: &trM.Params{
					Input: map[string]interface{}{
						"quota_id": n3,
					},
					User:  &user.SdkUser{Id: "3"},
					Ev:    consts.Test,
					AppId: 3,
				},
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "duplicate try",
			r:    &reward{TxId: "1"},
			args: args{
				ctx: context.Background(),
				params: &trM.Params{
					Input: map[string]interface{}{
						"quota_id": n1,
						"group_id": n1,
					},
					User:  &user.SdkUser{Id: "1"},
					Ev:    consts.Test,
					AppId: 1,
				},
			},
			want:    &trM.Resp{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Try(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("reward.Try() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reward.Try() = %v, want %v", got, tt.want)
			}
		})
	}
}
