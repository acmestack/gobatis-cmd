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

func genXml(config Config, tableName string, model []common.ModelInfo) {
    if config.MapperFile == "xml" {
        xmlDir := config.Path + "xml/"
        if !io.IsPathExists(xmlDir) {
            io.Mkdir(xmlDir)
        }
        xmlFile, err := io.OpenAppend(xmlDir + strings.ToLower(tableName) + "_mapper.xml")
        if err == nil {
            defer xmlFile.Close()

            builder := strings.Builder{}
            buildMapper(&builder, config, tableName, model,
                formatXmlColumns, formatBackQuoteXml, formatBackQuoteXml)
            io.Write(xmlFile, []byte(builder.String()))
        }
    } else if config.MapperFile == "go" {
        xmlDir := config.Path
        if !io.IsPathExists(xmlDir) {
            io.Mkdir(xmlDir)
        }
        xmlFile, err := io.OpenAppend(xmlDir + strings.ToLower(tableName) + "_mapper.go")
        if err == nil {
            defer xmlFile.Close()

            builder := strings.Builder{}

            builder.WriteString("package ")
            builder.WriteString(config.PackageName)
            builder.WriteString(common.Newline())
            builder.WriteString(common.Newline())

            builder.WriteString(fmt.Sprintf("var %sMapper = `", common.TableName2ModelName(tableName)))
            builder.WriteString(common.Newline())

            buildMapper(&builder, config, tableName, model,
                formatGoColumns, formatBackQuoteGo, formatBackQuoteGo)
            builder.WriteString("`")
            builder.WriteString(common.Newline())

            io.Write(xmlFile, []byte(builder.String()))
        }
    }
}

type fomatter func(string) string

func buildMapper(builder *strings.Builder, config Config, tableName string, model []common.ModelInfo,
        columnsFunc func (string, []common.ModelInfo) string, tableFunc, columnFunc fomatter) {
    modelName := common.TableName2ModelName(tableName)
    builder.WriteString(fmt.Sprintf("<mapper namespace=\"%s.%s\">", config.PackageName, modelName))
    builder.WriteString(common.Newline())

    builder.WriteString(common.ColumnSpace())
    builder.WriteString("<sql id=\"columns_id\">")
    columns := columnsFunc(tableName, model)

    builder.WriteString(columns)
    builder.WriteString("</sql>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())

    tableName = tableFunc(tableName)

    //select
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("<select id=\"select%s\">", modelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("SELECT <include refid=\"columns_id\"> </include> FROM %s", tableName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("<where>")
    builder.WriteString(common.Newline())
    for _, f := range model {
        fieldName := common.Column2DynamicName(modelName, f.ColumnName)
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>",
            getIfStr(f.DataType, fieldName), columnFunc(f.ColumnName), fieldName))
        builder.WriteString(common.Newline())
    }
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</where>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</select>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    //select end

    //select count
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("<select id=\"select%sCount\">", modelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("<where>")
    builder.WriteString(common.Newline())
    for _, f := range model {
        fieldName := common.Column2DynamicName(modelName, f.ColumnName)
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>",
            getIfStr(f.DataType, fieldName), columnFunc(f.ColumnName), fieldName))
        builder.WriteString(common.Newline())
    }
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</where>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</select>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    //select count

    //insert
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("<insert id=\"insert%s\">", modelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("INSERT INTO %s (%s)", tableName, columns))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("VALUES(")
    builder.WriteString(common.Newline())
    for i := range model {
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        //builder.WriteString(fmt.Sprintf("#{%s}", common.Column2Modelfield(model[i].ColumnName)))
        builder.WriteString(fmt.Sprintf("#{%s}", common.Column2DynamicName(modelName, model[i].ColumnName)))
        if i < len(model)-1 {
            builder.WriteString(",")
        }
        builder.WriteString(common.Newline())
    }
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(")")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</insert>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    //insert end

    //update
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("<update id=\"update%s\">", modelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("UPDATE %s", tableName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("<set>")
    builder.WriteString(common.Newline())
    index := -1
    for i, f := range model {
        if strings.ToUpper(f.ColumnKey) == "PRI" {
            index = i
            continue
        }
        fieldName := common.Column2DynamicName(modelName, f.ColumnName)
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\"> %s = #{%s} </if>",
            getIfStr(f.DataType, fieldName), columnFunc(f.ColumnName), fieldName))
        builder.WriteString(common.Newline())
    }
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</set>")
    builder.WriteString(common.Newline())
    if index != -1 {
        f := model[index]
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(fmt.Sprintf("WHERE %s = #{%s}",
            columnFunc(f.ColumnName), common.Column2DynamicName(modelName, f.ColumnName)))
        builder.WriteString(common.Newline())
    }
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</update>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.Newline())
    //update end

    //delete
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("<delete id=\"delete%s\">", modelName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(fmt.Sprintf("DELETE FROM %s", tableName))
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("<where>")
    builder.WriteString(common.Newline())
    for _, f := range model {
        fieldName := common.Column2DynamicName(modelName, f.ColumnName)
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(common.ColumnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>",
            getIfStr(f.DataType, fieldName), columnFunc(f.ColumnName), fieldName))
        builder.WriteString(common.Newline())
    }
    builder.WriteString(common.ColumnSpace())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</where>")
    builder.WriteString(common.Newline())
    builder.WriteString(common.ColumnSpace())
    builder.WriteString("</delete>")
    builder.WriteString(common.Newline())
    //delete end

    builder.WriteString("</mapper>")
    builder.WriteString(common.Newline())
}

func getIfStr(ctype, name string) string {
    return strings.Replace(sqlType2IfFormatMap[ctype], "%s", fmt.Sprintf("{%s}", name), -1)
}

func formatGoColumns(tableName string, model []common.ModelInfo) string {
    columns := "` + \""
    for i := range model {
        columns += formatColumnName(tableName, model[i].ColumnName)
        if i < len(model)-1 {
            columns += ","
        }
    }
    columns += "\" + `"
    return columns
}

func formatXmlColumns(tableName string, model []common.ModelInfo) string {
    columns := ""
    tableName = fmt.Sprintf("`%s`", tableName)
    for i := range model {
        columns += formatColumnName(tableName, model[i].ColumnName)
        if i < len(model)-1 {
            columns += ","
        }
    }
    return columns
}

func formatBackQuoteGo(src string) string {
    return "` + \"`" + src + "`\" + `"
}

func formatBackQuoteXml(src string) string {
    return fmt.Sprintf("`%s`", src)
}

func formatColumnName(tableName, columnName string) string {
    return fmt.Sprintf("`%s`", columnName)
}
