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

func genXml(config config, tableName string, model []modelInfo) {
    xmlDir := config.path + "xml/"
    if !io.IsPathExists(xmlDir) {
        io.Mkdir(xmlDir)
    }
    xmlFile, err := io.OpenAppend(xmlDir + tableName + ".xml")
    if err == nil {
        defer xmlFile.Close()
        modelName := tableName2ModelName(tableName)
        builder := strings.Builder{}
        builder.WriteString(fmt.Sprintf("<mapper namespace=\"%s.%s\">", config.packageName, modelName))
        builder.WriteString(newline())

        builder.WriteString(columnSpace())
        builder.WriteString("<sql id=\"columns_id\">")
        columns := ""
        for i := range model {
            columns += model[i].columnName
            if i < len(model) - 1 {
                columns += ","
            }
        }
        builder.WriteString(columns)
        builder.WriteString("</sql>")
        builder.WriteString(newline())
        builder.WriteString(newline())

        //select
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("<select id=\"select%s\">",modelName))
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
            fieldName := f.columnName//column2Modelfield(f.columnName)
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>", getIfStr(f.dataType, fieldName), f.columnName, fieldName))
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
        builder.WriteString(fmt.Sprintf("<select id=\"select%sCount\">",modelName))
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
            fieldName := f.columnName//column2Modelfield(f.columnName)
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>", getIfStr(f.dataType, fieldName), f.columnName, fieldName))
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
        builder.WriteString(fmt.Sprintf("<insert id=\"insert%s\">",modelName))
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
            builder.WriteString(fmt.Sprintf("#{%s}", model[i].columnName))
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
        builder.WriteString(fmt.Sprintf("<update id=\"update%s\">",modelName))
        builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("UPDATE %s", tableName))
        builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString("<set>")
        builder.WriteString(newline())
        for _, f := range model {
            fieldName := f.columnName//column2Modelfield(f.columnName)
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("<if test=\"%s\"> %s = #{%s} </if>", getIfStr(f.dataType, fieldName), f.columnName, fieldName))
            builder.WriteString(newline())
        }
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString("</set>")
        builder.WriteString(newline())
        for _, f := range model {
            if strings.ToUpper(f.columnKey) == "PRI" {
                builder.WriteString(columnSpace())
                builder.WriteString(columnSpace())
                builder.WriteString(fmt.Sprintf("WHERE %s = #{%s}", f.columnName, f.columnName))
                builder.WriteString(newline())
                break
            }
        }
        builder.WriteString(columnSpace())
        builder.WriteString("</update>")
        builder.WriteString(newline())
        builder.WriteString(newline())
        //update end

        //delete
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("<delete id=\"delete%s\">",modelName))
        builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString(fmt.Sprintf("DELETE FROM %s",tableName))
        builder.WriteString(newline())
        builder.WriteString(columnSpace())
        builder.WriteString(columnSpace())
        builder.WriteString("<where>")
        builder.WriteString(newline())
        for _, f := range model {
            fieldName := f.columnName//column2Modelfield(f.columnName)
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(columnSpace())
            builder.WriteString(fmt.Sprintf("<if test=\"%s\">AND %s = #{%s} </if>", getIfStr(f.dataType, fieldName), f.columnName, fieldName))
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
        io.Write(xmlFile, []byte(builder.String()))
    }
}

func getIfStr(ctype, name string) string {
    return strings.Replace(sqlType2IfFormatMap[ctype], "%s", name, -1)
}