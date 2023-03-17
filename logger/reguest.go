package logger

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func Request(l LogContent) {
	logFile, err := Init()
	if err != nil {
		panic(err)
	}

	defer func() {
		logFile.Close()
		p := recover()
		if p != nil {
			fmt.Println("LOG FILE ERROR", p)
		}
		return
	}()

	whilte := color.New(color.FgWhite)
	bold := whilte.Add(color.BgGreen)

	level := "[DEBUG]"
	message := fmt.Sprintf("{project:%s, handler:%s, request:%s, time:%s}", l.Project, l.Handler, l.Message, l.Time)

	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(fmt.Sprintf("%s", bold.Sprint(level)))

	logFile.WriteString(fmt.Sprintf("%s %s\n", level, message))
	log.Println(message)
}
