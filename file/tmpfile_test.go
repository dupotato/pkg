/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2023-01-06 14:48:19
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2023-01-06 16:53:36
 * @FilePath: /pkg/file/tmpfile_test.go
 * @Description:
 *
 * Copyright (c) 2023 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package file

import (
	"sync"
	"testing"
)

func Test_Delete1(t *testing.T) {
	Path := "./test/data"
	NewTmpFileManager(Path, 3)
	sg := &sync.WaitGroup{}
	sg.Add(2)
	// go func(sg *sync.WaitGroup) {
	// 	id := 1
	// 	for {
	// 		fname := fmt.Sprintf("%s/1_%d", Path, id)

	// 		os.WriteFile(fname, []byte("123"), os.ModeAppend)
	// 		AppendTmpFile(fname)
	// 		time.Sleep(1 * time.Second)
	// 		id = id + 2
	// 	}
	// 	sg.Done()
	// }(sg)

	// go func(sg *sync.WaitGroup) {
	// 	id := 0
	// 	for {
	// 		fname := fmt.Sprintf("%s/2_%d", Path, id)

	// 		os.WriteFile(fname, []byte("123"), os.ModeAppend)
	// 		AppendTmpFile(fname)
	// 		time.Sleep(1 * time.Second)
	// 		id = id + 2
	// 	}
	// 	sg.Done()
	// }(sg)

	sg.Wait()
}
