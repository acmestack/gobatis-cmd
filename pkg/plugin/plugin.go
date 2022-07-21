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

package plugin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/acmestack/gobatis-cmd/pkg/common"
	"github.com/acmestack/gobatis-cmd/pkg/config"
	mio "github.com/acmestack/gobatis-cmd/pkg/io"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunPlugin(config config.Config, tableName string, model []common.ModelInfo) error {
	if config.Plugin == "" {
		return nil
	}
	b, e := ExecPluginMethod(config.Plugin, common.OutPutSuffixMethod, nil)
	if e != nil {
		return e
	}

	info := common.GenerateInfo{
		Driver:  config.Driver,
		Table:   tableName,
		Package: config.PackageName,
		Models:  model,
	}
	d, _ := json.Marshal(info)
	gendata, errGen := ExecPluginMethod(config.Plugin, common.GenerateMethod, d)
	if errGen != nil {
		return errGen
	}

	outputDir := config.Path
	if !mio.IsPathExists(outputDir) {
		mio.Mkdir(outputDir)
	}
	output := strings.ToLower(tableName) + strings.TrimSpace(string(b))
	outputFile, err := mio.OpenAppend(filepath.Join(outputDir, output))
	if err == nil {
		defer outputFile.Close()
		return mio.Write(outputFile, gendata)
	}
	return nil
}

func ExecPluginMethod(path string, method string, data []byte) ([]byte, error) {
	args := []string{"-" + common.MethodFlag, method}

	cmd := exec.Command(path, args...)
	if cmd == nil {
		return nil, errors.New("exec plugin failed")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if data != nil {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		stdin.Write(data)
		stdin.Write([]byte{byte('\n')})
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	n, errR := io.Copy(buf, stdout)
	fmt.Println(n)
	if errR != nil {
		return nil, errR
	}

	errW := cmd.Wait()
	if errW != nil {
		return nil, errW
	}

	code := cmd.ProcessState.ExitCode()
	if code != 0 {
		return nil, errors.New("exec plugin exit not 0")
	}
	return buf.Bytes(), nil
}
