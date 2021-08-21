package logger

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// ErrorRetryLog to store retry log
type ErrorRetryLog struct {
	Group   string
	PID     string
	Msg     string
	Content interface{}
}

// ErrorLog to store error log
type ErrorLog struct {
	Error error
	Msg   string
}

// ErrorFeatureLog to store feature log error
type ErrorFeatureLog struct {
	Feature string
	Request interface{}
	Error   interface{}
	Msg     string
}

// InfoProcessLog to store processing log time
type InfoProcessLog struct {
	Content interface{}
	Start   time.Time
	End     time.Time
	Msg     string
}

// InfoLog represents log for info
type InfoLog struct {
	Content interface{}
	Msg     string
}

var (
	// consoleLogger represents active logger to os.Stdout
	consoleLogger zerolog.Logger
)

// InitLogger create new logger
func InitLogger(maxAge int, debug bool) {
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000000"
	cw := zerolog.New(os.Stdout)

	consoleLogger = zerolog.New(cw).With().Timestamp().Logger()
}

// LogError create error log
func LogError(l ErrorLog) {
	consoleLogger.Error().Str("Error", l.Error.Error()).Msg(l.Msg)
}

// LogFeatureError create feature log error
func LogFeatureError(l ErrorFeatureLog) {
	e, _ := json.Marshal(l.Error)
	p, _ := json.Marshal(l.Request)

	consoleLogger.Error().
		RawJSON("Request", p).
		RawJSON("Error", e).
		Str("Feature", l.Feature).
		Msg(l.Msg)
}

// LogProcessInfo create processing feature log info
func LogProcessInfo(p InfoProcessLog) {
	s := (p.End.UnixNano() - p.Start.UnixNano()) / int64(time.Millisecond)
	content, _ := json.Marshal(p.Content)

	consoleLogger.Info().
		RawJSON("Content", content).
		Str("Start", p.Start.Format(time.RFC3339)).
		Str("End", p.End.Format(time.RFC3339)).
		Int64("Duration (ms)", s).
		Msg("Process Completed")
}

// Info create info log
func Info(l InfoLog) {
	content, _ := json.Marshal(l.Content)

	consoleLogger.Info().
		RawJSON("Content", content).
		Msg(l.Msg)
}

// Msg log
func Msg(msg string) {
	consoleLogger.Info().Msg(msg)
}

// LogRetry log error retry
func LogRetry(l ErrorRetryLog) {
	content, _ := json.Marshal(l.Content)

	consoleLogger.Error().
		Str("Group", l.Group).
		Str("PID", l.PID).
		RawJSON("Content", content).
		Msg(l.Msg)
}
