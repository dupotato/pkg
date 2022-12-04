/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-03 22:32:42
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-03 22:53:10
 * @FilePath: /pkg/main.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"pkg/logger"
)

func main() {
	fmt.Println("main begin")
	var lc logger.LConfig
	lc.Type = "zerolog"
	logger.NewLogger(lc)

}
