// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

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
