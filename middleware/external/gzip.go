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

package external

import (
	"bytes"
	gz "compress/gzip"
	"strings"

	"github.com/go-the-way/anoweb/context"
)

type gzip struct {
	ContentType []string
	MinSize     int
	Level       int
}

func Gzip() *gzip {
	return &gzip{
		ContentType: []string{
			"application/javascript",
			"application/json",
			"application/xml",
			"text/javascript",
			"text/json",
			"text/xml",
			"text/plain",
			"text/xml",
			"html/css",
		},
		MinSize: 2 << 9, //1KB
		Level:   gz.BestSpeed,
	}
}

func (g *gzip) accept(ctx *context.Context) bool {
	acceptEncoding := ctx.Request.Header.Get("Accept-Encoding")
	return strings.Contains(acceptEncoding, "gzip")
}

func (g *gzip) Handler() func(c *context.Context) {
	return func(ctx *context.Context) {
		ctx.Chain()
		if !g.accept(ctx) {
			return
		}
		r := ctx.Response
		if odata := r.Data; odata != nil && len(odata) >= g.MinSize {
			ctx.Response.Header.Set("Vary", "Content-Encoding")
			ctx.Response.Header.Set("Content-Encoding", "gzip")
			var buffers bytes.Buffer
			gw, _ := gz.NewWriterLevel(&buffers, g.Level)
			defer func() { _ = gw.Close() }()
			_, _ = gw.Write(odata)
			_ = gw.Flush()
			ctx.Data(buffers.Bytes())
		}
	}
}
