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
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	// test for no date dir
	{
		req, _ := http.NewRequest(http.MethodGet, "/?file=test-download.txt", nil)
		tmpFile := filepath.Join(os.TempDir(), time.Now().Format("20060102"), "test-download.txt")
		tDir := filepath.Dir(tmpFile)
		_ = os.MkdirAll(tDir, 0700)
		err := ioutil.WriteFile(tmpFile, []byte("hello world"), 0700)
		if err != nil {
			t.Error(err)
		}
		defer func() {
			_ = os.RemoveAll(tDir)
		}()
		d := Download()
		ctx := context.New(req, &config.Template{})
		ctx.Add(d.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, "attachment;filename="+url.QueryEscape("test-download.txt"), r.Header.Get(headers.ContentDisposition))
		require.Equal(t, mime.BINARY, r.ContentType)
		require.Equal(t, []byte("hello world"), r.Data)
	}
	// test for date dir
	{
		req, _ := http.NewRequest(http.MethodGet, "/?file=test-download.txt", nil)
		d := Download()
		tmpFile := filepath.Join(os.TempDir(), time.Now().Format("20060102"), "test-download.txt")
		err := ioutil.WriteFile(tmpFile, []byte("hello world"), 0700)
		if err != nil {
			t.Error(err)
		}
		defer func() {
			_ = os.Remove(tmpFile)
		}()
		ctx := context.New(req, &config.Template{})
		ctx.Add(d.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, "attachment;filename="+url.QueryEscape("test-download.txt"), r.Header.Get(headers.ContentDisposition))
		require.Equal(t, mime.BINARY, r.ContentType)
		require.Equal(t, []byte("hello world"), r.Data)
	}
}

func TestDownloadParamEmpty(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?file0=test-download.txt", nil)
	d := Download()
	ctx := context.New(req, &config.Template{})
	ctx.Add(d.Handler())
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, mime.JSON, r.ContentType)
	require.Equal(t, []byte(`{"msg":"file is empty","code":1}`), r.Data)
}

func TestDownloadFileNotExists(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?file=test-download.txt", nil)
	d := Download()
	ctx := context.New(req, &config.Template{})
	ctx.Add(d.Handler())
	ctx.Chain()
	r := ctx.Response
	require.Equal(t, mime.JSON, r.ContentType)
	require.Equal(t, []byte(`{"msg":"file is not exists","code":1}`), r.Data)
}

func TestDownloadPath(t *testing.T) {
	{
		dir1 := filepath.Join(os.TempDir(), "test-download.txt")
		d := Download()
		d.DateDir = false
		dir2 := d.Path("test-download.txt")
		require.Equal(t, dir1, dir2)
	}
	{
		dir1 := filepath.Join(os.TempDir(), time.Now().Format("20060102"), "test-download.txt")
		dir2 := Download().Path("test-download.txt")
		require.Equal(t, dir1, dir2)
	}
}

func TestDownloadFile(t *testing.T) {
	file, err := Download().File("test-download.txt")
	require.Nil(t, file)
	require.NotNil(t, err)
}

func TestDownloadBuf(t *testing.T) {
	buf, err := Download().Buf("test-download.txt")
	require.Nil(t, buf)
	require.NotNil(t, err)
}
