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
	Register    bool   `json:"register"`
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
