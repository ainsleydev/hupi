// Copyright 2023 ainsley.dev. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

// DefaultLogger is an alias for the standard logrus Logger.
var DefaultLogger = logrus.New()

// Bootstrap creates a new Logger.
func Bootstrap(prefix string) {
	DefaultLogger.SetFormatter(&localFormatter{
		Prefix: prefix,
	})
	DefaultLogger.SetLevel(logrus.TraceLevel)
}

// WithField logs with field, sets a new map containing
// "fields".
func WithField(key string, value interface{}) *logrus.Entry {
	return DefaultLogger.WithFields(logrus.Fields{"fields": logrus.Fields{
		key: value,
	}})
}

// WithFields logs with fields, sets a new map containing
// "fields".
func WithFields(fields logrus.Fields) *logrus.Entry {
	return DefaultLogger.WithFields(logrus.Fields{"fields": fields})
}

// WithError - Logs with a Verbis error.
func WithError(err interface{}) *logrus.Entry {
	return DefaultLogger.WithField("error", err)
}

// SetOutput sets the output of the DefaultLogger to an io.Writer,
// useful for testing.
func SetOutput(writer io.Writer) {
	DefaultLogger.SetOutput(writer)
}

// SetLevel sets the level of the DefaultLogger.
func SetLevel(level logrus.Level) {
	DefaultLogger.SetLevel(level)
}

// SetLogger sets the application DefaultLogger.
func SetLogger(l *logrus.Logger) {
	DefaultLogger = l
}

// Trace logs a trace message with args.
func Trace(args ...any) {
	DefaultLogger.Trace(args...)
}

// Debug logs a debug message with args.
func Debug(args ...any) {
	DefaultLogger.Debug(args...)
}

// Info logs ab info message with args.
func Info(args ...any) {
	DefaultLogger.Info(args...)
}

// Warn logs a warn message with args.
func Warn(args ...any) {
	DefaultLogger.Warn(args...)
}

// Error logs an error message with args.
func Error(args ...any) {
	DefaultLogger.Error(args...)
}

// Fatal logs a fatal message with args.
func Fatal(args ...any) {
	DefaultLogger.Fatal(args...)
}

// Panic logs a panic message with args.
func Panic(args ...any) {
	DefaultLogger.Panic(args...)
}

// Tracef logs a trace message with a format and args.
func Tracef(format string, args ...any) {
	DefaultLogger.Tracef(format, args...)
}

// Debugf logs a debug message with a format and args.
func Debugf(format string, args ...any) {
	DefaultLogger.Debugf(format, args...)
}

// Infof logs ab info message with a format and args.
func Infof(format string, args ...any) {
	DefaultLogger.Infof(format, args...)
}

// Warnf logs a warn message with a format and args.
func Warnf(format string, args ...any) {
	DefaultLogger.Warnf(format, args...)
}

// Errorf logs an error message with a format and args.
func Errorf(format string, args ...any) {
	DefaultLogger.Errorf(format, args...)
}
