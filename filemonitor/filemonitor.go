/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-06 17:10:23
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-06 17:10:36
 * @FilePath: /pkg/filemonitor/filemonitor.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package filemonitor

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

type FileMonitor struct {
	filepath string
	beginpos int
	readpos  int
	Lines    chan string
}

var GFileMonitor *FileMonitor

func NewFileMonitor(filepath string, pos int) *FileMonitor {
	fm := &FileMonitor{
		filepath: filepath,
		beginpos: pos,
		readpos:  pos,
		Lines:    make(chan string, 1024),
	}
	go fm.fileMonitor()
	return fm
}

// func fileMonitor(filepath string) {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal("watcher new err: [%+v]", err)
// 		return
// 	}
// 	defer watcher.Close()
// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-watcher.Events:
// 				if !ok {
// 					return
// 				}
// 				log.Print("event:", event)

// 			case _, ok := <-watcher.Errors:
// 				if !ok {
// 					return
// 				}
// 			}
// 		}
// 	}()
// 	err = watcher.Add("filepath")
// 	if err != nil {
// 		log.Fatal("err add dir: [%+v]", err)
// 	}
// 	<-done
// }

func (f *FileMonitor) fileMonitor() {
	fileName := f.filepath
	config := tail.Config{
		ReOpen:    true,                                          // 重新打开
		Follow:    true,                                          // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: f.beginpos}, // 从文件的哪个地方开始读
		MustExist: false,                                         // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	//event := make(map[string]string)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		f.readpos++

		fmt.Printf("line_%d : %s\n", f.readpos, line.Text)
		f.Lines <- line.Text
	}
}
