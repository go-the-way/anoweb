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
	"errors"
	"strconv"
)

var (
	ErrNoSuchQueryParam = errors.New("query: no such query param")
)

func (ctx *Context) Query(name string) string {
	return ctx.QueryDef(name, "")
}

func (ctx *Context) QueryDef(name, defVal string) string {
	if v, have := ctx.Request.Form[name]; have {
		if v != nil && len(v) > 0 {
			return v[0]
		}
	}
	return defVal
}

func (ctx *Context) Queries(name string) []string {
	if v, have := ctx.Request.Form[name]; have {
		if v != nil {
			return v
		}
	}
	return nil
}

func (ctx *Context) QueriesDef(name string, defVal []string) []string {
	if v, have := ctx.Request.Form[name]; have {
		if v != nil {
			return v
		}
	}
	return defVal
}

func (ctx *Context) IntQueryParam(name string) (int, error) {
	if !ctx.HaveQueryParam(name) {
		return 0, ErrNoSuchQueryParam
	}
	v := ctx.Query(name)
	return strconv.Atoi(v)
}

func (ctx *Context) IntQueryParamDef(name string, defVal int) int {
	if v, err := ctx.IntQueryParam(name); err != nil {
		return defVal
	} else {
		return v
	}
}

func (ctx *Context) HaveQueryParam(name string) bool {
	_, have := ctx.Request.Form[name]
	return have
}

func (ctx *Context) Rest() string {
	return ctx.Param("RESTFUL_KEY")
}

func (ctx *Context) IntRest() (int64, rror) {
	return strconv.ParseInt(ctx.Rest(), 10, 64)
}
