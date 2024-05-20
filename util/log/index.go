package log

import (
	"fmt"
	"time"
)

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func Error(err error) {
	if err != nil {
		fmt.Println(getTime(), " [Error]   :", err.Error())
	}
}

func Info(info string) {
	fmt.Println(getTime(), " [Info]    :", info)
}

func Debug(debug string) {
	fmt.Println(getTime(), " [Debug]   :", debug)
}

func Warning(warning string) {
	fmt.Println(getTime(), " [Warning] :", warning)
}

func Fatal(fatal error) {
	Error(fatal)
	if fatal != nil {
		panic(fatal)
	}
}
