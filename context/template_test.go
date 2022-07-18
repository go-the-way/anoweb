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

package context

import (
	"embed"
	"html/template"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestAddFunc(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.AddFunc("sum", func(a, b int) int { return a + b })
	require.NotNil(t, ctx.funcMap["sum"])
}

func TestAddFuncMap(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.AddFuncMap(template.FuncMap{"sum": func(a, b int) int { return a + b }})
	require.NotNil(t, ctx.funcMap["sum"])
}

func TestTemplate(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.Template("<h1>hello world</h1>", nil)
	require.Equal(t, []byte("<h1>hello world</h1>"), ctx.Response.Data)
	require.Equal(t, mime.HTML, ctx.Response.ContentType)
}

func TestTemplateNewPanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.Template("<h1>{{.HELLO</h1>", nil)
}

func TestTemplateExecutePanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{
		FuncMap: map[string]any{
			"sum": func(a, b int) int {
				return a + b
			},
		},
	})
	ctx.Template(`<h1>{{sum "hello" "world"}}</h1>`, nil)
}

func TestTemplateFile(t *testing.T) {
	_testFunc := func() {
		root, _ := os.Getwd()
		tplDir, _ := filepath.Abs(filepath.Join(root, "testdata"))
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{Cache: true, Root: tplDir, Suffix: ".html"})
		// test for no data
		{
			ctx.TemplateFile("test", nil)
			require.Equal(t, []byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><h1>hello world</h1></body></html>`), ctx.Response.Data)
			require.Equal(t, mime.HTML, ctx.Response.ContentType)
		}
		// test for data
		{
			ctx.TemplateFile("test_with_data", map[string]any{"Apple": "100"})
			require.Equal(t, []byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><h1>hello world</h1><h1>100</h1></body></html>`), ctx.Response.Data)
			require.Equal(t, mime.HTML, ctx.Response.ContentType)
		}
	}
	_testFunc()
	_testFunc()
}

func TestTemplateFileReadFilePanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	root, _ := os.Getwd()
	tplDir, _ := filepath.Abs(filepath.Join(root, "testdata"))
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{Root: tplDir, Suffix: ".html", FuncMap: map[string]any{}})
	ctx.TemplateFile("test-haha", nil)
}

//go:embed testdata
var tFS embed.FS

func TestTemplateFS(t *testing.T) {
	_testFunc := func() {
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{Cache: true, Suffix: ".html", FuncMap: map[string]any{}})
		// test for no data
		{
			ctx.TemplateFS(&tFS, "testdata/test", nil)
			require.Equal(t, []byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><h1>hello world</h1></body></html>`), ctx.Response.Data)
			require.Equal(t, mime.HTML, ctx.Response.ContentType)
		}
		// test for data
		{
			ctx.TemplateFS(&tFS, "testdata/test_with_data", map[string]any{"Apple": "100"})
			require.Equal(t, []byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><h1>hello world</h1><h1>100</h1></body></html>`), ctx.Response.Data)
			require.Equal(t, mime.HTML, ctx.Response.ContentType)
		}
	}
	_testFunc()
	_testFunc()
}

func TestTemplateFSReadFilePanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	root, _ := os.Getwd()
	tplDir, _ := filepath.Abs(filepath.Join(root, "testdata"))
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{Root: tplDir, Suffix: ".html", FuncMap: map[string]any{}})
	ctx.TemplateFS(&tFS, "test-haha", nil)
}
