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
	"fmt"
	"strconv"
	"strings"
)

// ParamMap return param map
func (ctx *Context) ParamMap() map[string][]string {
	return ctx.paramMap
}

// Param return named param
func (ctx *Context) Param(name string) string {
	return ctx.ParamDefault(name, "")
}

// ParamDefault return named param
func (ctx *Context) ParamDefault(name, defaultVal string) string {
	params, have := ctx.paramMap[name]
	if !have || params == nil || len(params) <= 0 || params[0] == "" {
		return defaultVal
	}
	return params[0]
}

// Params return named param values
func (ctx *Context) Params(name string, defaultVal []string) []string {
	params := ctx.paramMap[name]
	if params != nil {
		return params
	}
	return defaultVal
}

// HasParam return true if named param in param map
func (ctx *Context) HasParam(name string) bool {
	return ctx.ParamMap()[name] != nil
}

// IntParam return named param of int with default `defaultVal`
func (ctx *Context) IntParam(name string, defaultVal int64) int64 {
	if !ctx.HasParam(name) {
		return defaultVal
	}
	intVal, err := strconv.ParseInt(ctx.ParamDefault(name, fmt.Sprintf("%d", defaultVal)), 10, 64)
	if err != nil {
		intVal = defaultVal
	}
	return intVal
}

// FloatParam return named param of float
func (ctx *Context) FloatParam(name string, defaultVal float64) float64 {
	if !ctx.HasParam(name) {
		return defaultVal
	}
	floatVal, err := strconv.ParseFloat(ctx.ParamDefault(name, fmt.Sprintf("%f", defaultVal)), 64)
	if err != nil {
		floatVal = defaultVal
	}
	return floatVal
}

// transformParamMap transform param map
func (ctx *Context) transformParamMap(multiFunc func(name string, params []string) string) map[string]string {
	sm := make(map[string]string, 0)
	for k, v := range ctx.paramMap {
		sm[k] = multiFunc(k, v)
	}
	return sm
}

// SingleParamMap return single value param map
func (ctx *Context) SingleParamMap() map[string]string {
	return ctx.transformParamMap(func(name string, params []string) string {
		if params != nil && len(params) > 0 {
			return params[0]
		}
		return ""
	})
}

// JoinedParamMap return joined single value param map
func (ctx *Context) JoinedParamMap(separator string) map[string]string {
	return ctx.transformParamMap(func(name string, params []string) string {
		if params != nil && len(params) > 0 {
			return strings.Join(params, separator)
		}
		return ""
	})
}

// SetParamMap set param map
func (ctx *Context) SetParamMap(paramMap map[string][]string, flush bool) *Context {
	if flush {
		for k := range ctx.paramMap {
			delete(ctx.paramMap, k)
		}
	}
	if paramMap != nil && len(paramMap) > 0 {
		for k, v := range paramMap {
			ctx.paramMap[k] = v
		}
	}
	return ctx
}

// Key return REST-ful key
func (ctx *Context) Key() string {
	return ctx.Param("RESTFUL_KEY")
}

// IntKey return int REST-ful key
func (ctx *Context) IntKey() int64 {
	intID, err := strconv.ParseInt(ctx.Key(), 10, 64)
	if err != nil {
		panic(err)
	}
	return intID
}
