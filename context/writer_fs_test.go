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
	"embed"
	"testing"

	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

//go:embed testdata
var rFS embed.FS

func TestTextFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.TextFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.TEXT, r.ContentType)
}

func TestJSONFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSONFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.JSON, r.ContentType)
}

func TestXMLFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.XMLFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.XML, r.ContentType)
}

func TestImageFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.ImageFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestICOFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.ICOFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.ICO, r.ContentType)
}

func TestBMPFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.BMPFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.BMP, r.ContentType)
}

func TestJPGFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JPGFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestJPEGFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JPEGFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestPNGFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.PNGFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.PNG, r.ContentType)
}

func TestGIFFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.GIFFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.GIF, r.ContentType)
}

func TestHTMLFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.HTMLFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.HTML, r.ContentType)
}

func TestCSSFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.CSSFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.CSS, r.ContentType)
}

func TestJSFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSFSFile(&rFS, "testdata/file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.JS, r.ContentType)
}

func TestFSFile(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.FSFile(&rFS, "testdata/file.txt", mime.TEXT)
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.TEXT, r.ContentType)
}

func TestFSFilePanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.FSFile(&rFS, "testdata/file-haha.txt", mime.TEXT)
}
