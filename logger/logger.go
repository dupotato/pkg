/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-03 17:24:37
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-27 11:09:24
 * @FilePath: /pkg/logger/logger.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package logger

import (
	"fmt"
)

type Severity int

const (
	DEBUG Severity = iota
	INFO
	WARN
	ERROR
	FATAL
)

var severityName = []string{
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}

const (
	numSeverity = 5
)

type LType int

const (
	DEFAULTLOG LType = iota
	ZEROLOG
	SYSLOG
)

type LConfig struct {
	Type            LType
	Level           int
	SyslogPriority  string
	SyslogSeverity  string
	LogPath         string
	FileName        string
	FileRotateCount int
	FileRotateSize  uint64
	RotateByHour    bool
	KeepHours       uint
	Dev             bool
}

var lcing LConfig

func init() {
	lcing.Type = DEFAULTLOG
}

func NewLogger(lc LConfig) {
	switch lc.Type {
	case ZEROLOG:
		InitZero(lc)
	case SYSLOG:
		InitSyslog("local0", lc.FileName)
	}
	lcing = lc
}

func Debugf(s string, v ...interface{}) {
	switch lcing.Type {
	case ZEROLOG:
		zDebugf(&s, &v)
	case SYSLOG:
		sysDebugf(&s, &v)
	default:
		fmt.Printf(s+"\n", v...)
	}
}
func Infof(s string, v ...interface{}) {
	switch lcing.Type {
	case ZEROLOG:
		zInfof(&s, &v)
	case SYSLOG:
		sysInfof(&s, &v)
	default:
		//fmt.Printf(s, v)
		fmt.Printf(s+"\n", v...)
	}
}
func Warnf(s string, v ...interface{}) {
	switch lcing.Type {
	case ZEROLOG:
		zWarnf(&s, &v)
	case SYSLOG:
		sysWarnf(&s, &v)
	default:
		fmt.Printf(s+"\n", v...)
		//fmt.Printf("\n")
	}
}
func Errorf(s string, v ...interface{}) {
	switch lcing.Type {
	case ZEROLOG:
		zErrorf(&s, &v)
	case SYSLOG:
		sysErrorf(&s, &v)
	default:
		//fmt.Printf(s, v)
		fmt.Printf(s+"\n", v...)
	}
}
func Fatalf(s string, v ...interface{}) {
	switch lcing.Type {
	case ZEROLOG:
		zFatalf(&s, &v)
	case SYSLOG:
		sysFatalf(&s, &v)
	default:
		//fmt.Printf(s, v)
		fmt.Printf(s+"\n", v...)
	}
}

// no format
