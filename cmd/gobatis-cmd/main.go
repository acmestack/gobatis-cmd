/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/acmestack/gobatis-cmd/pkg/config"
	"github.com/acmestack/gobatis-cmd/pkg/db"
	"github.com/acmestack/gobatis-cmd/pkg/generator"
	"github.com/acmestack/gobatis-cmd/pkg/io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	driver := flag.String("driver", "mysql", "driver of db")
	packageName := flag.String("pkg", "acmestack.gobatis.default", "Set the package name of .go file")
	dbName := flag.String("db", "", "the name of db instance used in model files")
	tableName := flag.String("table", "", "the name of table to be generated")
	host := flag.String("host", "localhost", "host of db")
	port := flag.Int("port", 3306, "port of db ")
	username := flag.String("user", "", "user name of db")
	pw := flag.String("pw", "", "password of db")
	path := flag.String("path", "", "root path to save files")
	modelfile := flag.String("model", "", "the name of model file")
	tagName := flag.String("tag", "column", "the name of field tag,eg: column,json  column,json,yaml")
	mapper := flag.String("mapper", "xml", "generate mapper file: xml | template | go")
	plugin := flag.String("plugin", "", "path of plugin")
	keyword := flag.Bool("keyword", false, "with Keyword escape")
	namespace := flag.String("namespace", "", "namespace")
	confFile := flag.String("f", "", "config file")
	register := flag.Bool("register", false, "add register code")
	flag.Parse()

	conf := config.FileConfig{}
	if *confFile != "" {
		err := loadFromFile(&conf, *confFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		conf.Path = formatPath(conf.Path)
	} else {
		conf.Driver = *driver
		conf.Path = formatPath(*path)
		conf.PackageName = *packageName
		conf.Namespace = *namespace
		conf.ModelFile = *modelfile
		conf.TagName = *tagName
		conf.MapperFile = *mapper
		conf.Plugin = *plugin
		conf.Keyword = *keyword
		conf.Register = *register
		conf.TableName = *tableName
		conf.DBName = *dbName
		conf.Host = *host
		conf.Port = *port
		conf.User = *username
		conf.Password = *pw
	}

	dbDriver := db.GetDriver(conf.Driver)
	if dbDriver == nil {
		log.Print("not support driver: ", conf.Driver)
		os.Exit(-1)
	}

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
	fmt.Printf("config file: %s\n", string(b))
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
