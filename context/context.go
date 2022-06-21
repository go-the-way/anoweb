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
	"html/template"
	"net/http"

	"github.com/go-the-way/anoweb/config"
)

// Context struct
type Context struct {
	Request        *http.Request
	Response       *Response
	pos            int
	handlers       []func(c *Context)
	paramMap       map[string][]string
	MultipartMap   map[string][]*MultipartFile
	dataMap        map[string]interface{}
	funcMap        template.FuncMap
	templateConfig *config.Template
}

// New context
func New() *Context {
	ctx := &Context{
		Response:     Builder().DefaultBuild(),
		handlers:     make([]func(c *Context), 0),
		paramMap:     make(map[string][]string, 0),
		MultipartMap: make(map[string][]*MultipartFile, 0),
		dataMap:      make(map[string]interface{}, 0),
	}
	ctx.onCreated()
	return ctx
}

func (ctx *Context) Allocate(req *http.Request, templateConfig *config.Template) {
	ctx.Request = req
	funcMap := make(map[string]interface{}, 0)
	if templateConfig != nil && templateConfig.FuncMap != nil {
		for k, v := range templateConfig.FuncMap {
			funcMap[k] = v
		}
	}
	_ = ctx.Request.ParseForm()
	if ctx.Request.Form != nil {
		for k := range ctx.Request.Form {
			ctx.paramMap[k] = ctx.Request.Form[k]
		}
	}
	ctx.templateConfig = templateConfig
	ctx.funcMap = funcMap
}

// Add context handler
func (ctx *Context) Add(handlers ...func(ctx *Context)) *Context {
	for _, h := range handlers {
		if h != nil {
			ctx.handlers = append(ctx.handlers, h)
		}
	}
	return ctx
}

// Chain execute context handler
func (ctx *Context) Chain() {
	if len(ctx.handlers) > 0 {
		if ctx.pos > len(ctx.handlers)-1 {
			ctx.onDestroyed()
			return
		}
		ctx.pos++
		ctx.handlers[ctx.pos-1](ctx)
	}
}

// Data buffer
func (ctx *Context) Data(data []byte) *Context {
	ctx.Response.Data = data
	return ctx
}

// Status code
func (ctx *Context) Status(status int) *Context {
	ctx.Response.Status = status
	return ctx
}

// Header return header
func (ctx *Context) Header(header http.Header) *Context {
	for k, v := range header {
		if _, have := ctx.Response.Header[k]; have {
			ctx.Response.Header[k] = append(ctx.Response.Header[k], v...)
		} else {
			ctx.Response.Header[k] = v
		}
	}
	return ctx
}

// AddCookie add cookie
func (ctx *Context) AddCookie(cookies ...*http.Cookie) *Context {
	ctx.Response.Cookies = append(ctx.Response.Cookies, cookies...)
	return ctx
}
