package main

import (
	"GITHUA.COM/student/0805/mylogger"
)

func main() {
	log := mylogger.NewLoger("debug")
	log.Debug()
	log.Info()
	log.Warning()
	log.Error()
	log.Fatal()
}
