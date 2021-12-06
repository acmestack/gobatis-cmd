// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xfali/gobatis-cmd/pkg/common"
	"github.com/xfali/gobatis-cmd/pkg/config"
	db2 "github.com/xfali/gobatis-cmd/pkg/db"
	"github.com/xfali/gobatis-cmd/pkg/generator"
	"testing"
)

func InitSqlite(t *testing.T) []common.ModelInfo {
    sql_table := "CREATE TABLE IF NOT EXISTS CITRON_META (" +
        "file_name VARCHAR(255) PRIMARY KEY," +
        "file_path VARCHAR(255) NULL," +
        "parent VARCHAR(255) NULL," +
        "file_from VARCHAR(255) NULL," +
        "file_to VARCHAR(255) NULL," +
        "hidden TINYINT ," +
        "file_state TINYINT ," +
        "is_dir TINYINT ," +
        "file_size BIGINT ," +
        "mod_time TIMESTAMP," +
        "checksum VARCHAR(128) NULL," +
        "checksum_type VARCHAR(16) NULL);"

    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        t.Fatal(err.Error())
    }
    defer db.Close()

    _, err = db.Exec(sql_table)
    if err != nil {
        t.Fatal(err.Error())
    }

    sqlDb := db2.SqliteDB{}
    err2 := sqlDb.Open("sqlite3", "./test.db")
    if err2 != nil {
        t.Fatal(err2)
    } else {
        defer sqlDb.Close()
    }
    s, err := sqlDb.QueryTableNames("test.db")
    if err != nil {
        t.Fatal(err)
    } else {
        t.Log(s)
    }

    for _, v := range s {
        m, err := sqlDb.QueryTableInfo("test.db", v)
        if err != nil {
            t.Fatal(err)
        } else {
            t.Log(m)
        }
    }

    m, _ := sqlDb.QueryTableInfo("test.db", "CITRON_META")
    return m
}

func TestSqlite(t *testing.T) {
    InitSqlite(t)
}

func TestSqliteGenAll(t *testing.T) {
    config := config.Config{
        PackageName: "mapper",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
        //ModelFile:   "model.go",
    }
    m := InitSqlite(t)
    generator.GenModel(config, "CITRON_META", m)
    generator.GenXml(config, "CITRON_META", m)
    generator.GenV2Proxy(config, "CITRON_META", m)
}
