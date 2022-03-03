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

	"github.com/go-the-way/anoweb/mime"
)

// ResponseBuilder struct
type ResponseBuilder struct {
	r *Response
}

// Build Response
func (r *ResponseBuilder) Build() *Response {
	return r.r
}

// Builder Response
func Builder() *ResponseBuilder {
	return &ResponseBuilder{r: &Response{Status: http.StatusOK, ContentType: mime.TEXT}}
}

// DefaultBuild Response
func (r *ResponseBuilder) DefaultBuild() *Response {
	return r.Header(http.Header{}).Cookies(make([]*http.Cookie, 0)).Status(http.StatusNotFound).Data([]byte("404 page not found")).ContentType(mime.TEXT).Build()
}

// Data Response
func (r *ResponseBuilder) Data(data []byte) *ResponseBuilder {
	r.r.Data = data
	return r
}

// Header Response
func (r *ResponseBuilder) Header(header http.Header) *ResponseBuilder {
	r.r.Header = header
	return r
}

// Cookies Response
func (r *ResponseBuilder) Cookies(cookies []*http.Cookie) *ResponseBuilder {
	r.r.Cookies = cookies
	return r
}

// Status Response
func (r *ResponseBuilder) Status(status int) *ResponseBuilder {
	r.r.Status = status
	return r
}

// ContentType Response
func (r *ResponseBuilder) ContentType(contentType string) *ResponseBuilder {
	r.r.ContentType = contentType
	return r
}
