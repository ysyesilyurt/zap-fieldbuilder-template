package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

/* Generic zap initializer template for Generic uber-go/zap FieldBuilder Template - ysyesilyurt 2021 */

var (
	coloredLevelStringMap = map[zapcore.Level]Color{
		zapcore.DebugLevel:  Green,
		zapcore.InfoLevel:   Blue,
		zapcore.WarnLevel:   Yellow,
		zapcore.ErrorLevel:  Red,
		zapcore.DPanicLevel: Red,
		zapcore.PanicLevel:  Red,
		zapcore.FatalLevel:  Red,
	}
	defaultLevelColor = Red
)

/* Being used as the ctx key to store zap.Logger */
type loggerKeyType string

const loggerKey loggerKeyType = "request_key"

func (lkt loggerKeyType) String() string {
	return string(lkt)
}

/* RootLogger Fallback/root logger for events without context */
var rootLogger *zap.Logger

/* RootLoggerf Should be used only for events that require no context and formatting */
var rootLoggerf *zap.SugaredLogger

func init() {
	rootLogger, _ = getLogger("debug", "console")
	rootLoggerf = rootLogger.Sugar()
}

/* logger.NewContextWithLogger creates a new zap.Logger with the specified logFields zap.Field
and stores that logger into given ctx context.Context using constant loggerKey */
func NewContextWithLogger(ctx context.Context, logFields ...zap.Field) context.Context {
	return context.WithValue(ctx, loggerKey, GetContextLogger(ctx).With(logFields...))
}

/* logger.GetContextLogger fetches zap.Logger inside given ctx context. Context if exists
otherwise returns fallback/root logger */
func GetContextLogger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return rootLogger
	}

	/* If the logger was not passed into the context, then just return the global one instead of exploding. */
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	} else {
		return rootLogger
	}
}

func getLogger(level, encoding string) (*zap.Logger, error) {
	environment := os.Getenv("ENV")

	var logConfig zap.Config
	if checkIsStringOneOfOrContained(environment, "dev", "local", "test") {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = customDevLevelEncoder
	} else if checkIsStringOneOfOrContained(environment, "prod") {
		logConfig = zap.NewProductionConfig()
		logConfig.EncoderConfig.EncodeLevel = customProdLevelEncoder
	} else {
		log.New(os.Stdout, Yellow.Add(" [ WARN ] "), log.Ldate|log.Ltime|log.Lshortfile).
			Println("Unknown logger environment config, using default logger config...")
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = customDevLevelEncoder
	}

	logLevel := getLogLevel(level)
	logConfig.Level = zap.NewAtomicLevelAt(logLevel)
	logConfig.OutputPaths = []string{"stdout"}
	logConfig.ErrorOutputPaths = []string{"stdout"}
	logConfig.Encoding = encoding
	logConfig.EncoderConfig.EncodeTime = timeEncoder
	return logConfig.Build()
}

func getLogLevel(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006 Jan 2 (Mon) 15:04:05"))
}

func defaultLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func customDevLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	levelColor, ok := coloredLevelStringMap[level]
	if !ok {
		levelColor = defaultLevelColor
	}
	enc.AppendString("[" + levelColor.Add(level.CapitalString()) + "]")
}

func customProdLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
