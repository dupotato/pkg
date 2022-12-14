/*
 * @Author: dueb dueb@channelsoft.com
 * @Date: 2022-12-27 09:53:59
 * @LastEditors: dueb dueb@channelsoft.com
 * @LastEditTime: 2023-01-06 16:30:36
 * @FilePath: /pkg/file/filehandle.go
 * @Description:
 *
 * Copyright (c) 2022 by dueb dueb@channelsoft.com, All Rights Reserved.
 */
package file

import (
	"fmt"
	"os"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func CreateMutiDir(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Deletefile(file string) {
	os.Remove(file)
}

func DeleteAllFile(path string) {
	os.RemoveAll(path)
}
