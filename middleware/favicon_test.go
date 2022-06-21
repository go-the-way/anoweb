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
	"embed"
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

//go:embed testdata
var fs embed.FS

func TestFavicon(t *testing.T) {
	{
		req, _ := http.NewRequest(http.MethodGet, "/favicon.ico", nil)
		c := Favicon("testdata/a.ico", nil)
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(c.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, []byte{}, r.Data)
		require.Equal(t, mime.ICO, r.ContentType)
	}
	{
		req, _ := http.NewRequest(http.MethodGet, "/favicon.ico", nil)
		c := Favicon("testdata/a.ico", &fs)
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(c.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, []byte{}, r.Data)
		require.Equal(t, mime.ICO, r.ContentType)
	}
}
