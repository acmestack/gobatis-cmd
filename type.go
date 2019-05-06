/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

var sqlType2GoMap = map[string]string{
    "int":                "int",
    "integer":            "int",
    "tinyint":            "int",
    "smallint":           "int",
    "mediumint":          "int",
    "bigint":             "int",
    "int unsigned":       "int",
    "integer unsigned":   "int",
    "tinyint unsigned":   "int",
    "smallint unsigned":  "int",
    "mediumint unsigned": "int",
    "bigint unsigned":    "int",
    "bit":                "int",
    "bool":               "bool",
    "enum":               "string",
    "set":                "string",
    "varchar":            "string",
    "char":               "string",
    "tinytext":           "string",
    "mediumtext":         "string",
    "text":               "string",
    "longtext":           "string",
    "blob":               "string",
    "tinyblob":           "string",
    "mediumblob":         "string",
    "longblob":           "string",
    "date":               "time.Time",
    "datetime":           "time.Time",
    "timestamp":          "time.Time",
    "time":               "time.Time",
    "float":              "float64",
    "double":             "float64",
    "decimal":            "float64",
    "binary":             "string",
    "varbinary":          "string",
}
