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
	"compress/flate"
	"strings"

	"github.com/go-the-way/anoweb/context"
)

type deflate struct {
	ContentType []string
	MinSize     int
	Level       int
}

// Deflate return new deflate Middleware
func Deflate() *deflate {
	return &deflate{
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
		Level:   flate.BestSpeed,
	}
}

func (d *deflate) accept(ctx *context.Context) bool {
	acceptEncoding := ctx.Request.Header.Get("Accept-Encoding")
	return strings.Contains(acceptEncoding, "deflate")
}

func (d *deflate) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		ctx.Chain()
		if !d.accept(ctx) {
			return
		}
		r := ctx.Response
		if odata := r.Data; nil != odata && len(odata) >= d.MinSize {
			ctx.Response.Header.Set("Vary", "Content-Encoding")
			ctx.Response.Header.Set("Content-Encoding", "deflate")
			var buffers bytes.Buffer
			fw, _ := flate.NewWriter(&buffers, d.Level)
			defer func() { _ = fw.Close() }()
			_, _ = fw.Write(odata)
			_ = fw.Flush()
			ctx.Data(buffers.Bytes())
		}
	}
}
