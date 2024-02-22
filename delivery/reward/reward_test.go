package reward

// func TestRewardConfirm(t *testing.T) {
// 	util.SetUnitTestMode()
// 	type args struct {
// 		ctx    context.Context
// 		params *model.Params
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "correct params",
// 			args: args{
// 				ctx: context.Background(),
// 				params: &model.Params{
// 					Input: map[string]interface{}{
// 						"quota_id": uint(1),
// 						"group_id": uint(1),
// 					},
// 					User: &user.SdkUser{
// 						Id: "tinshine",
// 					},
// 					Ev:    consts.Test,
// 					AppId: 100,
// 				},
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "missing quota_id in Input",
// 			args: args{
// 				ctx: context.Background(),
// 				params: &model.Params{
// 					Input: map[string]interface{}{
// 						"group_id": uint(1),
// 					},
// 					User: &user.SdkUser{
// 						Id: "tinshine",
// 					},
// 					Ev:    consts.Test,
// 					AppId: 100,
// 				},
// 			},
// 			wantErr: true,
// 		}, {
// 			name: "missing group_id in Input",
// 			args: args{
// 				ctx: context.Background(),
// 				params: &model.Params{
// 					Input: map[string]interface{}{
// 						"quota_id": uint(1),
// 					},
// 					User: &user.SdkUser{
// 						Id: "tinshine",
// 					},
// 					Ev:    consts.Test,
// 					AppId: 100,
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			txId, err := idgen.Gen(tt.args.ctx)
// 			if err != nil {
// 				t.Errorf("error generating, err %v", err)
// 				return
// 			}
// 			r := NewReward(txId)
// 			_, err = r.Try(tt.args.ctx, tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Try() error = %v", err)
// 				return
// 			}
// 			_, err = r.Confirm(tt.args.ctx, tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Confirm() error = %v", err)
// 				return
// 			}
// 		})
// 	}
// }

// func TestTryCancel(t *testing.T) {
// 	conf.Init()
// 	log.Init()
// 	rds.Init()
// 	type args struct {
// 		ctx    context.Context
// 		params *model.Params
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "correct params",
// 			args: args{
// 				ctx: context.Background(),
// 				params: &model.Params{
// 					Input: map[string]interface{}{
// 						"quota_id": uint(1),
// 						"group_id": uint(2),
// 					},
// 					User: &user.SdkUser{
// 						Id: "tinshine",
// 					},
// 					Ev:    consts.Test,
// 					AppId: 100,
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			txId, err := idgen.Gen(tt.args.ctx)
// 			if err != nil {
// 				t.Errorf("error generating, err %v", err)
// 				return
// 			}
// 			r := NewReward(txId)
// 			_, err = r.Try(tt.args.ctx, tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Try() error = %v", err)
// 				return
// 			}
// 			_, err = r.Cancel(tt.args.ctx, tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Confirm() error = %v", err)
// 				return
// 			}
// 		})
// 	}
// }
