package logger

type Logger interface {
	Debug(tag, message string)
	Info(tag, message string)
	Warn(tag, message string)
	Error(tag, message string)
}
