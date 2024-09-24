package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugar *zap.SugaredLogger

type messageType string

const (
	Info  messageType = "Info"
	Error messageType = "Error"
)

func InitLogger() {
	// file, err := os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic("Failed to open log file")
	// }

	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:     "time",
		LevelKey:    "level",
		MessageKey:  "msg",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.TimeEncoderOfLayout("[2006-01-02 15:04]"),
	})

	// fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	// 	TimeKey:     "time",
	// 	LevelKey:    "level",
	// 	MessageKey:  "msg",
	// 	EncodeLevel: zapcore.CapitalLevelEncoder,
	// 	EncodeTime:  zapcore.TimeEncoderOfLayout("2006-01-02 15:04"),
	// })

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		// zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	sugar = logger.Sugar()
}

func SyncLogger() {
	err := sugar.Sync()
	if err != nil {
		sugar.Error("Error syncing logger: ", err)
	}
}

func Message(message string, msgType messageType) {
	switch msgType {
	case Info:
		sugar.Infof("Message: %s", message)
	case Error:
		sugar.Errorf("Error: %s", message)
	default:
		sugar.Infof("Info: %s", message)
	}
}
