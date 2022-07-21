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

package db

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/acmestack/gobatis-cmd/pkg/common"
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
	var a1, a2, a3 string
	var a4, a5 interface{}
	var pk bool
	for row.Next() {
		var info common.ModelInfo
		err := row.Scan(&a1, &a2, &a3, &a4, &a5, &pk)
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
