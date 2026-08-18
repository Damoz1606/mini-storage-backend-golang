package logger

import (
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logIdKey = "id"
const tagKey = "tag"
const timeKey = "timestamp"

type zapLogger struct {
	client *zap.Logger
}

func NewZap() Logger {
	return NewZapWithSink(zapcore.AddSync(os.Stdout))
}

func NewZapWithSink(sink zapcore.WriteSyncer) Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = timeKey
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true

	encoder := zapcore.NewJSONEncoder(config.EncoderConfig)

	core := zapcore.NewCore(encoder, sink, zapcore.DebugLevel)

	var opts []zap.Option
	if config.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(zapcore.Level(100)))
	}

	zapInstance := zap.New(core, opts...)

	defer func() {
		_ = zapInstance.Sync()
	}()

	logID := uuid.NewString()

	zapInstance = zapInstance.With(
		zap.String(logIdKey, logID),
	)

	return &zapLogger{
		client: zapInstance,
	}
}

func (z zapLogger) Debug(tag, message string) {
	z.client.Debug(message, zap.String(tagKey, tag))
}

func (z zapLogger) Info(tag, message string) {
	z.client.Info(message, zap.String(tagKey, tag))
}

func (z zapLogger) Warn(tag, message string) {
	z.client.Warn(message, zap.String(tagKey, tag))
}

func (z zapLogger) Error(tag, message string) {
	z.client.Error(message, zap.String(tagKey, tag))
}
