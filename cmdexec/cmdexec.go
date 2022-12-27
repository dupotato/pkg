/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-07 11:36:29
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-07 11:39:58
 * @FilePath: /pkg/cmdexec/cmdexec.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package cmdexec

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"

	"github.com/dupotato/pkg/logger"
)

type CmdExec struct {
	cmdstr string
	para   []string
	cmd    *exec.Cmd
	result string
}

func NewCmdExec(cmd string, para []string) *CmdExec {
	return &CmdExec{
		cmdstr: cmd,
		para:   para,
	}
}

func (c *CmdExec) BeginExec() bool {
	c.cmd = exec.Command(c.cmdstr, c.para...)

	var stdout io.ReadCloser
	var err error
	if stdout, err = c.cmd.StdoutPipe(); err != nil {
		logger.Errorf(err.Error())
		return false
	}
	defer stdout.Close()

	if err := c.cmd.Start(); err != nil {
		logger.Errorf(err.Error())
		return false
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		logger.Errorf(err.Error())
		return false
	} else {
		fmt.Println(string(opBytes))
		c.result = string(opBytes)
	}
	return true
}
