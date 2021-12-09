/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

package stringutils

import (
	"bytes"
	"strings"
)

// snake string, XxYy to xx_yy , XxYY to xx_yy
func Camel2snake(s string) string {
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
func Snake2camel(s string) string {
	s = strings.Trim(s, "_")
	if strings.Index(s, "_") != -1 {
	   return Snake2camel3(s)
	} else {
	   d := []byte(s)
	   if len(d) > 0 {
	       if d[0] >= 'a' && d[0] <= 'z' {
	           d[0] -= 32
	           return string(d)
	       }
	   }
	}
	return s
}

func Snake2camel2(s string) string {
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

func Snake2camel3(s string) string {
    strs := strings.Split(s, "_")
    buf := bytes.Buffer{}
    buf.Grow(len(s))
    for _, v := range strs {
        if v == "" {
            continue
        }
        // 第一个字符必须大写
        d := v[0]
        if d >= 'a' && d <= 'z' {
            d -= 32
        }
        buf.WriteByte(d)
        low := false
        for i := 1; i < len(v); i++ {
            d := v[i]
            // 中间夹杂小写
            if d >= 'a' && d <= 'z' {
                low = true
            } else if d >= 'A' && d <= 'Z' && !low {
                // 大写且中间没有经过小写，则表明是连续大写，则改写为小写
                buf.WriteByte(d + 32)
                continue
            }
            buf.WriteByte(d)
        }
    }
    return buf.String()
}
