package mylogger

import (
	"fmt"
	"time"
)

//Logger ....
type Logger struct {
	Level LogLevel
}

//NewLoger .....
func NewLoger(levelStr string) Logger {
	Level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: Level,
	}
}

//Enble 判断是否满足写入的需要
func (l Logger) Enble(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}
func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

//Debug ...
func (l Logger) Debug() {
	if l.Enble(DEBUG) {
		log(DEBUG, "我是一个debug错误")
	}
}

//Trace ..
func (l Logger) Trace() {
	if l.Enble(TRACE) {
		log(TRACE, "我是一个TRACE错误")
	}
}

//Info ...
func (l Logger) Info() {
	if l.Enble(INFO) {
		log(INFO, "我是一个INFO错误")
	}
}

//Warning ...
func (l Logger) Warning() {
	if l.Enble(WARNING) {
		log(WARNING, "我是一个WARNING错误")
	}
}

//Error ...
func (l Logger) Error() {
	if l.Enble(ERROR) {
		log(ERROR, "我是一个ERROR错误")
	}
}

//Fatal ...
func (l Logger) Fatal() {
	if l.Enble(FATAL) {
		log(FATAL, "我是一个FATAL错误")
	}
}
