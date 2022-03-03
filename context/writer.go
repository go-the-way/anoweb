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
	"net/http"
)

// Response struct
type Response struct {
	Data        []byte
	Header      http.Header
	Cookies     []*http.Cookie
	Status      int
	ContentType string
}

// Write from r
func (ctx *Context) Write(r *Response) {
	if r.Data != nil {
		ctx.Response.Data = r.Data
	}
	if r.ContentType != "" {
		ctx.Response.ContentType = r.ContentType
	}
	if r.Header != nil && len(r.Header) > 0 {
		for k, v := range r.Header {
			ctx.Response.Header[k] = v
		}
	}
	if r.Cookies != nil && len(r.Cookies) > 0 {
		for _, cookie := range r.Cookies {
			ctx.Response.Cookies = append(ctx.Response.Cookies, cookie)
		}
	}
	if r.Status > 0 {
		ctx.Response.Status = r.Status
	}
}
