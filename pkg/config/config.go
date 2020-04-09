// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description: 

package config

type Config struct {
    Driver      string `json:"driver"`
    Path        string `json:"path"`
    PackageName string `json:"package"`
    Namespace   string `json:"namespace"`
    ModelFile   string `json:"modelFile"`
    TagName     string `json:"tagName"`
    MapperFile  string `json:"mapperFile"`
    Plugin      string `json:"plugin"`
    Keyword     bool   `json:"keyword"`
}

type FileConfig struct {
    Config
    TableName string `json:"tableName"`
    DBName    string `json:"dbName"`
    Host      string `json:"host"`
    Port      int    `json:"port"`
    User      string `json:"user"`
    Password  string `json:"password"`
}
