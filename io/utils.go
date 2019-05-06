/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package io

import "os"

// 判断文件夹是否存在
func IsPathExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}

func Mkdir(path string) error {
    return os.MkdirAll(path, os.ModePerm)
}

func OpenAppend(filePath string) (*os.File, error) {
    return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
}

func Write(file *os.File, data []byte) error {
    n := 0
    for n < len(data) {
        c, err := file.Write(data)
        if err != nil {
            return err
        }
        n += c
    }
    return nil
}
