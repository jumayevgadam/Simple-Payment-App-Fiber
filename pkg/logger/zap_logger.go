// We use logger package for tracking errors also.
// One of the best logger in golang is zap.logger.
package logger

import (
	"os"

	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Ensuring apiLogger struct implements Logger interface.
var _ Logger = (*apiLogger)(nil)

// Logger interface keeps needed methods for logging.
type Logger interface {
	InitLogger()
	Debug(args ...any)
	Debugf(template string, args ...any)
	Info(args ...any)
	Infof(template string, args ...any)
	Warn(args ...any)
	Warnf(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
	DPanic(args ...any)
	DPanicf(template string, args ...any)
	Fatal(args ...any)
	Fatalf(template string, args ...any)
}

// apiLogger struct must implement Logger methods now.
type apiLogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

// apiLogger constructor.
func NewApiLogger(cfg *config.Config) *apiLogger {
	return &apiLogger{cfg: cfg}
}

// Creating logger levels.
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// getLoggerLevel func useful func for getting level.
func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// InitLogger.
func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)
	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// Logger methods.

// Debug.
func (l *apiLogger) Debug(args ...any) {
	l.sugarLogger.Debug(args...)
}

// Debugf.
func (l *apiLogger) Debugf(template string, args ...any) {
	l.sugarLogger.Debugf(template, args...)
}

// Info.
func (l *apiLogger) Info(args ...any) {
	l.sugarLogger.Info(args...)
}

// Infof.
func (l *apiLogger) Infof(template string, args ...any) {
	l.sugarLogger.Infof(template, args...)
}

// Warn.
func (l *apiLogger) Warn(args ...any) {
	l.sugarLogger.Warn(args...)
}

// Warnf.
func (l *apiLogger) Warnf(template string, args ...any) {
	l.sugarLogger.Warnf(template, args...)
}

// Error.
func (l *apiLogger) Error(args ...any) {
	l.sugarLogger.Error(args...)
}

// Errorf.
func (l *apiLogger) Errorf(template string, args ...any) {
	l.sugarLogger.Errorf(template, args...)
}

// DPanic.
func (l *apiLogger) DPanic(args ...any) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf.
func (l *apiLogger) DPanicf(template string, args ...any) {
	l.sugarLogger.DPanicf(template, args...)
}

// Fatal.
func (l *apiLogger) Fatal(args ...any) {
	l.sugarLogger.Fatal(args...)
}

// Fatalf.
func (l *apiLogger) Fatalf(tempate string, args ...any) {
	l.sugarLogger.Fatalf(tempate, args...)
}
