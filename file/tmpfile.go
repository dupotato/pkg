/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2023-01-06 10:34:22
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2023-01-06 16:42:39
 * @FilePath: /pkg/file/tmpfile.go
 * @Description:
 *
 * Copyright (c) 2023 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package file

import (
	"fmt"
	"sync"
	"time"
)

type TmpFileManager struct {
	tmpdir        string
	deleteSeconds int
	fileQue       map[string]time.Time
	mutexFileQue  sync.RWMutex

	fileList      []string
	fileListBak   []string
	mutexFileList sync.RWMutex

	deleteQueue chan string
}

var tmpFm *TmpFileManager

func NewTmpFileManager(tmpdir string, deleteSeconds int) bool {
	if isExist(tmpdir) == false {
		CreateMutiDir(tmpdir)
	} else {
		DeleteAllFile(tmpdir)
		CreateMutiDir(tmpdir)
	}

	if IsDir(tmpdir) == false {
		return false
	}

	tmpFm = &TmpFileManager{
		tmpdir:        tmpdir,
		deleteSeconds: deleteSeconds,
		fileQue:       make(map[string]time.Time),
		mutexFileQue:  sync.RWMutex{},

		fileList:      make([]string, 0),
		fileListBak:   make([]string, 0),
		mutexFileList: sync.RWMutex{},

		deleteQueue: make(chan string, 2000),
	}
	tmpFm.StartWork()
	return true
}

func (t *TmpFileManager) StartWork() {
	go func() {
		for {
			time.Sleep(time.Duration(t.deleteSeconds) * time.Second)
			t.checkTime()
		}
	}()

	go func() {
		// for v := range t.deleteQueue {
		// 	fmt.Printf("delete file %s\n", v)
		// }
		var count uint64
		for {
			v := <-t.deleteQueue
			Deletefile(v)
			count++
			fmt.Printf("delete file %s count=%d \n", v, count)
		}
	}()
}

func (t *TmpFileManager) checkTime() {
	t.mutexFileList.Lock()
	defer t.mutexFileList.Unlock()
	if len(t.fileList) > 0 {
		if len(t.fileListBak) > 0 {
			for _, v := range t.fileListBak {
				t.deleteQueue <- v
			}
			t.fileListBak = t.fileListBak[0:0]
		}
		t.fileListBak = append(t.fileListBak, t.fileList...)
		t.fileList = t.fileList[0:0]
	}
	fmt.Printf("delete queue len %d %d %d \n", len(t.fileList), len(t.fileListBak), len(t.deleteQueue))
}

func (t *TmpFileManager) AddTmpFile(filepath string) {
	t.mutexFileList.RLock()
	defer t.mutexFileList.RUnlock()
	fmt.Printf("Add TmpFile %s\n", filepath)
	t.fileList = append(t.fileList, filepath)
}

func AppendTmpFile(filepath string) {
	tmpFm.AddTmpFile(filepath)
}
