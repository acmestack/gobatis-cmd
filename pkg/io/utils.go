/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
