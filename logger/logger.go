/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-03 17:24:37
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-03 22:48:07
 * @FilePath: /pkg/logger/logger.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package logger

func init() {

}

type LConfig struct {
	Type            string
	Level           int
	SyslogPriority  string
	SyslogSeverity  string
	LogPath         string
	FileName        string
	FileRotateCount int
	FileRotateSize  uint64
	RotateByHour    bool
	KeepHours       uint
}

func NewLogger(lc LConfig) {
	if lc.Type == "zerolog" {
		InitZero(lc)
	}
}
