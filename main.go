/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-03 22:32:42
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-07 14:45:28
 * @FilePath: /pkg/main.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// fmt.Println("main begin")
	// var lc logger.LConfig
	// logger.NewLogger(lc)
	// logger.Debugf("111%s", "ok")

	//escape := "https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F69rH_u4IQFb_Swf-5-sAlw"
	escape1 := "https://mp.weixin.qq.com/s/69rH_u4IQFb_Swf-5-sAlw"
	// url decode
	unescape, _ := url.QueryUnescape(escape1)
	fmt.Println("un:", unescape)

}
