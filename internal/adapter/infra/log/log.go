package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hamgit.ir/novin-backend/trader-bot/config"
	"os"
)

func Init() {
	logger := configLogger()
	zap.ReplaceGlobals(logger)
}

func configLogger() *zap.Logger {

	consoleEncoderConfig := zap.NewProductionEncoderConfig()
	consoleEncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	terminalZapCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	core := zapcore.NewTee(terminalZapCore)
	logger := zap.New(core, zap.AddCaller())
	return logger.With(zap.String("service", config.C().App.Name))
}
