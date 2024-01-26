package conf

import (
	"io"
	"marketing/consts"
	"marketing/consts/errs"
	"os"
	"sync"

	"github.com/bytedance/sonic"
)

var configure = make(map[consts.Env]map[string]string)

var Init = sync.OnceFunc(initConf)

func initConf() {
	loadConf(consts.Test, "~/go/src/marketing/conf/log.json")
	loadConf(consts.Test, "~/go/src/marketing/conf/mysql.json")
}

func loadConf(env consts.Env, path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bts, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if _, ok := configure[env]; !ok {
		configure[env] = make(map[string]string)
	}

	cfg := make(map[string]string)
	if err := sonic.Unmarshal(bts, &cfg); err != nil {
		panic(err)
	}

	for k, v := range cfg {
		configure[env][k] = v
	}
}

func GetConf(env consts.Env, key string) (string, error) {
	if _, ok := configure[env]; !ok {
		return "", errs.COnfEnvNotExist
	}
	val, ok := configure[env][key]
	if !ok {
		return "", errs.ConfNotFound
	}
	return val, nil
}
