/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import (
    "fmt"
    "github.com/xfali/gobatis-cmd/common"
    "github.com/xfali/gobatis-cmd/io"
    "strings"
)

func genProxy(config Config, tableName string, models []common.ModelInfo) {
    mapperDir := config.Path
    if !io.IsPathExists(mapperDir) {
        io.Mkdir(mapperDir)
    }
    mapperFile, err := io.OpenAppend(mapperDir + strings.ToLower(tableName) + "_proxy.go")
    if err == nil {
        defer mapperFile.Close()

        modelName := tableName2ModelName(tableName)
        builder := strings.Builder{}
        builder.WriteString("package ")
        builder.WriteString(config.PackageName)
        builder.WriteString(Newline())
        builder.WriteString(Newline())

        builder.WriteString("import (")
        builder.WriteString(Newline())
        builder.WriteString(ColumnSpace())
        builder.WriteString(`"context"`)
        builder.WriteString(Newline())
        builder.WriteString(ColumnSpace())
        builder.WriteString(`"github.com/xfali/gobatis"`)
        builder.WriteString(Newline())
        //builder.WriteString(ColumnSpace())
        //builder.WriteString(`"github.com/xfali/gobatis/factory"`)
        //builder.WriteString(Newline())
        //builder.WriteString(ColumnSpace())
        //builder.WriteString(`"github.com/xfali/gobatis/session/runner"`)
        //builder.WriteString(Newline())
        builder.WriteString(")")
        builder.WriteString(Newline())
        builder.WriteString(Newline())

        proxyName := fmt.Sprintf("%sCallProxy", modelName)
        builder.WriteString(fmt.Sprintf("type %s gobatis.Session", proxyName))
        builder.WriteString(Newline())
        builder.WriteString(Newline())

        builder.WriteString("func init() {")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf("modelV := %s{}", modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("gobatis.RegisterModel(&modelV)")
        builder.WriteString(Newline())

        if config.MapperFile == "xml" {
            builder.WriteString(ColumnSpace())
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperFile(\"%sxml/%s.xml\")", config.Path, tableName))
            builder.WriteString(Newline())
        } else if config.MapperFile == "go" {
            builder.WriteString(ColumnSpace())
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperData([]byte(%sMapper))", modelName))
            builder.WriteString(Newline())
        }

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())

        builder.WriteString(fmt.Sprintf("func New%s(proxyMrg *gobatis.SessionManager) *%s {", proxyName, proxyName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf("return (*%s)(proxyMrg.NewSession())", proxyName))
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())

        //Tx
        builder.WriteString(fmt.Sprintf("func (proxy *%s) Tx(txFunc func(s *%s) error) {", proxyName, proxyName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(`sess := (*gobatis.Session)(proxy)`)
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(`sess.Tx(func(session *gobatis.Session) error {`)
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(ColumnSpace())
        builder.WriteString("return txFunc(proxy)")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(`})`)
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //Tx end

        //select
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%s(model %s) ([]%s, error) {", proxyName, modelName, modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf("var dataList []%s", modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Select("select%s").Context(context.Background()).Param(model).Result(&dataList)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return dataList, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //select end

        //select count
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%sCount(model %s) (int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Select("select%sCount").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //select count end

        //insert
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Insert%s(model %s) (int64, int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`runner := (*gobatis.Session)(proxy).Insert("insert%s").Context(context.Background()).Param(model)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(`err := runner.Result(&ret)`)
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(`id := runner.LastInsertId()`)
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return ret, id, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //insert end

        //update
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Update%s(model %s) (int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Update("update%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //update end

        //delete
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Delete%s(model %s) (int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Delete("delete%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //delete end

        io.Write(mapperFile, []byte(builder.String()))
    }
}
