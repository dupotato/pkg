/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-07 11:40:42
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-07 11:42:15
 * @FilePath: /pkg/cmdexec/cmdexec_test.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package cmdexec

import "testing"

func Test_NewCmdExec(t *testing.T) {
	var para []string
	para = append(para, "-al")
	lscmd := NewCmdExec("ll ", para)
	lscmd.BeginExec()
}
