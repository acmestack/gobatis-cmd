// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package common

type DBDriver interface {
	Open(driver, info string) error
	Close() error

	QueryTableInfo(dbName, tableName string) ([]ModelInfo, error)
	QueryTableNames(dbName string) ([]string, error)
}
