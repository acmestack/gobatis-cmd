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

package generator

import (
	"fmt"
	"github.com/acmestack/gobatis-cmd/pkg/common"
	"github.com/acmestack/gobatis-cmd/pkg/config"
	"github.com/acmestack/gobatis-cmd/pkg/io"
	"strings"
)

func GenProxy(config config.Config, tableName string, models []common.ModelInfo) {
	mapperDir := config.Path
	if !io.IsPathExists(mapperDir) {
		io.Mkdir(mapperDir)
	}
	mapperFile, err := io.OpenAppend(mapperDir + strings.ToLower(tableName) + "_proxy.go")
	if err == nil {
		defer mapperFile.Close()

		modelName := common.TableName2ModelName(tableName)
		builder := strings.Builder{}
		builder.WriteString("package ")
		builder.WriteString(config.PackageName)
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())

		builder.WriteString("import (")
		builder.WriteString(common.Newline())
		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`"context"`)
		builder.WriteString(common.Newline())
		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`"github.com/acmestack/gobatis"`)
		builder.WriteString(common.Newline())
		//builder.WriteString(common.ColumnSpace())
		//builder.WriteString(`"github.com/acmestack/gobatis/factory"`)
		//builder.WriteString(common.Newline())
		//builder.WriteString(common.ColumnSpace())
		//builder.WriteString(`"github.com/acmestack/gobatis/session/runner"`)
		//builder.WriteString(common.Newline())
		builder.WriteString(")")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())

		proxyName := fmt.Sprintf("%sCallProxy", modelName)
		builder.WriteString(fmt.Sprintf("type %s gobatis.Session", proxyName))
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())

		builder.WriteString("func init() {")
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf("modelV := %s{}", modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("gobatis.RegisterModel(&modelV)")
		builder.WriteString(common.Newline())

		if config.MapperFile == "xml" {
			builder.WriteString(common.ColumnSpace())
			builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperFile(\"%sxml/%s.xml\")", config.Path, tableName))
			builder.WriteString(common.Newline())
		} else if config.MapperFile == "go" {
			builder.WriteString(common.ColumnSpace())
			builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperData([]byte(%sMapper))", modelName))
			builder.WriteString(common.Newline())
		}

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())

		builder.WriteString(fmt.Sprintf("func New%s(proxyMrg *gobatis.SessionManager) *%s {", proxyName, proxyName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf("return (*%s)(proxyMrg.NewSession())", proxyName))
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())

		//Tx
		builder.WriteString(fmt.Sprintf("func (proxy *%s) Tx(txFunc func(s *%s) error) {", proxyName, proxyName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`sess := (*gobatis.Session)(proxy)`)
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`sess.Tx(func(session *gobatis.Session) error {`)
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(common.ColumnSpace())
		builder.WriteString("return txFunc(proxy)")
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`})`)
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())
		//Tx end

		//select
		builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%s(model %s) ([]%s, error) {", proxyName, modelName, modelName, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf("var dataList []%s", modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Select("select%s").Context(context.Background()).Param(model).Result(&dataList)`, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("return dataList, err")
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())
		//select end

		//select count
		builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%sCount(model %s) (int64, error) {", proxyName, modelName, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("var ret int64")
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Select("select%sCount").Context(context.Background()).Param(model).Result(&ret)`, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("return ret, err")
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())
		//select count end

		//insert
		builder.WriteString(fmt.Sprintf("func (proxy *%s)Insert%s(model %s) (int64, int64, error) {", proxyName, modelName, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("var ret int64")
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf(`runner := (*gobatis.Session)(proxy).Insert("insert%s").Context(context.Background()).Param(model)`, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`err := runner.Result(&ret)`)
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(`id := runner.LastInsertId()`)
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("return ret, id, err")
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())
		//insert end

		//update
		builder.WriteString(fmt.Sprintf("func (proxy *%s)Update%s(model %s) (int64, error) {", proxyName, modelName, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("var ret int64")
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Update("update%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("return ret, err")
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())
		//update end

		//delete
		builder.WriteString(fmt.Sprintf("func (proxy *%s)Delete%s(model %s) (int64, error) {", proxyName, modelName, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("var ret int64")
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Delete("delete%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
		builder.WriteString(common.Newline())

		builder.WriteString(common.ColumnSpace())
		builder.WriteString("return ret, err")
		builder.WriteString(common.Newline())

		builder.WriteString("}")
		builder.WriteString(common.Newline())
		builder.WriteString(common.Newline())
		//delete end

		io.Write(mapperFile, []byte(builder.String()))
	}
}
