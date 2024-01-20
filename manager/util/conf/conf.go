package conf

import (
	"io"
	"os"

	"github.com/bytedance/sonic"
)

var configure = make(map[string]string)

func InitConf() {
	loadLogConf()
}

func loadLogConf() {
	f, err := os.Open("~/go/src/marketing/manager/conf/log.conf")
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
