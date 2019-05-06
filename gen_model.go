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

func findTime(model []modelInfo) bool {
    for _, info := range model {
        if info.dataType == "date" || info.dataType == "datetime" || info.dataType == "timestamp" || info.dataType == "time" {
            return true
        }
    }
    return false
}

func genModel(config config, tableName string, model []modelInfo) {
    modelDir := config.path
    if !io.IsPathExists(modelDir) {
        io.Mkdir(modelDir)
    }
    exist := io.IsPathExists(modelDir + config.modelFile)
    modelFile, err := io.OpenAppend(modelDir + config.modelFile)
    if err == nil {
        defer modelFile.Close()

        modelName := tableName2ModelName(tableName)
        builder := strings.Builder{}
        if !exist {
            builder.WriteString("package ")
            builder.WriteString(config.packageName)
            builder.WriteString(newline())
            builder.WriteString(newline())

            if findTime(model) {
                builder.WriteString("import \"time\"")
                builder.WriteString(newline())
                builder.WriteString(newline())
            }
        }

        builder.WriteString("type ")
        builder.WriteString(modelName)
        builder.WriteString(" struct {")
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("//TableName gobatis.ModelName `%s`", tableName))
        builder.WriteString(newline())

        for _, info := range model {
            builder.WriteString(columnSpace())
            builder.WriteString(column2Modelfield(info.columnName))
            builder.WriteString(" ")
            builder.WriteString(sqlType2GoMap[info.dataType])
            builder.WriteString(" ")
            builder.WriteString(fmt.Sprintf("`%s:\"%s\"`", config.tagName, info.columnName))
            builder.WriteString(newline())
        }
        builder.WriteString("}")
        builder.WriteString(newline())

        io.Write(modelFile, []byte(builder.String()))
    }
}
