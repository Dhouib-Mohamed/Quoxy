package log

import (
	"api-authenticator-proxy/util/env"
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
)

var logLevel = env.GetLogLevel()

func SetLogLevel(level int) {
	if level >= env.NONE && level <= env.DEBUG {
		logLevel = level
	} else {
		Warning("Invalid log level")
	}
}

func getTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func print(colorType color.Attribute, desiredLogLevel int, logPrefix string, a ...any) {
	if len(a) > 0 && logLevel >= desiredLogLevel {
		c := color.New(colorType)
		msg := fmt.Sprintf("%s %s : ", getTime(), logPrefix)
		msg += strings.TrimRight(fmt.Sprintln(a...), "\n")
		c.Print(msg)
	}
}

func Error(err error) {
	if err != nil && logLevel >= env.ERROR {
		c := color.New(color.FgRed)
		c.Print(getTime(), " [Error]  : ", err.Error())
	}
}

func Info(a ...any) {
	print(color.FgGreen, env.INFO, "[Info]   ", a...)
}

func Debug(a ...any) {
	print(color.FgBlue, env.DEBUG, "[Debug]  ", a...)
}

func Warning(a ...any) {
	print(color.FgYellow, env.INFO, "[Warning]", a...)
}

func Fatal(fatal error) {
	Error(fatal)
	fmt.Println("")
	if fatal != nil {
		panic(fatal)
	}
}

func CLI(message ...any) {
	if logLevel >= env.CLI {
		c := color.New(color.FgMagenta)
		c.Println(message...)
	}
}
