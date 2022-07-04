/** CGo interface for VILLASnode logger
 *
 * @author Steffen Vogel <svogel2@eonerc.rwth-aachen.de>
 * @copyright 2014-2022, Institute for Automation of Complex Power Systems, EONERC
 * @license Apache 2.0
 *********************************************************************************/

package main

// #include <stdlib.h>
// #include <villas/nodes/go.h>
// void bridge_go_logger_log(_go_logger_log_cb cb, _go_logger l, int level, char *msg);
import "C"
import (
	"fmt"
	"os"
)

type LogLevel int

const (
	// https://github.com/gabime/spdlog/blob/a51b4856377a71f81b6d74b9af459305c4c644f8/include/spdlog/common.h#L76
	LogLevelTrace    LogLevel = iota
	LogLevelDebug    LogLevel = iota
	LogLevelInfo     LogLevel = iota
	LogLevelWarn     LogLevel = iota
	LogLevelError    LogLevel = iota
	LogLevelCritical LogLevel = iota
	LogLevelOff      LogLevel = iota
)

type VillasLogger struct {
	inst C._go_logger
	cb   C._go_logger_log_cb
}

func NewVillasLogger(cb C._go_logger_log_cb, l C._go_logger) *VillasLogger {
	return &VillasLogger{
		cb:   cb,
		inst: l,
	}
}

func (l *VillasLogger) log(lvl LogLevel, msg string) {
	C.bridge_go_logger_log(l.cb, l.inst, C.int(lvl), C.CString(msg))
}

func (l *VillasLogger) Trace(msg string) {
	l.log(LogLevelTrace, msg)
}

func (l *VillasLogger) Debug(msg string) {
	l.log(LogLevelDebug, msg)
}

func (l *VillasLogger) Info(msg string) {
	l.log(LogLevelInfo, msg)
}

func (l *VillasLogger) Warn(msg string) {
	l.log(LogLevelWarn, msg)
}

func (l *VillasLogger) Error(msg string) {
	l.log(LogLevelError, msg)
}

func (l *VillasLogger) Critical(msg string) {
	l.log(LogLevelCritical, msg)
}

func (l *VillasLogger) Tracef(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelTrace, msg)
}

func (l *VillasLogger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelDebug, msg)
}

func (l *VillasLogger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelInfo, msg)
}

func (l *VillasLogger) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelWarn, msg)
}

func (l *VillasLogger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelError, msg)
}

func (l *VillasLogger) Criticalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.log(LogLevelCritical, msg)
}

func (l *VillasLogger) Panic(msg string) {
	l.Critical(msg)
	fmt.Println("Paniced")
	os.Exit(-1)
}
func (l *VillasLogger) Panicf(format string, args ...interface{}) {
	l.Criticalf(format, args...)
	fmt.Println("Paniced")
	os.Exit(-1)
}
