package logger

import "github.com/charmbracelet/log"

func SetLogLevel(level log.Level) {
	log.SetLevel(level)
}

func Info(msg interface{}, keyvals ...interface{}) {
	log.Info(msg, keyvals)
}

func Warn(msg interface{}, keyvals ...interface{}) {
	log.Warn(msg, keyvals)
}

func Debug(msg interface{}, keyvals ...interface{}) {
	log.Debug(msg, keyvals)
}

func Error(msg interface{}, keyvals ...interface{}) {
	log.Error(msg, keyvals)
}

func Fatal(msg interface{}, keyvals ...interface{}) {
	log.Fatal(msg, keyvals)
}
