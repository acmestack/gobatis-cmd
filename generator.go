/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import (
    "github.com/xfali/gobatis-cmd/common"
    "log"
    "os"
)


func genOneTable(config Config, db common.DBDriver, dbName, table string) {
    models, err := db.QueryTableInfo(dbName, table)
    if err != nil {
        log.Print(err)
        os.Exit(-3)
    }
    //conf := config
    //if conf.Namespace == "" {
    //    conf.Namespace = config.PackageName + "." + common.TableName2ModelName(table)
    //}
    err2 := generate(config, models, table)
    if err2 != nil {
        log.Print(err2)
        os.Exit(-2)
    }
}

func generate(config Config, models []common.ModelInfo, tableName string) (err error) {
    genModel(config, tableName, models)
    if config.MapperFile == "template" {
        genTemplate(config, tableName, models)
    } else {
        genXml(config, tableName, models)
    }

    genV2Proxy(config, tableName, models)

    return RunPlugin(config, tableName, models)
}

