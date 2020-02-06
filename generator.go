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
)

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

