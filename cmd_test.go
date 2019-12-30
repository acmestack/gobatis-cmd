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
    "testing"
)

func createModeInfo() *[]common.ModelInfo {
    model := []common.ModelInfo{}
    info := common.ModelInfo{
        ColumnName: "id",
        DataType:   "bigint",
        Nullable:   "NO",
        Comment:    "id",
        Tag:        "id",
        ColumnKey:  "PRI",
    }
    model = append(model, info)

    info = common.ModelInfo{
        ColumnName: "username",
        DataType:   "varchar",
        Nullable:   "NO",
        Comment:    "username",
        Tag:        "username",
    }
    model = append(model, info)

    info = common.ModelInfo{
        ColumnName: "password",
        DataType:   "varchar",
        Nullable:   "NO",
        Comment:    "password",
        Tag:        "password",
    }
    model = append(model, info)

    info = common.ModelInfo{
        ColumnName: "update_time",
        DataType:   "timestamp",
        Nullable:   "YES",
        Comment:    "update_time",
        Tag:        "update_time",
    }
    model = append(model, info)
    return &model
}

func TestMode(t *testing.T) {
    t.Run("0", func(t *testing.T) {
        config := Config{
            PackageName: "test_package",
            Path:        "c:/tmp/",
            TagName:     "",
            MapperFile:  "xml",
            ModelFile:   "model.go",
        }

        genModel(config, "test_table", *createModeInfo())
    })

    t.Run("1", func(t *testing.T) {
        config := Config{
            PackageName: "test_package",
            Path:        "c:/tmp/",
            TagName:     "xfield",
            MapperFile:  "xml",
            ModelFile:   "model.go",
        }

        genModel(config, "test_table", *createModeInfo())
    })

    t.Run("2", func(t *testing.T) {
        config := Config{
            PackageName: "test_package",
            Path:        "c:/tmp/",
            TagName:     "xfield,json",
            MapperFile:  "xml",
            ModelFile:   "model.go",
        }

        genModel(config, "test_table", *createModeInfo())
    })

    t.Run("2.1", func(t *testing.T) {
        config := Config{
            PackageName: "test_package",
            Path:        "c:/tmp/",
            TagName:     "xfield,json,",
            MapperFile:  "xml",
            ModelFile:   "model.go",
        }

        genModel(config, "test_table", *createModeInfo())
    })

    t.Run("2.2", func(t *testing.T) {
        config := Config{
            PackageName: "test_package",
            Path:        "c:/tmp/",
            TagName:     "xml,json,",
            MapperFile:  "xml",
            ModelFile:   "model.go",
        }

        genModel(config, "test_table", *createModeInfo())
    })

    t.Run("3", func(t *testing.T) {
        config := Config{
            PackageName: "test_package",
            Path:        "c:/tmp/",
            TagName:     "xfield,json,xml",
            MapperFile:  "xml",
            ModelFile:   "model.go",
        }

        genModel(config, "test_table", *createModeInfo())
    })
}

func TestXml(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
    }

    genXml(config, "test_table", *createModeInfo())
    config.MapperFile = "go"
    genXml(config, "test_table", *createModeInfo())
}

func TestProxy(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
    }

    genProxy(config, "test_table", *createModeInfo())
}

func TestV2Proxy(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
    }

    genV2Proxy(config, "test_table", *createModeInfo())
}

func TestAll1(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
        //ModelFile:   "model.go",
    }
    info := *createModeInfo()
    genModel(config, "test_table", info)
    genXml(config, "test_table", info)
    genProxy(config, "test_table", info)
}

func TestAll2(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "go",
        //ModelFile:   "model.go",
    }
    info := *createModeInfo()
    genModel(config, "test_table", info)
    genXml(config, "test_table", info)
    genProxy(config, "test_table", info)
}

func TestAll3(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
        //ModelFile:   "model.go",
    }
    info := *createModeInfo()
    genModel(config, "TEST_TABLE", info)
    genXml(config, "TEST_TABLE", info)
    genV2Proxy(config, "TEST_TABLE", info)
}

func TestPlugin(t *testing.T) {
    config := Config{
        PackageName: "test_package",
        Path:        "c:/tmp/",
        TagName:     "xfield",
        MapperFile:  "xml",
        Plugin: "c:/tmp/webplugin.exe",
        //ModelFile:   "model.go",
    }
    info := *createModeInfo()
    RunPlugin(config, "TEST_TABLE", info)
}
