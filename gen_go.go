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

func genGo(config config, tableName string, model []modelInfo) {
    if config.mapperFile == "go"{
        xmlDir := config.path
        if !io.IsPathExists(xmlDir) {
            io.Mkdir(xmlDir)
        }
        xmlFile, err := io.OpenAppend(xmlDir + tableName + "_mapper.go")
        if err == nil {
            defer xmlFile.Close()

            builder := strings.Builder{}

            builder.WriteString("package ")
            builder.WriteString(config.packageName)
            builder.WriteString(newline())
            builder.WriteString(newline())

            builder.WriteString(fmt.Sprintf("var %sMapper = `", tableName2ModelName(tableName)))
            builder.WriteString(newline())

            buildGoMapper(&builder, config, tableName, model)
            builder.WriteString("`")
            builder.WriteString(newline())

            io.Write(xmlFile, []byte(builder.String()))
        }
    }
}

func buildGoMapper(builder *strings.Builder, config config, tableName string, model []modelInfo) {
    modelName := tableName2ModelName(tableName)
    builder.WriteString(fmt.Sprintf("<mapper namespace=\"%s.%s\">", config.packageName, modelName))
    builder.WriteString(newline())

    builder.WriteString(columnSpace())
    builder.WriteString("<sql id=\"columns_id\">")
    columns := "` + \""
    for i := range model {
        columns += formatColumnName(tableName, model[i].columnName)
        if i < len(model)-1 {
            columns += ","
        }
    }
    columns += "\" + `"
    builder.WriteString(columns)
    builder.WriteString("</sql>")
    builder.WriteString(newline())
    builder.WriteString(newline())

    tableName = writeBackQuote(tableName)

    //select
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("<select id=\"select%s\">", modelName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("SELECT <include refid=\"columns_id\"> </include> FROM %s", tableName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("<where>")
    builder.WriteString(newline())
    for _, f := range model {
        fieldName := column2DynamicName(modelName, f.columnName)
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>",
            getIfStr(f.dataType, fieldName), writeBackQuote(f.columnName), fieldName))
        builder.WriteString(newline())
    }
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("</where>")
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString("</select>")
    builder.WriteString(newline())
    builder.WriteString(newline())
    //select end

    //select count
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("<select id=\"select%sCount\">", modelName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("<where>")
    builder.WriteString(newline())
    for _, f := range model {
        fieldName := column2DynamicName(modelName, f.columnName)
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>",
            getIfStr(f.dataType, fieldName), writeBackQuote(f.columnName), fieldName))
        builder.WriteString(newline())
    }
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("</where>")
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString("</select>")
    builder.WriteString(newline())
    builder.WriteString(newline())
    //select count

    //insert
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("<insert id=\"insert%s\">", modelName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("INSERT INTO %s (%s)", tableName, columns))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("VALUES(")
    builder.WriteString(newline())
    for i := range model {
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        //builder.WriteString(fmt.Sprintf("#{%s}", column2Modelfield(model[i].columnName)))
        builder.WriteString(fmt.Sprintf("#{%s}", column2DynamicName(modelName, model[i].columnName)))
        if i < len(model)-1 {
            builder.WriteString(",")
        }
        builder.WriteString(newline())
    }
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString(")")
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString("</insert>")
    builder.WriteString(newline())
    builder.WriteString(newline())
    //insert end

    //update
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("<update id=\"update%s\">", modelName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("UPDATE %s", tableName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("<set>")
    builder.WriteString(newline())
    index := -1
    for i, f := range model {
        if strings.ToUpper(f.columnKey) == "PRI" {
            index = i
            continue
        }
        fieldName := column2DynamicName(modelName, f.columnName)
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\"> %s = #{%s} </if>",
            getIfStr(f.dataType, fieldName), writeBackQuote(f.columnName), fieldName))
        builder.WriteString(newline())
    }
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("</set>")
    builder.WriteString(newline())
    if index != -1 {
        f := model[index]
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("WHERE %s = #{%s}",
            writeBackQuote(f.columnName), column2DynamicName(modelName, f.columnName)))
        builder.WriteString(newline())
    }
    builder.WriteString(columnSpace())
    builder.WriteString("</update>")
    builder.WriteString(newline())
    builder.WriteString(newline())
    //update end

    //delete
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("<delete id=\"delete%s\">", modelName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString(fmt.Sprintf("DELETE FROM %s", tableName))
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("<where>")
    builder.WriteString(newline())
    for _, f := range model {
        fieldName := column2DynamicName(modelName, f.columnName)
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>",
            getIfStr(f.dataType, fieldName), writeBackQuote(f.columnName), fieldName))
        builder.WriteString(newline())
    }
    builder.WriteString(columnSpace())
    builder.WriteString(columnSpace())
    builder.WriteString("</where>")
    builder.WriteString(newline())
    builder.WriteString(columnSpace())
    builder.WriteString("</delete>")
    builder.WriteString(newline())
    //delete end

    builder.WriteString("</mapper>")
    builder.WriteString(newline())
}

func getIfStr(ctype, name string) string {
    return strings.Replace(sqlType2IfFormatMap[ctype], "%s", fmt.Sprintf("{%s}", name), -1)
}

func writeBackQuote(src string) string {
    return "` + \"`" + src + "`\" + `"
}

func formatColumnName(tableName, columnName string) string {
    return fmt.Sprintf("`%s`", columnName)
}
