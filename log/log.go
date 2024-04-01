package log

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Loggers struct {
	Trace *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

var Logger Loggers

func LogInit() {
	// 로그 파일 열기 또는 생성 (기존 로그는 덮어쓰기)
	logDir := os.Getenv("ROOT_PATH") + "/log"
	logFileName := getLogFileName()
	logFile, err := os.OpenFile(logDir+"/"+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Cannot create log file:", err)
	}

	// 로그 레벨 및 포맷 설정
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// 표준 출력 및 로그 파일에 출력 설정
	logWriter := io.MultiWriter(os.Stdout, logFile)
	Logger.Trace = log.New(logWriter, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Info = log.New(logWriter, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Warn = log.New(logWriter, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Error = log.New(logWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func getLogFileName() string {
	// 현재 날짜를 기반으로 로그 파일 이름 생성 (예: "2023-10-19.log")
	today := time.Now()
	return today.Format("2006-01-02") + ".log"
}
