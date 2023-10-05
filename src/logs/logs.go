package logs

import (
	"ecommerce-backend/src/pkg"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logz *zap.Logger

func init() {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	discordCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(NewDiscordWriter()),
		zap.ErrorLevel, // Send logs at the error level and above
	)
	// config.Core
	logz = zap.New(discordCore, zap.AddCaller(), zap.AddCallerSkip(1))
	// zap.ReplaceGlobals(log)
	defer logz.Sync()
}

func Info(message string, fields ...zap.Field) {
	logz.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logz.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		logz.Error(v.Error(), fields...)
	case string:
		logz.Error(v, fields...)
	}
}

type DiscordWriter struct{}

func NewDiscordWriter() *DiscordWriter {
	return &DiscordWriter{}
}

func (w *DiscordWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	c := pkg.GetCtx()
	if c != nil {
		ev := pkg.ServerEnvironment{
			Hostname: c.Hostname,
			Url:      c.Url,
			Method:   c.Method,
		}
		go pkg.WebhookSend(message, ev)
	}
	return len(p), nil
}
