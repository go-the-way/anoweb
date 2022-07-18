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
	"bytes"
	"embed"
	"html/template"
	"os"
	"path/filepath"
)

var templateCaches = make(map[string]string, 0)

// AddFunc add func
func (ctx *Context) AddFunc(name string, funcMap any) *Context {
	if name != "" && funcMap != nil {
		ctx.funcMap[name] = funcMap
	}
	return ctx
}

// AddFuncMap add funcMap
func (ctx *Context) AddFuncMap(funcMap template.FuncMap) *Context {
	if funcMap != nil && len(funcMap) > 0 {
		for k, v := range funcMap {
			ctx.AddFunc(k, v)
		}
	}
	return ctx
}

// Template Response template
func (ctx *Context) Template(tpl string, data map[string]any) {
	t, err := template.New("HTML").Funcs(ctx.funcMap).Parse(tpl)
	if err != nil {
		panic(err)
	}
	ctx.SetDataMap(data, false)
	var w bytes.Buffer
	err = t.Execute(&w, ctx.dataMap)
	if err != nil {
		panic(err)
	}
	ctx.HTML(w.String())
}

// TemplateFile Response template with file
func (ctx *Context) TemplateFile(prefix string, data map[string]any) {
	fileName := prefix + ctx.templateConfig.Suffix
	tpl, have := templateCaches[fileName]
	realPath := filepath.Join(ctx.templateConfig.Root, fileName)
	if !have {
		bytes2, err := os.ReadFile(realPath)
		if err != nil {
			panic(err)
		}
		tpl = string(bytes2)
		if ctx.templateConfig.Cache {
			templateCaches[fileName] = tpl
		}
	}
	ctx.Template(tpl, data)
}

// TemplateFS Response template with rFS
func (ctx *Context) TemplateFS(FS *embed.FS, prefix string, data map[string]any) {
	fileName := prefix + ctx.templateConfig.Suffix
	tpl, have := templateCaches[fileName]
	if !have {
		bytes2, err := FS.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		tpl = string(bytes2)
		if ctx.templateConfig.Cache {
			templateCaches[fileName] = tpl
		}
	}
	ctx.Template(tpl, data)
}
