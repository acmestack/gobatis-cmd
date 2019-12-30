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

func genV2Proxy(config Config, tableName string, models []common.ModelInfo) {
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
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperFile(\"%sxml/%s_mapper.xml\")", config.Path, strings.ToLower(tableName)))
            builder.WriteString(Newline())
        } else if config.MapperFile == "go" {
            builder.WriteString(ColumnSpace())
            builder.WriteString(fmt.Sprintf("gobatis.RegisterMapperData([]byte(%sMapper))", modelName))
            builder.WriteString(Newline())
        }

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())

        //select
        builder.WriteString(fmt.Sprintf("func Select%s(sess *gobatis.Session, model %s) ([]%s, error) {", modelName, modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf("var dataList []%s", modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Select("select%s").Param(model).Result(&dataList)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return dataList, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //select end

        //select count
        builder.WriteString(fmt.Sprintf("func Select%sCount(sess *gobatis.Session, model %s) (int64, error) {", modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Select("select%sCount").Param(model).Result(&ret)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //select count end

        //insert
        builder.WriteString(fmt.Sprintf("func Insert%s(sess *gobatis.Session, model %s) (int64, int64, error) {", modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`runner := sess.Insert("insert%s").Param(model)`, modelName))
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
        builder.WriteString(fmt.Sprintf("func Update%s(sess *gobatis.Session, model %s) (int64, error) {", modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Update("update%s").Param(model).Result(&ret)`, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("return ret, err")
        builder.WriteString(Newline())

        builder.WriteString("}")
        builder.WriteString(Newline())
        builder.WriteString(Newline())
        //update end

        //delete
        builder.WriteString(fmt.Sprintf("func Delete%s(sess *gobatis.Session, model %s) (int64, error) {", modelName, modelName))
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString("var ret int64")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf(`err := sess.Delete("delete%s").Param(model).Result(&ret)`, modelName))
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
