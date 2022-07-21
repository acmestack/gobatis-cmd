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

package common

type ModelInfo struct {
	ColumnName string
	DataType   string
	Nullable   string
	ColumnKey  string
	Comment    string
	Tag        string
}

const (
	MethodFlag         = "method"
	OutPutSuffixMethod = "output"
	GenerateMethod     = "generate"
)

type GenerateInfo struct {
	Driver  string      `json:"driver"`
	Package string      `json:"package"`
	Table   string      `json:"table"`
	Models  []ModelInfo `json:"models"`
}

type PluginResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
