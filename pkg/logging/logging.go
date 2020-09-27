package logging

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	DefaultLevel = 0
)

// New constructs a new logger from the provided Options.
//
// The default options are:
//   - Level: DefaultLevel
//   - Stdout: true
//
func New(fileName string, opts ...Option) (*zap.Logger, error) {
	var (
		err    error
		logger *zap.Logger
	)

	o := Options{
		FileName: fileName,
		Level:    DefaultLevel,
		Stdout:   true,
	}

	for _, opt := range opts {
		opt(&o)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     jsonTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(o.Level))

	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(encoderConfig)
	hook := lumberjack.Logger{
		Filename:   o.FileName,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}
	fileCore := zapcore.NewCore(je, zapcore.AddSync(&hook), atomicLevel)
	cores = append(cores, fileCore)
	var options []zap.Option
	if o.Stdout {
		ce := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(ce, zapcore.AddSync(os.Stdout), atomicLevel)
		cores = append(cores, consoleCore)
		caller := zap.AddCaller()
		development := zap.Development()
		options = append(options, caller, development)
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core, options...)

	zap.ReplaceGlobals(logger)

	if err == nil {
		return logger, nil
	} else {
		return nil, err
	}
}

// jsonTimeEncoder is the custom json time encoder.
func jsonTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("2006/01/05 15:04:05:000"))
}
