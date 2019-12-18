/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package main

import "testing"

func createModeInfo() *[]modelInfo {
    model := []modelInfo{}
    info := modelInfo{
        columnName: "id",
        dataType:   "bigint",
        nullable:   "NO",
        comment:    "id",
        tag:        "id",
        columnKey:  "PRI",
    }
    model = append(model, info)

    info = modelInfo{
        columnName: "username",
        dataType:   "varchar",
        nullable:   "NO",
        comment:    "username",
        tag:        "username",
    }
    model = append(model, info)

    info = modelInfo{
        columnName: "password",
        dataType:   "varchar",
        nullable:   "NO",
        comment:    "password",
        tag:        "password",
    }
    model = append(model, info)

    info = modelInfo{
        columnName: "update_time",
        dataType:   "timestamp",
        nullable:   "YES",
        comment:    "update_time",
        tag:        "update_time",
    }
    model = append(model, info)
    return &model
}

func TestMode(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:  "xml",
        modelFile:   "model.go",
    }

    genModel(config, "test_table", *createModeInfo())
}

func TestXml(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:  "xml",
    }

    genXml(config, "test_table", *createModeInfo())
    config.mapperFile = "go"
    genXml(config, "test_table", *createModeInfo())
}

func TestProxy(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:   "xml",
    }

    genProxy(config, "test_table", *createModeInfo())
}

func TestV2Proxy(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:   "xml",
    }

    genV2Proxy(config, "test_table", *createModeInfo())
}

func TestAll1(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:   "xml",
        //modelFile:   "model.go",
    }
    info := *createModeInfo()
    genModel(config, "test_table", info)
    genXml(config, "test_table", info)
    genProxy(config, "test_table", info)
}

func TestAll2(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:   "go",
        //modelFile:   "model.go",
    }
    info := *createModeInfo()
    genModel(config, "test_table", info)
    genXml(config, "test_table", info)
    genProxy(config, "test_table", info)
}

func TestAll3(t *testing.T) {
    config := config{
        packageName: "test_package",
        path:        "c:/tmp/",
        tagName:     "xfield",
        mapperFile:   "xml",
        //modelFile:   "model.go",
    }
    info := *createModeInfo()
    genModel(config, "TEST_TABLE", info)
    genXml(config, "TEST_TABLE", info)
    genV2Proxy(config, "TEST_TABLE", info)
}
