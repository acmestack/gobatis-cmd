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

func genTemplate(config Config, tableName string, model []common.ModelInfo) {
	targetDir := config.Path + "template/"
	if !io.IsPathExists(targetDir) {
		io.Mkdir(targetDir)
	}
	targetFile, err := io.OpenAppend(targetDir + strings.ToLower(tableName) + "_mapper.tmpl")
	if err == nil {
		defer targetFile.Close()

		builder := strings.Builder{}
		buildTmplMapper(&builder, config, tableName, model)
		io.Write(targetFile, []byte(builder.String()))
	}
}

func buildTmplMapper(builder *strings.Builder, config Config, tableName string, model []common.ModelInfo) {
	modelName := common.TableName2ModelName(tableName)
    columns := formatXmlColumns(tableName, model)
    tableName = formatBackQuoteXml(tableName)

    //select
	builder.WriteString(fmt.Sprintf(`{{define "select%s"}}`, modelName))
	builder.WriteString(common.Newline())

    builder.WriteString(fmt.Sprintf(`{{$COLUMNS := "%s"}}`, columns))
    builder.WriteString(common.Newline())

    builder.WriteString(fmt.Sprintf(`SELECT {{$COLUMNS}} FROM %s`, tableName))
    builder.WriteString(common.Newline())

    builder.WriteString(genTmplWhere(modelName, model))
    builder.WriteString(common.Newline())

    builder.WriteString(`{{end}}`)
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())

    //insert
    builder.WriteString(fmt.Sprintf(`{{define "insert%s"}}`, modelName))
    builder.WriteString(common.Newline())

    builder.WriteString(fmt.Sprintf(`{{$COLUMNS := "%s"}}`, columns))
    builder.WriteString(common.Newline())

    builder.WriteString(fmt.Sprintf(`INSERT INTO %s({{$COLUMNS}})`, tableName))
    builder.WriteString(common.Newline())

    builder.WriteString("VALUES(")
    builder.WriteString(common.Newline())

    builder.WriteString(genTmplValues(modelName, model))

    builder.WriteString(")")
    builder.WriteString(common.Newline())

    builder.WriteString(`{{end}}`)
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())

    //update
    builder.WriteString(fmt.Sprintf(`{{define "update%s"}}`, modelName))
    builder.WriteString(common.Newline())

    builder.WriteString(fmt.Sprintf(`UPDATE %s`, tableName))
    builder.WriteString(common.Newline())

    setStr, index := genTmplSet(modelName, model)
    builder.WriteString(setStr)
    builder.WriteString(common.Newline())

    if index != -1 {
        builder.WriteString(genTmplWhere(modelName, model[index:index+1]))
        builder.WriteString(common.Newline())
    }

    builder.WriteString(`{{end}}`)
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())

    //delete
    builder.WriteString(fmt.Sprintf(`{{define "delete%s"}}`, modelName))
    builder.WriteString(common.Newline())

    builder.WriteString(fmt.Sprintf(`DELETE FROM %s`, tableName))
    builder.WriteString(common.Newline())

    builder.WriteString(genTmplWhere(modelName, model))
    builder.WriteString(common.Newline())

    builder.WriteString(`{{end}}`)
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
}

func genTmplWhere(modelName string, model []common.ModelInfo) string {
    builder := strings.Builder{}

    builder.WriteString("{{")
    for i := range model {
        field := common.Column2Modelfield(model[i].ColumnName)
        if i == 0 {
            builder.WriteString(fmt.Sprintf(`where (ne .%s %s) "AND" "%s" .%s ""`, field, getTmplCond(model[i].DataType), model[i].ColumnName, field))
        } else {
            builder.WriteString(fmt.Sprintf(` | where (ne .%s %s) "AND" "%s" .%s`, field, getTmplCond(model[i].DataType), model[i].ColumnName, field))
        }
    }
    builder.WriteString("}}")

    return builder.String()
}

func genTmplSet(modelName string, model []common.ModelInfo) (string, int) {
    builder := strings.Builder{}

    index := -1
    builder.WriteString("{{")
    for i := range model {
        field := common.Column2Modelfield(model[i].ColumnName)
        if i == 0 {
            builder.WriteString(fmt.Sprintf(`set (ne .%s %s) "%s" .%s ""`, field, getTmplCond(model[i].DataType), model[i].ColumnName, field))
        } else {
            builder.WriteString(fmt.Sprintf(` | set (ne .%s %s) "%s" .%s`, field, getTmplCond(model[i].DataType), model[i].ColumnName, field))
        }
        if strings.ToUpper(model[i].ColumnKey) == "PRI" {
            index = i
            continue
        }
    }
    builder.WriteString("}}")

    return builder.String(), index
}

func genTmplValues(modelName string, model []common.ModelInfo) string {
    builder := strings.Builder{}

    size := len(model)
    for i := range model {
        if sqlType2GoMap[model[i].DataType] == "string" {
            builder.WriteString(fmt.Sprintf("'{{.%s}}'", common.Column2Modelfield(model[i].ColumnName)))
        } else {
            builder.WriteString(fmt.Sprintf("{{.%s}}", common.Column2Modelfield(model[i].ColumnName)))
        }

        size--
        if size > 0 {
            builder.WriteString(", ")
        }
    }

    return builder.String()
}

func getTmplCond(ctype string) string {
    return sqlType2IfCondMap[ctype]
}
