package log

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hamgit.ir/novin-backend/trader-bot/config"
	"log"
	"os"
)

func Init(lc fx.Lifecycle) error {
	logger := configLogger()
	zap.ReplaceGlobals(logger)

	lc.Append(fx.Hook{
		OnStop: func(c context.Context) error {
			if err := zap.L().Sync(); err != nil {
				log.Println("logger failed to sync:", err)
			}
			return nil
		},
	})
	return nil
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
