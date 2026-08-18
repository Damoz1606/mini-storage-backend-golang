package logger_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/Damoz1606/mini-storage-backend-golang/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

const mockedTag = "testing-tag"
const mockedMessage = "Test log message"

func TestNewZap(t *testing.T) {
	t.Run("Should return an Logger instance when NewZap is called", func(t *testing.T) {
		value := logger.NewZap()

		assert.Implements(t, (*logger.Logger)(nil), value)
	})
}

func TestZapLogger_Debug(t *testing.T) {
	t.Run("Should log a Debug message when Debug is called", func(t *testing.T) {
		buf := &bytes.Buffer{}
		log := logger.NewZapWithSink(zapcore.AddSync(buf))

		log.Debug(mockedTag, mockedMessage)

		assertLogger(t, buf, "debug")
	})
}

func TestZapLogger_Info(t *testing.T) {
	t.Run("Should log a Info message when Info is called", func(t *testing.T) {
		buf := &bytes.Buffer{}
		log := logger.NewZapWithSink(zapcore.AddSync(buf))

		log.Info(mockedTag, mockedMessage)

		assertLogger(t, buf, "info")
	})
}

func TestZapLogger_Warn(t *testing.T) {
	t.Run("Should log a Warn message when Warn is called", func(t *testing.T) {
		buf := &bytes.Buffer{}
		log := logger.NewZapWithSink(zapcore.AddSync(buf))

		log.Warn(mockedTag, mockedMessage)

		assertLogger(t, buf, "warn")
	})
}

func TestZapLogger_Error(t *testing.T) {
	t.Run("Should log a Error message when Error is called", func(t *testing.T) {
		buf := &bytes.Buffer{}
		log := logger.NewZapWithSink(zapcore.AddSync(buf))

		log.Error(mockedTag, mockedMessage)

		assertLogger(t, buf, "error")
	})
}

func assertLogger(t *testing.T, buf *bytes.Buffer, level string) {
	var result map[string]interface{}
	err := json.Unmarshal(buf.Bytes(), &result)
	require.NoError(t, err)

	assert.Equal(t, level, result["level"])
	assert.Equal(t, mockedMessage, result["msg"])
	assert.Equal(t, mockedTag, result["tag"])
	assert.NotEmpty(t, result["timestamp"])
	assert.NotEmpty(t, result["id"])
}
