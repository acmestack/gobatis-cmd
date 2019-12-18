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

func genV2Proxy(config config, tableName string, models []modelInfo) {
    mapperDir := config.path
    if !io.IsPathExists(mapperDir) {
        io.Mkdir(mapperDir)
    }
    mapperFile, err := io.OpenAppend(mapperDir + strings.ToLower(tableName) + "_proxy.go")
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
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperFile(\"%sxml/%s_mapper.xml\")", config.path, strings.ToLower(tableName)))
            builder.WriteString(newline())
        } else if config.mapperFile == "go" {
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperData([]byte(%sMapper))", modelName))
            builder.WriteString(newline())
        }

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())

        //select
        builder.WriteString(fmt.Sprintf("func Select%s(sess *gobatis.Session, model %s) ([]%s, error) {", modelName, modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("var dataList []%s", modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Select("select%s").Context(context.Background()).Param(model).Result(&dataList)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return dataList, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //select end

        //select count
        builder.WriteString(fmt.Sprintf("func Select%sCount(sess *gobatis.Session, model %s) (int64, error) {", modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Select("select%sCount").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //select count end

        //insert
        builder.WriteString(fmt.Sprintf("func Insert%s(sess *gobatis.Session, model %s) (int64, int64, error) {", modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`runner := sess.Insert("insert%s").Context(context.Background()).Param(model)`, modelName))
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
        builder.WriteString(fmt.Sprintf("func Update%s(sess *gobatis.Session, model %s) (int64, error) {", modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Update("update%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(newline())

        builder.WriteString("}")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //update end

        //delete
        builder.WriteString(fmt.Sprintf("func Delete%s(sess *gobatis.Session, model %s) (int64, error) {", modelName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Delete("delete%s").Context(context.Background()).Param(model).Result(&ret)`, modelName))
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
