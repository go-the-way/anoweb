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

package anoweb

import (
	"embed"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"
	"github.com/go-the-way/anoweb/router"
	"github.com/stretchr/testify/require"
)

type testRouterCase struct {
	a       *App
	method  string
	pattern string
	handler func(ctx *context.Context)
}

func TestAppRequest(t *testing.T) {
	test := func(t *testing.T, method, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Request(pattern, helloHandler).parseRouters(), method, pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", method, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	// test for Request
	{
		var (
			pattern      = "/index"
			method       = "*"
			val          = 0
			changeVal    = 10
			helloHandler = func(ctx *context.Context) {
				val = changeVal
			}
		)
		cases := make([]*testRouterCase, 0)
		{
			a := New()
			a.Request(pattern, helloHandler).parseRouters()
			cases = append(cases, &testRouterCase{a, method, pattern, helloHandler})
		}
		for _, trc := range cases {
			for _, m := range []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodPatch,
				http.MethodHead,
				http.MethodOptions,
			} {
				trc.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", m, pattern)].Handler(nil)
				require.Equal(t, val, changeVal)
			}
		}
	}
	// test for Get
	{
		test(t, http.MethodGet, "/index", 0, 10)
		test(t, http.MethodGet, "/index/1", 1, 100)
		test(t, http.MethodGet, "/index/2", 2, 200)
		test(t, http.MethodGet, "/index/3", 3, 300)
	}
	// test for Post
	{
		test(t, http.MethodPost, "/index", 0, 10)
		test(t, http.MethodPost, "/index/1", 1, 100)
		test(t, http.MethodPost, "/index/2", 2, 200)
		test(t, http.MethodPost, "/index/3", 3, 300)
	}
	// test for Put
	{
		test(t, http.MethodPut, "/index", 0, 10)
		test(t, http.MethodPut, "/index/1", 1, 100)
		test(t, http.MethodPut, "/index/2", 2, 200)
		test(t, http.MethodPut, "/index/3", 3, 300)
	}
	// test for Patch
	{
		test(t, http.MethodPatch, "/index", 0, 10)
		test(t, http.MethodPatch, "/index/1", 1, 100)
		test(t, http.MethodPatch, "/index/2", 2, 200)
		test(t, http.MethodPatch, "/index/3", 3, 300)
	}
	// test for Delete
	{
		test(t, http.MethodDelete, "/index", 0, 10)
		test(t, http.MethodDelete, "/index/1", 1, 100)
		test(t, http.MethodDelete, "/index/2", 2, 200)
		test(t, http.MethodDelete, "/index/3", 3, 300)
	}
	// test for Head
	{
		test(t, http.MethodHead, "/index", 0, 10)
		test(t, http.MethodHead, "/index/1", 1, 100)
		test(t, http.MethodHead, "/index/2", 2, 200)
		test(t, http.MethodHead, "/index/3", 3, 300)
	}
	// test for Options
	{
		test(t, http.MethodOptions, "/index", 0, 10)
		test(t, http.MethodOptions, "/index/1", 1, 100)
		test(t, http.MethodOptions, "/index/2", 2, 200)
		test(t, http.MethodOptions, "/index/3", 3, 300)
	}
}

func TestAppGet(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Get(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodGet, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppPost(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Post(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodPost, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppPut(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Put(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodPut, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppDelete(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Delete(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodDelete, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppPatch(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Patch(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodPatch, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppHead(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Head(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodHead, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppOptions(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		valV := val
		changeValV := changeVal
		helloHandler := func(ctx *context.Context) {
			valV = changeValV
		}
		testCase := &testRouterCase{New().Options(pattern, helloHandler).parseRouters(), "", pattern, helloHandler}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodOptions, pattern)].Handler(nil)
		require.Equal(t, valV, changeValV)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

func TestAppResource(t *testing.T) {
	test := func(t *testing.T, pattern string, val, changeVal int) {
		req, _ := http.NewRequest(http.MethodGet, pattern, nil)
		ctx := context.New(req, &config.Template{})
		testCase := &testRouterCase{New().Resource(pattern, "context/testdata/file.txt", mime.TEXT).parseRouters(), "", pattern, nil}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodGet, pattern)].Handler(ctx)
		r := ctx.Response
		require.Equal(t, []byte(`hello world`), r.Data)
		require.Equal(t, mime.TEXT, r.ContentType)
	}
	test(t, "/index", 0, 10)
	test(t, "/index/1", 1, 100)
	test(t, "/index/2", 2, 200)
	test(t, "/index/3", 3, 300)
}

//go:embed context/testdata
var fs embed.FS

func TestAppFSResource(t *testing.T) {
	test := func(t *testing.T, pattern string) {
		req, _ := http.NewRequest(http.MethodGet, pattern, nil)
		ctx := context.New(req, &config.Template{})
		testCase := &testRouterCase{New().FSResource(&fs, pattern, "context/testdata/file.txt", mime.TEXT).parseRouters(), "", pattern, nil}
		testCase.a.parsedRouters.Simples[fmt.Sprintf("%s:%s", http.MethodGet, pattern)].Handler(ctx)
		r := ctx.Response
		require.Equal(t, []byte(`hello world`), r.Data)
		require.Equal(t, mime.TEXT, r.ContentType)
	}
	test(t, "/index")
	test(t, "/index/1")
	test(t, "/index/2")
	test(t, "/index/3")
}

func TestAddRouter(t *testing.T) {
	{
		r := router.NewRouter().Get("/haha", func(ctx *context.Context) {})
		a := New().AddRouter(r).parseRouters()
		require.Equal(t, 1, len(a.parsedRouters.Simples))
	}
	{
		r := router.NewRouter().
			Get("/haha/{id}", func(ctx *context.Context) {}).
			Get("/haha/{id}/{name}", func(ctx *context.Context) {})
		a := New().AddRouter(r).parseRouters()
		require.Equal(t, 1, len(a.parsedRouters.Dynamics))
	}
}

func TestAddRouterGroup(t *testing.T) {
	{
		r := router.NewRouter().Get("/haha", func(ctx *context.Context) {})
		rg := router.NewGroup("").Add(r)
		a := New().AddRouterGroup(rg).parseRouters()
		require.Equal(t, 1, len(a.parsedRouters.Simples))
	}
	{
		r := router.NewRouter().Get("/haha", func(ctx *context.Context) {})
		rg := router.NewGroup("/v1").Add(r)
		a := New().AddRouterGroup(rg).parseRouters()
		require.Equal(t, 1, len(a.parsedRouters.Simples))
	}
}
