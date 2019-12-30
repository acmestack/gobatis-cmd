/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import (
    "database/sql"
    "github.com/xfali/gobatis-cmd/common"
)

func generate(config Config, db *db, dbName, tableName string) (err error) {
    sqlStr := `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_KEY,COLUMN_COMMENT
		FROM COLUMNS 
		WHERE table_schema = ? AND TABLE_NAME = ?`
    row, err := (*sql.DB)(db).Query(sqlStr, dbName, tableName)
    if err != nil {
        return err
    }
    defer row.Close()

    var models []common.ModelInfo
    var info common.ModelInfo
    for row.Next() {
        err := row.Scan(&info.ColumnName, &info.DataType, &info.Nullable, &info.ColumnKey, &info.Comment)
        if err != nil {
            continue
        }
        info.Tag = info.ColumnName
        models = append(models, info)
    }

    genModel(config, tableName, models)
    genXml(config, tableName, models)
    genV2Proxy(config, tableName, models)

    return RunPlugin(config, tableName, models)
}

func generateAll(config Config, db *db, dbName string) error {
    sqlStr := `SELECT table_name from tables where table_schema = ?`
    row, err := (*sql.DB)(db).Query(sqlStr, dbName)
    if err != nil {
        return err
    }

    defer row.Close()

    var tableName string
    for row.Next() {
        err := row.Scan(&tableName)
        if err != nil {
            continue
        }
        generate(config, db, dbName, tableName)
    }
    return nil
}
