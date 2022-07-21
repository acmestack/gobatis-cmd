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
	"github.com/acmestack/gobatis-cmd/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type db sql.DB

type MysqlDB struct {
	db *sql.DB
}

type PostgresDB struct {
	db *sql.DB
}

var buildinDrivers = map[string]common.DBDriver{
	"mysql":    &MysqlDB{},
	"postgres": &PostgresDB{},
	"sqlite3":  &SqliteDB{},
}

func GetDriver(name string) common.DBDriver {
	return buildinDrivers[name]
}

func (db *MysqlDB) Open(driver, info string) error {
	d, err := sql.Open(driver, info)
	if err != nil {
		return err
	}
	db.db = d
	return nil
}

func (db *MysqlDB) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

func (db *MysqlDB) QueryTableInfo(dbName, tableName string) ([]common.ModelInfo, error) {
	sqlStr := `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,COLUMN_KEY,COLUMN_COMMENT
		FROM COLUMNS 
		WHERE table_schema = ? AND TABLE_NAME = ?`
	row, err := db.db.Query(sqlStr, dbName, tableName)
	if err != nil {
		return nil, err
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
	return models, nil
}

func (db *MysqlDB) QueryTableNames(dbName string) ([]string, error) {
	sqlStr := `SELECT table_name from tables where table_schema = ?`
	row, err := db.db.Query(sqlStr, dbName)
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

func (db *PostgresDB) Open(driver, info string) error {
	d, err := sql.Open(driver, info)
	if err != nil {
		return err
	}
	db.db = d
	return nil
}

func (db *PostgresDB) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

func (db *PostgresDB) QueryTableInfo(dbName, tableName string) ([]common.ModelInfo, error) {
	//FIXME: primary key not support
	sqlStr := `SELECT column_name, udt_name, is_nullable
		   FROM information_schema.columns sc
		   WHERE table_name = $1`
	row, err := db.db.Query(sqlStr, tableName)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var models []common.ModelInfo
	var info common.ModelInfo
	for row.Next() {
		err := row.Scan(&info.ColumnName, &info.DataType, &info.Nullable)
		if err != nil {
			continue
		}
		info.Tag = info.ColumnName
		models = append(models, info)
	}

	models[0].ColumnKey = "PRI"
	return models, nil
}

func (db *PostgresDB) QueryTableNames(dbName string) ([]string, error) {
	sqlStr := `SELECT tablename
				FROM pg_tables
				WHERE tablename NOT LIKE 'pg%'
					  AND tablename NOT LIKE 'sql_%'
				ORDER BY  tablename`
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

func GenDBInfo(driver, db, username, pw, host string, port int) string {
	if driver == "mysql" {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, pw, host, port, "information_schema")
	} else if driver == "postgres" {
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, username, pw, db, "disable")
	} else if driver == "sqlite" {
		return host
	}
	return ""
}
