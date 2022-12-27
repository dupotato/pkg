/*
 * @Author: dueb duerbin@126.com
 * @Date: 2022-05-31 15:24:29
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-27 11:02:18
 * @FilePath: /pkg/utils/retry.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb duerbin@126.com, All Rights Reserved.
 */
package utils

import (
	"time"

	"github.com/dupotato/pkg/logger"
)

func Retry(attempts int, sleep time.Duration, fn func() error) error {
	//fmt.Println(attempts)
	//logger.Debugf("retry begin %d\n", attempts)
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			//fmt.Println(err)
			return s.error
		}

		if attempts--; attempts > 0 {
			logger.Warnf("retry func error: %s. attemps #%d after %v.", err.Error(), attempts, sleep)
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func NoRetryError(err error) stop {
	return stop{err}
}
