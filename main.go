/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-03 22:32:42
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-06 10:40:29
 * @FilePath: /pkg/main.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package main

import (
	"fmt"

	"github.com/dupotato/pkg/logger"
)

func main() {
	fmt.Println("main begin")
	var lc logger.LConfig
	logger.NewLogger(lc)
	logger.Debugf("111%s", "ok")
}
