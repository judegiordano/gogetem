package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func level() log.Level {
	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
		return log.InfoLevel
	}
	switch strings.ToUpper(lvl) {
	case "DEBUG":
		return log.DebugLevel
	case "WARN":
		return log.WarnLevel
	case "INFO":
		return log.InfoLevel
	default:
		return log.ErrorLevel
	}
}

func init() {
	log.SetLevel(level())
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
}

func SetLogLevel(lvl log.Level) {
	log.SetLevel(lvl)
}

func GetLogLevel() log.Level {
	return log.GetLevel()
}

func Info(msg interface{}, keyvals ...interface{}) {
	log.Infoln(msg, keyvals)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	log.Warnln(msg, keyvals)
}

func Debug(msg interface{}, keyvals ...interface{}) {
	log.Debugln(msg, keyvals)
}

func Error(msg interface{}, keyvals ...interface{}) {
	log.Errorln(msg, keyvals)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	log.Fatalln(msg, keyvals)
}
