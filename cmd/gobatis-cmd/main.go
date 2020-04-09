/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

package main

import (
    "encoding/json"
    "flag"
    "github.com/xfali/gobatis-cmd/internal/pkg/db"
    "github.com/xfali/gobatis-cmd/internal/pkg/generator"
    "github.com/xfali/gobatis-cmd/pkg/config"
    "github.com/xfali/gobatis-cmd/pkg/io"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

func main() {
    driver := flag.String("driver", "mysql", "driver of db")
    packageName := flag.String("pkg", "xfali.gobatis.default", "Set the package name of .go file")
    dbName := flag.String("db", "", "the name of db instance used in model files")
    tableName := flag.String("table", "", "the name of table to be generated")
    host := flag.String("host", "localhost", "host of db")
    port := flag.Int("port", 3306, "port of db ")
    username := flag.String("user", "", "user name of db")
    pw := flag.String("pw", "", "password of db")
    path := flag.String("path", "", "root path to save files")
    modelfile := flag.String("model", "", "the name of model file")
    tagName := flag.String("tag", "xfield", "the name of field tag,eg: xfield,json  xfield,json,yaml")
    mapper := flag.String("mapper", "xml", "generate mapper file: xml | template | go")
    plugin := flag.String("plugin", "", "path of plugin")
    keyword := flag.Bool("keyword", false, "with Keyword escape")
    namespace := flag.String("namespace", "", "namespace")
    confFile := flag.String("f", "", "config file")
    flag.Parse()

    dbDriver := db.GetDriver(*driver)
    if dbDriver == nil {
        log.Print("not support driver: ", *driver)
        os.Exit(-1)
    }

    conf := config.FileConfig{}
    if *confFile != "" {
        err := loadFromFile(&conf, *confFile)
        if err != nil {
            os.Exit(1)
        }
    }

    root := formatPath(*path)
    conf.Driver = *driver
    conf.Path = root
    conf.PackageName = *packageName
    conf.Namespace = *namespace
    conf.ModelFile = *modelfile
    conf.TagName = *tagName
    conf.MapperFile = *mapper
    conf.Plugin = *plugin
    conf.Keyword = *keyword
    conf.TableName = *tableName
    conf.DBName = *dbName
    conf.Host = *host
    conf.Port = *port
    conf.User = *username
    conf.Password = *pw

    err := dbDriver.Open(conf.Driver, db.GenDBInfo(conf.Driver, conf.DBName, conf.User, conf.Password, conf.Host, conf.Port))
    if err != nil {
        log.Print(err)
        os.Exit(-1)
    }
    defer dbDriver.Close()

    if conf.TableName == "" {
        tables, err2 := dbDriver.QueryTableNames(conf.DBName)
        if err2 != nil {
            log.Print(err2)
            os.Exit(-2)
        }
        for _, v := range tables {
            generator.GenOneTable(conf.Config, dbDriver, conf.DBName, v)
        }
    } else {
        generator.GenOneTable(conf.Config, dbDriver, conf.DBName, conf.TableName)
    }
    os.Exit(0)
}

func loadFromFile(conf *config.FileConfig, path string) error {
    b, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(b, conf)
}

func formatPath(path string) string {
    root := strings.TrimSpace(path)
    if root == "" {
        root = "./"
    } else {
        if !io.IsPathExists(path) {
            io.Mkdir(path)
        }
        if root[len(root)-1:] != "/" {
            root = root + "/"
        }
    }
    return root
}
