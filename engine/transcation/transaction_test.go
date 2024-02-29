package transaction

import (
	"context"
	"marketing/common"
	"marketing/consts"
	"marketing/consts/engine"
	"marketing/engine/transcation/model"
	tsM "marketing/task/model"
	"marketing/user"
	"marketing/util"
	"testing"
)

func Test_tx_Execute(t *testing.T) {
	util.SetUnitTestMode()
	var n3 uint = 3
	type args struct {
		ctx    context.Context
		tasks  []*tsM.Task
		params *model.Params
	}
	tests := []struct {
		name    string
		tr      *tx
		args    args
		wantErr bool
	}{
		{
			name: "try and confirm",
			tr:   NewTx(consts.Test),
			args: args{
				ctx: context.Background(),
				tasks: []*tsM.Task{
					{
						Model: common.Model{
							Id:    engine.TaskId_Reward,
							AppId: 1,
						},
						Name: "test",
						Type: engine.Tr_Local,
					},
				},
				params: &model.Params{
					Input: map[string]interface{}{
						"quota_id": n3,
						"group_id": n3,
					},
					User: &user.SdkUser{Id: "3"},
				},
			},
		}, {
			name: "try and cancel",
			tr:   NewTx(consts.Test),
			args: args{
				ctx: context.Background(),
				tasks: []*tsM.Task{
					{
						Model: common.Model{
							Id:    engine.TaskId_Reward,
							AppId: 1,
						},
						Name: "test",
						Type: engine.Tr_Local,
					},
				},
				params: &model.Params{
					Input: map[string]interface{}{
						"quota_id": n3,
					},
					User: &user.SdkUser{Id: "3"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Execute(tt.args.ctx, tt.args.tasks, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("tx.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
