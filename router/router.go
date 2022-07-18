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
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/util"

	"github.com/gorilla/mux"
)

var (
	dynamicRouteRe   = regexp.MustCompile(`{([^{]+)}`)
	dynamicParamRep  = `([\w_.-]+)`
	supportedMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodHead,
		http.MethodOptions,
	}
)

type router = mux.Router

// Router struct
type Router struct{ *router }

// NewRouter return new router
func NewRouter() *Router { return &Router{mux.NewRouter()} }

func (r *Router) Req(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route("*", pattern, handler)
}

// Get Route Get Method
func (r *Router) Get(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodGet, pattern, handler)
}

// Post Route Post Method
func (r *Router) Post(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodPost, pattern, handler)
}

// Put Route Put Method
func (r *Router) Put(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodPut, pattern, handler)
}

// Delete Route Delete Method
func (r *Router) Delete(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodDelete, pattern, handler)
}

// Patch Route Patch Method
func (r *Router) Patch(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodPatch, pattern, handler)
}

// Head Route Head Method
func (r *Router) Head(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodHead, pattern, handler)
}

// Options Route Options Method
func (r *Router) Options(pattern string, handler func(ctx *context.Context)) *Router {
	return r.Route(http.MethodOptions, pattern, handler)
}

// Resource Route a resource
func (r *Router) Resource(pattern, file, contentType string) *Router {
	return r.Get(pattern, func(ctx *context.Context) {
		ctx.File(file, contentType)
	})
}

// FSResource Route a resource
func (r *Router) FSResource(fs *embed.FS, pattern, file, contentType string) *Router {
	return r.Get(pattern, func(ctx *context.Context) {
		ctx.FSFile(fs, file, contentType)
	})
}

// Route Route DIY Method
func (r *Router) Route(method, pattern string, handler func(ctx *context.Context)) *Router {
	r.mustSupport(method)
	pattern = util.TrimSpecialChars(pattern)
	if dynamicRouteRe.MatchString(pattern) {
		r.dynamicRoute(method, pattern, handler)
	} else {
		r.simpleRoute(method, pattern, handler)
	}

	return r
}

func (r *Router) mustSupport(method string) {
	if method == "*" {
		return
	}
	for _, m := range supportedMethods {
		if m == method {
			return
		}
	}
	panic(errors.New("method not supported : " + method))
}

func (r *Router) simpleRoute(method, pattern string, handler func(ctx *context.Context)) *Router {
	var methods []string
	if method == "*" {
		methods = append(methods, supportedMethods...)
	} else {
		methods = append(methods, method)
	}
	for _, m := range methods {
		r.Simples = append(r.Simples, &Simple{m, pattern, handler})
	}
	return r
}

func (r *Router) dynamicRoute(method, pattern string, handler func(ctx *context.Context)) *Router {
	finds := dynamicRouteRe.FindAllStringSubmatch(pattern, -1)
	params := make([]string, len(finds))
	for i, f := range finds {
		f0 := f[0]
		f1 := f[1]
		pattern = strings.Replace(pattern, f0, dynamicParamRep, 1)
		params[i] = f1
	}
	methods := make([]string, 0)
	if method == "*" {
		methods = append(methods, supportedMethods...)
	} else {
		methods = append(methods, method)
	}
	for _, m := range methods {
		r.Dynamics = append(r.Dynamics, &Dynamic{params, &Simple{m, pattern, handler}})
	}
	return r
}
