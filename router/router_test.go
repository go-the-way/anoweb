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

package router

import (
	"embed"
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	require.Equal(t, make([]*Simple, 0), r.Simples)
	require.Equal(t, make([]*Dynamic, 0), r.Dynamics)
}

func TestRouterAll(t *testing.T) {
	type testCase struct {
		*Router
		method  string
		pattern string
	}

	cases := make([]*testCase, 0)
	// test for all methods
	{
		r := NewRouter()
		r.Request("/hello_world", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello_world"})

		r.Request("/hello/world", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello/world"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello/world"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello/world"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello/world"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello/world"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello/world"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello/world"})

		r.Request("/hello/world/apple", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello/world/apple"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello/world/apple"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello/world/apple"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello/world/apple"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello/world/apple"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello/world/apple"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello/world/apple"})

		r.Request("/hello/world/apple/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello/world/apple/pear"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello/world/apple/pear"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello/world/apple/pear"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello/world/apple/pear"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello/world/apple/pear"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello/world/apple/pear"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello/world/apple/pear"})
	}

	// test for GET
	{
		r := NewRouter()
		r.Get("/hello_world", func(ctx *context.Context) {})
		r.Get("/hello_world/apple", func(ctx *context.Context) {})
		r.Get("/hello_world/orange", func(ctx *context.Context) {})
		r.Get("/hello_world/banana", func(ctx *context.Context) {})
		r.Get("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/pear"})
	}
	// test for POST
	{
		r := NewRouter()
		r.Post("/hello_world", func(ctx *context.Context) {})
		r.Post("/hello_world/apple", func(ctx *context.Context) {})
		r.Post("/hello_world/orange", func(ctx *context.Context) {})
		r.Post("/hello_world/banana", func(ctx *context.Context) {})
		r.Post("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodPost, "/hello_world/pear"})
	}
	// test for PUT
	{
		r := NewRouter()
		r.Put("/hello_world", func(ctx *context.Context) {})
		r.Put("/hello_world/apple", func(ctx *context.Context) {})
		r.Put("/hello_world/orange", func(ctx *context.Context) {})
		r.Put("/hello_world/banana", func(ctx *context.Context) {})
		r.Put("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodPut, "/hello_world/pear"})
	}
	// test for DELETE
	{
		r := NewRouter()
		r.Delete("/hello_world", func(ctx *context.Context) {})
		r.Delete("/hello_world/apple", func(ctx *context.Context) {})
		r.Delete("/hello_world/orange", func(ctx *context.Context) {})
		r.Delete("/hello_world/banana", func(ctx *context.Context) {})
		r.Delete("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodDelete, "/hello_world/pear"})
	}
	// test for PATCH
	{
		r := NewRouter()
		r.Patch("/hello_world", func(ctx *context.Context) {})
		r.Patch("/hello_world/apple", func(ctx *context.Context) {})
		r.Patch("/hello_world/orange", func(ctx *context.Context) {})
		r.Patch("/hello_world/banana", func(ctx *context.Context) {})
		r.Patch("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodPatch, "/hello_world/pear"})
	}
	// test for HEAD
	{
		r := NewRouter()
		r.Head("/hello_world", func(ctx *context.Context) {})
		r.Head("/hello_world/apple", func(ctx *context.Context) {})
		r.Head("/hello_world/orange", func(ctx *context.Context) {})
		r.Head("/hello_world/banana", func(ctx *context.Context) {})
		r.Head("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodHead, "/hello_world/pear"})
	}
	// test for OPTIONS
	{
		r := NewRouter()
		r.Options("/hello_world", func(ctx *context.Context) {})
		r.Options("/hello_world/apple", func(ctx *context.Context) {})
		r.Options("/hello_world/orange", func(ctx *context.Context) {})
		r.Options("/hello_world/banana", func(ctx *context.Context) {})
		r.Options("/hello_world/pear", func(ctx *context.Context) {})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodOptions, "/hello_world/pear"})
	}
	// test for Resource
	{
		r := NewRouter()
		r.Resource("/hello_world", "index.html", mime.HTML)
		r.Resource("/hello_world/apple", "index.html", mime.HTML)
		r.Resource("/hello_world/orange", "index.html", mime.HTML)
		r.Resource("/hello_world/banana", "index.html", mime.HTML)
		r.Resource("/hello_world/pear", "index.html", mime.HTML)
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/pear"})
	}
	// test for FSResource
	{
		var fs embed.FS
		r := NewRouter()
		r.FSResource(&fs, "/hello_world", "index.html", mime.HTML)
		r.FSResource(&fs, "/hello_world/apple", "index.html", mime.HTML)
		r.FSResource(&fs, "/hello_world/orange", "index.html", mime.HTML)
		r.FSResource(&fs, "/hello_world/banana", "index.html", mime.HTML)
		r.FSResource(&fs, "/hello_world/pear", "index.html", mime.HTML)
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/apple"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/orange"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/banana"})
		cases = append(cases, &testCase{r, http.MethodGet, "/hello_world/pear"})
	}

	for _, c := range cases {
		matched := false
		for _, s := range c.Router.Simples {
			if s.Method == c.method && s.Pattern == c.pattern {
				matched = true
			}
		}
		require.Equal(t, true, matched)
	}
}

func buildReq() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	return req
}

func TestRouterResource(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	r := NewRouter()
	r.Resource("/res", "res.txt", mime.TEXT)
	ctx := context.New()
	ctx.Allocate(buildReq(), &config.Template{})
	ctx.Add(r.Simples[0].Handler)
	ctx.Chain()
}

func TestRouterSimpleAll(t *testing.T) {
	r := NewRouter()
	r.Route("*", "/id", func(ctx *context.Context) {})
	for i, m := range supportedMethods {
		t.Run(m, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, m, r.Simples[i].Method)
			require.Equal(t, "/id", r.Simples[i].Pattern)
		})
	}
}

var fs embed.FS

func TestRouterFSResource(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	r := NewRouter()
	r.FSResource(&fs, "/res", "res.txt", mime.TEXT)
	ctx := context.New()
	ctx.Allocate(buildReq(), &config.Template{})
	ctx.Add(r.Simples[0].Handler)
	ctx.Chain()
}

func TestRouterDynamic(t *testing.T) {
	r := NewRouter()
	r.Route(http.MethodGet, "/{id}/{name}", func(ctx *context.Context) {})
	require.Equal(t, http.MethodGet, r.Dynamics[0].Method)
	require.Equal(t, []string{"id", "name"}, r.Dynamics[0].Params)
	require.Equal(t, `/([\w_.-]+)/([\w_.-]+)`, r.Dynamics[0].Pattern)
}

func TestRouterDynamicAll(t *testing.T) {
	r := NewRouter()
	r.Route("*", "/{id}/{name}", func(ctx *context.Context) {})
	for i, m := range supportedMethods {
		t.Run(m, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, m, r.Dynamics[i].Method)
			require.Equal(t, []string{"id", "name"}, r.Dynamics[i].Params)
			require.Equal(t, `/([\w_.-]+)/([\w_.-]+)`, r.Dynamics[i].Pattern)
		})
	}
}

func TestRouterSupport(t *testing.T) {
	defer func() {
		if re := recover(); re != nil && re.(error) != nil {
			require.Equal(t, "method not supported : BODY", re.(error).Error())
		}
	}()
	r := NewRouter()
	r.Route("BODY", "/", func(ctx *context.Context) {})
}
