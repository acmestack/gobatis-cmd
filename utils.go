/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import "strings"

func newline() string {
    return "\n"
}

func columnSpace() string {
    return "    "
}

func tableName2ModelName(tablename string) string {
    return snake2camel(tablename)
}

func column2Modelfield(column string) string {
    return snake2camel(column)
}

func column2DynamicName(tableName, column string) string {
    return tableName + "." + column
}

// snake string, XxYy to xx_yy , XxYY to xx_yy
func camel2snake(s string) string {
    data := make([]byte, 0, len(s)*2)
    j := false
    num := len(s)
    for i := 0; i < num; i++ {
        d := s[i]
        if i > 0 && d >= 'A' && d <= 'Z' && j {
            data = append(data, '_')
        }
        if d != '_' {
            j = true
        }
        data = append(data, d)
    }
    return strings.ToLower(string(data))
}

// camel string, xx_yy to XxYy
func snake2camel(s string) string {
    data := make([]byte, 0, len(s))
    j := false
    k := false
    num := len(s) - 1
    for i := 0; i <= num; i++ {
        d := s[i]
        if k == false && d >= 'A' && d <= 'Z' {
            k = true
        }
        if d >= 'a' && d <= 'z' && (j || k == false) {
            d = d - 32
            j = false
            k = true
        }
        if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
            j = true
            continue
        }
        data = append(data, d)
    }
    return string(data)
}
