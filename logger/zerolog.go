/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-03 22:17:17
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-03 22:53:53
 * @FilePath: /pkg/logger/zerolog.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package logger

import (
	"io"
	"os"
	"path"

	//"github.com/nacos-group/nacos-sdk-go/common/logger"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	zz "github.com/rs/zerolog/log"
)

func InitZero(c LConfig) {
	zerolog.SetGlobalLevel(zerolog.Level(c.Level))
	zz.Logger = zz.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05.999",
		},
	)
	zero_loadConfig(c)
}

func zero_loadConfig(lc LConfig) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999"
	zerolog.CallerSkipFrameCount = 4
	//RewriteStderrFile()
	logdev := 1
	if logdev == 1 {
		zz.Logger = zz.Output(
			io.MultiWriter(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				NoColor:    false,
				TimeFormat: "2006-01-02 15:04:05.999",
			}, &lumberjack.Logger{
				Filename:   path.Join(lc.LogPath, lc.FileName),
				MaxSize:    int(lc.FileRotateSize), // megabytes
				MaxBackups: int(lc.FileRotateCount),
				MaxAge:     int((lc.KeepHours + 23) / 24), // days
				Compress:   true,                          // disabled by default
			}),
		).With().Caller().Logger()
	} else {
		zz.Logger = zz.Output(
			io.MultiWriter(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				NoColor:    false,
				TimeFormat: "2006-01-02 15:04:05.999",
			}, &lumberjack.Logger{
				Filename:   path.Join(lc.LogPath, lc.FileName),
				MaxSize:    int(lc.FileRotateSize), // megabytes
				MaxBackups: int(lc.FileRotateCount),
				MaxAge:     int((lc.KeepHours + 23) / 24), // days
				Compress:   true,                          // disabled by default                  // disabled by default
			}),
		).With().Logger()
	}
	Infof("init zerolog over")
}

func Debugf(s string, v ...interface{}) {
	logEvent(zz.Debug(), &s, &v)
}
func Infof(s string, v ...interface{}) {
	logEvent(zz.Info(), &s, &v)
}
func Warnf(s string, v ...interface{}) {
	logEvent(zz.Warn(), &s, &v)
}
func Errorf(s string, v ...interface{}) {
	logEvent(zz.Error(), &s, &v)
}
func Fatalf(s string, v ...interface{}) {
	logEvent(zz.Fatal(), &s, &v)
}

func logEvent(e *zerolog.Event, s *string, v *[]interface{}) {
	if v != nil {
		e.Msgf(*s, (*v)...)
	} else {
		e.Msgf(*s)
	}
}
