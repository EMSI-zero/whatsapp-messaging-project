package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLog(baseDir string) (err error) {
	Logger = logrus.New()

	logDirPath := os.Getenv("LOG_DIR_PATH")
	if logDirPath == "" {
		Logger.SetOutput(os.Stdout)
		return nil
	}

	if !filepath.IsAbs(logDirPath) {
		logDirPath = filepath.Join(baseDir, logDirPath)
	}

	logDirPath, err = filepath.Abs(logDirPath)
	if err != nil {
		return err
	}

	logFileName := fmt.Sprintf("%d", os.Getpid())

	logFilePath, err := filepath.Abs(filepath.Join(logDirPath, logFileName))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	Logger.SetFormatter(formatter)
	Logger.SetOutput(io.MultiWriter(os.Stdout, file))
	Logger.SetLevel(logrus.InfoLevel)

	return nil
}

func WithField(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

func Info(ctx context.Context, args ...interface{}) {
	userID := ctx.Value("user_id")
	Logger.WithField("user_id", userID).Info(args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	userID := ctx.Value("user_id")
	Logger.WithField("user_id", userID).Debug(args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	userID := ctx.Value("user_id")
	Logger.WithField("user_id", userID).Warn(args...)
}

func Error(ctx context.Context, args ...interface{}) {
	userID := ctx.Value("user_id")
	Logger.WithField("user_id", userID).Error(args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	userID := ctx.Value("user_id")
	Logger.WithField("user_id", userID).Panic(args...)
}
