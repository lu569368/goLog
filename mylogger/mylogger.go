package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//LogLevel 66666
type LogLevel uint16

const (
	//UNKNOWN ...
	UNKNOWN LogLevel = iota
	//DEBUG ...
	DEBUG
	//TRACE ...
	TRACE
	//INFO ...
	INFO
	//WARNING ...
	WARNING
	//ERROR ...
	ERROR
	//FATAL ...
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志类别")
		return UNKNOWN, err
	}
}
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "debug"
	case TRACE:
		return "trace"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		return "debug"
	}
}
func getInfo(skg int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skg)
	if !ok {
		fmt.Println(" runtime.Caller()函数运行错误")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}
