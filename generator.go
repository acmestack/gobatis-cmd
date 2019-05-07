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
)

type modelInfo struct {
    columnName string
    dataType   string
    nullable   string
    columnKey  string
    comment    string
    tag        string
}

func generate(config config, db *db, dbName, tableName string) error {
    sqlStr := `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_KEY,COLUMN_COMMENT
		FROM COLUMNS 
		WHERE table_schema = ? AND TABLE_NAME = ?`
    row, err := (*sql.DB)(db).Query(sqlStr, dbName, tableName)
    if err != nil {
        return err
    }
    defer row.Close()

    var models []modelInfo
    var info modelInfo
    for row.Next() {
        err := row.Scan(&info.columnName, &info.dataType, &info.nullable, &info.columnKey, &info.comment)
        if err != nil {
            continue
        }
        info.tag = info.columnName
        models = append(models, info)
    }

    genModel(config, tableName, models)
    genXml(config, tableName, models)
    genProxy(config, tableName, models)

    return nil
}

func generateAll(config config, db *db, dbName string) error {
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
