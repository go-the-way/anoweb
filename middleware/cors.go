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

package middleware

import (
	"net/http"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"
)

type cors struct {
	Origin       string
	Methods      []string
	AllowHeaders []string
	Header       http.Header
}

// Cors return cors
func Cors() *cors {
	return &cors{
		Origin:       "*",
		Methods:      []string{"GET, POST, DELETE, PUT, PATCH, HEAD, OPTIONS"},
		AllowHeaders: make([]string, 0),
		Header:       make(http.Header, 0),
	}
}

// Handler implements
func (cs *cors) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		cs.Header.Set(headers.AccessControlAllowOrigin, cs.Origin)
		cs.Header[headers.Allow] = cs.Methods
		cs.Header[headers.AccessControlAllowHeaders] = cs.AllowHeaders
		cs.Header[headers.AccessControlAllowMethods] = cs.Methods
		for k, v := range cs.Header {
			for _, vv := range v {
				ctx.Response.Header.Add(k, vv)
			}
		}
		if ctx.Request.Method != http.MethodHead && ctx.Request.Method != http.MethodOptions {
			ctx.Chain()
		} else {
			ctx.Data([]byte{}).Status(http.StatusOK)
		}
	}
}
