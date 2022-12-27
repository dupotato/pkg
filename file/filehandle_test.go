/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-27 09:54:24
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-27 09:54:38
 * @FilePath: /pkg/file/filehandle_test.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package file

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_Rename(t *testing.T) {
	d1 := "/data/1.txt"
	d2 := "/data/bak/1/3/4/2.txt"
	d, _ := filepath.Split(d2)
	CreateMutiDir(d)
	err := os.Rename(d1, d2)
	fmt.Println(err)
}
