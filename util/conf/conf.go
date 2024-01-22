package conf

import (
	"io"
	"os"
	"sync"

	"github.com/bytedance/sonic"
)

var configure = make(map[string]string)

var Init = sync.OnceFunc(initConf)

func initConf() {
	loadConf("~/go/src/marketing/conf/log.json")
	loadConf("~/go/src/marketing/conf/mysql.json")
}

func loadConf(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bts, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	cfg := make(map[string]string)
	if err := sonic.Unmarshal(bts, &cfg); err != nil {
		panic(err)
	}

	for k, v := range cfg {
		configure[k] = v
	}
}

func GetConf(key string) (string, error) {
	val, ok := configure[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}
