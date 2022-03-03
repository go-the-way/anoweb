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
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestStatic(t *testing.T) {
	// test for *.txt
	{
		req, _ := http.NewRequest(http.MethodGet, "/static/a.txt", nil)
		s := Static(false, "./testdata", "")
		ctx := context.New(req, &config.Template{})
		ctx.Add(s.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, mime.TEXT, r.ContentType)
		require.Equal(t, []byte("hello world"), r.Data)
	}
	// test for *.jpg
	{
		req, _ := http.NewRequest(http.MethodGet, "/static/a.jpg", nil)
		s := Static(false, "./testdata", "")
		ctx := context.New(req, &config.Template{})
		ctx.Add(s.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, mime.JPG, r.ContentType)
		require.Equal(t, []byte{}, r.Data)
	}
	// test for *.zip
	{
		req, _ := http.NewRequest(http.MethodGet, "/static/a.zip", nil)
		s := Static(false, "./testdata", "")
		ctx := context.New(req, &config.Template{})
		ctx.Add(s.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, mime.ZIP, r.ContentType)
		require.Equal(t, []byte{}, r.Data)
	}
}

func TestStaticMimeNotExists(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/static/a.txt", nil)
	s := Static(false, "./testdata", "")
	s.Mimes = map[string]string{}
	ctx := context.New(req, &config.Template{})
	ctx.Add(s.Handler())
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, mime.BINARY, r.ContentType)
	require.Equal(t, []byte("hello world"), r.Data)
}

func TestStaticCache(t *testing.T) {
	s := Static(true, "./testdata", "")
	{
		req, _ := http.NewRequest(http.MethodGet, "/static/a.txt", nil)
		ctx := context.New(req, &config.Template{})
		ctx.Add(s.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, mime.TEXT, r.ContentType)
		require.Equal(t, []byte("hello world"), r.Data)
	}
	{
		req, _ := http.NewRequest(http.MethodGet, "/static/a.txt", nil)
		ctx := context.New(req, &config.Template{})
		ctx.Add(s.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, mime.TEXT, r.ContentType)
		require.Equal(t, []byte("hello world"), r.Data)
	}
}

func TestStaticReadFile(t *testing.T) {
	s := Static(true, "./testdata", "")
	req, _ := http.NewRequest(http.MethodGet, "/static/a-haha.txt", nil)
	ctx := context.New(req, &config.Template{})
	ctx.Add(s.Handler())
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, mime.TEXT, r.ContentType)
	require.Equal(t, []byte("404 page not found"), r.Data)
}

func TestStaticAdd(t *testing.T) {
	s := Static(true, "./testdata", "")
	s.Add("polo", mime.TEXT)
	require.Equal(t, mime.TEXT, s.Mimes["polo"])
}

func TestStaticAdds(t *testing.T) {
	s := Static(true, "./testdata", "")
	s.Adds(map[string]string{"polo": mime.TEXT})
	require.Equal(t, mime.TEXT, s.Mimes["polo"])
}
