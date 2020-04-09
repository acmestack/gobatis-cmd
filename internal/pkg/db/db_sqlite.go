// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package db

import (
    "database/sql"
    "fmt"
    //_ "github.com/mattn/go-sqlite3"
    "github.com/xfali/gobatis-cmd/pkg/common"
    "strings"
)

type SqliteDB struct {
    db *sql.DB
}

func (db *SqliteDB) Open(driver, info string) error {
    d, err := sql.Open(driver, info)
    if err != nil {
        return err
    }
    db.db = d
    return nil
}

func (db *SqliteDB) Close() error {
    if db.db != nil {
        return db.db.Close()
    }
    return nil
}

func (db *SqliteDB) QueryTableInfo(dbName, tableName string) ([]common.ModelInfo, error) {
    sqlStr := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
    row, err := db.db.Query(sqlStr)
    if err != nil {
        return nil, err
    }

    defer row.Close()

    var models []common.ModelInfo
    var a1,a2,a3 string
    var a4,a5 interface{}
    var pk bool
    for row.Next() {
        var info common.ModelInfo
        err := row.Scan(&a1,&a2,&a3, &a4, &a5, &pk)
        if err != nil {
            continue
        }
        info.ColumnName = a2
        info.DataType = formatSqliteType(a3)
        if pk {
            info.ColumnKey = "PRI"
        }
        info.Tag = info.ColumnName
        models = append(models, info)
    }
    return models, nil
}

func (db *SqliteDB) QueryTableNames(dbName string) ([]string, error) {
    sqlStr := `SELECT name FROM sqlite_master WHERE type='table'`
    row, err := db.db.Query(sqlStr)
    if err != nil {
        return nil, err
    }

    defer row.Close()

    var ret []string
    var tableName string
    for row.Next() {
        err := row.Scan(&tableName)
        if err != nil {
            continue
        }
        ret = append(ret, tableName)
    }

    return ret, nil
}

func formatSqliteType(origin string) string {
    origin = strings.ToLower(origin)
    i := strings.Index(origin, "(")
    if i != -1 {
        return origin[:i]
    }
    return origin
}
