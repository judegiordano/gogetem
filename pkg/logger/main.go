package logger

import (
	log "github.com/sirupsen/logrus"
)

func init() {
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
