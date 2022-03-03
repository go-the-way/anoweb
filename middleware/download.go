// Copyright 2022 anoweb Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"
	"github.com/go-the-way/anoweb/mime"
)

type download struct {
	Root    string
	DateDir bool
}

// Download return download
func Download() *download {
	return &download{Root: os.TempDir(), DateDir: true}
}

// Handler implements
func (d *download) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		fileName := ctx.Param("file")
		if fileName == "" {
			ctx.JSONText(`{"msg":"file is empty","code":1}`)
			return
		}
		root := d.Root
		dateDir := ""
		if d.DateDir {
			dateDir = time.Now().Format("20060102")
		}
		filePath := filepath.Join(root, dateDir, fileName)
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			ctx.JSONText(`{"msg":"file is not exists","code":1}`)
			return
		}
		ctx.Write(context.Builder().Header(map[string][]string{
			headers.ContentDisposition: {"attachment;filename=" + url.QueryEscape(fileName)},
		}).ContentType(mime.BINARY).Data(bytes).Build())
	}
}

func (d *download) Path(file string) string {
	dd := ""
	if d.DateDir {
		dd = time.Now().Format("20060102")
	}
	return filepath.Join(d.Root, dd, file)
}

func (d *download) File(file string) (*os.File, error) {
	return os.Open(d.Path(file))
}

func (d *download) Buf(file string) ([]byte, error) {
	return ioutil.ReadFile(d.Path(file))
}
