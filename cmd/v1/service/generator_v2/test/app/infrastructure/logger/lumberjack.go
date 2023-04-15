package logger

import (
	"github.com/wirnat/axara/cmd/v1/service/generator_v2/test/app/infrastructure/env"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)
const DateFormat = "2006-01-02"


func SetupLogger(logPath string) {
	if logPath == "" {
		logPath = env.DefaultLogPath
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			log.Fatalf("create Logger dir %v fail: %v", logPath, err)
			panic(err)
		}
	}

	t := time.Now()
	lumberjackLogger := &lumberjack.Logger{
		// Log file abbsolute path, os agnostic
		Filename:   filepath.ToSlash(logPath + "/" + t.Format(DateFormat) + ".log"),
		MaxSize:    env.ENV.Log.MaxSize, // MB
		MaxBackups: env.ENV.Log.MaxBackup,
		MaxAge:     env.ENV.Log.MaxAge,   // days
		Compress:   env.ENV.Log.Compress, // disabled by default
	}

	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)

	logFormatter := new(log.TextFormatter)
	logFormatter.TimestampFormat = time.RFC1123Z // or RFC3339
	logFormatter.FullTimestamp = true

	log.SetFormatter(logFormatter)
	log.SetLevel(log.InfoLevel)
	log.SetOutput(multiWriter)
}
