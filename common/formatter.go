// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package common

type KeyWordFormatter func(string) string

var KwFormatter = MysqlKeyWordFormatter
var kwFormatterMap = map[string]KeyWordFormatter{
	"mysql":    MysqlKeyWordFormatter,
	"postgres": PostgresKeyWordFormatter,
}

func MysqlKeyWordFormatter(src string) string {
	return "`" + src + "`"
}

func PostgresKeyWordFormatter(src string) string {
	return `"` + src + `"`
}

func SelectKeyWordFormatter(driver string) {
	v := kwFormatterMap[driver]
	if v != nil {
		KwFormatter = v
	}
}
