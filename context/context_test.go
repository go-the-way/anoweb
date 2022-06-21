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
	"github.com/go-the-way/anoweb/config"
	"github.com/stretchr/testify/require"
	"html/template"
	"net/http"
	"testing"
)

func TestContextNew(t *testing.T) {
	req := buildReq(`{"apple":100}`)
	ctx := New()
	ctx.Allocate(req, &config.Template{
		FuncMap: template.FuncMap{
			"sum": func(a, b int) int { return a + b },
		},
	})
	require.Equal(t, req, ctx.Request)
	ctx.templateConfig.FuncMap = nil
	require.Equal(t, &config.Template{}, ctx.templateConfig)
	type _apple struct {
		Apple int `json:"apple"`
	}
	a := _apple{}
	ctx.Bind(&a)
	require.Equal(t, &_apple{100}, &a)
}

func TestContextAdd(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.Add(func(ctx *Context) {})
	ctx.Add(func(ctx *Context) {})
	ctx.Add(func(ctx *Context) {})
	require.Equal(t, 3, len(ctx.handlers))
}

func TestContextChain(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	apple := 0
	ctx.Add(func(ctx *Context) { apple++; ctx.Chain() })
	ctx.Add(func(ctx *Context) { apple++; ctx.Chain() })
	ctx.Add(func(ctx *Context) { apple++ })
	ctx.Chain()
	require.Equal(t, 3, apple)
}

func TestContextWrite(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.Data([]byte(`hello world`))
	require.Equal(t, []byte(`hello world`), ctx.Response.Data)
}

func TestContextStatus(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.Status(150)
	require.Equal(t, 150, ctx.Response.Status)
}

func TestContextHeader(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.Header(http.Header{"nobody": {"apple"}})
	ctx.Header(http.Header{"nobody": {"pear"}})
	require.Equal(t, http.Header{"nobody": {"apple", "pear"}}, ctx.Response.Header)
}

func TestContextAddCookie(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.AddCookie(&http.Cookie{Name: "apple", Value: "100"})
	require.Equal(t, []*http.Cookie{{Name: "apple", Value: "100"}}, ctx.Response.Cookies)
}
