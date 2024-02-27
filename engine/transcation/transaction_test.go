package transaction

// func Test_tx_Execute(t *testing.T) {
// 	type args struct {
// 		ctx    context.Context
// 		tasks  []*tsM.Task
// 		params *model.Params
// 	}
// 	tests := []struct {
// 		name    string
// 		tr      *tx
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "try and cancel",
// 			args: args{
// 				ctx: context.Background(),
// 				tasks: []*tsM.Task{
// 					{
// 						Model: common.Model{Id: 1, AppId: 1},
// 						Name:  "test",
// 					},
// 				},
// 				params: &model.Params{
// 					Input: map[string]interface{}{
// 						"quota_id": n3,
// 						"group_id": n3,
// 					},
// 					User: &user.SdkUser{Id: "3"},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.tr.Execute(tt.args.ctx, tt.args.tasks, tt.args.params); (err != nil) != tt.wantErr {
// 				t.Errorf("tx.Execute() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
