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

package pkg

import (
	"github.com/acmestack/gobatis-cmd/pkg/stringutils"
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
