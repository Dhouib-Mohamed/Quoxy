package log

import (
	"api-authenticator-proxy/util/env"
	"fmt"
	"github.com/fatih/color"
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

func Error(err error) {
	if err != nil && logLevel >= env.ERROR {
		c := color.New(color.FgRed)
		c.Print(getTime(), " [Error]  : ", err.Error())
	}
}

func Info(info ...any) {
	if len(info) != 0 && logLevel >= env.INFO {
		c := color.New(color.FgGreen)
		msg := fmt.Sprintf("%s [Info]   : ", getTime())
		msg += fmt.Sprintln(info...)
		c.Print(msg)
	}
}

func Debug(debug ...any) {
	if len(debug) != 0 && logLevel >= env.DEBUG {
		c := color.New(color.FgBlue)
		msg := fmt.Sprintf("%s [Debug]  : ", getTime())
		msg += fmt.Sprintln(debug...)
		c.Print(msg)
	}
}

func Warning(warning ...any) {
	if len(warning) != 0 && logLevel >= env.INFO {
		c := color.New(color.FgYellow)
		msg := fmt.Sprintf("%s [Warning]: ", getTime())
		msg += fmt.Sprintln(warning...)
		c.Print(msg)
	}
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
