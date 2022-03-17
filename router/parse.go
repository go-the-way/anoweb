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
	"fmt"
	"regexp"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/util"
)

type (
	// SimpleM type define map[string]*Simple
	SimpleM = map[string]*Simple
	// DynamicM type define map[string]*Simple
	DynamicM = map[string]map[string]*Dynamic
	// ParsedRouter defines parsed router
	ParsedRouter struct {
		// Simples routers
		Simples SimpleM
		// Dynamics routers K<Method> V< K<Pattern> V<Simple> >
		Dynamics DynamicM
	}
)

// Handler found simple or dynamic handler
func (pr *ParsedRouter) Handler(ctx *context.Context) func(ctx *context.Context) {
	simple := pr.simple(ctx)
	if simple != nil {
		return simple
	}
	return pr.dynamic(ctx)
}

func (pr *ParsedRouter) simple(ctx *context.Context) func(ctx *context.Context) {
	if len(pr.Simples) <= 0 {
		return nil
	}
	routeKey := fmt.Sprintf("%s:%s", ctx.Request.Method, util.ReBuildPath(ctx.Request.URL.Path))
	if simple, have := pr.Simples[routeKey]; have {
		return simple.Handler
	}
	return nil
}

func (pr *ParsedRouter) dynamic(ctx *context.Context) func(ctx *context.Context) {
	if len(pr.Dynamics) <= 0 {
		return nil
	}
	if mp, have := pr.Dynamics[ctx.Request.Method]; have {
		for k, v := range mp {
			re := regexp.MustCompile(k)
			reqPath := util.ReBuildPath(ctx.Request.URL.Path)
			finds := re.FindAllStringSubmatch(reqPath, -1)
			if len(finds) > 0 {
				subFinds := finds[0]
				if len(subFinds) == len(v.Params)+1 {
					paramsMap := make(map[string][]string, len(v.Params))
					for i, f := range subFinds[1:] {
						paramsMap[v.Params[i]] = []string{f}
					}
					ctx.SetParamMap(paramsMap, false)
					return v.Handler
				}
			}
		}
	}
	return nil
}
