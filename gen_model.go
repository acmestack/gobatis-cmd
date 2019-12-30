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

func findTime(model []common.ModelInfo) bool {
    for _, info := range model {
        if info.DataType == "date" || info.DataType == "datetime" || info.DataType == "timestamp" || info.DataType == "time" {
            return true
        }
    }
    return false
}

func genModel(config Config, tableName string, model []common.ModelInfo) {
    modelDir := config.Path
    if !io.IsPathExists(modelDir) {
        io.Mkdir(modelDir)
    }
    fileName := config.ModelFile
    if fileName == "" {
        fileName = strings.ToLower(tableName) + ".go"
    }
    exist := io.IsPathExists(modelDir + fileName)
    modelFile, err := io.OpenAppend(modelDir + fileName)
    if err == nil {
        defer modelFile.Close()

        modelName := tableName2ModelName(tableName)
        builder := strings.Builder{}
        if !exist {
            builder.WriteString("package ")
            builder.WriteString(config.PackageName)
            builder.WriteString(Newline())
            builder.WriteString(Newline())

            if findTime(model) {
                builder.WriteString("import \"time\"")
                builder.WriteString(Newline())
                builder.WriteString(Newline())
            }
        }

        builder.WriteString("type ")
        builder.WriteString(modelName)
        builder.WriteString(" struct {")
        builder.WriteString(Newline())

        builder.WriteString(ColumnSpace())
        builder.WriteString(fmt.Sprintf("//TableName gobatis.ModelName `%s`", tableName))
        builder.WriteString(Newline())

        for _, info := range model {
            builder.WriteString(ColumnSpace())
            builder.WriteString(column2Modelfield(info.ColumnName))
            builder.WriteString(" ")
            builder.WriteString(sqlType2GoMap[info.DataType])
            builder.WriteString(" ")
            //builder.WriteString(fmt.Sprintf("`%s:\"%s\"`", config.TagName, info.ColumnName))
            writeTag(&builder, config.TagName, info.ColumnName)
            builder.WriteString(Newline())
        }
        builder.WriteString("}")
        builder.WriteString(Newline())

        io.Write(modelFile, []byte(builder.String()))
    }
}

const(
    defaultTag = "xfield"
)

func writeTag(b *strings.Builder, tagName, columnName string) string {
    oriTags := strings.Split(tagName, ",")
    var tags []string
    found := false
    for _, v := range oriTags {
        if v != "" {
            if v == defaultTag {
                found = true
            }
            tags = append(tags, v)
        }
    }

    if !found {
        tags = append(tags, defaultTag)
    }

    l := len(tags)
    b.WriteString("`")
    for _, tag := range tags {
        if tag == "" {
            continue
        }
        b.WriteString(fmt.Sprintf("%s:\"%s\"", tag, columnName))
        l--
        if l > 0 {
            b.WriteString(" ")
        }
    }
    b.WriteString("`")
    return b.String()
}
