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

package common

const (
	MysqlKeywordEscapeCharStart       = "`"
	MysqlKeywordEscapeCharEnd         = "`"
	MysqlEscapeKeywordEscapeCharStart = "`"
	MysqlEscapeKeywordEscapeCharEnd   = "`"

	PostgresKeywordEscapeCharStart       = `"`
	PostgresKeywordEscapeCharEnd         = `"`
	PostgresEscapeKeywordEscapeCharStart = `\"`
	PostgresEscapeKeywordEscapeCharEnd   = `\"`

	SqlServerKeywordEscapeCharStart       = `[`
	SqlServerKeywordEscapeCharEnd         = `]`
	SqlServerEscapeKeywordEscapeCharStart = `[`
	SqlServerEscapeKeywordEscapeCharEnd   = `]`

	DummyKeywordEscapeCharStart       = ""
	DummyKeywordEscapeCharEnd         = ""
	DummyEscapeKeywordEscapeCharStart = ""
	DummyEscapeKeywordEscapeCharEnd   = ""
)

var KeywordEscapeCharStart = DummyKeywordEscapeCharStart
var KeywordEscapeCharEnd = DummyKeywordEscapeCharEnd
var EscapeKeywordEscapeCharStart = DummyEscapeKeywordEscapeCharStart
var EscapeKeywordEscapeCharEnd = DummyEscapeKeywordEscapeCharEnd

type KeywordFormatter func(string) string

var KwFormatter = CommonKeywordFormatter

func CommonKeywordFormatter(src string) string {
	return KeywordEscapeCharStart + src + KeywordEscapeCharEnd
}

func CommonEscapeKeywordFormatter(src string) string {
	return EscapeKeywordEscapeCharStart + src + EscapeKeywordEscapeCharEnd
}

func SelectKeywordFormatter(driver string) {
	if driver == "mysql" {
		KeywordEscapeCharStart = MysqlKeywordEscapeCharStart
		KeywordEscapeCharEnd = MysqlKeywordEscapeCharEnd
		EscapeKeywordEscapeCharStart = MysqlEscapeKeywordEscapeCharStart
		EscapeKeywordEscapeCharEnd = MysqlEscapeKeywordEscapeCharEnd
	} else if driver == "mssql" || driver == "adodb" {
		KeywordEscapeCharStart = SqlServerKeywordEscapeCharStart
		KeywordEscapeCharEnd = SqlServerKeywordEscapeCharEnd
		EscapeKeywordEscapeCharStart = SqlServerEscapeKeywordEscapeCharStart
		EscapeKeywordEscapeCharEnd = SqlServerEscapeKeywordEscapeCharEnd
	} else {
		KeywordEscapeCharStart = PostgresKeywordEscapeCharStart
		KeywordEscapeCharEnd = PostgresKeywordEscapeCharEnd
		EscapeKeywordEscapeCharStart = PostgresEscapeKeywordEscapeCharStart
		EscapeKeywordEscapeCharEnd = PostgresEscapeKeywordEscapeCharEnd
	}
}
