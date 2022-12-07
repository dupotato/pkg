/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-04 10:21:26
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2022-12-06 16:38:55
 * @FilePath: /pkg/file/file.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package file

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var osType string
var path string

func init() {
	osType = runtime.GOOS
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
}

func MkdirIfNecessary(createDir string) (err error) {
	s := strings.Split(createDir, path)
	startIndex := 0
	dir := ""
	if s[0] == "" {
		startIndex = 1
	} else {
		dir, _ = os.Getwd() //当前的目录
	}
	for i := startIndex; i < len(s); i++ {
		var d string
		if osType == WINDOWS && filepath.IsAbs(createDir) {
			d = strings.Join(s[startIndex:i+1], path)
		} else {
			d = dir + path + strings.Join(s[startIndex:i+1], path)
		}
		if _, e := os.Stat(d); os.IsNotExist(e) {
			err = os.Mkdir(d, os.ModePerm) //在当前目录下生成md目录
			if err != nil {
				break
			}
		}
	}

	return err
}

func GetCurrentPath() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("can not get current path")
	}
	return dir
}

func IsExistFile(filePath string) bool {
	if len(filePath) == 0 {
		return false
	}
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}
