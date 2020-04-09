// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package test

import (
	db2 "github.com/xfali/gobatis-cmd/internal/pkg/db"
	"log"
	"os"
	"testing"
)

func TestMysqlInfo(t *testing.T) {
	t.Log(db2.GenDBInfo("mysql","testdb", "test", "test", "localhost", 3306))
}

func TestMysql(t *testing.T) {
	db := db2.GetDriver("postgres")
	if db == nil {
		log.Print("not support driver: ", "postgres")
		os.Exit(-1)
	}

	err := db.Open("postgres", db2.GenDBInfo("postgres", "testdb", "test", "test", "localhost", 5432))
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
	defer db.Close()

	tables, err := db.QueryTableNames("testdb")
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range tables {
		t.Log(v)
		model, err := db.QueryTableInfo("testdb", v)
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range model {
			t.Log(m)
		}
	}

}
