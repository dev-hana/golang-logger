package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type LogContent struct {
	Project  string
	Handler  string
	Function string
	CodeLine int
	LogType  int
	Message  string
	Time     string
}

const Request = 0  //debug
const Response = 1 //debug
const Debug = 2
const Info = 3
const Warn = 4
const Error = 5
const Fatal = 6

func Logger(l LogContent) {
	now := time.Now()
	timeCustom := now.Format("2006-01-02")

	var level string
	var message string

	whilte := color.New(color.FgWhite)
	bold := whilte.Add(color.BgGreen)

	if l.LogType == 0 {
		message = fmt.Sprintf("{project:%s, handler:%s, request:%s, time:%s}", l.Project, l.Handler, l.Message, l.Time)
	} else if l.LogType == 1 {
		message = fmt.Sprintf("{project:%s, handler:%s, response:%s, time:%s}", l.Project, l.Handler, l.Message, l.Time)
	} else {
		message = fmt.Sprintf("{project:%s, handler:%s, function:%s, message:%s, line:%d, time:%s}", l.Project, l.Handler, l.Function, l.Message, l.CodeLine, l.Time)
	}
	switch l.LogType {

	// Request
	case 0:
		level = "[DUBUG] "
	case 1:
		level = "[DUBUG] "
	case 2:
		level = "[DUBUG] "
	case 3:
		level = "[INFO] "
	case 4:
		level = "[WARN] "
		bold = whilte.Add(color.BgHiYellow)
	case 5:
		level = "[ERROR] "
		bold = whilte.Add(color.BgHiRed)
	case 6:
		level = "[FATL] "
		bold = whilte.Add(color.BgHiRed)
	}

	file_name := "logs/" + timeCustom + ".log"

	logFile, err := os.OpenFile(file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	defer func() {
		logFile.Close()
		p := recover()
		if p != nil {
			fmt.Println("LOG FILE ERROR", p)
		}
		return
	}()

	if err != nil {
		panic(err)
	}

	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(level)

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)
	bold.Print(message)
	fmt.Println()
}
