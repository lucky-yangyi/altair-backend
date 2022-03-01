package log

import (
	"altair-backend/config"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"runtime"
	"time"
)

var Log *Logger

func init() {
	fileName := config.AppConfig.LogSavePath + "/" +
		config.AppConfig.LogFileName + config.AppConfig.LogFileExt
	Log = NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) string() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""

}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	Fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}

}
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) withFields(f Fields) *Logger {
	ll := l.clone()
	if ll.Fields == nil {
		ll.Fields = make(Fields)
	}
	for k, v := range f {
		ll.Fields[k] = v
	}
	return ll

}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}
	return ll

}

func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depath := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depath])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll

}

func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.Fields)+4)
	data["level"] = level.string()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.Fields) > 0 {
		for k, v := range l.Fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func Output(level Level, message string) {
	body, _ := json.Marshal(Log.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		Log.newLogger.Print(content)
	case LevelInfo:
		Log.newLogger.Print(content)
	case LevelWarn:
		Log.newLogger.Print(content)
	case LevelError:
		Log.newLogger.Print(content)
	case LevelFatal:
		Log.newLogger.Print(content)
	case LevelPanic:
		Log.newLogger.Print(content)

	}

}

func Info(v ...interface{}) {
	Output(LevelInfo, fmt.Sprint(v...))

}
func Infof(format string, v ...interface{}) {
	Output(LevelInfo, fmt.Sprintf(format, v...))
}
func Fatal(v ...interface{}) {
	Output(LevelFatal, fmt.Sprint(v...))

}

func Fatalf(format string, v ...interface{}) {
	Output(LevelFatal, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	Output(LevelError, fmt.Sprint(v...))
}
