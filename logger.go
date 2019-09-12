// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sugarzap

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _globalLogger Logger
var _globalHasher Hasher

type Option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

type Hasher interface {
	Hash(interface{}) string
}

// logger delegates all calls to the underlying zap.Logger
type Logger struct {
	*zap.SugaredLogger
	config         *zap.Config
	replaceGlobals bool
	hasher         Hasher
	callerSkip     int
}

func SetGlobalHasher(h Hasher) {
	_globalHasher = h
}

func AddCallerSkip(skip int) Option {
	return optionFunc(func(log *Logger) {
		log.callerSkip += skip
	})
}

func AddHasher(h Hasher) Option {
	return optionFunc(func(l *Logger) {
		l.hasher = h
	})
}

func WithConfig(c zap.Config) Option {
	return optionFunc(func(l *Logger) {
		l.config = &c
	})
}

func ReplaceGlobals() Option {
	return optionFunc(func(l *Logger) {
		l.replaceGlobals = true
	})
}

func (l Logger) With(key string, value interface{}) Logger {
	n := l
	n.SugaredLogger = l.SugaredLogger.With(key, value)
	return n
}

func (l Logger) WithHash(key string, value interface{}) Logger {
	h := _globalHasher
	if l.hasher != nil {
		h = l.hasher
	}
	if h == nil {
		return l
	}
	hashed := l.hasher.Hash(value)
	n := l
	n.SugaredLogger = l.SugaredLogger.With(key, hashed)
	return n
}

func WithOptions(opts ...Option) Logger {
	l := &Logger{}
	for _, opt := range opts {
		opt.apply(l)
	}

	c := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
		//EncoderConfig: zap.NewProductionEncoderConfig(),
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	if l.config != nil {
		c = *l.config
	}
	zapL, err := c.Build()
	if err != nil {
		log.Fatal(err)
	}

	zapL = zapL.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCaller(), zap.AddCallerSkip(l.callerSkip))
	l.SugaredLogger = zapL.Sugar()
	if l.replaceGlobals {
		zap.ReplaceGlobals(zapL)
		_globalLogger = *l
	}
	return *l
}

func init() {
	WithOptions(ReplaceGlobals(), AddCallerSkip(1))
}

func With(key string, value interface{}) Logger {
	n := _globalLogger
	n.SugaredLogger = _globalLogger.SugaredLogger.With(key, value)
	return n
}
func WithHash(key string, value interface{}) Logger {
	if _globalLogger.hasher == nil {
		return _globalLogger
	}
	hashed := _globalLogger.hasher.Hash(value)
	n := _globalLogger
	n.SugaredLogger = _globalLogger.SugaredLogger.With(zap.String(key, hashed))
	return n
}

func Debug(args ...interface{}) {
	_globalLogger.Debug(args...)
}
func Info(args ...interface{}) {
	_globalLogger.Info(args...)
}
func Error(args ...interface{}) {
	_globalLogger.Error(args...)
}
func Warn(args ...interface{}) {
	_globalLogger.Warn(args...)
}
func Fatal(args ...interface{}) {
	_globalLogger.Fatal(args...)
}

func Debugf(t string, args ...interface{}) {
	_globalLogger.Debugf(t, args...)
}
func Infof(t string, args ...interface{}) {
	_globalLogger.Infof(t, args...)
}
func Errorf(t string, args ...interface{}) {
	_globalLogger.Errorf(t, args...)
}
func Warnf(t string, args ...interface{}) {
	_globalLogger.Warnf(t, args...)
}
func Fatalf(t string, args ...interface{}) {
	_globalLogger.Fatalf(t, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	_globalLogger.Debugw(msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	_globalLogger.Infow(msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	_globalLogger.Warnw(msg, keysAndValues...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	_globalLogger.Errorw(msg, keysAndValues...)
}
func Faterw(msg string, keysAndValues ...interface{}) {
	_globalLogger.Fatalw(msg, keysAndValues...)
}