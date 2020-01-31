package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/f1renze/the-architect/common/constant"
)

var (
	logger      *zap.Logger
	once        sync.Once
	lowPriority zap.LevelEnablerFunc
)

func init() {
	lowPriority = func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	}
}

type Any map[string]interface{}

func Init() {
	once.Do(func() {
		stdEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		stdCore := zapcore.NewCore(stdEnc, zapcore.Lock(os.Stdout), lowPriority)

		logger = zap.New(stdCore).WithOptions(
			zap.AddCaller(), zap.AddCallerSkip(1),
		)
	})
}

// redis io writer for logger
type redisWriter struct {
	cli *redis.Client
	key string
}

func (w *redisWriter) Write(p []byte) (int, error) {
	n, err := w.cli.RPush(w.key, p).Result()
	return int(n), err
}

func NewRedisWriter(key string, cli *redis.Client) io.Writer {
	return &redisWriter{
		cli: cli,
		key: key,
	}
}

func Load(cli *redis.Client) {
	if cli == nil {
		return
	}

	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	redisSyncer := zapcore.AddSync(
		NewRedisWriter(constant.RedisKey4Log, cli),
	)
	redisCore := zapcore.NewCore(enc, redisSyncer, lowPriority)

	logger = logger.WithOptions(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(c, redisCore)
	}))
}

func unpack(any Any) (fields []zap.Field) {
	if any == nil {
		return
	}
	for k, v := range any {
		fields = append(fields, zap.Any(k, v))
	}
	return
}

func Info(msg string, any Any) {
	logger.Info(msg, unpack(any)...)
}

func InfoF(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}
func Debug(msg string, any Any) {
	logger.Debug(msg, unpack(any)...)
}

func DebugF(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func Warn(msg string, any Any) {
	logger.Warn(msg, unpack(any)...)
}

func WarnF(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

func Error(msg string, any Any) {
	logger.Error(msg, unpack(any)...)
}

func ErrorF(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

func Fatal(msg string, err error) {
	logger.Fatal(msg, zap.Error(err))
}
