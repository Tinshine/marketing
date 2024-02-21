package conf

import (
	"fmt"
	"io"
	"marketing/consts"
	"marketing/consts/errs"
	"marketing/util/common"
	"os"
	"path"
	"sync"

	"github.com/bytedance/sonic"
)

var configure = make(map[consts.Env]map[string]string)

var Init = sync.OnceFunc(initConf)

func initConf() {
	loadConf(consts.Test, "log.json")
	loadConf(consts.Test, "mysql_test.json")
	loadConf(consts.Prod, "mysql_prod.json")
}

func loadConf(env consts.Env, fileName string) {
	rPath, err := common.GetRelativePath()
	if err != nil {
		panic(err)
	}
	filePath := path.Join(rPath, "conf", fileName)
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("err is: %v, path is: %s", err, fileName))
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
