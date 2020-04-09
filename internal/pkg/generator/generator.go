/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package generator

import (
    "github.com/xfali/gobatis-cmd/pkg/common"
    "github.com/xfali/gobatis-cmd/pkg/config"
    "github.com/xfali/gobatis-cmd/pkg/plugin"
    "log"
    "os"
)


func GenOneTable(config config.Config, db common.DBDriver, dbName, table string) {
    models, err := db.QueryTableInfo(dbName, table)
    if err != nil {
        log.Print(err)
        os.Exit(-3)
    }
    conf := config
    if conf.Namespace == "" {
       conf.Namespace = config.PackageName + "." + common.TableName2ModelName(table)
    }
    err2 := Generate(conf, models, table)
    if err2 != nil {
        log.Print(err2)
        os.Exit(-2)
    }
}

func Generate(config config.Config, models []common.ModelInfo, tableName string) (err error) {
    GenModel(config, tableName, models)
    if config.MapperFile == "template" {
        GenTemplate(config, tableName, models)
    } else {
        GenXml(config, tableName, models)
    }

    GenV2Proxy(config, tableName, models)

    return plugin.RunPlugin(config, tableName, models)
}

