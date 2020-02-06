// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package main

import (
	"log"
	"os"
	"testing"
)

func TestMysql(t *testing.T) {
	db := buildinDrivers["postgres"]
	if db == nil {
		log.Print("not support driver: ", "postgres")
		os.Exit(-1)
	}

	err := db.Open("postgres", genDBInfo("postgres", "testdb", "test", "test", "localhost", 5432))
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
