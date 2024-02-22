package conf

// func TestGetConf(t *testing.T) {
// 	Init()
// 	type args struct {
// 		env consts.Env
// 		key string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		{
// 			name: "log dir",
// 			args: args{
// 				env: consts.Test,
// 				key: conf.ConfKeyAppLogFile,
// 			},
// 			want:    "/output/app.log",
// 			wantErr: false,
// 		},
// 		{
// 			name: "mysql ip",
// 			args: args{
// 				env: consts.Test,
// 				key: conf.ConfKeyMysqlIP,
// 			},
// 			want:    "127.0.0.1",
// 			wantErr: false,
// 		},
// 		{
// 			name: "conf env not exist",
// 			args: args{
// 				env: consts.Dev,
// 				key: conf.ConfKeyAppLogFile,
// 			},
// 			want:    "",
// 			wantErr: true,
// 		},
// 		{
// 			name: "conf not found",
// 			args: args{
// 				env: consts.Test,
// 				key: "not exist",
// 			},
// 			want:    "",
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetConf(tt.args.env, tt.args.key)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetConf() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("GetConf() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
