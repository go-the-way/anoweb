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
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func buildMultipartReq(ignoreExt bool) *http.Request {
	filename := "part"
	if !ignoreExt {
		filename += ".txt"
	}
	_ = ioutil.WriteFile(filename, []byte("hello"), 0700)
	defer func() { _ = os.Remove(filename) }()
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, _ := w.CreateFormFile("file", filename)
	file, _ := os.Open(filename)
	defer func() { _ = file.Close() }()
	_, _ = io.Copy(fw, file)
	_ = w.Close()
	req, _ := http.NewRequest(http.MethodPost, "", &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req
}

func TestUpload(t *testing.T) {
	req := buildMultipartReq(false)
	u := Upload()
	ctx := context.New()
	ctx.Allocate(req, &config.Template{})
	ctx.Add(u.Handler())
	ctx.Chain()
	_ = ctx.ParseMultipart(0)
	ff := ctx.MultipartFile("file")
	_ = ff.Copy("cp.txt")
	defer func() { _ = os.Remove("cp.txt") }()
	buf, _ := ioutil.ReadFile("cp.txt")
	require.Equal(t, []byte("hello"), buf)
}

func TestUploadVerifySize(t *testing.T) {
	req := buildMultipartReq(false)
	u := Upload()
	u.Size = 0
	ctx := context.New()
	ctx.Allocate(req, &config.Template{})
	ctx.Add(u.Handler())
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, []byte(fmt.Sprintf("{\"msg\":\"the file size exceed limit[max:%d, current:%d]\",\"code\":1}", 0, 5)), r.Data)
	require.Equal(t, mime.JSON, r.ContentType)
}

func TestUploadVerifyMime(t *testing.T) {
	req := buildMultipartReq(false)
	u := Upload()
	u.Mimes = []string{}
	ctx := context.New()
	ctx.Allocate(req, &config.Template{})
	ctx.Add(u.Handler())
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, []byte(fmt.Sprintf("{\"msg\":\"the file mime type not supported[supports:%s, current:%s]\",\"code\":1}", "", mime.BINARY)), r.Data)
	require.Equal(t, mime.JSON, r.ContentType)
}

func TestUploadVerifyExt(t *testing.T) {
	{
		req := buildMultipartReq(true)
		u := Upload()
		u.Extensions = []string{}
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(u.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, []byte(fmt.Sprintf("{\"msg\":\"the file ext not supported[supports:%s, current:%s]\",\"code\":1}", "", "")), r.Data)
		require.Equal(t, mime.JSON, r.ContentType)
	}
	{
		req := buildMultipartReq(false)
		u := Upload()
		u.Extensions = []string{}
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(u.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, []byte(fmt.Sprintf("{\"msg\":\"the file ext not supported[supports:%s, current:%s]\",\"code\":1}", "", ".txt")), r.Data)
		require.Equal(t, mime.JSON, r.ContentType)
	}
}

func TestUploadCreate(t *testing.T) {
	u := Upload()
	_ = u.Create("yo.txt", []byte("hello world"))
	require.Nil(t, u.Remove("yo.txt"))
}
