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

package generator

import (
	"github.com/acmestack/gobatis-cmd/pkg/common"
	"github.com/acmestack/gobatis-cmd/pkg/config"
	"github.com/acmestack/gobatis-cmd/pkg/plugin"
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
