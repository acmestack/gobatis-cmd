// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package common

const (
	MysqlKeywordEscapeChar    = "`"
	PostgresKeywordEscapeChar = `"`

	MysqlEscapeKeywordEscapeChar    = "`"
	PostgresEscapeKeywordEscapeChar = `\"`
)

var KeywordEscapeChar = MysqlKeywordEscapeChar
var EscapeKeywordEscapeChar = MysqlEscapeKeywordEscapeChar

type KeywordFormatter func(string) string

var KwFormatter = CommonKeywordFormatter

func CommonKeywordFormatter(src string) string {
	return KeywordEscapeChar + src + KeywordEscapeChar
}

func CommonEscapeKeywordFormatter(src string) string {
	return EscapeKeywordEscapeChar + src + EscapeKeywordEscapeChar
}

func SelectKeywordFormatter(driver string) {
	if driver == "postgres" {
		KeywordEscapeChar = PostgresKeywordEscapeChar
		EscapeKeywordEscapeChar = PostgresEscapeKeywordEscapeChar
	} else {
		KeywordEscapeChar = MysqlKeywordEscapeChar
		EscapeKeywordEscapeChar = MysqlEscapeKeywordEscapeChar
	}
}
