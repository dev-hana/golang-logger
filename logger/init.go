package logger

import (
	"os"
	"time"
)

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func Init() (*os.File, error) {
	file_name := "logs/" + NowDate() + ".log"
	logFile, err := os.OpenFile(file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return logFile, err
}
