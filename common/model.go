/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package common

type ModelInfo struct {
    ColumnName string
    DataType   string
    Nullable   string
    ColumnKey  string
    Comment    string
    Tag        string
}


const (
    MethodFlag         = "method"
    OutPutSuffixMethod = "output"
    GenerateMethod     = "generate"
)

type GenerateInfo struct {
    Package string             `json:"package"`
    Table   string             `json:"table"`
    Models  []ModelInfo `json:"models"`
}

type PluginResult struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
