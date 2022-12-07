/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-06 17:11:13
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-06 18:03:13
 * @FilePath: /pkg/filemonitor/filemonitor_test.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package filemonitor

import (
	"fmt"
	"testing"
	"time"
)

func Test_NewfileMonitor(t *testing.T) {
	fm := NewFileMonitor("./1", 0)
	for {
		select {
		case line := <-fm.Lines:
			fmt.Println(line)
		case <-time.Tick(5 * time.Second):
			return
		} //dohanlde(line)
	}
}
