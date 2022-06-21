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
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.Text(`hello world`)
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.TEXT, r.ContentType)
}

func TestTextFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte("hello world"), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.TextFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.TEXT, r.ContentType)
}

func TestJSON(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSON(map[string]interface{}{"apple": 100})
	r := ctx.Response
	require.Equal(t, []byte(`{"apple":100}`), r.Data)
	require.Equal(t, mime.JSON, r.ContentType)
}

func TestJSONPanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSON(map[string]interface{}{"_": func() {}})
}

func TestJSONText(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSONText(`{"apple":100}`)
	r := ctx.Response
	require.Equal(t, []byte(`{"apple":100}`), r.Data)
	require.Equal(t, mime.JSON, r.ContentType)
}

func TestJSONFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`{"apple":100}`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSONFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`{"apple":100}`), r.Data)
	require.Equal(t, mime.JSON, r.ContentType)
}

func TestXML(t *testing.T) {
	type _xml struct {
		Apple string `xml:"apple"`
	}
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.XML(&_xml{"100"})
	r := ctx.Response
	require.Equal(t, []byte(`<_xml><apple>100</apple></_xml>`), r.Data)
	require.Equal(t, mime.XML, r.ContentType)
}

func TestXMLPanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.XML([]byte(`hello world`))
}

func TestXMLText(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.XMLText(`<_xml><apple>100</apple></_xml>`)
	r := ctx.Response
	require.Equal(t, []byte(`<_xml><apple>100</apple></_xml>`), r.Data)
	require.Equal(t, mime.XML, r.ContentType)
}

func TestXMLFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`<_xml><apple>100</apple></_xml>`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.XMLFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`<_xml><apple>100</apple></_xml>`), r.Data)
	require.Equal(t, mime.XML, r.ContentType)
}

func TestHTML(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.HTML(`<html><body>hello world</body></html>`)
	r := ctx.Response
	require.Equal(t, []byte(`<html><body>hello world</body></html>`), r.Data)
	require.Equal(t, mime.HTML, r.ContentType)
}

func TestHTMLFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`<html><body>hello world</body></html>`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.HTMLFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`<html><body>hello world</body></html>`), r.Data)
	require.Equal(t, mime.HTML, r.ContentType)
}

func TestCSS(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.CSS(`div{color:red;}`)
	r := ctx.Response
	require.Equal(t, []byte(`div{color:red;}`), r.Data)
	require.Equal(t, mime.CSS, r.ContentType)
}

func TestCSSFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`div{color:red;}`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.CSSFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`div{color:red;}`), r.Data)
	require.Equal(t, mime.CSS, r.ContentType)
}

func TestJS(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JS(`alert(100);`)
	r := ctx.Response
	require.Equal(t, []byte(`alert(100);`), r.Data)
	require.Equal(t, mime.JS, r.ContentType)
}

func TestJSFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`alert(100);`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.JSFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte(`alert(100);`), r.Data)
	require.Equal(t, mime.JS, r.ContentType)
}
