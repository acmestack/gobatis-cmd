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

package test

import (
	db2 "github.com/acmestack/gobatis-cmd/pkg/db"
	"log"
	"os"
	"testing"
)

func TestMysqlInfo(t *testing.T) {
	t.Log(db2.GenDBInfo("mysql", "testdb", "test", "test", "localhost", 3306))
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
