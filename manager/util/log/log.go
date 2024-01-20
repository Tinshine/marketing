package log

import (
	"log"
	consts "marketing/item_manager/const/conf"
	"marketing/item_manager/util/conf"
	"os"
	"strings"
)

var logger *log.Logger

func SetLogger() {
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

func Info(event string, kvs ...string) {
	sB := &strings.Builder{}
	sB.WriteString("[INFO] ")
	writeKVs(sB, kvs...)
	logger.Printf(sB.String())
}

func Error(event string, err error, kvs ...string) {
	sB := &strings.Builder{}
	sB.WriteString("[ERROR] error=")
	sB.WriteString(err.Error())
	sB.WriteString(" ")
	writeKVs(sB, kvs...)
	logger.Printf(sB.String())
}

func Warn(event string, kvs ...string) {
	sB := &strings.Builder{}
	sB.WriteString("[WARN] ")
	writeKVs(sB, kvs...)
	logger.Printf(sB.String())
}

func writeKVs(sB *strings.Builder, kvs ...string) {
	if len(kvs)%2 == 1 {
		kvs = append(kvs, "")
	}
	for i := 0; i < len(kvs); i += 2 {
		sB.WriteString(kvs[i])
		sB.WriteString("=")
		sB.WriteString(kvs[i+1])
		sB.WriteString(" ")
	}
}
