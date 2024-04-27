package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger(dir string, level logrus.Level) (*logrus.Logger, *os.File, error) {
	logFile, err := OpenLogFile(dir)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetOutput(io.MultiWriter(os.Stderr, logFile))
	return logger, logFile, nil
}

func OpenLogFile(dir string) (*os.File, error) {
	now := time.Now()
	filename := filepath.Join(dir, fmt.Sprintf("%s.log", now.Format(time.DateOnly)))
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}
