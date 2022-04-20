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
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/go-the-way/anoweb/config"

	"github.com/stretchr/testify/require"
)

func buildMultipartReq() *http.Request {
	filename := "part.txt"
	_ = ioutil.WriteFile(filename, []byte("hello"), 0700)
	defer func() { _ = os.Remove(filename) }()
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	vw, _ := w.CreateFormField("value")
	_, _ = vw.Write([]byte(`hello world`))
	fw, _ := w.CreateFormFile("file", filename)
	file, _ := os.Open(filename)
	defer func() { _ = file.Close() }()
	_, _ = io.Copy(fw, file)
	_ = w.Close()
	req, _ := http.NewRequest(http.MethodPost, "", &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req
}

func TestMultipartFileNil(t *testing.T) {
	ctx := New(buildReq(""), &config.Template{})
	_ = ctx.ParseMultipart(0)
	require.Nil(t, ctx.MultipartFile("file"))
}

func TestMultipartFileCopy(t *testing.T) {
	ctx := New(buildMultipartReq(), &config.Template{})
	defer func() { _ = ctx.Request.MultipartForm.RemoveAll() }()
	_ = ctx.ParseMultipart(0)
	ff := ctx.MultipartFile("file")
	_ = ff.Copy("cp.txt")
	defer func() { _ = os.Remove("cp.txt") }()
	bs, _ := ioutil.ReadFile("cp.txt")
	require.Equal(t, []byte("hello"), bs)
}

func TestMultipartFileCopyErr(t *testing.T) {
	ctx := New(buildMultipartReq(), &config.Template{})
	defer func() { _ = ctx.Request.MultipartForm.RemoveAll() }()
	_ = ctx.ParseMultipart(0)
	ff := ctx.MultipartFile("file")
	_ = ctx.Request.MultipartForm.RemoveAll()
	if err := ff.Copy("cp.txt"); err == nil {
		t.Error("test failed")
	}
}

func TestParseMultipart(t *testing.T) {
	filename := "part.txt"
	_ = ioutil.WriteFile(filename, []byte("hello"), 0700)
	defer func() { _ = os.Remove(filename) }()
	ctx := New(buildMultipartReq(), &config.Template{})
	defer func() { _ = ctx.Request.MultipartForm.RemoveAll() }()
	_ = ctx.ParseMultipart(0)
	ff := ctx.MultipartFile("file")
	require.NotNil(t, ff)
}

func TestMultipartFile(t *testing.T) {
	filename := "part.txt"
	_ = ioutil.WriteFile(filename, []byte("hello"), 0700)
	defer func() { _ = os.Remove(filename) }()
	ctx := New(buildMultipartReq(), &config.Template{})
	defer func() { _ = ctx.Request.MultipartForm.RemoveAll() }()
	_ = ctx.ParseMultipart(0)
	ff := ctx.MultipartFile("file")
	require.NotNil(t, ff)
}

func TestMultipartFiles(t *testing.T) {
	filename := "part.txt"
	_ = ioutil.WriteFile(filename, []byte("hello"), 0700)
	defer func() { _ = os.Remove(filename) }()
	ctx := New(buildMultipartReq(), &config.Template{})
	defer func() { _ = ctx.Request.MultipartForm.RemoveAll() }()
	_ = ctx.ParseMultipart(0)
	ff := ctx.MultipartFiles("file")
	require.Equal(t, 1, len(ff))
}
