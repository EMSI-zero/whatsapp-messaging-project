package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLog(baseDir string) (err error){
	Logger = logrus.New()

	logDirPath := os.Getenv("LOG_DIR_PATH")
	if logDirPath == ""{
		log.SetOutput(os.Stdout)
		return nil
	}

	if !filepath.IsAbs(logDirPath){
		logDirPath = filepath.Join(baseDir, logDirPath)
	}

	logDirPath, err = filepath.Abs(logDirPath)
	if err != nil {
		return err
	}

	logFileName := fmt.Sprintf("%d",os.Getpid())

	logFilePath, err:= filepath.Abs(filepath.Join(logDirPath, logFileName))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(io.MultiWriter(os.Stdout, file))

	return nil
}