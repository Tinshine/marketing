package log

import (
	"fmt"
	"log"
	"marketing/consts"
	confConst "marketing/consts/conf"
	"marketing/util/conf"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/bytedance/sonic"
)

var logger *log.Logger

var Init = sync.OnceFunc(setLogger)

func setLogger() {
	fileName, err := conf.GetConf(consts.Test, confConst.ConfKeyAppLogFile)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger = log.New(f, "AppLog ", log.Ldate|log.Ltime)
}

func Info(event string, kvs ...interface{}) {
	sB := &strings.Builder{}
	writeKVs(sB, info, event, kvs...)
	logger.Printf(sB.String())
}

func Error(event string, e error, kvs ...interface{}) {
	sB := &strings.Builder{}
	kvs = append([]interface{}{"err", e}, kvs...)
	writeKVs(sB, err, event, kvs...)
	logger.Printf(sB.String())
}

func Warn(event string, kvs ...interface{}) {
	sB := &strings.Builder{}
	writeKVs(sB, warn, event, kvs...)
	logger.Printf(sB.String())
}

type level string

const (
	info level = "Info"
	err  level = "Error"
	warn level = "Warn"
)

func writeKVs(sB *strings.Builder, l level, event string, kvs ...interface{}) {
	sB.WriteString(string(l))
	_, file, line, ok := runtime.Caller(2)
	if ok {
		if l == info || l == warn {
			sB.WriteString("\t")
		}
		sB.WriteString(fmt.Sprintf("\t%s:%d ", file, line))
	}
	if event != "" {
		sB.WriteString("Event=")
		sB.WriteString(event)
		sB.WriteString(" | ")
	}
	if len(kvs)%2 == 1 {
		kvs = append(kvs, "")
	}
	for i := 0; i < len(kvs); i += 2 {
		sB.WriteString(formatInterface(kvs[i]))
		sB.WriteString("=")
		sB.WriteString(formatInterface(kvs[i+1]))
		if i+2 < len(kvs) {
			sB.WriteString(" | ")
		}
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
