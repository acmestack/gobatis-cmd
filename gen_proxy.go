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
        builder.WriteString(`"github.com/xfali/gobatis/config"`)
        builder.WriteString(newline())
        //builder.WriteString(columnSpace())
        //builder.WriteString(`"github.com/xfali/gobatis/factory"`)
        //builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(`"github.com/xfali/gobatis/session/runner"`)
        builder.WriteString(newline())
        builder.WriteString(")")
        builder.WriteString(newline())
        builder.WriteString(newline())

        proxyName := fmt.Sprintf("%sCallProxy", modelName)
        builder.WriteString(fmt.Sprintf("type %s runner.RunnerSession", proxyName))
        builder.WriteString(newline())
        builder.WriteString(newline())

        builder.WriteString("func init() {")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("modelV := %s{}", modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("config.RegisterModel(&modelV)")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("config.RegisterMapperFile(\"%sxml/%s.xml\")", config.path,tableName))
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())

        builder.WriteString(fmt.Sprintf("func New(proxyMrg *runner.SessionManager) *%s {", proxyName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("return (*%s)(proxyMrg.NewSession())", proxyName))
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())

        //Tx
        builder.WriteString(fmt.Sprintf("func (proxy *%s) Tx(txFunc func(s *%s) bool) {", proxyName, proxyName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`sess := (*runner.RunnerSession)(proxy)`)
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(`sess.Tx(func(session *runner.RunnerSession) bool {`)
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
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Select%s(model %s) []%s {", proxyName, modelName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("var dataList []%s", modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`(*runner.RunnerSession)(proxy).Select("select%s").Param(model).Result(&dataList)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return dataList")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //select end

        //insert
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Insert%s(model %s) int64 {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`(*runner.RunnerSession)(proxy).Insert("insert%s").Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //insert end

        //update
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Update%s(model %s) int64 {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`(*runner.RunnerSession)(proxy).Update("update%s").Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //update end

        //delete
        builder.WriteString(fmt.Sprintf("func (proxy *%s)Delete%s(model %s) int64 {", proxyName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`(*runner.RunnerSession)(proxy).Delete("delete%s").Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //delete end

        io.Write(mapperFile, []byte(builder.String()))
    }
}
