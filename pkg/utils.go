/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

package pkg

import (
	"github.com/xfali/gobatis-cmd/pkg/stringutils"
	"strings"
)

func Newline() string {
	return "\n"
}

func ColumnSpace() string {
	return "    "
}

func TableName2ModelName(tableName string) string {
	return Snake2camel(strings.ToLower(tableName))
}

func Column2Modelfield(column string) string {
	return Snake2camel(strings.ToLower(column))
}

func Column2DynamicName(tableName, column string) string {
	return tableName + "." + column
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func Camel2snake(s string) string {
	return stringutils.Camel2snake(s)
}

// camel string, xx_yy to XxYy
func Snake2camel(s string) string {
	return stringutils.Snake2camel(s)
}
