/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description:
 */

package main

import (
	"flag"
	"github.com/xfali/gobatis-cmd/io"
	"log"
	"os"
	"strings"
)

type Config struct {
	Driver      string
	Path        string
	PackageName string
	ModelFile   string
	TagName     string
	MapperFile  string
	Plugin      string
	Keyword     bool
}

func main() {
	driver := flag.String("driver", "mysql", "driver of db")
	packageName := flag.String("pkg", "xfali/gobatis/default", "Set the package name of .go file")
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
	flag.Parse()

	db := buildinDrivers[*driver]
	if db == nil {
		log.Print("not support driver: ", *driver)
		os.Exit(-1)
	}

	err := db.Open(*driver, genDBInfo(*driver, *dbName, *username, *pw, *host, *port))
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
	defer db.Close()

	root := formatPath(*path)

	config := Config{
		Driver:      *driver,
		Path:        root,
		PackageName: *packageName,
		ModelFile:   *modelfile,
		TagName:     *tagName,
		MapperFile:  *mapper,
		Plugin:      *plugin,
		Keyword:     *keyword,
	}

	if *tableName == "" {
		tables, err2 := db.QueryTableNames(*dbName)
		if err2 != nil {
			log.Print(err2)
			os.Exit(-2)
		}
		for _, v := range tables {
			models, err := db.QueryTableInfo(*dbName, v)
			if err != nil {
				log.Print(err)
				os.Exit(-3)
			}
			err = generate(config, models, v)
			if err != nil {
				log.Print(err)
				os.Exit(-3)
			}
		}
	} else {
		models, err := db.QueryTableInfo(*dbName, *tableName)
		if err != nil {
			log.Print(err)
			os.Exit(-3)
		}
		err2 := generate(config, models, *tableName)
		if err2 != nil {
			log.Print(err2)
			os.Exit(-2)
		}
	}
	os.Exit(0)
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
