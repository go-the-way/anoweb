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
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/router"
	"net/http"
)

// Request Route all Methods
func (a *App) Request(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route("*", pattern, handler)
}

// Get Route Get Method
func (a *App) Get(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodGet, pattern, handler)
}

// Post Route Post Method
func (a *App) Post(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodPost, pattern, handler)
}

// Put Route Put Method
func (a *App) Put(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodPut, pattern, handler)
}

// Delete Route Delete Method
func (a *App) Delete(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodDelete, pattern, handler)
}

// Patch Route Patch Method
func (a *App) Patch(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodPatch, pattern, handler)
}

// Head Route Head Method
func (a *App) Head(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodHead, pattern, handler)
}

// Options Route Options Method
func (a *App) Options(pattern string, handler func(ctx *context.Context)) *App {
	return a.Route(http.MethodOptions, pattern, handler)
}

// Route Route DIY Method
func (a *App) Route(method, pattern string, handler func(ctx *context.Context)) *App {
	a.routers[0].Route(method, pattern, handler)
	return a
}

// Resource Route a resource
func (a *App) Resource(pattern, file, contentType string) *App {
	return a.Route(http.MethodGet, pattern, func(ctx *context.Context) {
		ctx.File(file, contentType)
	})
}

// FSResource Route a resource
func (a *App) FSResource(fs *embed.FS, pattern, file, contentType string) *App {
	return a.Route(http.MethodGet, pattern, func(ctx *context.Context) {
		ctx.FSFile(fs, file, contentType)
	})
}

// AddRouter Add Routers
func (a *App) AddRouter(r ...*router.Router) *App {
	a.routers = append(a.routers, r...)
	return a
}

// AddRouterGroup Add Group Routers
func (a *App) AddRouterGroup(g ...*router.Group) *App {
	a.groups = append(a.groups, g...)
	return a
}

func (a *App) simpleParseFunc(prefix string, simples []*router.Simple) {
	for _, simple := range simples {
		routeKey := fmt.Sprintf("%s:%s%s", simple.Method, prefix, simple.Pattern)
		a.parsedRouters.Simples[routeKey] = simple
	}
}

func (a *App) dynamicParseFunc(prefix string, dynamics []*router.Dynamic) {
	for _, d := range dynamics {
		pattern := fmt.Sprintf("^%s%s$", prefix, d.Pattern)
		if mm, have := a.parsedRouters.Dynamics[d.Method]; have {
			mm[pattern] = d
		} else {
			a.parsedRouters.Dynamics[d.Method] = map[string]*router.Dynamic{pattern: d}
		}
	}
}

func (a *App) parseRouters() *App {
	for _, g := range a.groups {
		for _, gr := range g.Routers() {
			if g.Prefix() == "" {
				a.routers = append(a.routers, gr)
			} else {
				a.simpleParseFunc(g.Prefix(), gr.Simples)
				a.dynamicParseFunc(g.Prefix(), gr.Dynamics)
			}
		}
	}
	for _, r := range a.routers {
		a.simpleParseFunc("", r.Simples)
		a.dynamicParseFunc("", r.Dynamics)
	}
	return a
}
