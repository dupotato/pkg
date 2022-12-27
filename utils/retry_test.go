/*
 * @Author: dueb duerbin@126.com
 * @Date: 2022-05-31 15:25:28
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-27 10:10:12
 * @FilePath: /pkg/utils/retry_test.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb duerbin@126.com, All Rights Reserved.
 */
package utils

import (
	"fmt"
	"testing"
	"time"
)

var cnt int = 2

func Test_retry(t *testing.T) {
	var ret string
	var ok bool

	a := func() error {
		ok, ret = upload()
		if ok == true {
			return nil
		} else {
			return fmt.Errorf("11")
		}
	}
	Retry(2, 1*time.Second, a)
	fmt.Println(ok, ret)
}

func upload() (bool, string) {
	if cnt == 0 {
		return true, "success"
	} else {
		cnt--
		return false, "failed"
	}
}
