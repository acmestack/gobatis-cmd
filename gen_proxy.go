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
    "github.com/xfali/gobatis-cmd/io"
    "strings"
)

func genProxy(config config, tableName string, models []modelInfo) {
    mapperDir := config.path
    if !io.IsPathExists(mapperDir) {
        io.Mkdir(mapperDir)
    }
    mapperFile, err := io.OpenAppend(mapperDir + tableName + "_proxy.go")
    if err == nil {
        defer mapperFile.Close()

        modelName := tableName2ModelName(tableName)
        builder := strings.Builder{}
        builder.WriteString("package ")
        builder.WriteString(config.packageName)
        builder.WriteString(newline())
        builder.WriteString(newline())

        builder.WriteString("import (")
        builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(`"context"`)
        builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(`"github.com/xfali/gobatis"`)
        builder.WriteString(newline())
        //builder.WriteString(columnSpace())
        //builder.WriteString(`"github.com/xfali/gobatis/factory"`)
        //builder.WriteString(newline())
        //builder.WriteString(columnSpace())
        //builder.WriteString(`"github.com/xfali/gobatis/session/runner"`)
        //builder.WriteString(newline())
        builder.WriteString(")")
        builder.WriteString(newline())
        builder.WriteString(newline())

        proxyName := fmt.Sprintf("%sCallProxy", modelName)
        builder.WriteString(fmt.Sprintf("type %s gobatis.Session", proxyName))
        builder.WriteString(newline())
        builder.WriteString(newline())

        builder.WriteString("func init() {")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("modelV := %s{}", modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("gobatis.RegisterModel(&modelV)")
        builder.WriteString(newline())

        if config.mapperFile == "xml" {
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperFile(\"%sxml/%s.xml\")", config.path, tableName))
            builder.WriteString(newline())
        } else if config.mapperFile == "go" {
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperData([]byte(%sMapper))", modelName))
            builder.WriteString(newline())
        }

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())

        builder.WriteString(fmt.Sprintf("func New%s(proxyMrg *gobatis.SessionManager) *%s {", proxyName, proxyName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("return (*%s)(proxyMrg.NewSession())", proxyName))
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())

        //Tx
        builder.WriteString(fmt.Sprintf("func (proxy *%s) Tx(txFunc func(s *%s) error) {", proxyName, proxyName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`sess := (*gobatis.Session)(proxy)`)
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`sess.Tx(func(session *gobatis.Session) error {`)
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString("return txFunc(proxy)")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`})`)
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //Tx end

        //select
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%s(model %s) ([]%s, error) {", proxyName, modelName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("var dataList []%s", modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Select("select%s").Context(context.Background()).Param(model).Result(&dataList)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return dataList, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //select end

        //select count
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%sCount(model %s) (int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Select("select%sCount").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //select count end

        //insert
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Insert%s(model %s) (int64, int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`runner := (*gobatis.Session)(proxy).Insert("insert%s").Context(context.Background()).Param(model)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`err := runner.Result(&ret)`)
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`id := runner.LastInsertId()`)
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret, id, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //insert end

        //update
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Update%s(model %s) (int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Update("update%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //update end

        //delete
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Delete%s(model %s) (int64, error) {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := (*gobatis.Session)(proxy).Delete("delete%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //delete end

        io.Write(mapperFile, []byte(builder.String()))
    }
}
