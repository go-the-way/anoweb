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
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"

	"github.com/stretchr/testify/require"
)

func buildReq(ignoreAE bool) *http.Request {
	req, _ := http.NewRequest("", "", nil)
	req.Header.Set("Accept", "*")
	if !ignoreAE {
		req.Header.Set("Accept-Encoding", "gzip, deflate")
	}
	req.Header.Set("Connection", "keep-alive")
	return req
}

func TestDeflate(t *testing.T) {
	data := `0123456789 0123456789 0123456789 0123456789 0123456789 0123456789`
	d := Deflate()
	d.MinSize = 10
	ctx := context.New(buildReq(false), &config.Template{})
	ctx.Add(d.Handler())
	ctx.Add(func(ctx *context.Context) {
		ctx.Text(data)
	})
	ctx.Chain()
	r := ctx.Response
	rd := flate.NewReader(bytes.NewReader(r.Data))
	defer func() { _ = rd.Close() }()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(rd)
	require.Equal(t, data, buf.String())
	require.Equal(t, "Content-Encoding", ctx.Response.Header.Get("Vary"))
	require.Equal(t, "deflate", ctx.Response.Header.Get("Content-Encoding"))
}

func TestDeflateWithoutDeflate(t *testing.T) {
	data := `0123456789 0123456789 0123456789 0123456789 0123456789 0123456789`
	d := Deflate()
	d.MinSize = 10
	ctx := context.New(buildReq(true), &config.Template{})
	ctx.Add(d.Handler())
	ctx.Add(func(ctx *context.Context) {
		ctx.Text(data)
	})
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, []byte(data), r.Data)
	require.Equal(t, "", ctx.Response.Header.Get("Vary"))
	require.Equal(t, "", ctx.Response.Header.Get("Content-Encoding"))
}
