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
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"

	"github.com/stretchr/testify/require"
)

func TestParsedRouterHandler(t *testing.T) {
	// test for no simple
	{
		req, _ := http.NewRequest(http.MethodGet, "/index", nil)
		pass := false
		pr := &ParsedRouter{Simples: SimpleM{"GET:": {http.MethodGet, "/", func(ctx *context.Context) { pass = true }}}}
		ctx := context.New(req, &config.Template{})
		handler := pr.Handler(ctx)
		if handler != nil {
			handler(ctx)
		}
		require.Equal(t, false, pass)
	}
	// test for simple
	{
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		pass := false
		pr := &ParsedRouter{Simples: SimpleM{"GET:": {http.MethodGet, "/", func(ctx *context.Context) { pass = true }}}}
		ctx := context.New(req, &config.Template{})
		handler := pr.Handler(ctx)
		if handler != nil {
			handler(ctx)
		}
		require.Equal(t, true, pass)
	}
	// test for dynamic
	{
		pass := false
		req, _ := http.NewRequest(http.MethodGet, "/index/apple", nil)
		pr := &ParsedRouter{Dynamics: DynamicM{http.MethodGet: {`^/index/([\w_.-]+)$`: {[]string{"id"}, &Simple{http.MethodGet, "", func(ctx *context.Context) { pass = true }}}}}}
		ctx := context.New(req, &config.Template{})
		handler := pr.Handler(ctx)
		if handler != nil {
			handler(ctx)
		}
		require.Equal(t, true, pass)
		require.Equal(t, "apple", ctx.Param("id"))
	}
	// test for dynamic #2
	{
		pass := false
		req, _ := http.NewRequest(http.MethodGet, "/index/apple/pear", nil)
		pr := &ParsedRouter{Dynamics: DynamicM{http.MethodGet: {`^/index/([\w_.-]+)/([\w_.-]+)$`: {[]string{"id1", "id2"}, &Simple{http.MethodGet, "", func(ctx *context.Context) { pass = true }}}}}}
		ctx := context.New(req, &config.Template{})
		handler := pr.Handler(ctx)
		if handler != nil {
			handler(ctx)
		}
		require.Equal(t, true, pass)
		require.Equal(t, "apple", ctx.Param("id1"))
		require.Equal(t, "pear", ctx.Param("id2"))
	}
	// test for no dynamic
	{
		pass := false
		req, _ := http.NewRequest(http.MethodGet, "/index/apple", nil)
		pr := &ParsedRouter{Dynamics: DynamicM{http.MethodGet: {`^/index/([\w_.-]+)/([\w_.-]+)$`: {[]string{"id1", "id2"}, &Simple{http.MethodGet, "", func(ctx *context.Context) { pass = true }}}}}}
		ctx := context.New(req, &config.Template{})
		handler := pr.Handler(ctx)
		if handler != nil {
			handler(ctx)
		}
		require.Equal(t, false, pass)
	}
}
