package log

import (
	"log"
	consts "marketing/consts/conf"
	"marketing/util/conf"
	"os"
	"strings"
	"sync"

	"github.com/bytedance/sonic"
)

var logger *log.Logger

var Init = sync.OnceFunc(setLogger)

func setLogger() {
	fileName, err := conf.GetConf(consts.LogDirConfKey)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger = log.New(f, "AppLog ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(event string, kvs ...interface{}) {
	sB := &strings.Builder{}
	sB.WriteString("[INFO] ")
	writeKVs(sB, kvs...)
	logger.Printf(sB.String())
}

func Error(event string, err error, kvs ...interface{}) {
	sB := &strings.Builder{}
	sB.WriteString("[ERROR] error=")
	sB.WriteString(err.Error())
	sB.WriteString(" ")
	writeKVs(sB, kvs...)
	logger.Printf(sB.String())
}

func Warn(event string, kvs ...interface{}) {
	sB := &strings.Builder{}
	sB.WriteString("[WARN] ")
	writeKVs(sB, kvs...)
	logger.Printf(sB.String())
}

func writeKVs(sB *strings.Builder, kvs ...interface{}) {
	if len(kvs)%2 == 1 {
		kvs = append(kvs, "")
	}
	for i := 0; i < len(kvs); i += 2 {
		sB.WriteString(formatInterface(kvs[i]))
		sB.WriteString("=")
		sB.WriteString(formatInterface(kvs[i+1]))
		sB.WriteString(" ")
	}
}

func formatInterface(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		s, _ := sonic.MarshalString(v)
		return s
	}
}
