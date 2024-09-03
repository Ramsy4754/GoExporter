package xLogger

import (
	"log"
	"os"
	"runtime"
	"sync"
)

var logger *log.Logger
var once sync.Once

func initLogger() {
	logFilePath := getLogFilePath()

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogFilePath() string {
	if runtime.GOOS == "windows" {
		return "./go_exporter.log"
	} else {
		return "/home/ubuntu/go_exporter.log"
	}
}

func GetLogger() *log.Logger {
	once.Do(initLogger)
	return logger
}
